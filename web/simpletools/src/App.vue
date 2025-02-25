<script setup>
import { ref, watch, onMounted, defineAsyncComponent, nextTick } from 'vue';
import { Message, Shop, Grid, Plus, ChatLineRound, Close } from '@element-plus/icons-vue';
import ToolMarket from './components/ToolMarket.vue'; // 工具市场组件
import LayoutModal from './components/LayoutModal.vue'; // 布局模态框组件
import { ElMessage } from 'element-plus';

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
    const parsedTools = JSON.parse(savedTools);
    tools.value = parsedTools.map(tool => {
      return {
        ...tool,
        component: defineAsyncComponent({
          loader: () => import(`./components/${tool.componentName}.vue`),
          loadingComponent: () => '加载中...', // 加载中的占位组件
          errorComponent: () => '加载失败', // 加载失败的占位组件
          delay: 200, // 延迟显示加载中组件
          timeout: 3000, // 加载超时时间
        }),
      };
    });
  }
  if (savedLayout) {
    currentLayout.value = savedLayout;
  }
};

// 保存工具和布局到 localStorage
const saveToLocalStorage = () => {
  const toolsToSave = tools.value.map(tool => ({
    ...tool,
    component: undefined, // 不存储组件对象
    componentName: tool.componentName, // 存储组件名称
  }));
  localStorage.setItem('tools', JSON.stringify(toolsToSave));
  localStorage.setItem('currentLayout', currentLayout.value);
};

// 页面加载时从 localStorage 加载数据
onMounted(() => {
  nextTick(() => {
    loadFromLocalStorage();
  });
});

// 添加工具组件
const addTool = (tool) => {
  if (tools.value.length < 4) {
    tools.value.push({
      ...tool,
      id: `tool-${Date.now()}`, // 为每个工具生成唯一 ID
      componentName: tool.componentName, // 存储组件名称
    });
    ElMessage.success(`添加工具成功: ${tool.name}`);
  } else {
    ElMessage.error('操作面板最多只能添加4个工具');
  }
};

// 移除工具
const removeTool = (index) => {
  tools.value.splice(index, 1);
  saveToLocalStorage(); // 保存到 localStorage
};

// 根据布局方式和工具数量计算每个工具的样式
const getToolStyle = (index) => {
  const toolCount = tools.value.length;
  const layout = currentLayout.value;

  if (toolCount === 1) {
    return {
      gridArea: '1 / 1 / 3 / 3',
    };
  } else if (toolCount === 2) {
    if (layout === 'lyt1' || layout === 'lyt2' || layout === 'lyt3' || layout === 'lyt4') {
      // 上下分布
      return {
        gridArea: index === 0 ? '1 / 1 / 2 / 3' : '2 / 1 / 3 / 3',
      };
    } else if (layout === 'lyt5' || layout === 'lyt6' || layout === 'lyt7' || layout === 'lyt8') {
      // 左右分布
      return {
        gridArea: index === 0 ? '1 / 1 / 3 / 2' : '1 / 2 / 3 / 3',
      };
    }
  } else if (toolCount === 3) {
    if (layout === 'lyt1' || layout === 'lyt5') {
      // 上下分布，上方两格，下方一格
      return {
        gridArea:
          index === 0
            ? '1 / 1 / 2 / 2'
            : index === 1
              ? '1 / 2 / 2 / 3'
              : '2 / 1 / 3 / 3',
      };
    } else if (layout === 'lyt2' || layout === 'lyt6') {
      // 上下分布 下方两格，上方一格
      return {
        gridArea:
          index === 0
            ? '1 / 1 / 2 / 3'
            : index === 1
              ? '2 / 1 / 3 / 2'
              : '2 / 2 / 3 / 3',
      };
    } else if (layout === 'lyt3' || layout === 'lyt7') {
      // 左右分布，左侧两格，右侧一格
      return {
        gridArea:
          index === 0
            ? '1 / 1 / 2 / 2'
            : index === 1
              ? '2 / 1 / 3 / 2'
              : '1 / 2 / 3 / 3',
      };
    } else if (layout === 'lyt4' || layout === 'lyt8') {
      // 左右分布，右侧两格，左侧一格
      return {
        gridArea:
          index === 0
            ? '1 / 1 / 3 / 1'
            : index === 1
              ? '1 / 2 / 2 / 3'
              : '2 / 2 / 3 / 3',
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

// 拖拽逻辑
let dragIndex = null;

// 新增：记录当前拖拽的目标位置
const dragOverIndex = ref(null);

const onDragStart = (index) => {
  dragIndex = index; // 记录被拖拽的工具索引
};

// 拖拽悬停
const onDragOver = (event, index) => {
  event.preventDefault(); // 允许放置
  dragOverIndex.value = index; // 记录当前悬停的工具索引
};

// 拖拽结束
const onDragEnd = () => {
  dragOverIndex.value = null; // 重置拖拽目标位置
};

// 放置
const onDrop = (index) => {
  if (dragIndex === null || dragIndex === index) return; // 如果拖拽的是自己，直接返回

  // 交换工具位置
  const newTools = [...tools.value];
  const draggedTool = newTools[dragIndex];
  newTools[dragIndex] = newTools[index];
  newTools[index] = draggedTool;

  // 更新工具数组
  tools.value = newTools;
  saveToLocalStorage(); // 保存到 localStorage

  dragIndex = null; // 重置拖拽索引
  dragOverIndex.value = null; // 重置拖拽目标位置
};

// 监听 tools 和 currentLayout 的变化，自动保存到 localStorage
watch(tools, () => {
  saveToLocalStorage();
}, { deep: true });

watch(currentLayout, () => {
  // 重新计算工具的布局
  tools.value = [...tools.value]; // 强制触发重新渲染
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

// 新增：控制按钮组的展开状态
const isButtonsExpanded = ref(false);

const tooltipOpenDelay = 300; // 延迟时间需要大于展开动画时间

</script>

<template>
  <div class="container">
    <!-- 加号和提示文本 -->
    <div v-if="tools.length === 0" class="center-content">
      <div class="add-button" @click="isToolMarketVisible = true">
        <el-icon>
          <Plus />
        </el-icon>
      </div>
      <div class="tip-text">
        当前操作面板无工具，请前往工具市场添加工具
      </div>
    </div>

    <!-- 添加的工具组件 -->
    <div v-else class="tool-container" :class="currentLayout">
      <div v-for="(tool, index) in tools" :key="index" class="tool-item" :style="getToolStyle(index)"
        :class="{ 'drag-over': dragOverIndex === index }" draggable="true" @dragstart="onDragStart(index)"
        @dragover="onDragOver($event, index)" @drop="onDrop(index)" @dragend="onDragEnd">
        <!-- 关闭按钮 -->
        <div class="close-button" @click="removeTool(index)">
          <el-icon>
            <Close />
          </el-icon>
        </div>
        <!-- 动态渲染组件 -->
        <component :is="tool.component" />
      </div>
    </div>
  </div>

  <!-- 修改后的按钮组 -->
  <div class="new-button-group">
    <!-- 主按钮 -->
    <div class="main-button" @click="isButtonsExpanded = !isButtonsExpanded">
      <el-icon :class="{ 'rotate-icon': isButtonsExpanded }">
        <Plus />
      </el-icon>
    </div>

    <!-- 子按钮组 -->
    <div class="sub-buttons" :class="{ expanded: isButtonsExpanded }">
      <el-tooltip :show-after="tooltipOpenDelay" :hide-after="0" content="系统消息" placement="left">
        <el-button class="action-button">
          <el-icon>
            <Message />
          </el-icon>
        </el-button>
      </el-tooltip>

      <el-tooltip :show-after="tooltipOpenDelay" :hide-after="0" content="联系作者1152947493@qq.com" placement="left">
        <el-button class="action-button">
          <el-icon>
            <ChatLineRound />
          </el-icon>
        </el-button>
      </el-tooltip>

      <el-tooltip :show-after="tooltipOpenDelay" :hide-after="0" content="工具市场" placement="left">
        <el-button class="action-button" @click="isToolMarketVisible = true">
          <el-icon>
            <Shop />
          </el-icon>
        </el-button>
      </el-tooltip>

      <el-tooltip :show-after="tooltipOpenDelay" :hide-after="0" content="网格布局" placement="left">
        <el-button class="action-button" @click="isLayoutModalVisible = true">
          <el-icon>
            <Grid />
          </el-icon>
        </el-button>
      </el-tooltip>
    </div>
  </div>

  <!-- 工具市场模态框 -->
  <el-dialog v-model="isToolMarketVisible" :width="'80%'" :style="{ height: '70vh' }" class="fixed-dialog">
    <!-- 自定义标题区域 -->
    <template #header>
      <div class="dialog-header">
        <el-input v-model="searchQuery" placeholder="按名字或类型搜索工具" clearable class="search-input" suffix-icon="Search" />
      </div>
    </template>

    <!-- 工具市场内容 -->
    <ToolMarket :search-query="searchQuery" @add-tool="addTool" />
  </el-dialog>

  <!-- 布局模态框 -->
  <el-dialog v-model="isLayoutModalVisible" title="选择布局方式" :width="'60%'" :style="{ height: '70vh' }"
    class="fixed-dialog">
    <LayoutModal :current-layout="currentLayout" @change-layout="changeLayout" />
  </el-dialog>
</template>

<style>
html,
body,
#app {
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
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.center-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  /* 水平居中 */
  gap: 0.5rem;
  /* 加号和文本之间的间距 */
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
  grid-template-columns: 1fr 1fr;
  /* 2列 */
  grid-template-rows: 1fr 1fr;
  /* 2行 */
}

.tool-item {
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
  position: relative;
  /* 相对定位，用于关闭按钮的绝对定位 */
  min-width: 0;
  /* 防止内容溢出 */
  min-height: 0;
  /* 防止内容溢出 */
  height: 100%;
  /* 确保子元素高度占满父元素 */
}

.tool-item:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  /* 添加阴影 */
}

/* 拖拽预览样式 */
.tool-item.drag-over {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  /* 添加阴影 */
  transform: scale(1.001);
  /* 轻微放大 */
  transition: transform 0.3s ease;
  /* 平滑过渡 */
}

.close-button {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  cursor: pointer;
  color: #666;
  font-size: 1rem;
}

.close-button:hover {
  color: #333;
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
  justify-content: center;
  /* 将内容居中 */
  width: 100%;
  position: relative;
  /* 相对定位 */
}

.search-input {
  width: 300px;
  /* 搜索框宽度 */
  margin-left: 20px;
  /* 与标题的间距 */
}
</style>

<style>
/* 全局样式，用于固定模态框大小和内容滚动 */
.fixed-dialog .el-dialog__body {
  height: calc(100% - 100px);
  /* 减去标题和底部的高度 */
  overflow-y: auto;
  /* 内容区域滚动 */
  padding: 16px;
}

.fixed-dialog .el-dialog {
  max-height: 70vh;
  /* 固定高度 */
  display: flex;
  flex-direction: column;
}

/* 强制覆盖深色模式的样式 */
body {
  background-color: white !important;
  color: black !important;
}
</style>

<style scoped>
.new-button-group {
  position: fixed;
  right: 1.5rem;
  bottom: 1.5rem;
  display: flex;
  flex-direction: column-reverse;
  align-items: flex-end;
  gap: 0.5rem;
  z-index: 1000;
}

.main-button {
  cursor: pointer;
  background-color: white;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.main-button:hover {
  background-color: #f5f5f5;
}

.main-button .el-icon {
  font-size: 1.5rem;
  transition: transform 0.3s ease;
}

.rotate-icon {
  transform: rotate(45deg);
}

.sub-buttons {
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
  max-height: 0;
  opacity: 0;
  visibility: hidden;
  transition: all 0.3s ease;
  transform: translateY(100%);
  pointer-events: none;
  /* 新增：禁止子按钮交互 */
}

.sub-buttons.expanded {
  max-height: 300px;
  opacity: 1;
  visibility: visible;
  transform: translateY(0);
  margin-bottom: 10px;
  pointer-events: auto;
  /* 新增：允许子按钮交互 */
}

.el-popper {
  transition: opacity 0.3s 0.2s !important;
  /* 添加0.2s延迟 */
}

.action-button {
  border: none;
  background-color: white;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.action-button:hover {
  background-color: #f5f5f5;
}

.action-button .el-icon {
  font-size: 1.8rem;
}
</style>