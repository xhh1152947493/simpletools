const COLORS = {
    ERROR: '#ff0000', // 红色，用于严重错误
    WARNING: '#ffa500', // 橙色，用于警告
    INFO: '#ffff00', // 黄色，用于提示
    DANGER: '#dc143c', // 深红色，用于危险操作
    CORAL: '#ff7f50', // 珊瑚色，用于温和警告
    TOMATO: '#ff6347', // 番茄红，用于警告或提示
    SUCCESS: '#4caf50', // 绿色，用于成功提示
    LIME: '#00ff00', // 亮绿色，用于强调成功
    FOREST_GREEN: '#228b22', // 深绿色，用于正式的成功提示
    TEAL: '#008080', // 蓝绿色，用于温和的成功提示
    BLUE: '#2196f3', // 蓝色，用于信息提示或操作成功
    SKY_BLUE: '#87ceeb', // 天蓝色，用于轻松的成功提示
  };

  
export const createFloatingText = (content, options = {}) => {
    // 默认配置
    const defaultOptions = {
      container: document.body, // 默认容器为 body
      color: COLORS.TOMATO, // 默认颜色为番茄红
      fontSize: '20px', // 默认文字大小
      duration: 2000, // 默认动画时长
      offsetY: -100, // 默认垂直偏移量
    };
  
    // 合并用户配置和默认配置
    const { container, color, fontSize, duration, offsetY } = {
      ...defaultOptions,
      ...options,
    };
  
    // 创建飘字元素
    const textElement = document.createElement('div');
    textElement.className = 'floating-text';
    textElement.textContent = content;
    Object.assign(textElement.style, {
      position: 'fixed', // 使用 fixed 定位，确保在最顶层
      color,
      fontSize,
      whiteSpace: 'nowrap',
      left: '50%', // 水平居中
      top: '50%', // 垂直居中
      transform: 'translate(-50%, -50%)', // 居中偏移
      zIndex: 9999, // 确保在最顶层
      animation: `floatUp ${duration}ms ease-out forwards`,
    });
  
    // 定义飘字动画
    const styleSheet = document.styleSheets[0];
    styleSheet.insertRule(
      `@keyframes floatUp {
        0% { opacity: 1; transform: translate(-50%, -50%); }
        100% { opacity: 0; transform: translate(-50%, ${offsetY}px); }
      }`,
      styleSheet.cssRules.length
    );
  
    // 将飘字元素添加到容器中
    container.appendChild(textElement);
  
    // 动画结束后移除飘字元素
    setTimeout(() => {
      container.removeChild(textElement);
    }, duration);
  };