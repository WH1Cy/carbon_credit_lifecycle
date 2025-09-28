<template>
  <div class="dashboard">
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="main-title">
            <span class="title-icon">🌱</span>
            碳信通平台
          </h1>
          <p class="subtitle">绿色未来，从碳信用开始</p>
        </div>
        <div class="header-actions">
          <div class="action-buttons">
            <ThemeToggle />
            <a-button 
              v-if="!walletInfo.isConnected" 
              type="primary" 
              size="large"
              class="btn-hover"
              @click="handleConnectWallet"
              :loading="isConnecting"
            >
              <WalletOutlined />
              连接钱包
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
              <UserOutlined />
            </div>
            <div class="stat-info">
              <a-statistic
                title="注册用户数"
                :value="summary?.totalUsers || 0"
                :value-style="{ color: '#52c41a', fontSize: '28px', fontWeight: '600' }"
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
              <ProjectOutlined />
            </div>
            <div class="stat-info">
              <a-statistic
                title="项目总数"
                :value="summary?.totalProjects || 0"
                :value-style="{ color: '#1890ff', fontSize: '28px', fontWeight: '600' }"
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
              <GoldOutlined />
            </div>
            <div class="stat-info">
              <a-statistic
                title="碳信用批次"
                :value="summary?.totalCreditBatches || 0"
                :value-style="{ color: '#fa8c16', fontSize: '28px', fontWeight: '600' }"
              />
              <div class="stat-trend">
                <span class="trend-text">可交易</span>
              </div>
            </div>
          </div>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="12" :lg="6">
        <a-card 
          :loading="loading" 
          class="stat-card stat-card-certificates card-hover"
          :bordered="false"
        >
          <div class="stat-content">
            <div class="stat-icon certificates-icon">
              <SafetyCertificateOutlined />
            </div>
            <div class="stat-info">
              <a-statistic
                title="已签发证书"
                :value="summary?.totalCertificates || 0"
                :value-style="{ color: '#722ed1', fontSize: '28px', fontWeight: '600' }"
              />
              <div class="stat-trend">
                <span class="trend-text">权威认证</span>
              </div>
            </div>
          </div>
        </a-card>
      </a-col>
    </a-row>

    <!-- 内容区域 -->
    <a-row :gutter="[24, 24]" class="content-section">
      <!-- 最新项目 -->
      <a-col :xs="24" :lg="12">
        <a-card 
          class="content-card card-hover"
          :bordered="false"
        >
          <template #title>
            <div class="card-title">
              <div class="title-icon-wrapper">
                <ProjectOutlined class="title-icon" />
              </div>
              <span>最新项目</span>
            </div>
          </template>
          <template #extra>
            <router-link to="/project/overview" class="more-link">
              查看更多 <span class="arrow">→</span>
            </router-link>
          </template>
          
          <a-list
            :data-source="recentProjects"
            size="small"
            class="modern-list"
          >
            <template #renderItem="{ item }">
              <a-list-item class="list-item-modern">
                <a-list-item-meta>
                  <template #title>
                    <router-link :to="`/project/${item.id}`" class="project-link">
                      {{ item.name }}
                    </router-link>
                  </template>
                  <template #description>
                    <a-space>
                      <a-tag 
                        :color="getProjectStatusColor(item.status)"
                        class="status-tag"
                      >
                        {{ getProjectStatusText(item.status) }}
                      </a-tag>
                      <span class="location-text">{{ item.location }}</span>
                    </a-space>
                  </template>
                </a-list-item-meta>
                
                <template #actions>
                  <span class="date-text">
                    {{ formatDate(item.timestamp) }}
                  </span>
                </template>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
      </a-col>

      <!-- 最新碳信用 -->
      <a-col :xs="24" :lg="12">
        <a-card 
          class="content-card card-hover"
          :bordered="false"
        >
          <template #title>
            <div class="card-title">
              <div class="title-icon-wrapper">
                <GoldOutlined class="title-icon" />
              </div>
              <span>最新碳信用</span>
            </div>
          </template>
          <template #extra>
            <router-link to="/credit/market" class="more-link">
              查看更多 <span class="arrow">→</span>
            </router-link>
          </template>
          
          <a-list
            :data-source="recentCredits"
            size="small"
            class="modern-list"
          >
            <template #renderItem="{ item }">
              <a-list-item class="list-item-modern">
                <a-list-item-meta>
                  <template #title>
                    <router-link :to="`/credit/${item.id || '1'}`" class="project-link">
                      批次 #{{ item.id || getBatchIdFromHistory(item) }}
                    </router-link>
                  </template>
                  <template #description>
                    <a-space>
                      <a-tag 
                        :color="getCreditEventColor(item.eventType)"
                        class="status-tag"
                      >
                        {{ getCreditEventText(item.eventType) }}
                      </a-tag>
                      <span class="quantity-text">{{ item.quantity }} 吨CO₂e</span>
                    </a-space>
                  </template>
                </a-list-item-meta>
                
                <template #actions>
                  <span class="date-text">
                    {{ formatDate(item.timestamp) }}
                  </span>
                </template>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
      </a-col>
    </a-row>

    <!-- 最新证书 -->
    <a-row :gutter="[24, 24]" class="content-section">
      <a-col :xs="24">
        <a-card 
          class="content-card card-hover"
          :bordered="false"
        >
          <template #title>
            <div class="card-title">
              <div class="title-icon-wrapper">
                <SafetyCertificateOutlined class="title-icon" />
              </div>
              <span>最新核发证书</span>
            </div>
          </template>
          <template #extra>
            <router-link to="/certificate/overview" class="more-link">
              查看更多 <span class="arrow">→</span>
            </router-link>
          </template>
          
          <a-list
            :data-source="recentCertificates"
            size="small"
            class="modern-list"
          >
            <template #renderItem="{ item }">
              <a-list-item class="list-item-modern">
                <a-list-item-meta>
                  <template #title>
                    <router-link :to="`/certificate/${item.id}`" class="project-link">
                      证书 #{{ item.id }}
                    </router-link>
                  </template>
                  <template #description>
                    <a-space>
                      <a-tag 
                        :color="getCertificateStatusColor(item.status)"
                        class="status-tag"
                      >
                        {{ getCertificateStatusText(item.status) }}
                      </a-tag>
                      <span class="quantity-text">{{ item.quantity }} 吨CO₂e</span>
                    </a-space>
                  </template>
                </a-list-item-meta>
                
                <template #actions>
                  <span class="date-text">
                    {{ formatDate(item.timestamp) }}
                  </span>
                </template>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
      </a-col>
    </a-row>

    <!-- 个人资料区域 -->
    <a-row :gutter="[16, 16]" class="profile-section">
      <!-- 未连接钱包的情况 -->
      <a-col :span="24" v-if="!walletInfo.isConnected">
        <a-card class="profile-card">
          <div class="profile-empty">
            <a-empty
              description="请连接钱包以开始使用碳信通平台"
            >
              <a-button
                type="primary"
                size="large"
                :loading="isConnecting"
                @click="handleConnectWallet"
              >
                <WalletOutlined />
                连接钱包
              </a-button>
            </a-empty>
          </div>
        </a-card>
      </a-col>

      <!-- 已连接钱包但未注册用户 -->
      <a-col :span="24" v-else-if="!currentUser">
        <a-card title="个人资料" class="profile-card">
          <div class="profile-empty">
            <a-empty
              description="请完善个人资料以使用平台功能"
            >
              <a-button
                type="primary"
                size="large"
                @click="showProfileModal = true"
              >
                <UserOutlined />
                完善个人资料
              </a-button>
            </a-empty>
          </div>
        </a-card>
      </a-col>

      <!-- 已注册用户的个人资料展示 -->
      <a-col :span="24" v-else>
        <a-card title="个人资料" class="profile-card">
          <template #extra>
            <a-space>
              <a-button size="small" @click="goToProfile">
                查看更多
              </a-button>
              <a-button type="primary" size="small" @click="showProfileModal = true">
                编辑个人资料
              </a-button>
            </a-space>
          </template>

          <div class="profile-content">
            <div class="profile-avatar">
              <DefaultAvatar :name="currentUser.name || '用户'" :size="80" />
            </div>
            <div class="profile-info">
              <a-descriptions :column="2" size="small">
                <a-descriptions-item label="钱包地址">
                  <a-typography-text copyable>{{ walletInfo.address }}</a-typography-text>
                </a-descriptions-item>
                <a-descriptions-item label="用户名">
                  {{ currentUser.name || '未设置' }}
                </a-descriptions-item>
                <a-descriptions-item label="角色" :span="2">
                  <a-space>
                    <a-tag
                      v-for="role in currentUser.roles"
                      :key="role"
                      :color="getRoleColor(role)"
                    >
                      {{ getRoleText(role) }}
                    </a-tag>
                  </a-space>
                </a-descriptions-item>
              </a-descriptions>
            </div>
          </div>
        </a-card>
      </a-col>
    </a-row>

    <!-- 个人资料完善弹窗 -->
    <ProfileModal 
      v-model:visible="showProfileModal" 
      :user="currentUser"
      @success="handleProfileUpdated"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, nextTick } from 'vue'
import { 
  UserOutlined, 
  ProjectOutlined, 
  GoldOutlined, 
  SafetyCertificateOutlined,
  WalletOutlined
} from '@ant-design/icons-vue'
import ThemeToggle from '@/components/ThemeToggle.vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { connectWallet, walletInfo, isConnecting } from '@/services/web3'
import { getSummary, getAllProjects, getAllCredits, getAllCertificates } from '@/services/api'
import type { PlatformSummary, ProjectRecord, CreditRecord, CertificateRecord } from '@/types'
import { ProjectStatus, CreditEventType, CertificateStatus, Role } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProfileModal from '@/components/ProfileModal.vue'

// 防抖函数
const debounce = (func: Function, wait: number) => {
  let timeout: NodeJS.Timeout
  return function executedFunction(...args: any[]) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 数据状态
const loading = ref(false)
const projectsLoading = ref(false)
const creditsLoading = ref(false)
const certificatesLoading = ref(false)
const showProfileModal = ref(false)
const summary = ref<PlatformSummary | null>(null)
const recentProjects = ref<ProjectRecord[]>([])
const recentCredits = ref<CreditRecord[]>([])
const recentCertificates = ref<CertificateRecord[]>([])

// 数据加载状态管理
const dataLoaded = ref({
  summary: false,
  projects: false,
  credits: false,
  certificates: false
})

// 获取平台摘要数据（带缓存检查）
const fetchSummary = async () => {
  if (dataLoaded.value.summary && summary.value) {
    return
  }
  
  try {
    loading.value = true
    summary.value = await getSummary()
    dataLoaded.value.summary = true
  } catch (error) {
    console.error('获取平台摘要失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取最新项目（懒加载）
const fetchRecentProjects = async () => {
  if (dataLoaded.value.projects && recentProjects.value.length > 0) {
    return
  }
  
  try {
    projectsLoading.value = true
    const projects = await getAllProjects()
    recentProjects.value = projects
      .sort((a, b) => b.timestamp - a.timestamp)
      .slice(0, 5)
    dataLoaded.value.projects = true
  } catch (error) {
    console.error('获取最新项目失败:', error)
  } finally {
    projectsLoading.value = false
  }
}

// 获取最新碳信用（懒加载）
const fetchRecentCredits = async () => {
  if (dataLoaded.value.credits && recentCredits.value.length > 0) {
    return
  }
  
  try {
    creditsLoading.value = true
    const credits = await getAllCredits()
    recentCredits.value = credits
      .sort((a, b) => b.timestamp - a.timestamp)
      .slice(0, 5)
    dataLoaded.value.credits = true
  } catch (error) {
    console.error('获取最新碳信用失败:', error)
  } finally {
    creditsLoading.value = false
  }
}

// 获取最新证书（懒加载）
const fetchRecentCertificates = async () => {
  if (dataLoaded.value.certificates && recentCertificates.value.length > 0) {
    return
  }
  
  try {
    certificatesLoading.value = true
    const certificates = await getAllCertificates()
    recentCertificates.value = certificates
      .sort((a, b) => b.timestamp - a.timestamp)
      .slice(0, 5)
    dataLoaded.value.certificates = true
  } catch (error) {
    console.error('获取最新证书失败:', error)
  } finally {
    certificatesLoading.value = false
  }
}

// 防抖的数据刷新函数
const debouncedRefresh = debounce(async () => {
  dataLoaded.value = {
    summary: false,
    projects: false,
    credits: false,
    certificates: false
  }
  
  // 并行加载数据
  await Promise.allSettled([
    fetchSummary(),
    fetchRecentProjects(),
    fetchRecentCredits(),
    fetchRecentCertificates()
  ])
}, 500)

// 连接钱包
const handleConnectWallet = async () => {
  try {
    await connectWallet()
    await userStore.fetchCurrentUser()
    message.success('钱包连接成功')
  } catch (error: any) {
    message.error(error.message || '钱包连接失败')
  }
}

// 跳转到个人资料页面
const goToProfile = () => {
  router.push('/user/profile')
}

// 个人资料更新成功处理
const handleProfileUpdated = () => {
  showProfileModal.value = false
  userStore.fetchCurrentUser()
  message.success('个人资料更新成功')
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm')
}

const getProjectStatusColor = (status: ProjectStatus) => {
  const colors = {
    [ProjectStatus.EDITING]: 'orange',
    [ProjectStatus.REVIEWING]: 'blue',
    [ProjectStatus.APPROVED]: 'green',
    [ProjectStatus.REVOKED]: 'red'
  }
  return colors[status] || 'default'
}

const getProjectStatusText = (status: ProjectStatus) => {
  const texts = {
    [ProjectStatus.EDITING]: '编辑中',
    [ProjectStatus.REVIEWING]: '审核中',
    [ProjectStatus.APPROVED]: '已批准',
    [ProjectStatus.REVOKED]: '已撤销'
  }
  return texts[status] || '未知'
}

const getCreditEventColor = (eventType: CreditEventType) => {
  const colors = {
    [CreditEventType.MINT]: 'green',
    [CreditEventType.TRANSFER]: 'blue',
    [CreditEventType.RETIRE]: 'purple',
    [CreditEventType.REVOKE]: 'red'
  }
  return colors[eventType] || 'default'
}

const getCreditEventText = (eventType: CreditEventType) => {
  const texts = {
    [CreditEventType.MINT]: '铸造',
    [CreditEventType.TRANSFER]: '转让',
    [CreditEventType.RETIRE]: '退役',
    [CreditEventType.REVOKE]: '撤销'
  }
  return texts[eventType] || '未知'
}

const getRoleColor = (role: number) => {
  const colors: Record<number, string> = {
    [Role.REGULATOR]: 'red',
    [Role.PROJECT_OWNER]: 'blue',
    [Role.VERIFIER]: 'green',
    [Role.BUYER]: 'orange'
  }
  return colors[role] || 'default'
}

const getRoleText = (role: number) => {
  const texts: Record<number, string> = {
    [Role.NONE]: '无',
    [Role.REGULATOR]: '监管者',
    [Role.PROJECT_OWNER]: '项目所有者',
    [Role.VERIFIER]: '验证者',
    [Role.BUYER]: '购买者'
  }
  return texts[role] || '未知'
}

const getCertificateStatusColor = (status: CertificateStatus) => {
  const colors = {
    [CertificateStatus.ACTIVE]: 'green',
    [CertificateStatus.REVOKED]: 'red'
  }
  return colors[status] || 'default'
}

const getCertificateStatusText = (status: CertificateStatus) => {
  const texts = {
    [CertificateStatus.ACTIVE]: '有效',
    [CertificateStatus.REVOKED]: '已撤销'
  }
  return texts[status] || '未知'
}

const getBatchIdFromHistory = (credit: CreditRecord) => {
  // 由于后端返回的是历史记录，需要从索引推断批次ID
  return 1 // 简化处理，实际应用中需要更复杂的逻辑
}

// 懒加载数据
const loadDataOnVisible = async () => {
  await nextTick()
  
  // 优先加载摘要数据
  await fetchSummary()
  
  // 延迟加载其他数据，避免阻塞初始渲染
  setTimeout(() => {
    Promise.allSettled([
      fetchRecentProjects(),
      fetchRecentCredits(),
      fetchRecentCertificates()
    ])
  }, 100)
}

onMounted(() => {
  loadDataOnVisible()
})
</script>

<style scoped>
.dashboard {
  max-width: 1400px;
  margin: 0 auto;
  padding: 32px 24px;
  min-height: 100vh;
}

/* 页面标题样式 */
.page-header {
  margin-bottom: 40px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  padding: 32px;
  border-radius: 8px;
  color: var(--color-text);
  box-shadow: var(--box-shadow);
}

.title-section {
  flex: 1;
}

.main-title {
  margin: 0;
  font-size: 36px;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 12px;
}

.title-icon {
  font-size: 40px;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
}

.subtitle {
  margin: 8px 0 0 0;
  font-size: 16px;
  opacity: 0.9;
  font-weight: 400;
}

.header-actions {
  flex-shrink: 0;
}

.action-buttons {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* 统计卡片样式 */
.stats-cards {
  margin-bottom: 32px;
}

.stat-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  overflow: hidden;
  position: relative;
  transition: all 0.15s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: var(--box-shadow);
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: var(--color-primary);
}

.stat-card-users::before {
  background: var(--color-success);
}

.stat-card-projects::before {
  background: var(--color-primary);
}

.stat-card-credits::before {
  background: var(--color-warning);
}

.stat-card-certificates::before {
  background: #8250df;
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 24px;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: #ffffff;
  flex-shrink: 0;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.users-icon {
  background: var(--color-success);
}

.projects-icon {
  background: var(--color-primary);
}

.credits-icon {
  background: var(--color-warning);
}

.certificates-icon {
  background: #8250df;
}

.stat-info {
  flex: 1;
}

.stat-trend {
  margin-top: 4px;
}

.trend-text {
  font-size: 12px;
  color: var(--color-text-secondary);
  font-weight: 500;
}

/* 内容卡片样式 */
.content-section {
  margin-bottom: 32px;
}

.content-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.15s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: var(--box-shadow);
}

.card-title {
  display: flex;
  align-items: center;
  gap: 12px;
  font-weight: 600;
  font-size: 16px;
}

.title-icon-wrapper {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: var(--color-primary);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 16px;
}

.more-link {
  color: var(--color-primary);
  text-decoration: none;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 4px;
  transition: all 0.2s ease;
}

.more-link:hover {
  color: var(--color-primary);
  opacity: 0.8;
  transform: translateX(2px);
}

.arrow {
  transition: transform 0.2s ease;
}

.more-link:hover .arrow {
  transform: translateX(4px);
}

/* 列表样式 */
.modern-list {
  margin-top: 16px;
}

.list-item-modern {
  padding: 16px 0;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.2s ease;
}

.list-item-modern:hover {
  background: var(--color-bg-elevated);
  border-radius: 8px;
  padding-left: 12px;
  padding-right: 12px;
}

.list-item-modern:last-child {
  border-bottom: none;
}

.project-link {
  color: #262626;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s ease;
}

.project-link:hover {
  color: var(--color-primary);
}

.status-tag {
  border-radius: 6px;
  font-weight: 500;
  font-size: 12px;
}

.location-text,
.quantity-text {
  color: var(--color-text-secondary);
  font-size: 13px;
}

.date-text {
  color: var(--color-text-tertiary);
  font-size: 12px;
  font-weight: 500;
}

/* 个人资料区域 */
.profile-section {
  margin-top: 32px;
}

.profile-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  min-height: 200px;
  overflow: hidden;
  box-shadow: var(--box-shadow);
}

.profile-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 200px;
  gap: 16px;
}

.profile-content {
  display: flex;
  align-items: center;
  gap: 24px;
  padding: 24px;
}

.profile-avatar {
  flex-shrink: 0;
}

.profile-info {
  flex: 1;
}

.profile-actions {
  display: flex;
  gap: 12px;
  margin-top: 16px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .dashboard {
    padding: 16px;
  }
  
  .header-content {
    flex-direction: column;
    gap: 20px;
    text-align: center;
    padding: 24px;
  }
  
  .main-title {
    font-size: 28px;
  }
  
  .stat-content {
    flex-direction: column;
    text-align: center;
    gap: 12px;
  }
  
  .profile-content {
    flex-direction: column;
    text-align: center;
    gap: 16px;
  }
  
  .profile-actions {
    justify-content: center;
  }
}

/* 深度样式覆盖 */
:deep(.ant-statistic-title) {
  font-size: 14px;
  color: #8c8c8c;
  font-weight: 500;
  margin-bottom: 8px;
}

:deep(.ant-statistic-content) {
  font-size: 28px;
  font-weight: 600;
  line-height: 1.2;
}

:deep(.ant-card-head) {
  border-bottom: 1px solid #f0f0f0;
  padding: 20px 24px 16px;
}

:deep(.ant-card-body) {
  padding: 24px;
}

:deep(.ant-list-item-meta-title) {
  margin-bottom: 4px;
}

:deep(.ant-list-item-meta-description) {
  margin-bottom: 0;
}

:deep(.ant-tag) {
  border-radius: 6px;
  font-weight: 500;
  font-size: 12px;
  padding: 2px 8px;
  border: none;
}

:deep(.ant-btn) {
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.2s ease;
}

:deep(.ant-btn-primary) {
  background: var(--color-primary);
  border-color: var(--color-primary);
  box-shadow: var(--box-shadow);
}

:deep(.ant-btn-primary:hover) {
  background: var(--color-primary);
  border-color: var(--color-primary);
  opacity: 0.9;
  transform: translateY(-1px);
  box-shadow: var(--box-shadow-secondary);
}
</style>
