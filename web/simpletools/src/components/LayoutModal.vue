<script setup>
defineProps({
  currentLayout: {
    type: String,
    required: true,
  },
});

const emit = defineEmits(['change-layout']);

const layouts = [
  {
    id: 'lyt1',
    label: '布局1',
    type: 'vertical', // 第2个示意图：上下分布
    subType: 'top-two-bottom-one', // 第3个示意图：上方两格，下方一格
  },
  {
    id: 'lyt2',
    label: '布局2',
    type: 'vertical', // 第2个示意图：上下分布
    subType: 'bottom-two-top-one', // 第3个示意图：下方两格，上方一格
  },
  {
    id: 'lyt3',
    label: '布局3',
    type: 'horizontal', // 第2个示意图：左右分布
    subType: 'left-two-right-one', // 第3个示意图：左侧两格，右侧一格
  },
  {
    id: 'lyt4',
    label: '布局4',
    type: 'horizontal', // 第2个示意图：左右分布
    subType: 'right-two-left-one', // 第3个示意图：右侧两格，左侧一格
  },
];

const selectLayout = (layout) => {
  emit('change-layout', layout);
};

const applyLayout = (layout) => {
  emit('change-layout', layout);
  // 这里可以添加其他逻辑，比如关闭模态框等
};
</script>

<template>
  <div class="layout-modal">
    <div
      v-for="layout in layouts"
      :key="layout.id"
      class="layout-item"
      :class="{ active: currentLayout === layout.id }"
      @click="selectLayout(layout.id)"
    >
      <div class="layout-label">{{ layout.label }}</div>
      <div class="layout-preview">
        <!-- 1 个工具 - 固定布局 -->
        <div class="preview-grid single">
          <div class="tool-icon tool-1"></div>
        </div>
        <!-- 2 个工具 - 动态布局 -->
        <div class="preview-grid double" :class="layout.type">
          <div class="tool-icon tool-1"></div>
          <div class="tool-icon tool-2"></div>
        </div>
        <!-- 3 个工具 - 动态布局 -->
        <div class="preview-grid triple" :class="layout.subType">
          <div class="tool-icon tool-1"></div>
          <div class="tool-icon tool-2"></div>
          <div class="tool-icon tool-3"></div>
        </div>
        <!-- 4 个工具 - 固定布局 -->
        <div class="preview-grid quad">
          <div class="tool-icon tool-1"></div>
          <div class="tool-icon tool-2"></div>
          <div class="tool-icon tool-3"></div>
          <div class="tool-icon tool-4"></div>
        </div>
      </div>
      <!-- 当前布局不显示应用按钮 -->
      <button
        v-if="currentLayout !== layout.id"
        class="apply-button"
        @click.stop="applyLayout(layout.id)"
      >
        应用
      </button>
    </div>
  </div>
</template>

<style scoped>
.layout-modal {
  display: flex;
  flex-direction: column;
  gap: 1.5rem; /* 增加行间距 */
}

.layout-item {
  display: flex;
  align-items: center;
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  padding: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.layout-item:hover {
  background-color: #f5f5f5;
}

.layout-item.active {
  border-color: #409eff;
  background-color: #e6f7ff;
}

.layout-label {
  font-size: 1rem;
  font-weight: bold;
  width: 80px; /* 固定宽度 */
}

.layout-preview {
  display: flex;
  gap: 6rem; /* 增加示意图之间的间距 */
  flex-grow: 1;
}

.preview-grid {
  display: grid;
  width: 50px;
  height: 50px;
  border: 1px dashed #ccc;
  border-radius: 0.5rem;
  overflow: hidden;
}

.tool-icon {
  border-radius: 4px;
}

/* 不同工具的颜色 */
.tool-1 {
  background-color: #ff6b6b; /* 红色 */
}

.tool-2 {
  background-color: #4ecdc4; /* 青色 */
}

.tool-3 {
  background-color: #ffe66d; /* 黄色 */
}

.tool-4 {
  background-color: #6b5b95; /* 紫色 */
}

/* 1 个工具 - 固定布局 */
.single {
  grid-template-columns: 1fr;
  grid-template-rows: 1fr;
}

/* 2 个工具 - 水平分布 */
.double.horizontal {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr;
}

/* 2 个工具 - 垂直分布 */
.double.vertical {
  grid-template-columns: 1fr;
  grid-template-rows: 1fr 1fr;
}

/* 3 个工具 - 上方两格，下方一格 */
.triple.top-two-bottom-one {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "tool1 tool2"
    "tool3 tool3";
}

.triple.top-two-bottom-one .tool-icon:nth-child(1) {
  grid-area: tool1;
}

.triple.top-two-bottom-one .tool-icon:nth-child(2) {
  grid-area: tool2;
}

.triple.top-two-bottom-one .tool-icon:nth-child(3) {
  grid-area: tool3;
}

/* 3 个工具 - 下方两格，上方一格 */
.triple.bottom-two-top-one {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "tool3 tool3"
    "tool1 tool2";
}

.triple.bottom-two-top-one .tool-icon:nth-child(1) {
  grid-area: tool1;
}

.triple.bottom-two-top-one .tool-icon:nth-child(2) {
  grid-area: tool2;
}

.triple.bottom-two-top-one .tool-icon:nth-child(3) {
  grid-area: tool3;
}

/* 3 个工具 - 左侧两格，右侧一格 */
.triple.left-two-right-one {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "tool1 tool2"
    "tool1 tool3";
}

.triple.left-two-right-one .tool-icon:nth-child(1) {
  grid-area: tool1;
}

.triple.left-two-right-one .tool-icon:nth-child(2) {
  grid-area: tool2;
}

.triple.left-two-right-one .tool-icon:nth-child(3) {
  grid-area: tool3;
}

/* 3 个工具 - 右侧两格，左侧一格 */
.triple.right-two-left-one {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "tool2 tool1"
    "tool3 tool1";
}

.triple.right-two-left-one .tool-icon:nth-child(1) {
  grid-area: tool1;
}

.triple.right-two-left-one .tool-icon:nth-child(2) {
  grid-area: tool2;
}

.triple.right-two-left-one .tool-icon:nth-child(3) {
  grid-area: tool3;
}

/* 4 个工具 - 固定布局 */
.quad {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
}

.apply-button {
  margin-left: auto;
  padding: 0.5rem 1rem;
  background-color: #409eff;
  color: white;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.apply-button:hover {
  background-color: #66b1ff;
}
</style>