<template>
  <div :class="themeClass">
    <a-config-provider :theme="antdThemeConfig">
      <RouterView />
    </a-config-provider>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted, computed } from 'vue'
import { useTheme } from '@/composables/useTheme'

// 使用主题系统
const { antdThemeConfig, initTheme, themeMode } = useTheme()

// 计算主题类名
const themeClass = computed(() => {
  return themeMode.value === 'dark' ? 'dark-theme' : 'light-theme'
})

// 初始化主题
onMounted(() => {
  try {
    initTheme()
  } catch (error) {
    console.error('主题初始化失败:', error)
  }
})
</script>

<style>
/* 导入现代化字体 */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&family=JetBrains+Mono:wght@400;500;600&display=swap');

/* 主题变量定义 */
:root {
  /* 浅色主题变量 */
  --color-primary: #0969da;
  --color-success: #1a7f37;
  --color-warning: #9a6700;
  --color-error: #d1242f;
  --color-info: #0969da;
  
  --color-bg-container: #ffffff;
  --color-bg-elevated: #f6f8fa;
  --color-bg-layout: #ffffff;
  --color-bg-mask: rgba(0, 0, 0, 0.5);
  
  --color-text: #24292f;
  --color-text-secondary: #656d76;
  --color-text-tertiary: #8b949e;
  --color-text-quaternary: #d0d7de;
  
  --color-border: #d0d7de;
  --color-border-secondary: #d8dee4;
  --color-split: #d0d7de;
  
  --box-shadow: 0 1px 0 rgba(27, 31, 36, 0.04);
  --box-shadow-secondary: 0 3px 6px rgba(140, 149, 159, 0.15);
  --box-shadow-tertiary: 0 8px 24px rgba(140, 149, 159, 0.2);
  
  /* 字体变量 */
  --font-family-primary: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif;
  --font-family-mono: 'JetBrains Mono', 'SF Mono', Monaco, Inconsolata, 'Roboto Mono', 'Source Code Pro', monospace;
  
  /* 字体大小 */
  --font-size-xs: 12px;
  --font-size-sm: 14px;
  --font-size-base: 16px;
  --font-size-lg: 18px;
  --font-size-xl: 20px;
  --font-size-2xl: 24px;
  --font-size-3xl: 30px;
  --font-size-4xl: 36px;
  --font-size-5xl: 48px;
  
  /* 字体粗细 */
  --font-weight-light: 300;
  --font-weight-normal: 400;
  --font-weight-medium: 500;
  --font-weight-semibold: 600;
  --font-weight-bold: 700;
  --font-weight-extrabold: 800;
  
  /* 行高 */
  --line-height-tight: 1.25;
  --line-height-normal: 1.5;
  --line-height-relaxed: 1.75;
  --line-height-loose: 2;
  
  /* 字母间距 */
  --letter-spacing-tight: -0.025em;
  --letter-spacing-normal: 0;
  --letter-spacing-wide: 0.025em;
  --letter-spacing-wider: 0.05em;
}

/* 浅色主题类 */
.light-theme {
  /* 浅色主题变量已经在:root中定义 */
}

/* 深色主题变量 */
.dark-theme {
  --color-primary: #7c3aed;
  --color-success: #10b981;
  --color-warning: #f59e0b;
  --color-error: #ef4444;
  --color-info: #3b82f6;
  
  --color-bg-container: #1a1a1a;
  --color-bg-elevated: #262626;
  --color-bg-layout: #0f0f0f;
  --color-bg-mask: rgba(0, 0, 0, 0.8);
  
  --color-text: #f8fafc;
  --color-text-secondary: #cbd5e1;
  --color-text-tertiary: #94a3b8;
  --color-text-quaternary: #64748b;
  
  --color-border: #374151;
  --color-border-secondary: #4b5563;
  --color-split: #374151;
  
  --box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
  --box-shadow-secondary: 0 4px 6px rgba(0, 0, 0, 0.4);
  --box-shadow-tertiary: 0 10px 15px rgba(0, 0, 0, 0.5);
}

/* 全局样式 */
#app {
  height: 100vh;
  font-family: var(--font-family-primary);
  font-size: var(--font-size-base);
  line-height: var(--line-height-normal);
  font-weight: var(--font-weight-normal);
  letter-spacing: var(--letter-spacing-normal);
  background: var(--color-bg-layout);
  min-height: 100vh;
  color: var(--color-text);
  transition: background-color 0.2s ease, color 0.2s ease;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
}

body {
  margin: 0;
  padding: 0;
  font-family: var(--font-family-primary);
  font-size: var(--font-size-base);
  line-height: var(--line-height-normal);
  font-weight: var(--font-weight-normal);
  letter-spacing: var(--letter-spacing-normal);
  background: var(--color-bg-layout);
  color: var(--color-text);
  transition: background-color 0.2s ease, color 0.2s ease;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
}

/* 确保主题类正确应用 */
.light-theme,
.dark-theme {
  background: var(--color-bg-layout);
  color: var(--color-text);
  transition: background-color 0.2s ease, color 0.2s ease;
}

/* 字体美化样式类 */
.font-display {
  font-family: var(--font-family-primary);
  font-weight: var(--font-weight-bold);
  letter-spacing: var(--letter-spacing-tight);
  line-height: var(--line-height-tight);
}

.font-heading {
  font-family: var(--font-family-primary);
  font-weight: var(--font-weight-semibold);
  letter-spacing: var(--letter-spacing-normal);
  line-height: var(--line-height-tight);
}

.font-body {
  font-family: var(--font-family-primary);
  font-weight: var(--font-weight-normal);
  letter-spacing: var(--letter-spacing-normal);
  line-height: var(--line-height-normal);
}

.font-caption {
  font-family: var(--font-family-primary);
  font-weight: var(--font-weight-normal);
  letter-spacing: var(--letter-spacing-wide);
  line-height: var(--line-height-normal);
}

.font-mono {
  font-family: var(--font-family-mono);
  font-weight: var(--font-weight-normal);
  letter-spacing: var(--letter-spacing-normal);
  line-height: var(--line-height-normal);
}

/* 标题样式美化 */
h1, .h1 {
  font-size: var(--font-size-4xl);
  font-weight: var(--font-weight-bold);
  line-height: var(--line-height-tight);
  letter-spacing: var(--letter-spacing-tight);
  margin: 0 0 1rem 0;
}

h2, .h2 {
  font-size: var(--font-size-3xl);
  font-weight: var(--font-weight-semibold);
  line-height: var(--line-height-tight);
  letter-spacing: var(--letter-spacing-normal);
  margin: 0 0 0.875rem 0;
}

h3, .h3 {
  font-size: var(--font-size-2xl);
  font-weight: var(--font-weight-semibold);
  line-height: var(--line-height-normal);
  letter-spacing: var(--letter-spacing-normal);
  margin: 0 0 0.75rem 0;
}

h4, .h4 {
  font-size: var(--font-size-xl);
  font-weight: var(--font-weight-medium);
  line-height: var(--line-height-normal);
  letter-spacing: var(--letter-spacing-normal);
  margin: 0 0 0.625rem 0;
}

h5, .h5 {
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-medium);
  line-height: var(--line-height-normal);
  letter-spacing: var(--letter-spacing-normal);
  margin: 0 0 0.5rem 0;
}

h6, .h6 {
  font-size: var(--font-size-base);
  font-weight: var(--font-weight-medium);
  line-height: var(--line-height-normal);
  letter-spacing: var(--letter-spacing-wide);
  margin: 0 0 0.5rem 0;
}

/* 段落和文本样式 */
p {
  font-size: var(--font-size-base);
  line-height: var(--line-height-relaxed);
  margin: 0 0 1rem 0;
}

.text-xs { font-size: var(--font-size-xs); }
.text-sm { font-size: var(--font-size-sm); }
.text-base { font-size: var(--font-size-base); }
.text-lg { font-size: var(--font-size-lg); }
.text-xl { font-size: var(--font-size-xl); }
.text-2xl { font-size: var(--font-size-2xl); }
.text-3xl { font-size: var(--font-size-3xl); }
.text-4xl { font-size: var(--font-size-4xl); }
.text-5xl { font-size: var(--font-size-5xl); }

.font-light { font-weight: var(--font-weight-light); }
.font-normal { font-weight: var(--font-weight-normal); }
.font-medium { font-weight: var(--font-weight-medium); }
.font-semibold { font-weight: var(--font-weight-semibold); }
.font-bold { font-weight: var(--font-weight-bold); }
.font-extrabold { font-weight: var(--font-weight-extrabold); }

.leading-tight { line-height: var(--line-height-tight); }
.leading-normal { line-height: var(--line-height-normal); }
.leading-relaxed { line-height: var(--line-height-relaxed); }
.leading-loose { line-height: var(--line-height-loose); }

.tracking-tight { letter-spacing: var(--letter-spacing-tight); }
.tracking-normal { letter-spacing: var(--letter-spacing-normal); }
.tracking-wide { letter-spacing: var(--letter-spacing-wide); }
.tracking-wider { letter-spacing: var(--letter-spacing-wider); }

/* 深色主题下的Ant Design组件样式 */
.dark-theme .ant-card-head-title {
  color: var(--color-text) !important;
}

.dark-theme .ant-card-body {
  color: var(--color-text);
}

.dark-theme .ant-statistic-title {
  color: var(--color-text-secondary) !important;
}

.dark-theme .ant-statistic-content {
  color: var(--color-text) !important;
}

.dark-theme .ant-tag {
  color: var(--color-text) !important;
  border: 1px solid var(--color-border) !important;
}

/* 深色主题下的特定颜色标签 */
.dark-theme .ant-tag-purple {
  background: #7c3aed !important;
  color: white !important;
  border-color: #7c3aed !important;
}

.dark-theme .ant-tag-blue {
  background: #3b82f6 !important;
  color: white !important;
  border-color: #3b82f6 !important;
}

.dark-theme .ant-tag-green {
  background: #10b981 !important;
  color: white !important;
  border-color: #10b981 !important;
}

.dark-theme .ant-tag-orange {
  background: #f59e0b !important;
  color: white !important;
  border-color: #f59e0b !important;
}

.dark-theme .ant-tag-red {
  background: #ef4444 !important;
  color: white !important;
  border-color: #ef4444 !important;
}

/* 深色主题下所有彩色标签的通用样式 */
.dark-theme .ant-tag[class*="ant-tag-"] {
  font-weight: 500 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

/* 确保标签文字在深色背景下清晰可见 */
.dark-theme .ant-tag:not(.ant-tag-default) {
  color: white !important;
  font-weight: 600 !important;
}

/* 深色主题下的导航栏样式 */
.dark-theme .ant-layout-header {
  background: var(--color-bg-container) !important;
  border-bottom: 1px solid var(--color-border) !important;
}

.dark-theme .ant-layout-sider {
  background: var(--color-bg-container) !important;
}

.dark-theme .ant-menu {
  background: var(--color-bg-container) !important;
  color: var(--color-text) !important;
}

.dark-theme .ant-menu-item {
  color: var(--color-text) !important;
}

.dark-theme .ant-menu-item:hover {
  background: var(--color-bg-elevated) !important;
  color: var(--color-text) !important;
}

.dark-theme .ant-menu-item-selected {
  background: var(--color-primary) !important;
  color: white !important;
}

.dark-theme .ant-breadcrumb {
  color: var(--color-text) !important;
}

.dark-theme .ant-breadcrumb-link {
  color: var(--color-text) !important;
}

.dark-theme .ant-breadcrumb-separator {
  color: var(--color-text-secondary) !important;
}

/* 用户信息样式 */
.dark-theme .user-info {
  color: var(--color-text) !important;
}

.dark-theme .user-role {
  color: var(--color-text-secondary) !important;
}

.dark-theme .user-role.regulator {
  color: var(--color-primary) !important;
  font-weight: 600;
}

.dark-theme .user-role.project-owner {
  color: var(--color-info) !important;
  font-weight: 600;
}

.dark-theme .user-role.verifier {
  color: var(--color-warning) !important;
  font-weight: 600;
}

.dark-theme .user-role.buyer {
  color: var(--color-success) !important;
  font-weight: 600;
}

/* 深色主题下的页面标题和内容区域 */
.dark-theme .ant-page-header {
  background: var(--color-bg-container) !important;
  border-bottom: 1px solid var(--color-border) !important;
}

.dark-theme .ant-page-header-heading-title {
  color: var(--color-text) !important;
}

.dark-theme .ant-page-header-content {
  color: var(--color-text) !important;
}

/* 深色主题下的内容区域 */
.dark-theme .ant-layout-content {
  background: var(--color-bg-layout) !important;
  color: var(--color-text) !important;
}

/* 深色主题下的页面容器 */
.dark-theme .page-container {
  background: var(--color-bg-layout) !important;
  color: var(--color-text) !important;
}

.dark-theme .page-header {
  background: var(--color-bg-container) !important;
  border-bottom: 1px solid var(--color-border) !important;
}

.dark-theme .page-title {
  color: var(--color-text) !important;
}

.dark-theme .page-content {
  background: var(--color-bg-layout) !important;
  color: var(--color-text) !important;
}

/* 深色主题下的通用页面样式 */
.dark-theme .user-management,
.dark-theme .user-profile,
.dark-theme .project-overview,
.dark-theme .credit-market,
.dark-theme .certificate-overview {
  background: var(--color-bg-layout) !important;
  color: var(--color-text) !important;
}

.dark-theme .user-management .page-header,
.dark-theme .user-profile .page-header,
.dark-theme .project-overview .page-header,
.dark-theme .credit-market .page-header,
.dark-theme .certificate-overview .page-header {
  background: var(--color-bg-container) !important;
  border-bottom: 1px solid var(--color-border) !important;
  color: var(--color-text) !important;
}

.dark-theme .user-management .page-header h1,
.dark-theme .user-profile .page-header h1,
.dark-theme .project-overview .page-header h1,
.dark-theme .credit-market .page-header h1,
.dark-theme .certificate-overview .page-header h1 {
  color: var(--color-text) !important;
}

.dark-theme .user-management .page-header p,
.dark-theme .user-profile .page-header p,
.dark-theme .project-overview .page-header p,
.dark-theme .credit-market .page-header p,
.dark-theme .certificate-overview .page-header p {
  color: var(--color-text-secondary) !important;
}

/* 深色主题下的更多Ant Design组件 */
.dark-theme .ant-result-title {
  color: var(--color-text) !important;
}

.dark-theme .ant-result-subtitle {
  color: var(--color-text-secondary) !important;
}

.dark-theme .ant-descriptions-title {
  color: var(--color-text) !important;
}

.dark-theme .ant-descriptions-item-label {
  color: var(--color-text-secondary) !important;
}

.dark-theme .ant-descriptions-item-content {
  color: var(--color-text) !important;
}

.dark-theme .ant-typography {
  color: var(--color-text) !important;
}

.dark-theme .ant-typography-text {
  color: var(--color-text) !important;
}

.dark-theme .ant-typography-paragraph {
  color: var(--color-text) !important;
}

.dark-theme .ant-typography-title {
  color: var(--color-text) !important;
}
</style>
