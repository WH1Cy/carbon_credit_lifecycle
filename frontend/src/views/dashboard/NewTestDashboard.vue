<template>
  <div class="new-test-dashboard">
    <div class="header">
      <h1>🚀 新版本Dashboard - 带钱包按钮</h1>
      <div class="header-actions">
        <a-button 
          :type="walletConnected ? 'default' : 'primary'" 
          size="large" 
          @click="testWallet"
        >
          {{ walletConnected ? `💰 ${walletAddress.slice(0, 6)}...${walletAddress.slice(-4)}` : '💰 连接钱包' }}
        </a-button>
      </div>
    </div>
    
    <div class="content">
      <h2>🎉 如果您看到这个页面，说明路由更新成功！</h2>
      <p>这个版本包含了钱包连接按钮</p>
      
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
        <a-button @click="goToTest">测试页面</a-button>
      </div>
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

// 钱包状态
const walletConnected = ref(false)
const walletAddress = ref('')

const testClick = () => {
  alert('按钮点击正常！')
}

const testWallet = async () => {
  try {
    // 导入Web3服务
    const { connectWallet, walletInfo } = await import('@/services/web3')
    
    if (walletInfo.isConnected) {
      walletConnected.value = true
      walletAddress.value = walletInfo.address
      alert(`钱包已连接: ${walletInfo.address}`)
    } else {
      await connectWallet()
      walletConnected.value = true
      walletAddress.value = walletInfo.address
      alert(`钱包连接成功: ${walletInfo.address}`)
    }
  } catch (error) {
    console.error('钱包连接失败:', error)
    alert('钱包连接失败，请检查MetaMask是否已安装')
  }
}

const handleToggleTheme = () => {
  toggleTheme()
  alert(`主题已切换到: ${themeMode.value}`)
}

const goToTest = () => {
  router.push('/test')
}

onMounted(async () => {
  currentTime.value = new Date().toLocaleString()
  
  // 检查钱包连接状态
  try {
    const { walletInfo } = await import('@/services/web3')
    if (walletInfo.isConnected && walletInfo.address) {
      walletConnected.value = true
      walletAddress.value = walletInfo.address
    }
  } catch (error) {
    console.log('钱包服务初始化失败:', error)
  }
})
</script>

<style scoped>
.new-test-dashboard {
  min-height: 100vh;
  background: var(--color-bg-layout);
}

.header {
  background: var(--color-bg-container);
  border-bottom: 1px solid var(--color-border);
  padding: 16px 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: var(--box-shadow);
}

.header h1 {
  margin: 0;
  color: var(--color-primary);
  font-size: 24px;
  font-weight: 600;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.content {
  padding: 40px;
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.content h2 {
  color: var(--color-text);
  font-size: 32px;
  margin-bottom: 16px;
}

.content p {
  color: var(--color-text-secondary);
  font-size: 18px;
  margin-bottom: 40px;
}

.test-sections {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 24px;
  margin: 40px 0;
}

.test-section {
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  padding: 24px;
  box-shadow: var(--box-shadow);
  transition: all 0.2s ease;
}

.test-section:hover {
  box-shadow: var(--box-shadow-secondary);
  border-color: var(--color-border-secondary);
  transform: translateY(-2px);
}

.test-section h3 {
  color: var(--color-text);
  margin-bottom: 16px;
  font-size: 18px;
  font-weight: 600;
}

.test-section p {
  color: var(--color-text-secondary);
  margin: 8px 0;
  line-height: 1.5;
}

.navigation {
  text-align: center;
  margin-top: 40px;
  display: flex;
  gap: 16px;
  justify-content: center;
}
</style>
