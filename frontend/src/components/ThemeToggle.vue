<template>
  <div class="theme-toggle">
    <a-tooltip :title="isDark ? '切换到浅色模式' : '切换到深色模式'">
      <a-button
        type="text"
        size="large"
        class="theme-toggle-btn"
        @click="handleToggleTheme"
        :loading="isTransitioning"
      >
        <template #icon>
          <transition name="theme-icon" mode="out-in">
            <span v-if="isDark" key="sun">☀️</span>
            <span v-else key="moon">🌙</span>
          </transition>
        </template>
      </a-button>
    </a-tooltip>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTheme } from '@/composables/useTheme'

// 使用主题系统
const { isDark, toggleTheme } = useTheme()

// 过渡状态
const isTransitioning = ref(false)

// 处理主题切换
const handleToggleTheme = () => {
  isTransitioning.value = true
  toggleTheme()
  
  // 短暂延迟后重置过渡状态
  setTimeout(() => {
    isTransitioning.value = false
  }, 200)
}
</script>

<style scoped>
.theme-toggle {
  display: flex;
  align-items: center;
  justify-content: center;
}

.theme-toggle-btn {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  border: 1px solid var(--color-border);
  background: var(--color-bg-container);
  color: var(--color-text);
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.theme-toggle-btn:hover {
  background: var(--color-bg-elevated);
  border-color: var(--color-border-secondary);
  transform: translateY(-1px);
  box-shadow: var(--box-shadow);
}

.theme-toggle-btn:active {
  transform: translateY(0);
}

/* 主题图标过渡动画 */
.theme-icon-enter-active,
.theme-icon-leave-active {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.theme-icon-enter-from {
  opacity: 0;
  transform: rotate(-90deg) scale(0.8);
}

.theme-icon-leave-to {
  opacity: 0;
  transform: rotate(90deg) scale(0.8);
}

.theme-icon-enter-to,
.theme-icon-leave-from {
  opacity: 1;
  transform: rotate(0deg) scale(1);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .theme-toggle-btn {
    width: 36px;
    height: 36px;
  }
}
</style>
