<template>
  <div class="enhanced-dashboard">
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="main-title font-display">
            <span class="title-icon">🌱</span>
            碳信通平台
          </h1>
          <p class="subtitle font-body text-lg">绿色未来，从碳信用开始</p>
        </div>
        <div class="header-actions">
          <div class="action-buttons">
            <a-button 
              v-if="!walletConnected" 
              type="primary" 
              size="large"
              class="btn-hover"
              @click="handleConnectWallet"
              :loading="isConnecting"
            >
              💰 连接钱包
            </a-button>
            <a-button 
              v-else
              type="default" 
              size="large"
              class="btn-hover"
              @click="handleConnectWallet"
            >
              💰 {{ walletAddress.slice(0, 6) }}...{{ walletAddress.slice(-4) }}
            </a-button>
          </div>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <a-row :gutter="[24, 24]" class="stats-cards">
      <a-col :xs="24" :sm="12" :lg="6">
        <a-card 
          :loading="loading" 
          class="stat-card stat-card-users card-hover"
          :bordered="false"
        >
          <div class="stat-content">
            <div class="stat-icon users-icon">
              👥
            </div>
            <div class="stat-info">
              <a-statistic
                title="注册用户数"
                :value="summary?.totalUsers || 0"
                :value-style="{ color: 'var(--color-success)', fontSize: '28px', fontWeight: '600' }"
              />
              <div class="stat-trend">
                <span class="trend-text">活跃用户</span>
              </div>
            </div>
          </div>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="12" :lg="6">
        <a-card 
          :loading="loading" 
          class="stat-card stat-card-projects card-hover"
          :bordered="false"
        >
          <div class="stat-content">
            <div class="stat-icon projects-icon">
              🌳
            </div>
            <div class="stat-info">
              <a-statistic
                title="项目总数"
                :value="summary?.totalProjects || 0"
                :value-style="{ color: 'var(--color-info)', fontSize: '28px', fontWeight: '600' }"
              />
              <div class="stat-trend">
                <span class="trend-text">绿色项目</span>
              </div>
            </div>
          </div>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="12" :lg="6">
        <a-card 
          :loading="loading" 
          class="stat-card stat-card-credits card-hover"
          :bordered="false"
        >
          <div class="stat-content">
            <div class="stat-icon credits-icon">
              💚
            </div>
            <div class="stat-info">
              <a-statistic
                title="碳信用总量"
                :value="summary?.totalCredits || 0"
                :value-style="{ color: 'var(--color-primary)', fontSize: '28px', fontWeight: '600' }"
              />
              <div class="stat-trend">
                <span class="trend-text">已认证</span>
              </div>
            </div>
          </div>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="12" :lg="6">
        <a-card 
          :loading="loading" 
          class="stat-card stat-card-transactions card-hover"
          :bordered="false"
        >
          <div class="stat-content">
            <div class="stat-icon transactions-icon">
              💰
            </div>
            <div class="stat-info">
              <a-statistic
                title="交易总额"
                :value="summary?.totalTransactions || 0"
                :value-style="{ color: 'var(--color-warning)', fontSize: '28px', fontWeight: '600' }"
              />
              <div class="stat-trend">
                <span class="trend-text">累计交易</span>
              </div>
            </div>
          </div>
        </a-card>
      </a-col>
    </a-row>

    <!-- 内容区域 -->
    <a-row :gutter="[24, 24]" class="content-section">
      <!-- 快速操作 -->
      <a-col :xs="24" :lg="12">
        <a-card title="快速操作" class="content-card card-hover" :bordered="false">
          <div class="quick-actions">
            <a-button 
              type="primary" 
              size="large" 
              class="action-btn"
              @click="goToProjects"
            >
              🌳 项目管理
            </a-button>
            <a-button 
              type="default" 
              size="large" 
              class="action-btn"
              @click="goToCredits"
            >
              💚 碳信用市场
            </a-button>
            <a-button 
              type="default" 
              size="large" 
              class="action-btn"
              @click="goToCertificates"
            >
              📜 证书管理
            </a-button>
            <a-button 
              type="default" 
              size="large" 
              class="action-btn"
              @click="goToProfile"
            >
              👤 个人资料
            </a-button>
          </div>
        </a-card>
      </a-col>
      
      <!-- 系统状态 -->
      <a-col :xs="24" :lg="12">
        <a-card title="系统状态" class="content-card card-hover" :bordered="false">
          <div class="system-status">
            <div class="status-item">
              <span class="status-label">区块链网络:</span>
              <a-tag color="green">Sepolia Testnet</a-tag>
            </div>
            <div class="status-item">
              <span class="status-label">钱包状态:</span>
              <a-tag :color="walletConnected ? 'green' : 'red'">
                {{ walletConnected ? '已连接' : '未连接' }}
              </a-tag>
            </div>
            <div class="status-item">
              <span class="status-label">当前时间:</span>
              <span class="status-value">{{ currentTime }}</span>
            </div>
            <div class="status-item">
              <span class="status-label">主题模式:</span>
              <a-tag :color="isDark ? 'blue' : 'orange'">
                {{ isDark ? '深色模式' : '浅色模式' }}
              </a-tag>
            </div>
          </div>
        </a-card>
      </a-col>
    </a-row>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from '@/composables/useTheme'
import { message } from 'ant-design-vue'

const router = useRouter()
const currentTime = ref('')
const { themeMode, toggleTheme, isDark } = useTheme()

// 钱包状态
const walletConnected = ref(false)
const walletAddress = ref('')
const isConnecting = ref(false)

// 数据状态
const loading = ref(false)
const summary = ref({
  totalUsers: 156,
  totalProjects: 23,
  totalCredits: 1250,
  totalTransactions: 89
})

// 主题切换功能
const handleToggleTheme = () => {
  toggleTheme()
  message.info(`主题已切换到: ${themeMode.value}`)
}

// 钱包功能
const handleConnectWallet = async () => {
  if (walletConnected.value) {
    message.info(`钱包已连接: ${walletAddress.value}`)
    return
  }

  isConnecting.value = true
  try {
    const { connectWallet, walletInfo } = await import('@/services/web3')
    await connectWallet()
    walletConnected.value = true
    walletAddress.value = walletInfo.address
    message.success(`钱包连接成功: ${walletInfo.address}`)
  } catch (error) {
    console.error('钱包连接失败:', error)
    message.error('钱包连接失败，请检查MetaMask是否已安装')
  } finally {
    isConnecting.value = false
  }
}

// 导航功能
const goToProjects = () => {
  router.push('/project/overview')
}

const goToCredits = () => {
  router.push('/credit/market')
}

const goToCertificates = () => {
  router.push('/certificate/overview')
}

const goToProfile = () => {
  router.push('/user/profile')
}


// 初始化
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
.enhanced-dashboard {
  padding: 24px;
  background: var(--color-bg-layout);
  min-height: 100vh;
  color: var(--color-text);
  transition: background-color 0.2s ease, color 0.2s ease;
}

/* 页面标题 */
.page-header {
  margin-bottom: 32px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  padding: 24px;
  box-shadow: var(--box-shadow);
  transition: all 0.2s ease;
}

.title-section {
  flex: 1;
}

.main-title {
  font-size: 32px;
  font-weight: 700;
  color: var(--color-text);
  margin: 0 0 8px 0;
  display: flex;
  align-items: center;
  gap: 12px;
}

.title-icon {
  font-size: 36px;
}

.subtitle {
  font-size: 16px;
  color: var(--color-text-secondary);
  margin: 0;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.action-buttons {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* 统计卡片 */
.stats-cards {
  margin-bottom: 32px;
}

.stat-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  box-shadow: var(--box-shadow);
  transition: all 0.2s ease;
  position: relative;
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: var(--color-primary);
  transition: all 0.2s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--box-shadow-secondary);
  border-color: var(--color-border-secondary);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 8px 0;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border);
  transition: all 0.2s ease;
}

.stat-info {
  flex: 1;
}

.stat-trend {
  margin-top: 4px;
}

.trend-text {
  font-size: 12px;
  color: var(--color-text-tertiary);
}

/* 内容区域 */
.content-section {
  margin-bottom: 32px;
}

.content-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  box-shadow: var(--box-shadow);
  transition: all 0.2s ease;
}

.content-card :deep(.ant-card-head-title) {
  color: var(--color-text) !important;
}

.content-card :deep(.ant-card-body) {
  color: var(--color-text);
}

.content-card:hover {
  transform: translateY(-1px);
  box-shadow: var(--box-shadow-secondary);
  border-color: var(--color-border-secondary);
}

/* 快速操作 */
.quick-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.action-btn {
  height: 48px;
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.action-btn:hover {
  transform: translateY(-1px);
  box-shadow: var(--box-shadow);
}

/* 系统状态 */
.system-status {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid var(--color-border);
}

.status-item:last-child {
  border-bottom: none;
}

.status-label {
  font-weight: 500;
  color: var(--color-text);
}

.status-value {
  color: var(--color-text-secondary);
}


/* 响应式设计 */
@media (max-width: 768px) {
  .enhanced-dashboard {
    padding: 16px;
  }
  
  .header-content {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
  
  .main-title {
    font-size: 24px;
  }
  
  .quick-actions {
    grid-template-columns: 1fr;
  }
}
</style>
