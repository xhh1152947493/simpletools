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
                                <el-option label="GMT" value="GMT" />
                                <el-option label="Asia/Tokyo" value="Asia/Tokyo" />
                                <el-option label="Asia/Dubai" value="Asia/Dubai" />
                                <el-option label="Europe/London" value="Europe/London" />
                                <el-option label="Europe/Paris" value="Europe/Paris" />
                                <el-option label="Europe/Berlin" value="Europe/Berlin" />
                                <el-option label="America/New_York" value="America/New_York" />
                                <el-option label="America/Los_Angeles" value="America/Los_Angeles" />
                                <el-option label="America/Chicago" value="America/Chicago" />
                                <el-option label="Australia/Sydney" value="Australia/Sydney" />
                                <el-option label="Pacific/Auckland" value="Pacific/Auckland" />
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
                            <el-input v-model="datetime" placeholder="输入日期时间" />
                            <el-select v-model="timezone" style="width:160px">
                                <el-option label="Asia/Shanghai" value="Asia/Shanghai" />
                                <el-option label="UTC" value="UTC" />
                                <el-option label="GMT" value="GMT" />
                                <el-option label="Asia/Tokyo" value="Asia/Tokyo" />
                                <el-option label="Asia/Dubai" value="Asia/Dubai" />
                                <el-option label="Europe/London" value="Europe/London" />
                                <el-option label="Europe/Paris" value="Europe/Paris" />
                                <el-option label="Europe/Berlin" value="Europe/Berlin" />
                                <el-option label="America/New_York" value="America/New_York" />
                                <el-option label="America/Los_Angeles" value="America/Los_Angeles" />
                                <el-option label="America/Chicago" value="America/Chicago" />
                                <el-option label="Australia/Sydney" value="Australia/Sydney" />
                                <el-option label="Pacific/Auckland" value="Pacific/Auckland" />
                            </el-select>
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

// 时间戳转换相关
const timestamp = ref('');
const convertUnit = ref('s');
const timezone = ref('Asia/Shanghai');
const convertResult = ref('');

// 日期时间转换相关
const datetime = ref('');
const timestampSeconds = ref('');
const timestampMilliseconds = ref('');

// 格式化当前时间为 YYYY-MM-DD HH:mm:ss
const formatCurrentTime = () => {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
};

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
    updateTimestamp();
};

// 复制时间戳
const copyTimestamp = async () => {
    try {
        await navigator.clipboard.writeText(currentTimestamp.value.toString());
        ElMessage.success('时间戳已复制到剪贴板');
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
        // 将输入的日期时间字符串解析为指定时区的时间
        const dateString = datetime.value;
        const timeZone = timezone.value;

        // 使用 Intl.DateTimeFormat 将时间字符串转换为指定时区的 Date 对象
        const formatter = new Intl.DateTimeFormat('en-US', {
            timeZone: timeZone,
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false
        });

        // 将日期时间字符串解析为 Date 对象
        const date = new Date(dateString);
        const formattedDateString = formatter.format(date);

        // 将格式化后的时间字符串转换为时间戳
        const timestamp = new Date(formattedDateString).getTime();
        timestampSeconds.value = Math.floor(timestamp / 1000).toString();
        timestampMilliseconds.value = timestamp.toString();
    } catch (error) {
        timestampSeconds.value = '转换失败';
        timestampMilliseconds.value = '转换失败';
    }
};

// 组件挂载时初始化
onMounted(() => {
    updateTimestamp();
    timer = setInterval(updateTimestamp, 1000);

    // 初始化时填充时间戳输入框
    const now = Date.now();
    timestamp.value = unit.value === 's' ? Math.floor(now / 1000).toString() : now.toString();

    // 初始化时填充日期时间输入框
    datetime.value = formatCurrentTime();
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
    box-sizing: border-box;
    border-radius: 0.5rem;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.timestamp-content {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    overflow: auto;
}

.current-timestamp {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    margin: 0;
    padding: 0;
    flex-shrink: 0; /* 防止第一个卡片被拉伸 */
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

.convert-section {
    background: #fff;
    border-radius: 0.5rem;
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    overflow: auto;
}

/* 第一个 convert-section 的高度自适应 */
.convert-section:first-child {
    flex: 0 0 auto; /* 不拉伸，高度根据内容自适应 */
    height: auto; /* 高度自适应 */
}

/* 第二个 convert-section 占据剩余空间 */
.convert-section:last-child {
    flex: 1; /* 占据剩余空间 */
    /* height: auto; 高度自适应 */
}

.converter-group {
    margin-bottom: 0.5rem;
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