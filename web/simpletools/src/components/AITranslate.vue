<template>
  <div class="translation-container">
    <el-card class="translation-card">
      <!-- 操作栏 -->
      <div class="operation-bar">
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
            :rows="12"
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
            :rows="12"
            resize="none"
            readonly
            placeholder="翻译结果将显示在这里"
            class="text-area result-area"
          />
          <div class="copy-wrapper">
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
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Search, Right, DocumentCopy } from '@element-plus/icons-vue'

// 响应式数据
const translationDirection = ref('toChinese')
const sourceText = ref('')
const translatedText = ref('')
const loading = ref(false)

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
</script>

<style scoped>
.translation-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
}

.translation-card {
  width: 100%;
  max-width: 1200px;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.operation-bar {
  display: flex;
  align-items: center;
  margin-bottom: 24px;
  padding: 8px 0;
}

.content-wrapper {
  display: flex;
  gap: 20px;
}

.text-panel {
  flex: 1;
  position: relative;
}

.text-area {
  width: 100%;
  font-size: 14px;
  line-height: 1.6;
  border-radius: 8px;
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
}

.copy-wrapper {
  margin-top: 15px;
  text-align: right;
}

@media (max-width: 768px) {
  .content-wrapper {
    flex-direction: column;
  }
  
  .divider {
    transform: rotate(90deg);
    margin: 10px 0;
  }
}
</style>