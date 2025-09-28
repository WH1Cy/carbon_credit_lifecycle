import { ref, computed, watch } from 'vue'
import type { ThemeConfig } from 'ant-design-vue/es/config-provider/context'

// 主题类型
export type ThemeMode = 'light' | 'dark'

// 主题配置接口
interface ThemeColors {
  // 主色调
  primary: string
  success: string
  warning: string
  error: string
  info: string
  
  // 背景色
  bgContainer: string
  bgElevated: string
  bgLayout: string
  bgMask: string
  
  // 文字颜色
  text: string
  textSecondary: string
  textTertiary: string
  textQuaternary: string
  
  // 边框颜色
  border: string
  borderSecondary: string
  split: string
}

// 浅色主题配色
const lightThemeColors: ThemeColors = {
  // GitHub风格配色
  primary: '#0969da',
  success: '#1a7f37',
  warning: '#9a6700',
  error: '#d1242f',
  info: '#0969da',
  
  // 背景色
  bgContainer: '#ffffff',
  bgElevated: '#f6f8fa',
  bgLayout: '#ffffff',
  bgMask: 'rgba(0, 0, 0, 0.5)',
  
  // 文字颜色
  text: '#24292f',
  textSecondary: '#656d76',
  textTertiary: '#8b949e',
  textQuaternary: '#d0d7de',
  
  // 边框颜色
  border: '#d0d7de',
  borderSecondary: '#d8dee4',
  split: '#d0d7de'
}

// 深色主题配色
const darkThemeColors: ThemeColors = {
  // 深色主题配色
  primary: '#58a6ff',
  success: '#3fb950',
  warning: '#d29922',
  error: '#f85149',
  info: '#58a6ff',
  
  // 背景色
  bgContainer: '#0d1117',
  bgElevated: '#161b22',
  bgLayout: '#0d1117',
  bgMask: 'rgba(0, 0, 0, 0.8)',
  
  // 文字颜色
  text: '#f0f6fc',
  textSecondary: '#8b949e',
  textTertiary: '#6e7681',
  textQuaternary: '#30363d',
  
  // 边框颜色
  border: '#30363d',
  borderSecondary: '#21262d',
  split: '#30363d'
}

// 主题状态
const themeMode = ref<ThemeMode>('light')

// 从localStorage读取主题设置
try {
  const savedTheme = localStorage.getItem('theme-mode') as ThemeMode
  if (savedTheme && ['light', 'dark'].includes(savedTheme)) {
    themeMode.value = savedTheme
  }
} catch (error) {
  console.warn('无法读取主题设置，使用默认主题:', error)
  themeMode.value = 'light'
}

// 当前主题配色
const currentThemeColors = computed(() => {
  return themeMode.value === 'light' ? lightThemeColors : darkThemeColors
})

// Ant Design主题配置
const antdThemeConfig = computed((): ThemeConfig => {
  const colors = currentThemeColors.value
  
  return {
    token: {
      // 主色调
      colorPrimary: colors.primary,
      colorSuccess: colors.success,
      colorWarning: colors.warning,
      colorError: colors.error,
      colorInfo: colors.info,
      
      // 背景色
      colorBgContainer: colors.bgContainer,
      colorBgElevated: colors.bgElevated,
      colorBgLayout: colors.bgLayout,
      colorBgMask: colors.bgMask,
      
      // 文字颜色
      colorText: colors.text,
      colorTextSecondary: colors.textSecondary,
      colorTextTertiary: colors.textTertiary,
      colorTextQuaternary: colors.textQuaternary,
      
      // 边框和分割线
      colorBorder: colors.border,
      colorBorderSecondary: colors.borderSecondary,
      colorSplit: colors.split,
      
      // 圆角
      borderRadius: 6,
      borderRadiusLG: 8,
      borderRadiusSM: 4,
      borderRadiusXS: 2,
      
      // 阴影
      boxShadow: themeMode.value === 'light' 
        ? '0 1px 0 rgba(27, 31, 36, 0.04)'
        : '0 1px 0 rgba(0, 0, 0, 0.2)',
      boxShadowSecondary: themeMode.value === 'light'
        ? '0 3px 6px rgba(140, 149, 159, 0.15)'
        : '0 3px 6px rgba(0, 0, 0, 0.3)',
      boxShadowTertiary: themeMode.value === 'light'
        ? '0 8px 24px rgba(140, 149, 159, 0.2)'
        : '0 8px 24px rgba(0, 0, 0, 0.4)',
      
      // 字体
      fontSize: 14,
      fontSizeLG: 16,
      fontSizeSM: 12,
      fontSizeXL: 20,
      fontSizeHeading1: 32,
      fontSizeHeading2: 24,
      fontSizeHeading3: 20,
      fontSizeHeading4: 16,
      fontSizeHeading5: 14,
      
      // 字体权重
      fontWeightStrong: 600,
      
      // 字体族
      fontFamily: "'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif",
      
      // 间距
      padding: 16,
      paddingLG: 24,
      paddingSM: 12,
      paddingXS: 8,
      paddingXXS: 4,
      
      // 动画
      motionDurationSlow: '0.2s',
      motionDurationMid: '0.15s',
      motionDurationFast: '0.1s',
      motionEaseInOut: 'cubic-bezier(0.4, 0, 0.2, 1)',
      motionEaseOut: 'cubic-bezier(0, 0, 0.2, 1)',
    },
    components: {
      Card: {
        borderRadiusLG: 8,
        boxShadow: themeMode.value === 'light' 
          ? '0 1px 0 rgba(27, 31, 36, 0.04)'
          : '0 1px 0 rgba(0, 0, 0, 0.2)',
        boxShadowTertiary: themeMode.value === 'light'
          ? '0 3px 6px rgba(140, 149, 159, 0.15)'
          : '0 3px 6px rgba(0, 0, 0, 0.3)',
        colorBorder: colors.border,
        colorBorderSecondary: colors.borderSecondary,
      },
      Button: {
        borderRadius: 6,
        controlHeight: 32,
        boxShadow: themeMode.value === 'light' 
          ? '0 1px 0 rgba(27, 31, 36, 0.04)'
          : '0 1px 0 rgba(0, 0, 0, 0.2)',
      },
      Input: {
        borderRadius: 6,
        controlHeight: 32,
        colorBorder: colors.border,
        boxShadow: themeMode.value === 'light' 
          ? '0 1px 0 rgba(27, 31, 36, 0.04)'
          : '0 1px 0 rgba(0, 0, 0, 0.2)',
      },
      Select: {
        borderRadius: 6,
        controlHeight: 32,
        colorBorder: colors.border,
      },
      Table: {
        borderRadius: 6,
        colorBorder: colors.border,
      },
      Statistic: {
        // contentFontSize: 20, // 暂时注释掉，Ant Design版本可能不支持
      },
      Tag: {
        borderRadiusSM: 4,
        fontSizeSM: 11,
      },
      Pagination: {
        borderRadius: 6,
      },
      Modal: {
        borderRadius: 8,
        boxShadow: themeMode.value === 'light'
          ? '0 8px 24px rgba(140, 149, 159, 0.2)'
          : '0 8px 24px rgba(0, 0, 0, 0.4)',
      },
      Drawer: {
        borderRadius: 8,
      },
      Dropdown: {
        borderRadius: 6,
        boxShadow: themeMode.value === 'light'
          ? '0 3px 6px rgba(140, 149, 159, 0.15)'
          : '0 3px 6px rgba(0, 0, 0, 0.3)',
      },
      Tooltip: {
        borderRadius: 6,
        boxShadow: themeMode.value === 'light'
          ? '0 3px 6px rgba(140, 149, 159, 0.15)'
          : '0 3px 6px rgba(0, 0, 0, 0.3)',
      },
      Menu: {
        borderRadius: 6,
      },
      Tabs: {
        borderRadius: 6,
      },
      Badge: {
        borderRadius: 10,
      }
    }
  }
})

// 切换主题
const toggleTheme = () => {
  themeMode.value = themeMode.value === 'light' ? 'dark' : 'light'
}

// 设置主题
const setTheme = (mode: ThemeMode) => {
  themeMode.value = mode
}

// 监听主题变化，保存到localStorage
watch(themeMode, (newMode) => {
  try {
    localStorage.setItem('theme-mode', newMode)
  } catch (error) {
    console.warn('无法保存主题设置:', error)
  }
  
  // 更新HTML根元素的class
  try {
    const html = document.documentElement
    if (newMode === 'dark') {
      html.classList.add('dark-theme')
      html.classList.remove('light-theme')
    } else {
      html.classList.add('light-theme')
      html.classList.remove('dark-theme')
    }
  } catch (error) {
    console.warn('无法更新HTML主题类:', error)
  }
}, { immediate: true })

// 初始化HTML根元素class
const initTheme = () => {
  try {
    const html = document.documentElement
    if (themeMode.value === 'dark') {
      html.classList.add('dark-theme')
      html.classList.remove('light-theme')
    } else {
      html.classList.add('light-theme')
      html.classList.remove('dark-theme')
    }
  } catch (error) {
    console.error('主题初始化失败:', error)
  }
}

// 导出主题管理功能
export function useTheme() {
  return {
    // 状态
    themeMode: computed(() => themeMode.value),
    isDark: computed(() => themeMode.value === 'dark'),
    isLight: computed(() => themeMode.value === 'light'),
    currentThemeColors,
    
    // 配置
    antdThemeConfig,
    
    // 方法
    toggleTheme,
    setTheme,
    initTheme
  }
}
