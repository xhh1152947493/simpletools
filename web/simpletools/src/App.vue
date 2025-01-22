<script setup>
import { ref, watch } from 'vue';
import { Message, Shop, Grid, Plus } from '@element-plus/icons-vue';
import ToolMarket from './components/ToolMarket.vue'; // 工具市场组件
import ToolComponent from './components/ToolComponent.vue'; // 工具组件
import LayoutModal from './components/LayoutModal.vue'; // 布局模态框组件

// 控制工具市场模态框的显示
const isToolMarketVisible = ref(false);

// 控制布局模态框的显示
const isLayoutModalVisible = ref(false);

// 存储添加的工具组件
const tools = ref([]);

// 当前布局方式
const currentLayout = ref('default');

// 搜索关键词
const searchQuery = ref('');

// 添加工具组件
const addTool = (tool) => {
  if (tools.value.length < 4) {
    tools.value.push(tool);
  } else {
    alert('操作台最多只能添加 4 个工具');
  }
};

// 监听模态框的显示状态，动态控制页面滚动条
watch(isToolMarketVisible, (newVal) => {
  if (newVal) {
    document.body.style.overflow = 'hidden'; // 隐藏页面滚动条
  } else {
    document.body.style.overflow = 'auto'; // 恢复页面滚动条
  }
});
</script>

<template>
  <div class="container">
    <!-- 虚线框内的加号 -->
    <div class="add-button" @click="isToolMarketVisible = true">
      <el-icon><Plus /></el-icon>
    </div>

    <!-- 提示文字 -->
    <div v-if="tools.length === 0" class="tip-text">
      当前操作面板无工具，请前往工具市场添加工具
    </div>

    <!-- 添加的工具组件 -->
    <div
      v-for="(tool, index) in tools"
      :key="index"
      class="tool-item"
      :style="{ gridArea: `tool${index + 1}` }"
    >
      <ToolComponent :type="tool" />
    </div>
  </div>

  <!-- 按钮组 -->
  <div class="button-group">
    <el-tooltip content="系统消息" placement="left">
      <el-button class="action-button">
        <el-icon><Message /></el-icon>
      </el-button>
    </el-tooltip>

    <el-tooltip content="工具市场" placement="left">
      <el-button class="action-button" @click="isToolMarketVisible = true">
        <el-icon><Shop /></el-icon>
      </el-button>
    </el-tooltip>

    <el-tooltip content="网格布局" placement="left">
      <el-button class="action-button" @click="isLayoutModalVisible = true">
        <el-icon><Grid /></el-icon>
      </el-button>
    </el-tooltip>
  </div>

  <!-- 工具市场模态框 -->
  <el-dialog
    v-model="isToolMarketVisible"
    :width="'80%'"
    :style="{ height: '70vh' }"
    class="fixed-dialog"
  >
    <!-- 自定义标题区域 -->
    <template #header>
      <div class="dialog-header">
        <el-input
          v-model="searchQuery"
          placeholder="按名字或类型搜索工具"
          clearable
          class="search-input"
          suffix-icon="Search"
        />
      </div>
    </template>

    <!-- 工具市场内容 -->
    <ToolMarket :search-query="searchQuery" @add-tool="addTool" />
  </el-dialog>

  <!-- 布局模态框 -->
  <el-dialog
    v-model="isLayoutModalVisible"
    title="选择布局方式"
    :width="'60%'"
    :style="{ height: '70vh' }"
    class="fixed-dialog"
  >
    <LayoutModal :current-layout="currentLayout" @change-layout="changeLayout" />
  </el-dialog>
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
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.add-button {
  cursor: pointer;
  color: #666;
  font-size: 2rem;
  margin-bottom: 0.5rem; /* 与提示文字的间距 */
}

.add-button:hover {
  color: #333;
}

.tip-text {
  color: #999;
  font-size: 0.875rem;
  text-align: center;
}

.tool-item {
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
}

.button-group {
  position: fixed;
  right: 1.5rem;
  bottom: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
  z-index: 1000;
  align-items: flex-end;
}

.action-button {
  border: none;
  color: #666;
  padding: 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0;
}

.action-button .el-icon {
  font-size: 1.5rem;
  width: 1.5rem;
  height: 1.5rem;
}

.action-button:hover {
  color: #333;
}

/* 自定义模态框标题区域 */
.dialog-header {
  display: flex;
  align-items: center;
  justify-content: center; /* 将内容居中 */
  width: 100%;
  position: relative; /* 相对定位 */
}


.search-input {
  width: 300px; /* 搜索框宽度 */
  margin-left: 20px; /* 与标题的间距 */
}
</style>

<style>
/* 全局样式，用于固定模态框大小和内容滚动 */
.fixed-dialog .el-dialog__body {
  height: calc(100% - 100px); /* 减去标题和底部的高度 */
  overflow-y: auto; /* 内容区域滚动 */
  padding: 16px;
}

.fixed-dialog .el-dialog {
  max-height: 70vh; /* 固定高度 */
  display: flex;
  flex-direction: column;
}
</style>