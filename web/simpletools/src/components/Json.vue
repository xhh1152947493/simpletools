<template>
    <div class="json-tool-container" ref="jsonToolContainer">
      <el-card class="json-tool-card">
        <!-- 操作栏 -->
        <div class="operation-bar" ref="operationBar">
          <el-button 
            type="primary" 
            :loading="loading"
            @click="handleFormat">
            <template #icon>
              <el-icon><MagicStick /></el-icon>
            </template>
            格式化
          </el-button>
  
          <el-button 
            type="info" 
            @click="toggleDisplayMode">
            <template #icon>
              <el-icon><Switch /></el-icon>
            </template>
            {{ displayMode === 'pretty' ? '紧凑显示' : '友好显示' }}
          </el-button>
        </div>
  
        <!-- 内容区域 -->
        <div class="content-wrapper">
          <div class="text-panel">
            <el-input
              v-model="jsonInput"
              type="textarea"
              resize="none"
              placeholder="请输入 JSON 字符串"
              show-word-limit
              class="text-area"
              :class="{ 'error': isError }"
              ref="jsonInputRef"
            />
            <div v-if="isError" class="error-message">
              解析错误：{{ errorMessage }}，位置：第 {{ errorLine }} 行，第 {{ errorColumn }} 列
            </div>
          </div>
        </div>
      </el-card>
    </div>
  </template>
  
  <script setup>
  import { ref, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import { MagicStick, Switch } from '@element-plus/icons-vue'
  
  // 响应式数据
  const jsonInput = ref('')
  const displayMode = ref('pretty') // pretty: 友好显示, compact: 紧凑显示
  const isError = ref(false)
  const errorMessage = ref('')
  const errorLine = ref(0)
  const errorColumn = ref(0)
  const jsonInputRef = ref(null) // 用于定位输入框
  
  // 切换显示模式
  const toggleDisplayMode = () => {
    displayMode.value = displayMode.value === 'pretty' ? 'compact' : 'pretty'
    handleFormat() // 切换后重新格式化
  }
  
  // 格式化 JSON
  const handleFormat = () => {
    if (!jsonInput.value.trim()) {
      ElMessage.warning('请输入 JSON 字符串')
      return
    }
  
    try {
      const jsonObj = JSON.parse(jsonInput.value)
      if (displayMode.value === 'pretty') {
        jsonInput.value = JSON.stringify(jsonObj, null, 4) // 友好显示，缩进 4 空格
      } else {
        jsonInput.value = JSON.stringify(jsonObj) // 紧凑显示
      }
      isError.value = false
      ElMessage.success('JSON 格式化成功')
    } catch (error) {
      isError.value = true
      errorMessage.value = error.message
      // 解析错误位置
      const match = error.message.match(/at position (\d+)/)
      if (match) {
        const position = parseInt(match[1], 10)
        const lines = jsonInput.value.substring(0, position).split('\n')
        errorLine.value = lines.length
        errorColumn.value = position - lines.slice(0, -1).join('\n').length
      }
      ElMessage.error('JSON 格式化失败，请输入合法的 JSON 字符串')
      focusErrorPosition()
    }
  }
  
  // 定位错误位置
  const focusErrorPosition = () => {
    if (jsonInputRef.value && jsonInputRef.value.textarea) {
      const textarea = jsonInputRef.value.textarea
      textarea.focus()
      const lines = jsonInput.value.split('\n')
      let position = 0
      for (let i = 0; i < errorLine.value - 1; i++) {
        position += lines[i].length + 1 // +1 是换行符
      }
      position += errorColumn.value - 1
      textarea.setSelectionRange(position, position)
    }
  }
  </script>
  
  <style scoped>
  .json-tool-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%; /* 高度自适应父元素 */
    width: 100%; /* 宽度自适应父元素 */
    box-sizing: border-box;
  }
  
  .json-tool-container :deep(.el-card__body) {
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 20px;
  }
  
  .json-tool-container :deep(.el-textarea__inner) {
    height: 100% !important;
    resize: none;
  }
  
  .json-tool-card {
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
    gap: 10px; /* 按钮间距 */
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
  
  .error {
    border-color: #f56c6c;
  }
  
  .error-message {
    color: #f56c6c;
    font-size: 12px;
    margin-top: 8px;
  }
  </style>