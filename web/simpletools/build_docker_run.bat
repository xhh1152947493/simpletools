@echo off
setlocal

rem 设置镜像名称
set IMAGE_NAME=simpletools-wen

for /f "tokens=2 delims==" %%i in ('"wmic os get localdatetime /value"') do set datetime=%%i

set year=%datetime:~0,4%
set month=%datetime:~4,2%
set day=%datetime:~6,2%
set hour=%datetime:~8,2%
set minute=%datetime:~10,2%
set second=%datetime:~12,2%

set timestamp=%year%%month%%day%%hour%%minute%%second%

echo %timestamp%

rem 构建 Docker 镜像
docker build -f ./Dockerfile --build-arg TIMESTAMP=%timestamp% -t %IMAGE_NAME% .

rem 检查构建是否成功
if ERRORLEVEL 1 (
    echo Docker build failed!
    exit /b 1
)

rem 重启更新后的镜像
set CONTAINER_NAME=simpletools-web
set PORT=1234

rem 获取容器信息
for /f "delims=" %%i in ('docker inspect --format="{{.Created}}" %CONTAINER_NAME%') do set CREATED=%%i

rem 获取旧容器的镜像ID
for /f "delims=" %%i in ('docker inspect --format="{{.Image}}" %CONTAINER_NAME%') do set OLD_IMAGE_ID=%%i

rem 打印创建时间和版本号
echo old container %CONTAINER_NAME% was created at: %CREATED%
rem 停止并删除旧容器-优雅退出，日志需要挂在在主机上，不然rm会删除容器丢失日志
echo stop old container:
docker stop %CONTAINER_NAME% || echo No container to stop.
echo remove old container:
docker rm %CONTAINER_NAME% || echo No container to remove.

rem 为旧镜像添加新的标签
echo Tagging old image with timestamp:%timestamp%
docker tag %OLD_IMAGE_ID% %IMAGE_NAME%:%timestamp%

rem 启动新的容器
echo start new container %CONTAINER_NAME%:
docker run -d --name %CONTAINER_NAME% -p %PORT%:%PORT% -v web-simpletools:/var/log/nginx/ %IMAGE_NAME%
echo new container %CONTAINER_NAME% started on port %PORT%.
