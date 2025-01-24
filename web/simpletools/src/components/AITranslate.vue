<template>
  <div class="translation-container" ref="translationContainer">
    <el-card class="translation-card">
      <!-- 操作栏 -->
      <div class="operation-bar" ref="operationBar">
        <el-select v-model="translationDirection" placeholder="选择翻译方向" style="width: 200px; margin-right: 15px;">
          <el-option label="文本译中" value="toChinese" />
          <el-option label="文本译英" value="toEnglish" />
        </el-select>

        <el-button type="primary" :loading="loading" @click="handleTranslate" title="使用AI能力进行智能翻译">
          <template #icon>
            <el-icon>
              <Search />
            </el-icon>
          </template>
          立即翻译
        </el-button>
      </div>

      <!-- 内容区域 -->
      <div class="content-wrapper">
        <!-- 原文 -->
        <div class="text-panel">
          <el-input v-model="sourceText" type="textarea" resize="none" :placeholder="inputPlaceholder" maxlength="10000"
            show-word-limit class="text-area" />
        </div>

        <!-- 分隔线 -->
        <div class="divider">
          <el-icon>
            <Right />
          </el-icon>
        </div>

        <!-- 译文 -->
        <div class="text-panel">
          <el-input v-model="translatedText" type="textarea" resize="none" readonly placeholder="翻译结果将显示在这里"
            class="text-area result-area" />
          <div class="copy-wrapper" ref="copyWrapper">
            <el-button type="success" :icon="DocumentCopy" @click="copyResult" :disabled="!translatedText">
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
import { httpPost } from "@/utils/http.js";

// 响应式数据
const translationDirection = ref('toChinese')
const sourceText = ref('')
const translatedText = ref('')
const loading = ref(false)

// 输入框提示语计算属性
const inputPlaceholder = computed(() => {
  return translationDirection.value === 'toChinese'
    ? '支持各种语言，网络用语、专业单词、缩写、混合表达等'
    : '支持各种语言，网络用语、专业单词、缩写、混合表达等'
})

// 处理翻译操作
const handleTranslate = async () => {
  const content = sourceText.value.trim()
  if (!content) {
    ElMessage.warning('请输入要翻译的内容')
    return
  }

  if (content.length > 10000) {
    ElMessage.error('翻译内容不能超过10000个字符')
    return
  }

  try {
    loading.value = true
    const direction = translationDirection.value; // 这里获取下拉框的值
    const result = await requestTranslateResult(content, direction)
    if (result.code !== 0) {
      ElMessage.error('服务暂不可用')
      return
    }
    translatedText.value = result.data
  } catch (error) {
    ElMessage.error('翻译失败，请稍后重试')
  } finally {
    loading.value = false
  }
}


const requestTranslateResult = (content, direction) => {
  return httpPost( import.meta.env.VITE_API_URL + "/api/aitranslate", {"content": content, "language":  direction==='toChinese' ? 1 : 2})
}

// 复制结果
const copyResult = async () => {
  try {
    await navigator.clipboard.writeText(translatedText.value)
    ElMessage.success('翻译结果已复制到剪贴板')
  } catch (err) {
    ElMessage.error('复制失败')
  }
}
</script>

<style scoped>
.translation-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  /* 高度自适应父元素 */
  width: 100%;
  /* 宽度自适应父元素 */
  box-sizing: border-box;
}

.translation-container :deep(.el-card__body) {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 20px;
}

.translation-container :deep(.el-textarea__inner) {
  height: 100% !important;
  resize: none;
}

.translation-card {
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
  flex-shrink: 0;
  /* 防止操作栏被压缩 */
}

.content-wrapper {
  display: flex;
  gap: 20px;
  flex: 1;
  /* 内容区域自适应剩余空间 */
  overflow: hidden;
  /* 防止内容溢出 */
}

.text-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100%;
  min-height: 0;
  position: relative;
  /* 添加相对定位 */
}

.text-area {
  flex: 1;
  width: 100%;
  font-size: 14px;
  line-height: 1.6;
  border-radius: 0.5rem;
  box-sizing: border-box;
  margin-bottom: 50px;
  /* 为底部按钮留出空间 */
}

.copy-wrapper {
  position: absolute;
  /* 使用绝对定位 */
  bottom: 0;
  /* 固定在底部 */
  right: 0;
  /* 靠右对齐 */
  width: 100%;
  /* 占满宽度 */
  padding: 1px 0;
  /* 上下padding */
  text-align: right;
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
  flex-shrink: 0;
  /* 防止分隔线被压缩 */
}
</style>