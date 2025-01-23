<script setup>
import { computed } from 'vue';
import Clock from '../components/Clock.vue'; // 导入 Clock 组件
import Timestamp from '../components/Timestamp.vue'; // 导入 Timestamp 组件
import AITranslate  from '../components/AITranslate.vue';
import Calculator from '../components/Calculator.vue';
import Hash from '../components/Hash.vue';
import AICodeNamed from '../components/AICodeNamed.vue';

// 定义工具市场中的工具列表
const tools = [
  {
    id: 'tool-clock',
    name: '在线时钟',
    type: '显示类',
    icon: 'Clock',
    description: '显示当前详细的实时时间',
    component: Clock, // 组件对象
    componentName: 'Clock', // 组件名称
  },
  {
    id: 'tool-timestamp',
    name: '时间戳转换',
    type: '工具类',
    icon: 'Clock',
    description: '时间戳与时间的相互转换',
    component: Timestamp, // 组件对象
    componentName: 'Timestamp', // 组件名称
  },
  {
    id: 'tool-aitranslate',
    name: 'AI翻译',
    type: '翻译类',
    icon: 'Connection',
    description: '支持中英互译的AI翻译工具',
    component: AITranslate, // 组件对象
    componentName: 'AITranslate', // 组件名称
  },
  {
    id: 'tool-calculator',
    name: '在线计算器',
    type: '计算类',
    icon: 'EditPen',
    description: '支持基本的数学表达式运算的在线计算器',
    component: Calculator, // 组件对象
    componentName: 'Calculator', // 组件名称
  },
  {
    id: 'tool-hash',
    name: '哈希计算',
    type: '计算类',
    icon: 'Tools',
    description: '支持常见哈希算法的在线哈希计算工具',
    component: Hash, // 组件对象
    componentName: 'Hash', // 组件名称
  },
  {
    id: 'tool-AICodeNamed',
    name: 'AI变量函数命名',
    type: '翻译类',
    icon: 'Monitor',
    description: '智能命名变量和函数名的工具',
    component: AICodeNamed, // 组件对象
    componentName: 'AICodeNamed', // 组件名称
  },
];

// 接收父组件传递的搜索关键词
const props = defineProps({
  searchQuery: String,
});

// 过滤后的工具列表
const filteredTools = computed(() => {
  if (!props.searchQuery) {
    return tools;
  }
  return tools.filter(
    (tool) =>
      tool.name.toLowerCase().includes(props.searchQuery.toLowerCase()) ||
      tool.type.toLowerCase().includes(props.searchQuery.toLowerCase())
  );
});

// 向父组件传递添加的工具
const emit = defineEmits(['add-tool']);
const handleAddTool = (tool) => {
  emit('add-tool', tool);
};
</script>

<template>
  <div class="tool-market">
    <!-- 工具列表 -->
    <div class="tool-grid">
      <div
        v-for="(tool, index) in filteredTools"
        :key="index"
        class="tool-card"
        @click="handleAddTool(tool)"
      >
        <!-- 图标和工具名称 -->
        <div class="tool-header">
          <el-icon class="tool-icon"><component :is="tool.icon" /></el-icon>
          <div class="tool-info">
            <div class="tool-title">{{ tool.name }}</div>
            <div class="tool-type">{{ tool.type }}</div>
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