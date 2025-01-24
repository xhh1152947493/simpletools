<template>
    <div class="http-tool-container">
        <el-card class="http-tool-card">
            <!-- 操作栏 -->
            <div class="operation-bar">
                <el-select v-model="requestMethod" placeholder="选择请求方法" style="width: 120px; margin-right: 15px;">
                    <el-option label="GET" value="GET" />
                    <el-option label="POST" value="POST" />
                </el-select>

                <el-input v-model="requestUrl" placeholder="请输入请求 URL" style="flex: 1; margin-right: 15px;" />

                <el-button type="primary" :loading="loading" @click="handleSendRequest">
                    <template #icon>
                        <el-icon>
                            <Promotion />
                        </el-icon>
                    </template>
                    发送请求
                </el-button>
            </div>

            <!-- 请求头 -->
            <div class="section">
                <div class="section-title">请求头</div>
                <el-input v-model="requestHeaders" type="textarea" placeholder="请输入请求头（JSON 格式）" :rows="3" resize="none"
                    @keydown="requestHeaders = handleJsonIndent($event, requestHeaders)" />
            </div>

            <!-- 请求体 -->
            <div v-if="requestMethod === 'POST'" class="section">
                <div class="section-title">请求体</div>
                <el-input v-model="requestBody" type="textarea" placeholder="请输入请求体（JSON 格式）" :rows="5" resize="none"
                    @keydown="requestHeaders = handleJsonIndent($event, requestHeaders)" />
            </div>

            <!-- 响应结果 -->
            <div class="section">
                <div class="section-title">响应结果</div>
                <el-input v-model="responseResult" type="textarea" placeholder="响应结果将显示在这里" :rows="10" resize="none"
                    readonly />
            </div>
        </el-card>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Promotion } from '@element-plus/icons-vue'

// 响应式数据
const requestMethod = ref('GET') // 请求方法
const requestUrl = ref('') // 请求 URL
const requestHeaders = ref('') // 请求头
const requestBody = ref('') // 请求体
const responseResult = ref('') // 响应结果
const loading = ref(false) // 加载状态

// 发送请求
const handleSendRequest = async () => {
    if (!requestUrl.value.trim()) {
        ElMessage.warning('请输入请求 URL')
        return
    }

    try {
        loading.value = true

        // 解析请求头
        let headers = {}
        if (requestHeaders.value.trim()) {
            try {
                headers = JSON.parse(requestHeaders.value)
            } catch (error) {
                ElMessage.error('请求头必须是合法的 JSON 格式')
                return
            }
        }

        // 解析请求体
        let body = null
        if (requestMethod.value === 'POST' && requestBody.value.trim()) {
            try {
                body = JSON.parse(requestBody.value)
            } catch (error) {
                ElMessage.error('请求体必须是合法的 JSON 格式')
                return
            }
        }

        // 发送请求
        const response = await fetch(requestUrl.value, {
            method: requestMethod.value,
            headers: {
                'Content-Type': 'application/json',
                ...headers,
            },
            body: requestMethod.value === 'POST' ? JSON.stringify(body) : undefined,
        })

        // 处理响应
        const contentType = response.headers.get('content-type')
        if (contentType && contentType.includes('application/json')) {
            const data = await response.json()
            responseResult.value = JSON.stringify(data, null, 2)
        } else if (contentType && contentType.includes('text/plain')) {
            responseResult.value = await response.text()
        } else if (contentType && contentType.includes('text/html')) {
            responseResult.value = await response.text()
        } else {
            responseResult.value = await response.text() // 默认处理为文本
        }

        ElMessage.success('请求成功')
    } catch (error) {
        ElMessage.error('请求失败：' + error.message)
        responseResult.value = '请求失败：' + error.message
    } finally {
        loading.value = false
    }
}

// JSON 输入框自动补全和缩进
const handleJsonIndent = (event, targetRef) => {
    const textarea = event.target
    const start = textarea.selectionStart
    const end = textarea.selectionEnd
    let value = targetRef // 直接使用字符串

    // 处理输入 { 或 [ 时的自动补全
    if (event.key === '{' || event.key === '[') {
        event.preventDefault() // 阻止默认输入行为

        const pair = event.key === '{' ? '{}' : '[]'
        value = value.substring(0, start) + pair + value.substring(end) // 更新字符串

        // 设置光标位置在括号中间
        const cursorPosition = start + 1
        textarea.setSelectionRange(cursorPosition, cursorPosition);
        textarea.focus(); // 让 textarea 获得焦点
    }

    // 处理回车时的自动缩进
    if (event.key === 'Enter') {
        event.preventDefault() // 阻止默认换行行为

        // 获取当前行的缩进
        const lineStart = value.lastIndexOf('\n', start - 1) + 1
        const currentLine = value.substring(lineStart, start)
        const indent = currentLine.match(/^\s*/)[0] || '' // 获取当前行的缩进

        // 检查当前光标是否在 {} 或 [] 中间
        const beforeCursor = value.substring(0, start)
        const afterCursor = value.substring(end)
        const isInsideBraces = beforeCursor.endsWith('{') && afterCursor.startsWith('}')
        const isInsideBrackets = beforeCursor.endsWith('[') && afterCursor.startsWith(']')

        if (isInsideBraces || isInsideBrackets) {
            // 在 {} 或 [] 中间按下回车
            const newIndent = indent + '    ' // 增加一级缩进
            const newValue =
                value.substring(0, start) + // 光标前的内容
                '\n' + newIndent + // 新行和缩进
                '\n' + indent + // 新行（右侧括号的缩进）
                value.substring(end) // 光标后的内容

            value = newValue // 更新字符串

            // 设置光标位置在新行的缩进位置
            const cursorPosition = start + newIndent.length + 1
            textarea.selectionStart = textarea.selectionEnd = cursorPosition
        } else {
            // 普通回车行为
            const newValue = value.substring(0, start) + '\n' + indent + '    ' + value.substring(end)
            value = newValue // 更新字符串

            // 设置光标位置
            const cursorPosition = start + indent.length + 5
            textarea.selectionStart = textarea.selectionEnd = cursorPosition
        }
    }

    // 返回更新后的值
    return value
}
</script>

<style scoped>
.http-tool-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    width: 100%;
    box-sizing: border-box;
}

.http-tool-card {
    width: 100%;
    height: 100%;
    border-radius: 0.5rem;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    display: flex;
    flex-direction: column;
    overflow: hidden;
    /* 防止内容溢出 */
}

.operation-bar {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
    flex-shrink: 0;
    /* 防止操作栏被压缩 */
}

.section {
    margin-bottom: 20px;
    flex: 1;
    display: flex;
    flex-direction: column;
    min-height: 0;
    /* 防止内容区域被压缩 */
}

.section-title {
    font-size: 14px;
    font-weight: bold;
    color: #333;
    margin-bottom: 10px;
}

.el-textarea {
    font-family: monospace;
    font-size: 14px;
    flex: 1;
    min-height: 0;
    /* 防止输入框被压缩 */
}

.el-textarea :deep(.el-textarea__inner) {
    height: 100% !important;
    resize: none;
    overflow-y: auto;
    /* 支持滚动 */
}

/* 整体滚动区域 */
.http-tool-card :deep(.el-card__body) {
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow-y: auto;
    /* 支持滚动 */
    padding: 20px;
}
</style>