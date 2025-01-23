<template>
    <div class="hash-calculator-tool">
        <div class="hash-calculator-content">
            <!-- 整合到一个卡片中 -->
            <div class="hash-calculator-card">
                <!-- 标题 -->
                <div class="section-header">
                    <el-icon class="icon">
                        <Edit />
                    </el-icon>
                    <h3>散列值计算工具</h3>
                </div>

                <!-- 选择散列算法 -->
                <div class="input-body">
                    <el-select v-model="selectedAlgorithm" placeholder="选择散列算法" style="width: 100%">
                        <el-option v-for="algo in algorithms" :key="algo.value" :label="algo.label"
                            :value="algo.value" />
                    </el-select>
                </div>

                <!-- 输入文本 -->
                <div class="input-body">
                    <el-input v-model="inputText" placeholder="输入需要计算散列值的文本" type="textarea" rows="5" />
                </div>

                <!-- 计算按钮 -->
                <div class="input-body">
                    <el-button type="primary" @click="calculateHash" style="width: 100%">计算散列值</el-button>
                </div>

                <!-- 输出结果 -->
                <div class="output-body">
                    <el-input v-model="hashResult" readonly placeholder="计算结果将显示在这里" />
                    <el-button type="success" @click="copyHashResult" :icon="CopyDocument">复制结果</el-button>
                </div>

                <!-- 验证目标哈希值 -->
                <div class="input-body">
                    <el-input v-model="targetHash" placeholder="输入目标哈希值以验证" />
                    <el-button type="warning" @click="verifyHash" style="width: 100%">验证哈希值</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { CopyDocument, Edit } from '@element-plus/icons-vue';
import CryptoJS from 'crypto-js'; // 使用 crypto-js 库

// 支持的散列算法列表
const algorithms = [
    { label: 'MD5', value: 'md5' },
    { label: 'SHA-1', value: 'sha1' },
    { label: 'SHA-224', value: 'sha224' },
    { label: 'SHA-256', value: 'sha256' },
    { label: 'SHA-384', value: 'sha384' },
    { label: 'SHA-512', value: 'sha512' },
    { label: 'SHA3-224', value: 'sha3-224' },
    { label: 'SHA3-256', value: 'sha3-256' },
    { label: 'SHA3-384', value: 'sha3-384' },
    { label: 'SHA3-512', value: 'sha3-512' },
    { label: 'RIPEMD-160', value: 'ripemd160' },
];

// 响应式状态
const selectedAlgorithm = ref('md5'); // 默认选择 MD5
const inputText = ref('');
const hashResult = ref('');
const targetHash = ref('');
const verificationResult = ref('');

// 计算散列值
const calculateHash = () => {
    if (!inputText.value.trim()) {
        ElMessage.error('请输入需要计算散列值的文本');
        return;
    }

    try {
        let hash;
        switch (selectedAlgorithm.value) {
            case 'md5':
                hash = CryptoJS.MD5(inputText.value).toString();
                break;
            case 'sha1':
                hash = CryptoJS.SHA1(inputText.value).toString();
                break;
            case 'sha224':
                hash = CryptoJS.SHA224(inputText.value).toString();
                break;
            case 'sha256':
                hash = CryptoJS.SHA256(inputText.value).toString();
                break;
            case 'sha384':
                hash = CryptoJS.SHA384(inputText.value).toString();
                break;
            case 'sha512':
                hash = CryptoJS.SHA512(inputText.value).toString();
                break;
            case 'sha3-224':
                hash = CryptoJS.SHA3(inputText.value, { outputLength: 224 }).toString();
                break;
            case 'sha3-256':
                hash = CryptoJS.SHA3(inputText.value, { outputLength: 256 }).toString();
                break;
            case 'sha3-384':
                hash = CryptoJS.SHA3(inputText.value, { outputLength: 384 }).toString();
                break;
            case 'sha3-512':
                hash = CryptoJS.SHA3(inputText.value, { outputLength: 512 }).toString();
                break;
            case 'ripemd160':
                hash = CryptoJS.RIPEMD160(inputText.value).toString();
                break;
            default:
                ElMessage.error('不支持的散列算法');
                return;
        }
        hashResult.value = hash;
    } catch (error) {
        ElMessage.error('计算散列值时出错');
        hashResult.value = '计算失败';
    }
};

// 复制散列结果
const copyHashResult = async () => {
    if (!hashResult.value) {
        ElMessage.warning('没有可复制的结果');
        return;
    }

    try {
        await navigator.clipboard.writeText(hashResult.value);
        ElMessage.success('已复制散列值');
    } catch (err) {
        ElMessage.error('复制失败');
    }
};

// 验证目标哈希值
const verifyHash = () => {
    if (!targetHash.value.trim()) {
        ElMessage.error('请输入目标哈希值');
        return;
    }

    if (!hashResult.value) {
        ElMessage.error('请先计算散列值');
        return;
    }

    if (targetHash.value.toLowerCase() === hashResult.value.toLowerCase()) {
        ElMessage.success('哈希值匹配');
    } else {
        ElMessage.error('哈希值不匹配');
    }
};
</script>

<style scoped>
.hash-calculator-tool {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    box-sizing: border-box;
    border-radius: 0.5rem;
    background-color: #f5f5f5;
    /* 父容器背景颜色 */
}

.hash-calculator-content {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    overflow: auto;
    /* 启用滚动条 */
    background-color: #fff;
    /* 内容区域背景颜色 */
    border-radius: 0.5rem;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    padding: 1.5rem;
    /* 内边距 */
}

.hash-calculator-card {
    background: #fff;
    /* 卡片背景颜色 */
    border-radius: 0.5rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
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

.input-body,
.output-body {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.el-input,
.el-select {
    flex: 1;
    min-width: 200px;
}

.el-button {
    width: 100%;
}

.help-content {
    font-size: 14px;
    line-height: 1.6;
}

.help-content ul {
    margin: 0;
    padding-left: 20px;
}

.help-content code {
    background-color: #f5f5f5;
    padding: 2px 4px;
    border-radius: 4px;
    font-family: monospace;
}
</style>