<script setup>
import { ref, watch, onMounted } from 'vue';
import { Message, Shop, Grid, Plus, ChatLineRound } from '@element-plus/icons-vue';
import ToolMarket from './components/ToolMarket.vue'; // 工具市场组件
import LayoutModal from './components/LayoutModal.vue'; // 布局模态框组件

// 控制工具市场模态框的显示
const isToolMarketVisible = ref(false);

// 控制布局模态框的显示
const isLayoutModalVisible = ref(false);

// 存储添加的工具组件
const tools = ref([]);

// 当前布局方式
const currentLayout = ref('lyt1');

// 搜索关键词
const searchQuery = ref('');

// 从 localStorage 加载保存的工具和布局
const loadFromLocalStorage = () => {
  const savedTools = localStorage.getItem('tools');
  const savedLayout = localStorage.getItem('currentLayout');

  if (savedTools) {
    tools.value = JSON.parse(savedTools);
  }
  if (savedLayout) {
    currentLayout.value = savedLayout;
  }
};

// 保存工具和布局到 localStorage
const saveToLocalStorage = () => {
  localStorage.setItem('tools', JSON.stringify(tools.value));
  localStorage.setItem('currentLayout', currentLayout.value);
};

// 页面加载时从 localStorage 加载数据
onMounted(() => {
  loadFromLocalStorage();
});

// 添加工具组件
const addTool = (tool) => {
  if (tools.value.length < 4) {
    tools.value.push(tool);
    saveToLocalStorage(); // 保存到 localStorage
  } else {
    alert('操作台最多只能添加 4 个工具');
  }
};

// 根据布局方式和工具数量计算每个工具的样式
const getToolStyle = (index) => {
  const toolCount = tools.value.length;
  const layout = currentLayout.value;

  if (toolCount === 1) {
    return {
      gridArea: '1 / 1 / 2 / 2',
    };
  } else if (toolCount === 2) {
    if (layout === 'lyt1' || layout === 'lyt2') {
      // 上下分布
      return {
        gridArea: index === 0 ? '1 / 1 / 2 / 2' : '2 / 1 / 3 / 2',
      };
    } else if (layout === 'lyt5' || layout === 'lyt6') {
      // 左右分布
      return {
        gridArea: index === 0 ? '1 / 1 / 2 / 2' : '1 / 2 / 2 / 3',
      };
    }
  } else if (toolCount === 3) {
    if (layout === 'lyt1' || layout === 'lyt2') {
      // 上下分布，上方两格，下方一格
      return {
        gridArea:
          index === 0
            ? '1 / 1 / 2 / 2'
            : index === 1
            ? '1 / 2 / 2 / 3'
            : '2 / 1 / 3 / 3',
      };
    } else if (layout === 'lyt5' || layout === 'lyt6') {
      // 左右分布，左侧两格，右侧一格
      return {
        gridArea:
          index === 0
            ? '1 / 1 / 2 / 2'
            : index === 1
            ? '2 / 1 / 3 / 2'
            : '1 / 2 / 3 / 3',
      };
    }
  } else if (toolCount === 4) {
    // 田字布局
    return {
      gridArea:
        index === 0
          ? '1 / 1 / 2 / 2'
          : index === 1
          ? '1 / 2 / 2 / 3'
          : index === 2
          ? '2 / 1 / 3 / 2'
          : '2 / 2 / 3 / 3',
    };
  }
};

// 切换布局
const changeLayout = (layoutId) => {
  currentLayout.value = layoutId;
  saveToLocalStorage(); // 保存到 localStorage
};

// 拖拽结束事件
const onDragEnd = () => {
  saveToLocalStorage(); // 保存到 localStorage
};

// 监听 tools 和 currentLayout 的变化，自动保存到 localStorage
watch(tools, () => {
  saveToLocalStorage();
}, { deep: true });

watch(currentLayout, () => {
  saveToLocalStorage();
});

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
    <!-- 加号和提示文本 -->
    <div v-if="tools.length === 0" class="center-content">
      <div class="add-button" @click="isToolMarketVisible = true">
        <el-icon><Plus /></el-icon>
      </div>
      <div class="tip-text">
        当前操作面板无工具，请前往工具市场添加工具
      </div>
    </div>

    <!-- 添加的工具组件 -->
    <div v-else class="tool-container" :class="currentLayout">
      <div
        v-for="(tool, index) in tools"
        :key="index"
        class="tool-item"
        :style="getToolStyle(index)"
      >
        {{ tool.name }}
      </div>
    </div>
  </div>

  <!-- 按钮组 -->
  <div class="button-group">
    <el-tooltip content="系统消息" placement="left">
      <el-button class="action-button">
        <el-icon><Message /></el-icon>
      </el-button>
    </el-tooltip>

    <el-tooltip content="联系作者1152947493@qq.com" placement="left">
      <el-button class="action-button">
        <el-icon><ChatLineRound /></el-icon>
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
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
}

.center-content {
  display: flex;
  flex-direction: column;
  align-items: center; /* 水平居中 */
  gap: 0.5rem; /* 加号和文本之间的间距 */
}

.add-button {
  cursor: pointer;
  color: #666;
  font-size: 2rem;
}

.add-button:hover {
  color: #333;
}

.tip-text {
  color: #999;
  font-size: 0.875rem;
  text-align: center;
}

.tool-container {
  display: grid;
  width: 100%;
  height: 100%;
  gap: 0.5rem;
  padding: 0.5rem;
}

.tool-item {
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
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

/* 强制覆盖深色模式的样式 */
body {
  background-color: white !important;
  color: black !important;
}
</style>


