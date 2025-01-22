<script setup>
// 定义工具市场中的工具列表
const tools = [
  {
    name: '工具A',
    type: 'tool-a',
    icon: 'el-icon-s-tools',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: 'tool-b',
    icon: 'el-icon-s-data',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具C',
    type: 'tool-c',
    icon: 'el-icon-s-opportunity',
    description: '这是一个机会管理工具C',
  },
  {
    name: '工具D',
    type: 'tool-d',
    icon: 'el-icon-s-check',
    description: '这是一个验证工具D',
  },
  {
    name: '工具E',
    type: 'tool-e',
    icon: 'el-icon-s-order',
    description: '这是一个订单管理工具E',
  },
  {
    name: '工具F',
    type: 'tool-f',
    icon: 'el-icon-s-platform',
    description: '这是一个平台管理工具F',
  },
  {
    name: '工具A',
    type: 'tool-a',
    icon: 'el-icon-s-tools',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: 'tool-b',
    icon: 'el-icon-s-data',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具C',
    type: 'tool-c',
    icon: 'el-icon-s-opportunity',
    description: '这是一个机会管理工具C',
  },
  {
    name: '工具D',
    type: 'tool-d',
    icon: 'el-icon-s-check',
    description: '这是一个验证工具D',
  },
  {
    name: '工具E',
    type: 'tool-e',
    icon: 'el-icon-s-order',
    description: '这是一个订单管理工具E',
  },
  {
    name: '工具F',
    type: 'tool-f',
    icon: 'el-icon-s-platform',
    description: '这是一个平台管理工具F',
  },
  {
    name: '工具A',
    type: 'tool-a',
    icon: 'el-icon-s-tools',
    description: '这是一个功能强大的工具A',
  },
  {
    name: '工具B',
    type: 'tool-b',
    icon: 'el-icon-s-data',
    description: '这是一个数据分析工具B',
  },
  {
    name: '工具C',
    type: 'tool-c',
    icon: 'el-icon-s-opportunity',
    description: '这是一个机会管理工具C',
  },
  {
    name: '工具D',
    type: 'tool-d',
    icon: 'el-icon-s-check',
    description: '这是一个验证工具D',
  },
  {
    name: '工具E',
    type: 'tool-e',
    icon: 'el-icon-s-order',
    description: '这是一个订单管理工具E',
  },
  {
    name: '工具F',
    type: 'tool-f',
    icon: 'el-icon-s-platform',
    description: '这是一个平台管理工具F',
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
        <el-icon class="tool-icon"><component :is="tool.icon" /></el-icon>
        <div class="tool-title">{{ tool.name }}</div>
        <div class="tool-description">{{ tool.description }}</div>
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
  text-align: center;
  cursor: pointer;
  transition: box-shadow 0.3s ease;
}

.tool-card:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.tool-icon {
  font-size: 2rem;
  color: #409eff;
  margin-bottom: 0.5rem;
}

.tool-title {
  font-size: 1rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}

.tool-description {
  font-size: 0.875rem;
  color: #666;
}
</style>