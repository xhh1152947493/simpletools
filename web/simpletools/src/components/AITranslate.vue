<template>
  <div class="translation-container" ref="translationContainer">
    <el-card class="translation-card">
      <!-- 操作栏 -->
      <div class="operation-bar" ref="operationBar">
        <el-select v-model="translationDirection" placeholder="选择翻译方向" style="width: 200px; margin-right: 15px;">
          <el-option label="外译中" value="toChinese" />
          <el-option label="中译英" value="toEnglish" />
        </el-select>

        <el-button type="primary" :loading="loading" @click="handleTranslate">
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
import OpenAI from "openai";

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
    const response = await requestTranslateAPI(content)
    translatedText.value = response
  } catch (error) {
    ElMessage.error('翻译失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

const openai = new OpenAI({
  baseURL: 'https://api.deepseek.com',
  apiKey: 'sk-efeb2ec0bb21405dbbf43338d9997a02',
  dangerouslyAllowBrowser: true // 允许在浏览器环境中运行
});


const requestTranslateAPI = (content) => {
  return new Promise((resolve, reject) => {
    openai.chat.completions.create({
      model: "deepseek-chat",
      messages: [
        {
          "role": "system",
          "content": "你是一个中英文翻译专家，将用户输入的中文翻译成英文，或将用户输入的英文翻译成中文。对于非中文内容，它将提供中文翻译结果。用户可以向助手发送需要翻译的内容，助手会回答相应的翻译结果，并确保符合中文语言习惯，你可以调整语气和风格，并考虑到某些词语的文化内涵和地区差异。同时作为翻译家，需将原文翻译成具有信达雅标准的译文。\"信\" 即忠实于原文的内容与意图；\"达\" 意味着译文应通顺易懂，表达清晰；\"雅\" 则追求译文的文化审美和语言的优美。目标是创作出既忠于原作精神，又符合目标语言文化和读者审美的翻译。"
        },
        {
          "role": "user",
          "content": content,
        }
      ],
    })
      .then(completion => {
        resolve(completion.choices[0].message.content);
      })
      .catch(error => {
        reject(error);
      });
  });
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