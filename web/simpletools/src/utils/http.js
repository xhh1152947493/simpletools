// src/utils/http.js
import axios from 'axios';

// 生成随机字符串
function generateRandomString(length) {
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return result;
}

// 封装 POST 请求
export async function httpPost(url, data) {
    try {
        const timestamp = Math.floor(Date.now() / 1000);
        const randomString = Date.now() + generateRandomString(32);

        const headers = {
            'Content-Type': 'application/json',
            'X-Timestamp': timestamp.toString(),
            'X-Nonce': randomString
        };

        const response = await axios.post(import.meta.env.VITE_API_URL + url, data, { headers });
        return response.data; // 返回响应数据
    } catch (error) {
        throw error; // 抛出错误以便调用方处理
    }
}
