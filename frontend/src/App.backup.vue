<template>
  <a-config-provider :theme="antdThemeConfig">
    <RouterView />
  </a-config-provider>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
import { useTheme } from '@/composables/useTheme'

// 使用主题系统
const { antdThemeConfig, initTheme } = useTheme()

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
}

/* 深色主题变量 */
.dark-theme {
  --color-primary: #58a6ff;
  --color-success: #3fb950;
  --color-warning: #d29922;
  --color-error: #f85149;
  --color-info: #58a6ff;
  
  --color-bg-container: #0d1117;
  --color-bg-elevated: #161b22;
  --color-bg-layout: #0d1117;
  --color-bg-mask: rgba(0, 0, 0, 0.8);
  
  --color-text: #f0f6fc;
  --color-text-secondary: #8b949e;
  --color-text-tertiary: #6e7681;
  --color-text-quaternary: #30363d;
  
  --color-border: #30363d;
  --color-border-secondary: #21262d;
  --color-split: #30363d;
  
  --box-shadow: 0 1px 0 rgba(0, 0, 0, 0.2);
  --box-shadow-secondary: 0 3px 6px rgba(0, 0, 0, 0.3);
  --box-shadow-tertiary: 0 8px 24px rgba(0, 0, 0, 0.4);
}

/* 全局样式 */
#app {
  height: 100vh;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB',
    'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif, 'Apple Color Emoji',
    'Segoe UI Emoji', 'Segoe UI Symbol';
  background: var(--color-bg-layout);
  min-height: 100vh;
  color: var(--color-text);
  transition: background-color 0.2s ease, color 0.2s ease;
}

body {
  margin: 0;
  padding: 0;
  background: var(--color-bg-layout);
  color: var(--color-text);
  transition: background-color 0.2s ease, color 0.2s ease;
}

/* 主题化滚动条 */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: var(--color-bg-elevated);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: var(--color-border);
  border-radius: 4px;
  border: 2px solid var(--color-bg-elevated);
}

::-webkit-scrollbar-thumb:hover {
  background: var(--color-text-tertiary);
}

/* GitHub风格动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.15s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.15s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-up-enter-from {
  transform: translateY(8px);
  opacity: 0;
}

.slide-up-leave-to {
  transform: translateY(-8px);
  opacity: 0;
}

/* 主题化卡片悬停效果 */
.card-hover {
  transition: all 0.15s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid var(--color-border);
}

.card-hover:hover {
  transform: translateY(-1px);
  box-shadow: var(--box-shadow-secondary) !important;
  border-color: var(--color-border-secondary);
}

/* 主题化按钮悬停效果 */
.btn-hover {
  transition: all 0.1s cubic-bezier(0.4, 0, 0.2, 1);
}

.btn-hover:hover {
  transform: translateY(-1px);
  box-shadow: var(--box-shadow);
}

/* 主题化链接 */
a {
  color: var(--color-primary);
  text-decoration: none;
  transition: color 0.1s cubic-bezier(0.4, 0, 0.2, 1);
}

a:hover {
  color: var(--color-primary);
  opacity: 0.8;
  text-decoration: underline;
}

/* 主题化代码块 */
code {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border);
  border-radius: 4px;
  padding: 2px 4px;
  font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
  font-size: 12px;
  color: var(--color-text);
}

pre {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border);
  border-radius: 6px;
  padding: 16px;
  overflow-x: auto;
  font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
  font-size: 12px;
  line-height: 1.45;
  color: var(--color-text);
}

/* 主题化分割线 */
hr {
  border: none;
  height: 1px;
  background: var(--color-border);
  margin: 24px 0;
}

/* 主题化标签 */
.github-tag {
  display: inline-block;
  padding: 2px 6px;
  font-size: 11px;
  font-weight: 500;
  line-height: 1.2;
  border-radius: 4px;
  background: var(--color-bg-elevated);
  color: var(--color-text-secondary);
  border: 1px solid var(--color-border);
}

.github-tag.primary {
  background: var(--color-primary);
  color: #ffffff;
  border-color: var(--color-primary);
}

.github-tag.success {
  background: var(--color-success);
  color: #ffffff;
  border-color: var(--color-success);
}

.github-tag.warning {
  background: var(--color-warning);
  color: #ffffff;
  border-color: var(--color-warning);
}

.github-tag.error {
  background: var(--color-error);
  color: #ffffff;
  border-color: var(--color-error);
}

/* 主题化徽章 */
.github-badge {
  display: inline-flex;
  align-items: center;
  padding: 2px 6px;
  font-size: 11px;
  font-weight: 500;
  line-height: 1.2;
  border-radius: 10px;
  background: var(--color-text-secondary);
  color: #ffffff;
}

.github-badge.primary {
  background: var(--color-primary);
}

.github-badge.success {
  background: var(--color-success);
}

.github-badge.warning {
  background: var(--color-warning);
}

.github-badge.error {
  background: var(--color-error);
}

/* 主题化输入框 */
.github-input {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 6px;
  padding: 8px 12px;
  font-size: 14px;
  line-height: 1.5;
  color: var(--color-text);
  transition: border-color 0.1s cubic-bezier(0.4, 0, 0.2, 1);
}

.github-input:focus {
  border-color: var(--color-primary);
  outline: none;
  box-shadow: 0 0 0 3px rgba(9, 105, 218, 0.1);
}

.github-input::placeholder {
  color: var(--color-text-tertiary);
}

/* 主题化按钮 */
.github-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 6px 16px;
  font-size: 14px;
  font-weight: 500;
  line-height: 1.5;
  border-radius: 6px;
  border: 1px solid var(--color-border);
  background: var(--color-bg-container);
  color: var(--color-text);
  cursor: pointer;
  transition: all 0.1s cubic-bezier(0.4, 0, 0.2, 1);
  text-decoration: none;
}

.github-btn:hover {
  background: var(--color-bg-elevated);
  border-color: var(--color-border-secondary);
  transform: translateY(-1px);
  box-shadow: var(--box-shadow);
}

.github-btn.primary {
  background: var(--color-primary);
  border-color: var(--color-primary);
  color: #ffffff;
}

.github-btn.primary:hover {
  background: var(--color-primary);
  opacity: 0.9;
  border-color: var(--color-primary);
}

.github-btn.success {
  background: var(--color-success);
  border-color: var(--color-success);
  color: #ffffff;
}

.github-btn.success:hover {
  background: var(--color-success);
  opacity: 0.9;
  border-color: var(--color-success);
}

/* 主题化卡片 */
.github-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  padding: 24px;
  box-shadow: var(--box-shadow);
}

.github-card:hover {
  border-color: var(--color-border-secondary);
  box-shadow: var(--box-shadow-secondary);
}

/* 主题化表格 */
.github-table {
  width: 100%;
  border-collapse: collapse;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  overflow: hidden;
}

.github-table th {
  background: var(--color-bg-elevated);
  border-bottom: 1px solid var(--color-border);
  padding: 12px 16px;
  text-align: left;
  font-weight: 600;
  color: var(--color-text);
  font-size: 14px;
}

.github-table td {
  border-bottom: 1px solid var(--color-border);
  padding: 12px 16px;
  color: var(--color-text);
  font-size: 14px;
}

.github-table tr:hover td {
  background: var(--color-bg-elevated);
}

.github-table tr:last-child td {
  border-bottom: none;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .github-card {
    padding: 16px;
  }
  
  .github-table th,
  .github-table td {
    padding: 8px 12px;
    font-size: 13px;
  }
}
</style>
