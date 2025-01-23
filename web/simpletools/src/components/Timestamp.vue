<template>
    <div class="timestamp-tool">
        <div class="timestamp-content">
            <!-- 当前时间戳卡片 -->
            <div class="convert-section">
                <div class="current-timestamp">
                    <div class="timestamp-header">
                        <div class="section-title">当前时间戳</div>
                        <div class="timestamp-display">
                            <div class="timestamp-wrapper">
                                <span class="timestamp-value">{{ currentTimestamp }}</span>
                                <span class="timestamp-unit">{{ unit === 's' ? '秒' : '毫秒' }}</span>
                            </div>
                            <div class="timestamp-controls">
                                <el-button-group>
                                    <el-button :type="unit === 's' ? 'primary' : ''"
                                        @click="changeUnit('s')">秒</el-button>
                                    <el-button :type="unit === 'ms' ? 'primary' : ''"
                                        @click="changeUnit('ms')">毫秒</el-button>
                                </el-button-group>
                                <el-button type="primary" @click="copyTimestamp" :icon="CopyDocument">复制</el-button>
                                <el-button :type="isPaused ? 'success' : 'danger'" @click="togglePause"
                                    :icon="isPaused ? 'VideoPlay' : 'VideoPause'">
                                    {{ isPaused ? '开始' : '暂停' }}
                                </el-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- 整合后的转换工具卡片 -->
            <div class="convert-section">
                <!-- 时间戳转日期时间 -->
                <div class="converter-group">
                    <div class="section-header">
                        <el-icon class="icon">
                            <Clock />
                        </el-icon>
                        <h3>时间戳转日期时间</h3>
                    </div>
                    <div class="converter-body">
                        <div class="input-row">
                            <el-input v-model="timestamp" placeholder="输入时间戳" />
                            <el-select v-model="convertUnit" style="width:120px">
                                <el-option label="秒(s)" value="s" />
                                <el-option label="毫秒(ms)" value="ms" />
                            </el-select>
                            <el-button type="primary" @click="startConvert">转换</el-button>
                        </div>
                        <div class="result-row">
                            <el-input v-model="convertResult" readonly placeholder="转换结果" />
                            <el-select v-model="timezone" style="width:160px">
                                <el-option label="Asia/Shanghai" value="Asia/Shanghai" />
                                <el-option label="UTC" value="UTC" />
                            </el-select>
                        </div>
                    </div>
                </div>

                <!-- 日期时间转时间戳 -->
                <div class="converter-group">
                    <div class="section-header">
                        <el-icon class="icon">
                            <Calendar />
                        </el-icon>
                        <h3>日期时间转时间戳</h3>
                    </div>
                    <div class="converter-body">
                        <div class="input-row">
                            <el-date-picker v-model="datetime" type="datetime" placeholder="选择日期时间"
                                format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" style="flex:1" />
                            <el-button type="primary" @click="convertToTimestamp">转换</el-button>
                        </div>
                        <div class="result-row">
                            <el-input v-model="timestampSeconds" readonly placeholder="秒级时间戳" />
                            <el-input v-model="timestampMilliseconds" readonly placeholder="毫秒级时间戳" />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { ElMessage } from 'element-plus';
import { CopyDocument } from '@element-plus/icons-vue';

// 响应式状态
const unit = ref('s');
const isPaused = ref(false);
const currentTimestamp = ref('');
const activeTab = ref('timestamp');

// 时间戳转换相关
const timestamp = ref('');
const convertUnit = ref('s');
const timezone = ref('Asia/Shanghai');
const convertResult = ref('');

// 日期时间转换相关
const datetime = ref('');
const timestampSeconds = ref('');
const timestampMilliseconds = ref('');

// 更新当前时间戳
let timer = null;
const updateTimestamp = () => {
    if (!isPaused.value) {
        const now = Date.now();
        currentTimestamp.value = unit.value === 's' ? Math.floor(now / 1000) : now;
    }
};

const changeUnit = (newUnit) => {
    unit.value = newUnit;
    updateTimestamp(); // 立即更新时间戳
};

// 复制时间戳
const copyTimestamp = async () => {
    try {
        await navigator.clipboard.writeText(currentTimestamp.value.toString());
        ElMessage.success('复制成功');
    } catch (err) {
        ElMessage.error('复制失败');
    }
};

// 切换暂停/开始
const togglePause = () => {
    isPaused.value = !isPaused.value;
    if (isPaused.value) {
        ElMessage.error('已暂停计时');
    } else {
        ElMessage.success('已开始计时');
    }
};

// 时间戳转日期时间
const startConvert = () => {
    if (!timestamp.value) return;

    try {
        let timestampMs = convertUnit.value === 's'
            ? Number(timestamp.value) * 1000
            : Number(timestamp.value);

        const date = new Date(timestampMs);
        const formatter = new Intl.DateTimeFormat('zh-CN', {
            timeZone: timezone.value,
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false
        });

        convertResult.value = formatter.format(date);
    } catch (error) {
        convertResult.value = '无效的时间戳';
    }
};

// 日期时间转时间戳
const convertToTimestamp = () => {
    if (!datetime.value) return;

    try {
        const date = new Date(datetime.value);
        const timestamp = date.getTime();
        timestampSeconds.value = Math.floor(timestamp / 1000).toString();
        timestampMilliseconds.value = timestamp.toString();
    } catch (error) {
        timestampSeconds.value = '转换失败';
        timestampMilliseconds.value = '转换失败';
    }
};

// 组件挂载时启动定时器
onMounted(() => {
    updateTimestamp();
    timer = setInterval(updateTimestamp, 1000);
});

// 组件卸载时清除定时器
onUnmounted(() => {
    if (timer) {
        clearInterval(timer);
    }
});
</script>

<style scoped>
/* 基础容器样式 */
.timestamp-tool {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1.5rem;
    box-sizing: border-box;
    background-color: #f5f5f5;
}

.timestamp-content {
    width: 100%;
    max-width: 800px;
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

/* 当前时间戳样式 */
.current-timestamp {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    margin: 0;
    padding: 0;
}

.timestamp-header {
    display: flex;
    flex-direction: column;
    gap: 0;
}

.section-title {
    font-size: 1.1rem;
    font-weight: bold;
    color: #333;
    margin-bottom: 0;
}

.timestamp-display {
    display: flex;
    flex-direction: column;
    gap: 0.1rem;
    align-items: flex-start;
}

.timestamp-wrapper {
    display: inline-flex;
    align-items: baseline;
}

.timestamp-value {
    font-size: 2rem;
    font-family: monospace;
    color: #000;
}

.timestamp-unit {
    font-size: 1rem;
    font-family: monospace;
    color: #666;
    margin-left: 0.25rem;
    line-height: 1;
    position: relative;
    bottom: 0.1rem;
}

.timestamp-controls {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
    justify-content: flex-start;
}

/* 转换工具公共样式 */
.convert-section {
    background: #fff;
    border-radius: 8px;
    padding: 1.5rem;
}

.converter-group {
    margin-bottom: 1.5rem;
}

.converter-group:last-child {
    margin-bottom: 0;
}

.section-header {
    display: flex;
    align-items: center;
    margin-bottom: 12px;
}

.section-header h3 {
    margin: 0;
    font-size: 16px;
    color: #333;
}

.icon {
    margin-right: 8px;
    font-size: 18px;
    color: #409EFF;
}

/* 输入输出行布局 */
.converter-body {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.input-row,
.result-row {
    display: flex;
    gap: 10px;
    align-items: center;
}

.input-row {
    margin-bottom: 8px;
}

.result-row {
    flex-wrap: wrap;
}

/* 表单元素样式 */
.el-input,
.el-date-picker {
    flex: 1;
    min-width: 200px;
}

.el-select {
    width: auto;
}

.el-button {
    height: 32px;
    padding: 0 15px;
    flex-shrink: 0;
}
</style>