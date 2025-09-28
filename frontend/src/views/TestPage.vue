<template>
  <div class="test-page">
    <h1>测试页面</h1>
    <p>如果您能看到这个页面，说明基本渲染功能正常。</p>
    <div class="test-info">
      <p><strong>当前时间:</strong> {{ currentTime }}</p>
      <p><strong>Vue版本:</strong> {{ vueVersion }}</p>
      <p><strong>主题模式:</strong> {{ themeMode }}</p>
    </div>
    <a-button type="primary" @click="testClick">测试按钮</a-button>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTheme } from '@/composables/useTheme'

const currentTime = ref('')
const vueVersion = ref('3.x')
const { themeMode } = useTheme()

const testClick = () => {
  alert('按钮点击正常！')
}

onMounted(() => {
  currentTime.value = new Date().toLocaleString()
})
</script>

<style scoped>
.test-page {
  padding: 40px;
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
}

.test-page h1 {
  color: var(--color-primary, #0969da);
  margin-bottom: 20px;
}

.test-info {
  background: var(--color-bg-elevated, #f6f8fa);
  border: 1px solid var(--color-border, #d0d7de);
  border-radius: 8px;
  padding: 20px;
  margin: 20px 0;
  text-align: left;
}

.test-info p {
  margin: 10px 0;
  color: var(--color-text, #24292f);
}
</style>
