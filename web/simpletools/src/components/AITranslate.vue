<template>
  <div class="translation-container" ref="translationContainer">
    <el-card class="translation-card">
      <!-- 操作栏 -->
      <div class="operation-bar" ref="operationBar">
        <el-select 
          v-model="translationDirection" 
          placeholder="选择翻译方向"
          style="width: 200px; margin-right: 15px;">
          <el-option label="外译中" value="toChinese" />
          <el-option label="中译英" value="toEnglish" />
        </el-select>
        
        <el-button 
          type="primary" 
          :loading="loading"
          @click="handleTranslate">
          <template #icon>
            <el-icon><Search /></el-icon>
          </template>
          立即翻译
        </el-button>
      </div>

      <!-- 内容区域 -->
      <div class="content-wrapper">
        <!-- 原文 -->
        <div class="text-panel">
          <el-input
            v-model="sourceText"
            type="textarea"
            :rows="dynamicRows"
            resize="none"
            :placeholder="inputPlaceholder"
            maxlength="10000"
            show-word-limit
            class="text-area"
          />
        </div>

        <!-- 分隔线 -->
        <div class="divider">
          <el-icon><Right /></el-icon>
        </div>

        <!-- 译文 -->
        <div class="text-panel">
          <el-input
            v-model="translatedText"
            type="textarea"
            :rows="dynamicRows"
            resize="none"
            readonly
            placeholder="翻译结果将显示在这里"
            class="text-area result-area"
          />
          <div class="copy-wrapper" ref="copyWrapper">
            <el-button 
              type="success" 
              :icon="DocumentCopy" 
              @click="copyResult"
              :disabled="!translatedText">
              复制结果
            </el-button>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Search, Right, DocumentCopy } from '@element-plus/icons-vue'

// 响应式数据
const translationDirection = ref('toChinese')
const sourceText = ref('')
const translatedText = ref('')
const loading = ref(false)
const dynamicRows = ref(12) // 动态计算的行数

// 使用 ref 获取 DOM 元素
const translationContainer = ref(null)
const operationBar = ref(null)
const copyWrapper = ref(null)

// 输入框提示语计算属性
const inputPlaceholder = computed(() => {
  return translationDirection.value === 'toChinese' 
    ? '请输入要翻译的外语内容' 
    : '请输入要翻译的中文内容'
})

// 处理翻译操作
const handleTranslate = async () => {
  if (!sourceText.value.trim()) {
    ElMessage.warning('请输入要翻译的内容')
    return
  }

  if (sourceText.value.length > 10000) {
    ElMessage.error('翻译内容不能超过10000个字符')
    return
  }

  try {
    loading.value = true
    // 模拟API调用，实际应替换为真实API请求
    const response = await mockTranslateAPI()
    translatedText.value = response.data.result
  } catch (error) {
    ElMessage.error('翻译失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

// 模拟翻译API
const mockTranslateAPI = () => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({
        data: {
          result: translationDirection.value === 'toChinese' 
            ? '这是模拟的中文翻译结果' 
            : 'This is a mock English translation result'
        }
      })
    }, 800)
  })
}

// 复制结果
const copyResult = async () => {
  try {
    await navigator.clipboard.writeText(translatedText.value)
    ElMessage.success('翻译结果已复制到剪贴板')
  } catch (err) {
    ElMessage.error('复制失败，请手动选择复制')
  }
}

// 动态计算行数
const calculateRows = () => {
  if (translationContainer.value && operationBar.value && copyWrapper.value) {
    const containerHeight = translationContainer.value.clientHeight
    const operationBarHeight = operationBar.value.clientHeight
    const copyWrapperHeight = copyWrapper.value.clientHeight
    const padding = 40 // 上下内边距
    const rowHeight = 24 // 每行的高度
    const availableHeight = containerHeight - operationBarHeight - copyWrapperHeight - padding
    dynamicRows.value = Math.floor(availableHeight / rowHeight)
  }
}

// 监听窗口大小变化
onMounted(() => {
  // 打印父元素信息
  if (translationContainer.value && translationContainer.value.parentElement) {
    const parentElement = translationContainer.value.parentElement
    console.log('父元素标签名:', parentElement.tagName)
    console.log('父元素类名:', parentElement.className)
    console.log('父元素宽度:', parentElement.clientWidth, 'px')
    console.log('父元素高度:', parentElement.clientHeight, 'px')
  } else {
    console.log('未找到父元素')
  }

  calculateRows()
  window.addEventListener('resize', calculateRows)
})

onUnmounted(() => {
  window.removeEventListener('resize', calculateRows)
})
</script>

<style scoped>
.translation-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%; /* 高度自适应父元素 */
  width: 100%; /* 宽度自适应父元素 */
  padding: 20px;
  box-sizing: border-box;
}

.translation-card {
  width: 100%;
  height: 100%;
  border-radius: 12px;
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
  min-height: 0; /* 防止内容溢出 */
}

.text-area {
  width: 100%;
  flex: 1; /* 输入框自适应父容器高度 */
  font-size: 14px;
  line-height: 1.6;
  border-radius: 8px;
  box-sizing: border-box;
  min-height: 0; /* 防止内容溢出 */
}

.result-area {
  :deep(.el-textarea__inner) {
    background-color: #f8f9fa;
  }
}

.divider {
  display: flex;
  align-items: center;
  color: #409eff;
  font-size: 24px;
  padding: 0 10px;
  flex-shrink: 0; /* 防止分隔线被压缩 */
}

.copy-wrapper {
  margin-top: 15px;
  text-align: right;
  flex-shrink: 0; /* 防止复制按钮被压缩 */
}
</style>