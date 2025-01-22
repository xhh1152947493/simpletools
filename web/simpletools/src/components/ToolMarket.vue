<script setup>
// 定义工具市场中的工具列表
const tools = [
  {
    name: '工具A',
    type: '开发类', // 新增工具类型
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类', // 新增工具类型
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类',
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类',
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类',
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类',
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类', // 新增工具类型
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类', // 新增工具类型
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类',
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类',
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类',
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类',
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类', // 新增工具类型
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类', // 新增工具类型
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类',
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类',
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具A',
    type: '开发类',
    icon: 'Delete',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: '视频类',
    icon: 'FolderChecked',
    description: '这是一个数据分析工具B',
  },
];

// 向父组件传递添加的工具
const emit = defineEmits(['add-tool']);
const handleAddTool = (tool) => {
  emit('add-tool', tool.type);
};
</script>

<template>
  <div class="tool-market">
    <div class="tool-grid">
      <div v-for="(tool, index) in tools" :key="index" class="tool-card" @click="handleAddTool(tool)">
        <!-- 图标和工具名称 -->
        <div class="tool-header">
          <el-icon class="tool-icon"><component :is="tool.icon" /></el-icon>
          <div class="tool-info">
            <div class="tool-title">{{ tool.name }}</div>
            <div class="tool-type">{{ tool.type }}</div> <!-- 新增工具类型 -->
          </div>
        </div>
        <!-- 工具描述 -->
        <div class="tool-description">{{ tool.description }}</div>
        <!-- 加号图标 -->
        <el-icon class="add-icon">+</el-icon>
      </div>
    </div>
  </div>
</template>

<style scoped>
.tool-market {
  padding: 1rem;
}

.tool-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr); /* 每行 4 个工具 */
  gap: 1rem; /* 工具卡片之间的间距 */
}

.tool-card {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 1rem;
  cursor: pointer;
  transition: box-shadow 0.3s ease;
  position: relative; /* 为加号图标定位 */
}

.tool-card:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

/* 图标和工具名称的布局 */
.tool-header {
  display: flex;
  align-items: flex-start; /* 顶部对齐 */
  margin-bottom: 0.5rem; /* 与描述之间的间距 */
}

.tool-icon {
  font-size: 3rem; /* 图标大小 */
  color: #409eff;
  margin-right: 0.5rem; /* 图标与标题之间的间距 */
}

.tool-info {
  display: flex;
  flex-direction: column;
  justify-content: center; /* 垂直居中 */
}

.tool-title {
  font-size: 1rem;
  font-weight: bold;
  margin-bottom: 0.25rem; /* 工具名称与类型之间的间距 */
}

.tool-type {
  font-size: 0.75rem;
  color: #999;
}

.tool-description {
  font-size: 0.875rem;
  color: #666;
}

/* 加号图标 */
.add-icon {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  font-size: 1.5rem;
  color: #999;
  opacity: 0; /* 默认隐藏 */
  transition: opacity 0.3s ease;
}

.tool-card:hover .add-icon {
  opacity: 1; /* 鼠标悬停时显示 */
}
</style>