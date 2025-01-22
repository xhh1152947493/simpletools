
<script setup>
import { Message, Shop, Grid } from '@element-plus/icons-vue'
</script>

<template>
  <div class="container">
  </div>
  <div class="button-group">
    <el-button class="action-button">
      <el-icon><Message /></el-icon>
    </el-button>
    <el-button class="action-button">
      <el-icon><Shop /></el-icon>
    </el-button>
    <el-button class="action-button">
      <el-icon><Grid /></el-icon>
    </el-button>
  </div>
</template>

<style>
html, body, #app {
  margin: 0;
  padding: 0;
}
</style>

<style scoped>
.container {
  border: 0.125rem dashed #ccc;
  border-radius: 0.5rem;
  box-sizing: border-box;
  height: calc(100vh - 2rem);
  width: calc(100vw - 2rem);
  margin: 1rem;
  overflow: hidden;
  position: relative;
}

.button-group {
  position: fixed;
  right: 1.5rem;
  bottom: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
  z-index: 1000;
  align-items: flex-end; /* 确保按钮在交叉轴上对齐 */
}

.action-button {
  border: none;
  color: #666;
  padding: 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0; /* 确保没有额外的 margin */
}

.action-button .el-icon {
  font-size: 1.5rem;
  width: 1.5rem; /* 确保图标宽度一致 */
  height: 1.5rem; /* 确保图标高度一致 */
}

.action-button:hover {
  color: #333;
}
</style>