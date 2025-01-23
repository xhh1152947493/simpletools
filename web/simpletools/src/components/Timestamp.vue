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

            <!-- 转换工具卡片 -->
            <div class="convert-section">
                <el-tabs v-model="activeTab" class="convert-tabs">
                    <!-- 时间戳转日期时间 -->
                    <el-tab-pane label="时间戳转日期时间" name="timestamp">
                        <div class="input-group">
                            <el-input v-model="timestamp" placeholder="请输入时间戳" class="timestamp-input" />
                            <el-select v-model="convertUnit" class="unit-select">
                                <el-option label="秒(s)" value="s" />
                                <el-option label="毫秒(ms)" value="ms" />
                            </el-select>
                        </div>
                        <div class="result-section">
                            <el-input v-model="convertResult" readonly placeholder="转换结果" class="result-input" />
                            <el-select v-model="timezone" class="timezone-select">
                                <el-option label="Asia/Shanghai" value="Asia/Shanghai" />
                                <el-option label="UTC" value="UTC" />
                            </el-select>
                        </div>
                        <el-button type="primary" class="convert-button" @click="startConvert">转换</el-button>
                    </el-tab-pane>

                    <!-- 日期时间转时间戳 -->
                    <el-tab-pane label="日期时间转时间戳" name="datetime">
                        <div class="input-group">
                            <el-date-picker v-model="datetime" type="datetime" placeholder="选择日期时间"
                                format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss"
                                class="datetime-picker" />
                        </div>
                        <div class="timestamp-results">
                            <div class="result-item">
                                <span class="result-label">秒级时间戳：</span>
                                <el-input v-model="timestampSeconds" readonly class="result-input" />
                            </div>
                            <div class="result-item">
                                <span class="result-label">毫秒级时间戳：</span>
                                <el-input v-model="timestampMilliseconds" readonly class="result-input" />
                            </div>
                        </div>
                        <el-button type="primary" class="convert-button" @click="convertToTimestamp">转换</el-button>
                    </el-tab-pane>
                </el-tabs>
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
    gap: 2rem;
}

.convert-section {
    background-color: #fff;
    padding: 1.5rem;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.current-timestamp {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    /* 控制整体内容的上下间距 */
    margin: 0;
    padding: 0;
}

.timestamp-header {
    display: flex;
    flex-direction: column;
    gap: 0rem;
    /* 控制标题和时间戳之间的间距 */
}

.section-title {
    font-size: 1.1rem;
    font-weight: bold;
    color: #333;
    margin-bottom: 0;
    /* 移除默认的 margin-bottom */
}

.timestamp-display {
    display: flex;
    flex-direction: column;
    gap: 0.1rem;
    /* 控制时间戳和按钮之间的间距 */
    align-items: left;
    margin: 0;
    padding: 0;
}

.timestamp-wrapper {
    display: inline-flex;
    align-items: baseline;
    /* 基于基线对齐 */
}

.timestamp-value {
    font-size: 2rem;
    font-family: monospace;
    color: #000000;
    /* 修改为黑色 */
}

.timestamp-unit {
    font-size: 1rem;
    /* 缩小单位字体 */
    font-family: monospace;
    color: #666;
    /* 单位颜色稍浅 */
    margin-left: 0.25rem;
    /* 与时间戳数字的间距 */
    line-height: 1;
    /* 将行高设置为 1，避免行高影响对齐 */
    position: relative;
    bottom: 0.1rem;
    /* 微调垂直位置，根据需要调整 */
}

.timestamp-controls {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
    justify-content: left;
}

.input-group {
    display: flex;
    gap: 1rem;
    margin-bottom: 1rem;
}

.timestamp-input,
.datetime-picker {
    flex: 1;
}

.unit-select {
    width: 120px;
}

.result-section {
    display: flex;
    gap: 1rem;
    margin-bottom: 1rem;
}

.result-input {
    flex: 1;
}

.timezone-select {
    width: 200px;
}

.timestamp-results {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-bottom: 1rem;
}

.result-item {
    display: flex;
    align-items: center;
    gap: 1rem;
}

.result-label {
    min-width: 100px;
    color: #666;
}

.convert-button {
    width: 100%;
}
</style>