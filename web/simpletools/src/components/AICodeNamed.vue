<template>
    <div class="naming-container" ref="namingContainer">
      <el-card class="naming-card">
        <!-- 操作栏 -->
        <div class="operation-bar" ref="operationBar">
          <el-select 
            v-model="namingStyle" 
            placeholder="选择命名风格"
            style="width: 200px; margin-right: 15px;">
            <el-option label="大驼峰" value="PascalCase" />
            <el-option label="小驼峰" value="camelCase" />
            <el-option label="下划线" value="snake_case" />
            <el-option label="全大写下划线常量" value="UPPER_CASE" />
          </el-select>
          
          <el-button 
            type="primary" 
            :loading="loading"
            @click="handleNaming">
            <template #icon>
              <el-icon><Search /></el-icon>
            </template>
            立即命名
          </el-button>
        </div>
  
        <!-- 内容区域 -->
        <div class="content-wrapper">
          <!-- 输入框 -->
          <div class="text-panel">
            <el-input
              v-model="descriptionText"
              type="textarea"
              resize="none"
              placeholder="请输入变量或函数的描述"
              maxlength="2000"
              show-word-limit
              class="text-area"
            />
          </div>
  
          <!-- 分隔线 -->
          <div class="divider">
            <el-icon><Right /></el-icon>
          </div>
  
          <!-- 命名结果 -->
          <div class="result-panel">
            <div 
              v-for="(result, index) in namingResults" 
              :key="index" 
              class="result-item"
              @click="copyResult(result)">
              {{ result }}
            </div>
          </div>
        </div>
  
        <!-- 换一批按钮 -->
        <div class="refresh-button">
          <el-button 
            type="info" 
            @click="refreshResults">
            换一批
          </el-button>
        </div>
      </el-card>
    </div>
  </template>
  
  <script setup>
  import { ref, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import { Search, Right } from '@element-plus/icons-vue'
  
  // 响应式数据
  const namingStyle = ref('PascalCase')
  const descriptionText = ref('')
  const namingResults = ref([])
  const loading = ref(false)
  
  // 处理命名操作
  const handleNaming = async () => {
    if (!descriptionText.value.trim()) {
      ElMessage.warning('请输入变量或函数的描述')
      return
    }
  
    if (descriptionText.value.length > 2000) {
      ElMessage.error('描述内容不能超过2000个字符')
      return
    }
  
    try {
      loading.value = true
      // 模拟API调用，实际应替换为真实API请求
      const response = await mockNamingAPI()
      namingResults.value = response.data.results
    } catch (error) {
      ElMessage.error('命名失败，请稍后重试')
    } finally {
      loading.value = false
    }
  }
  
  // 模拟命名API
  const mockNamingAPI = () => {
    return new Promise((resolve) => {
      setTimeout(() => {
        const results = [
          'ExampleVariableName',
          'exampleVariableName',
          'example_variable_name',
          'EXAMPLE_VARIABLE_NAME',
          'AnotherExampleName'
        ]
        resolve({
          data: {
            results: results.map(result => convertNamingStyle(result, namingStyle.value))
          }
        })
      }, 800)
    })
  }
  
  // 转换命名风格
  const convertNamingStyle = (text, style) => {
    switch (style) {
      case 'PascalCase':
        return text.replace(/(?:^\w|[A-Z]|\b\w)/g, (word, index) => index === 0 ? word.toUpperCase() : word.toLowerCase()).replace(/\s+/g, '')
      case 'camelCase':
        return text.replace(/(?:^\w|[A-Z]|\b\w)/g, (word, index) => index === 0 ? word.toLowerCase() : word.toUpperCase()).replace(/\s+/g, '')
      case 'snake_case':
        return text.replace(/\s+/g, '_').toLowerCase()
      case 'UPPER_CASE':
        return text.replace(/\s+/g, '_').toUpperCase()
      default:
        return text
    }
  }
  
  // 复制结果
  const copyResult = async (text) => {
    try {
      await navigator.clipboard.writeText(text)
      ElMessage.success('命名结果已复制到剪贴板')
    } catch (err) {
      ElMessage.error('复制失败，请手动选择复制')
    }
  }
  
  // 换一批结果
  const refreshResults = async () => {
    try {
      loading.value = true
      const response = await mockNamingAPI()
      namingResults.value = response.data.results
    } catch (error) {
      ElMessage.error('获取新一批命名失败，请稍后重试')
    } finally {
      loading.value = false
    }
  }
  </script>
  
  <style scoped>
  .naming-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%; /* 高度自适应父元素 */
    width: 100%; /* 宽度自适应父元素 */
    box-sizing: border-box;
  }
  
  .naming-container :deep(.el-card__body) {
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 20px;
  }
  
  .naming-container :deep(.el-textarea__inner) {
    height: 100% !important;
    resize: none;
  }
  
  .naming-card {
    width: 100%;
    height: 100%;
    border-radius: 0.5rem;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    display: flex;
    flex-direction: column;
  }
  
  .operation-bar {
    display: flex;
    align-items: center;
    margin-bottom: 24px;
    padding: 8px 0;
    flex-shrink: 0; /* 防止操作栏被压缩 */
  }
  
  .content-wrapper {
    display: flex;
    gap: 20px;
    flex: 1; /* 内容区域自适应剩余空间 */
    overflow: hidden; /* 防止内容溢出 */
  }
  
  .text-panel {
    flex: 1;
    display: flex;
    flex-direction: column;
    height: 100%;
    min-height: 0;
  }
  
  .text-area {
    flex: 1;
    width: 100%;
    font-size: 14px;
    line-height: 1.6;
    border-radius: 0.5rem;
    box-sizing: border-box;
  }
  
  .result-panel {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 10px;
    overflow-y: auto;
  }
  
  .result-item {
    padding: 10px;
    border: 1px solid #ebeef5;
    border-radius: 0.5rem;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .result-item:hover {
    background-color: #f5f7fa;
  }
  
  .divider {
    display: flex;
    align-items: center;
    color: #409eff;
    font-size: 24px;
    padding: 0 10px;
    flex-shrink: 0; /* 防止分隔线被压缩 */
  }
  
  .refresh-button {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }
  </style>