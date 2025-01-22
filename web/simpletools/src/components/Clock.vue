<template>
    <div class="clock-tool">
      <div class="tool-header">
        <h3>时钟</h3>
        <el-button type="text" @click="$emit('close')">
          <el-icon><Close /></el-icon>
        </el-button>
      </div>
      <div class="clock-content">
        <div class="time-display">
          {{ currentTime }}
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, onUnmounted } from 'vue'
  import { Close } from '@element-plus/icons-vue'
  
  const currentTime = ref('')
  let timer = null
  
  const updateTime = () => {
    const now = new Date()
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    const seconds = String(now.getSeconds()).padStart(2, '0')
    currentTime.value = `${hours}:${minutes}:${seconds}`
  }
  
  onMounted(() => {
    updateTime()
    timer = setInterval(updateTime, 1000)
  })
  
  onUnmounted(() => {
    if (timer) {
      clearInterval(timer)
    }
  })
  </script>
  
  <style lang="scss" scoped>
  .clock-tool {
    height: 100%;
    display: flex;
    flex-direction: column;
    
    .tool-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 1rem;
  
      h3 {
        margin: 0;
      }
    }
  
    .clock-content {
      flex: 1;
      display: flex;
      align-items: center;
      justify-content: center;
  
      .time-display {
        font-size: 2.5rem;
        font-weight: bold;
        font-family: monospace;
      }
    }
  }
  </style>