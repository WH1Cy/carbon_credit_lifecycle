<template>
  <div class="test-dashboard">
    <h1>Dashboard - 测试版本</h1>
    <p>如果您能看到这个页面，说明基础功能正常。</p>
    
    <div class="test-sections">
      <div class="test-section">
        <h3>基础测试</h3>
        <p>当前时间: {{ currentTime }}</p>
        <a-button type="primary" @click="testClick">测试按钮</a-button>
      </div>
      
      <div class="test-section">
        <h3>主题测试</h3>
        <p>主题模式: {{ themeMode }}</p>
        <a-button @click="handleToggleTheme">切换主题</a-button>
      </div>
    </div>
    
    <div class="navigation">
      <a-button @click="goToTest">返回测试页面</a-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from '@/composables/useTheme'

const router = useRouter()
const currentTime = ref('')
const { themeMode, toggleTheme } = useTheme()

const testClick = () => {
  alert('按钮点击正常！')
}

const handleToggleTheme = () => {
  toggleTheme()
  alert(`主题已切换到: ${themeMode.value}`)
}

const goToTest = () => {
  router.push('/test')
}

onMounted(() => {
  currentTime.value = new Date().toLocaleString()
})
</script>

<style scoped>
.test-dashboard {
  padding: 40px;
  max-width: 1000px;
  margin: 0 auto;
  background: var(--color-bg-layout);
  min-height: 100vh;
}

.test-dashboard h1 {
  color: var(--color-primary);
  margin-bottom: 20px;
  text-align: center;
}

.test-sections {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  margin: 30px 0;
}

.test-section {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  padding: 20px;
  box-shadow: var(--box-shadow);
  transition: all 0.2s ease;
}

.test-section:hover {
  box-shadow: var(--box-shadow-secondary);
  border-color: var(--color-border-secondary);
}

.test-section h3 {
  color: var(--color-text);
  margin-bottom: 15px;
}

.test-section p {
  color: var(--color-text-secondary);
  margin: 10px 0;
}

.navigation {
  text-align: center;
  margin-top: 30px;
}
</style>
