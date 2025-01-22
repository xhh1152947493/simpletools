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
    id: 'default',
    label: '布局1',
  },
  {
    id: 'horizontal',
    label: '布局2',
  },
  {
    id: 'vertical',
    label: '布局3',
  },
  {
    id: 'left-two-right-one',
    label: '布局4',
  },
  {
    id: 'right-two-left-one',
    label: '布局5',
  },
  {
    id: 'top-two-bottom-one',
    label: '布局6',
  },
  {
    id: 'bottom-two-top-one',
    label: '布局7',
  },
  {
    id: 'four-grid',
    label: '布局8',
  },
];

const selectLayout = (layout) => {
  emit('change-layout', layout);
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
        <!-- 1 个工具 -->
        <div class="preview-grid single">
          <div class="tool-icon"></div>
        </div>
        <!-- 2 个工具 -->
        <div class="preview-grid double">
          <div class="tool-icon"></div>
          <div class="tool-icon"></div>
        </div>
        <!-- 3 个工具 -->
        <div class="preview-grid triple">
          <div class="tool-icon"></div>
          <div class="tool-icon"></div>
          <div class="tool-icon"></div>
        </div>
        <!-- 4 个工具 -->
        <div class="preview-grid quad">
          <div class="tool-icon"></div>
          <div class="tool-icon"></div>
          <div class="tool-icon"></div>
          <div class="tool-icon"></div>
        </div>
      </div>
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
  gap: 1.5rem; /* 增加示意图之间的间距 */
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
  background-color: #e0e0e0; /* 灰色方块 */
  border-radius: 4px;
}

/* 1 个工具 */
.single {
  grid-template-columns: 1fr;
  grid-template-rows: 1fr;
}

/* 2 个工具 */
.double {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr;
}

/* 3 个工具 */
.triple {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "tool1 tool2"
    "tool3 tool3";
}

.triple .tool-icon:nth-child(1) {
  grid-area: tool1;
}

.triple .tool-icon:nth-child(2) {
  grid-area: tool2;
}

.triple .tool-icon:nth-child(3) {
  grid-area: tool3;
}

/* 4 个工具 */
.quad {
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
}
</style>