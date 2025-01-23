<template>
    <div class="calculator-tool">
        <div class="calculator-content">
            <!-- 整合到一个卡片中 -->
            <div class="calculator-card">
                <!-- 标题和帮助按钮 -->
                <div class="section-header">
                    <el-icon class="icon">
                        <Edit />
                    </el-icon>
                    <h3>计算工具</h3>
                    <!-- 帮助图标按钮 -->
                    <el-popover
                        placement="bottom"
                        title=""
                        :width="300"
                        trigger="click"
                    >
                        <template #reference>
                            <el-button type="text" :icon="QuestionFilled" circle />
                        </template>
                        <div class="help-content">
                            <p>支持的运算符和符号：</p>
                            <ul>
                                <li><strong>+</strong> : 加法</li>
                                <li><strong>-</strong> : 减法</li>
                                <li><strong>*</strong> : 乘法</li>
                                <li><strong>/</strong> : 除法</li>
                                <li><strong>%</strong> : 取余</li>
                                <li><strong>**</strong> : 幂运算</li>
                                <li><strong>( )</strong> : 括号（用于优先级）</li>
                            </ul>
                            <p>示例：</p>
                            <ul>
                                <li><code>2 + 3 * 4</code> → 14</li>
                                <li><code>(2 + 3) * 4</code> → 20</li>
                                <li><code>10 % 3</code> → 1</li>
                                <li><code>2 ** 3</code> → 8</li>
                            </ul>
                        </div>
                    </el-popover>
                </div>

                <!-- 输入表达式 -->
                <div class="input-body">
                    <el-input
                        v-model="expression"
                        placeholder="例如：2 + 3 * 4"
                        @keyup.enter="calculate"
                    />
                </div>

                <!-- 输出结果 -->
                <div class="output-body">
                    <el-input
                        v-model="result"
                        readonly
                        placeholder="结果将显示在这里"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { Edit, QuestionFilled } from '@element-plus/icons-vue';

// 响应式状态
const expression = ref('');
const result = ref('');

// 计算表达式
const calculate = () => {
    // 定义允许的字符集：数字、基本运算符和括号
    const allowedChars = /^[\d+\-*/%().\s]+$/;

    // 检查表达式是否为空
    if (!expression.value.trim()) {
        ElMessage.error('表达式不能为空');
        result.value = '表达式为空';
        return;
    }

    // 检查表达式是否只包含允许的字符
    if (!allowedChars.test(expression.value)) {
        ElMessage.error('表达式包含非法字符，请检查输入');
        result.value = '非法字符';
        return;
    }

    try {
        // 使用 eval 计算表达式
        const calculatedResult = eval(expression.value);
        result.value = calculatedResult.toString();
    } catch (error) {
        ElMessage.error('表达式无效，请检查输入');
        result.value = '无效的表达式';
    }
};
</script>

<style scoped>
.calculator-tool {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    box-sizing: border-box;
    border-radius: 0.5rem;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.calculator-content {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    overflow: auto;
}

.calculator-card {
    background: #fff;
    border-radius: 0.5rem;
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    height: 100%;
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

.el-input {
    flex: 1;
    min-width: 200px;
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