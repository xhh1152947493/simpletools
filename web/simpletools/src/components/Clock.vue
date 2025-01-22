<template>
  <div class="clock-tool">
    <div class="clock-content">
      <!-- 显示年月日和星期 -->
      <div class="date-display">
        {{ currentDate }}
      </div>
      <!-- 显示时间 -->
      <div class="time-display">
        {{ currentTime }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';

// 当前时间
const currentTime = ref('');
// 当前日期（年月日和星期）
const currentDate = ref('');
let timer = null;

// 更新时间
const updateTime = () => {
  const now = new Date();

  // 格式化时间
  const hours = String(now.getHours()).padStart(2, '0');
  const minutes = String(now.getMinutes()).padStart(2, '0');
  const seconds = String(now.getSeconds()).padStart(2, '0');
  currentTime.value = `${hours}:${minutes}:${seconds}`;

  // 格式化日期和星期
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0'); // 月份从 0 开始，需要加 1
  const day = String(now.getDate()).padStart(2, '0');
  const weekdays = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
  const weekday = weekdays[now.getDay()]; // 获取星期几
  currentDate.value = `${year}年${month}月${day}日 ${weekday}`;
};

// 组件挂载时启动定时器
onMounted(() => {
  updateTime(); // 立即更新
  timer = setInterval(updateTime, 1000); // 每秒更新一次
});

// 组件卸载时清除定时器
onUnmounted(() => {
  if (timer) {
    clearInterval(timer);
  }
});
</script>

<style lang="scss" scoped>
.clock-tool {
  height: 100%;
  width: 100%; /* 占满父元素的宽度和高度 */
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 1em; /* 使用相对单位 */
}

.clock-content {
  text-align: center;
  width: 100%; /* 占满父元素的宽度 */
}

.date-display {
  font-size: clamp(1rem, 3vw, 1.5rem); /* 动态字体大小 */
  font-weight: bold;
  color: #333;
  margin-bottom: 0.5em; /* 使用相对单位 */
}

.time-display {
  font-size: clamp(2rem, 6vw, 3rem); /* 动态字体大小 */
  font-weight: bold;
  font-family: monospace;
  color: #000;
}

/* 媒体查询：在小屏幕上调整布局 */
@media (max-width: 600px) {
  .date-display {
    font-size: clamp(0.8rem, 4vw, 1.2rem); /* 更小的字体大小 */
  }

  .time-display {
    font-size: clamp(1.5rem, 8vw, 2.5rem); /* 更小的字体大小 */
  }
}
</style>