<template>
  <div class="credit-detail">
    <!-- 页面标题和导航 -->
    <div class="page-header">
      <a-page-header
        :title="`碳信用详情 #${creditId}`"
        @back="goBack"
      >
        <template #extra>
          <a-space>
            <a-button @click="refreshData">
              <ReloadOutlined />
              刷新
            </a-button>
            <a-button 
              v-if="canTransfer" 
              type="primary" 
              @click="showTransferModal"
            >
              <SwapOutlined />
              转让
            </a-button>
            <a-button 
              v-if="canRetire" 
              type="primary" 
              danger 
              @click="showRetireModal"
            >
              <DeleteOutlined />
              注销
            </a-button>
          </a-space>
        </template>
      </a-page-header>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <a-spin size="large">
        <template #tip>正在加载碳信用信息...</template>
      </a-spin>
    </div>

    <!-- 碳信用不存在 -->
    <template v-else-if="!credit">
      <a-result
        status="404"
        title="碳信用不存在"
        sub-title="抱歉，您访问的碳信用不存在或已被删除"
      >
        <template #extra>
          <a-button type="primary" @click="goBack">
            返回
          </a-button>
        </template>
      </a-result>
    </template>

    <!-- 碳信用详情内容 -->
    <template v-else>
      <!-- 碳信用基本信息 -->
      <a-card title="碳信用基本信息" class="detail-card">
        <template #extra>
          <a-tag :color="getCreditEventColor(credit.eventType)" size="large">
            {{ getCreditEventText(credit.eventType) }}
          </a-tag>
        </template>

        <a-descriptions :column="2" size="middle">
          <a-descriptions-item label="批次ID">
            <a-tag color="blue" size="large">#{{ getBatchId(credit) }}</a-tag>
          </a-descriptions-item>
          
          <a-descriptions-item label="年份">
            {{ credit.vintageYear }}年
          </a-descriptions-item>
          
          <a-descriptions-item label="数量">
            <span class="quantity-amount">{{ credit.quantity.toLocaleString() }} 吨CO₂e</span>
          </a-descriptions-item>
          
          <a-descriptions-item label="当前所有者">
            <div class="owner-info">
              <DefaultAvatar :name="getOwnerName(credit.currentOwner)" :size="24" />
              <span class="owner-name">{{ getOwnerName(credit.currentOwner) || formatAddress(credit.currentOwner) }}</span>
            </div>
          </a-descriptions-item>
          
          <a-descriptions-item label="所有者记录索引">
            <a-tag size="small" color="blue">#{{ credit.ownerRecordIndex }}</a-tag>
          </a-descriptions-item>
          
          <a-descriptions-item label="事件时间">
            {{ formatDate(credit.timestamp) }}
          </a-descriptions-item>
          
          <a-descriptions-item label="关联项目" :span="2">
            <router-link :to="`/project/${credit.projectId}`" class="project-link">
              <ProjectOutlined />
              项目 #{{ credit.projectId }}
            </router-link>
          </a-descriptions-item>
        </a-descriptions>

        <!-- 转让信息（仅在转让事件时显示） -->
        <template v-if="credit.eventType === CreditEventType.TRANSFER">
          <a-divider>转让信息</a-divider>
          <a-descriptions :column="2" size="middle">
            <a-descriptions-item label="转出方">
              <div class="transfer-info">
                <DefaultAvatar :name="getOwnerName(credit.from)" :size="24" />
                <span class="owner-name">{{ getOwnerName(credit.from) || formatAddress(credit.from) }}</span>
                <a-tag size="small" color="orange">记录索引: {{ credit.fromRecordIndex }}</a-tag>
              </div>
            </a-descriptions-item>
            
            <a-descriptions-item label="转入方">
              <div class="transfer-info">
                <DefaultAvatar :name="getOwnerName(credit.currentOwner)" :size="24" />
                <span class="owner-name">{{ getOwnerName(credit.currentOwner) || formatAddress(credit.currentOwner) }}</span>
                <a-tag size="small" color="green">记录索引: {{ credit.ownerRecordIndex }}</a-tag>
              </div>
            </a-descriptions-item>
          </a-descriptions>
        </template>
      </a-card>

      <!-- 相关项目信息 -->
      <a-card title="相关项目信息" class="detail-card" v-if="project">
        <a-descriptions :column="2" size="middle">
          <a-descriptions-item label="项目名称">
            {{ project.name }}
          </a-descriptions-item>
          
          <a-descriptions-item label="项目技术">
            <a-tag :color="getTechnologyColor(project.technology)">
              {{ getTechnologyText(project.technology) }}
            </a-tag>
          </a-descriptions-item>
          
          <a-descriptions-item label="项目状态">
            <a-tag :color="getProjectStatusColor(project.status)">
              {{ getProjectStatusText(project.status) }}
            </a-tag>
          </a-descriptions-item>
          
          <a-descriptions-item label="项目所有者">
            <div class="owner-info">
              <DefaultAvatar :name="getOwnerName(project.owner)" :size="24" />
              <span class="owner-name">{{ getOwnerName(project.owner) || formatAddress(project.owner) }}</span>
            </div>
          </a-descriptions-item>
          
          <a-descriptions-item label="项目描述" :span="2">
            {{ project.description }}
          </a-descriptions-item>
        </a-descriptions>

        <template #extra>
          <router-link :to="`/project/${credit.projectId}`">
            <a-button type="link">
              <EyeOutlined />
              查看项目详情
            </a-button>
          </router-link>
        </template>
      </a-card>

      <!-- 操作历史 -->
      <a-card title="操作历史" class="detail-card">
        <a-timeline>
          <a-timeline-item
            v-for="(historyCredit, index) in creditHistory"
            :key="index"
            :color="getCreditEventColor(historyCredit.eventType)"
          >
            <template #dot>
              <component 
                :is="getCreditEventIcon(historyCredit.eventType)" 
                style="font-size: 16px;"
              />
            </template>
            
            <div class="timeline-content">
              <div class="timeline-header">
                <span class="event-type">{{ getCreditEventText(historyCredit.eventType) }}</span>
                <span class="event-time">{{ formatDate(historyCredit.timestamp) }}</span>
              </div>
              
              <div class="timeline-details">
                <p><strong>数量：</strong>{{ historyCredit.quantity.toLocaleString() }} 吨CO₂e</p>
                <p><strong>所有者：</strong>{{ getOwnerName(historyCredit.currentOwner) || formatAddress(historyCredit.currentOwner) }}</p>
                
                <template v-if="historyCredit.eventType === CreditEventType.TRANSFER">
                  <p><strong>转出方：</strong>{{ getOwnerName(historyCredit.from) || formatAddress(historyCredit.from) }}</p>
                  <p><strong>转入方：</strong>{{ getOwnerName(historyCredit.currentOwner) || formatAddress(historyCredit.currentOwner) }}</p>
                </template>
                
                <template v-if="historyCredit.eventType === CreditEventType.MINT">
                  <p><strong>年份：</strong>{{ historyCredit.vintageYear }}年</p>
                  <p><strong>关联项目：</strong>项目 #{{ historyCredit.projectId }}</p>
                </template>
              </div>
            </div>
          </a-timeline-item>
        </a-timeline>
      </a-card>
    </template>

    <!-- 转让模态框 -->
    <CreditTransferModal
      v-if="credit"
      :visible="transferModalVisible"
      :credit="credit"
      @cancel="transferModalVisible = false"
      @success="handleTransferSuccess"
    />

    <!-- 注销确认模态框 -->
    <a-modal
      v-model:visible="retireModalVisible"
      title="确认注销碳信用"
      @ok="handleRetire"
      :confirm-loading="retireLoading"
    >
      <p>您确定要注销这个碳信用吗？</p>
      <p><strong>注意：此操作不可逆！</strong></p>
      <a-descriptions :column="1" size="small" bordered>
        <a-descriptions-item label="批次ID">#{{ getBatchId(credit) }}</a-descriptions-item>
        <a-descriptions-item label="数量">{{ credit?.quantity.toLocaleString() }} 吨CO₂e</a-descriptions-item>
        <a-descriptions-item label="年份">{{ credit?.vintageYear }}年</a-descriptions-item>
      </a-descriptions>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import {
  ReloadOutlined,
  SwapOutlined,
  DeleteOutlined,
  ProjectOutlined,
  EyeOutlined,
  TransactionOutlined,
  PlusOutlined,
  MinusOutlined,
  StopOutlined
} from '@ant-design/icons-vue'
import type { CreditRecord, ProjectRecord } from '@/types'
import { CreditEventType, ProjectStatus, Technology } from '@/types'
import { getCreditHistory, getProjectHistory } from '@/services/api'
import { retireCredits } from '@/services/web3'
import { useUserStore } from '@/stores/user'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import CreditTransferModal from '@/components/CreditTransferModal.vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

// 数据状态
const loading = ref(true)
const credit = ref<CreditRecord | null>(null)
const project = ref<ProjectRecord | null>(null)
const creditHistory = ref<CreditRecord[]>([])

// 模态框状态
const transferModalVisible = ref(false)
const retireModalVisible = ref(false)
const retireLoading = ref(false)

// 路由参数
const creditId = computed(() => route.params.id as string)

// 权限检查
const canTransfer = computed(() => {
  if (!credit.value || !userStore.currentUser) return false
  return credit.value.currentOwner.toLowerCase() === userStore.currentUser.addr.toLowerCase() &&
         credit.value.eventType !== CreditEventType.RETIRE &&
         credit.value.eventType !== CreditEventType.REVOKE
})

const canRetire = computed(() => {
  if (!credit.value || !userStore.currentUser) return false
  return credit.value.currentOwner.toLowerCase() === userStore.currentUser.addr.toLowerCase() &&
         credit.value.eventType !== CreditEventType.RETIRE &&
         credit.value.eventType !== CreditEventType.REVOKE
})

// 工具函数
const formatAddress = (address: string): string => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const formatDate = (timestamp: number): string => {
  return new Date(timestamp * 1000).toLocaleString('zh-CN')
}

const getBatchId = (creditRecord: CreditRecord): string => {
  return `${creditRecord.projectId}-${creditRecord.vintageYear}-${creditRecord.ownerRecordIndex}`
}

const getOwnerName = (address: string): string => {
  const user = userStore.getUserByAddress(address)
  return user?.name || ''
}

// 样式和状态工具函数
const getCreditEventColor = (eventType: CreditEventType): string => {
  switch (eventType) {
    case CreditEventType.MINT: return 'green'
    case CreditEventType.TRANSFER: return 'blue'
    case CreditEventType.RETIRE: return 'orange'
    case CreditEventType.REVOKE: return 'red'
    default: return 'default'
  }
}

const getCreditEventText = (eventType: CreditEventType): string => {
  switch (eventType) {
    case CreditEventType.MINT: return '生成'
    case CreditEventType.TRANSFER: return '转让'
    case CreditEventType.RETIRE: return '注销'
    case CreditEventType.REVOKE: return '撤销'
    default: return '未知'
  }
}

const getCreditEventIcon = (eventType: CreditEventType) => {
  switch (eventType) {
    case CreditEventType.MINT: return PlusOutlined
    case CreditEventType.TRANSFER: return TransactionOutlined
    case CreditEventType.RETIRE: return MinusOutlined
    case CreditEventType.REVOKE: return StopOutlined
    default: return TransactionOutlined
  }
}

const getProjectStatusColor = (status: ProjectStatus): string => {
  switch (status) {
    case ProjectStatus.EDITING: return 'orange'
    case ProjectStatus.REVIEWING: return 'blue'
    case ProjectStatus.APPROVED: return 'green'
    case ProjectStatus.REVOKED: return 'red'
    default: return 'default'
  }
}

const getProjectStatusText = (status: ProjectStatus): string => {
  switch (status) {
    case ProjectStatus.EDITING: return '编辑中'
    case ProjectStatus.REVIEWING: return '审核中'
    case ProjectStatus.APPROVED: return '已批准'
    case ProjectStatus.REVOKED: return '已撤销'
    default: return '未知'
  }
}

const getTechnologyColor = (technology: Technology): string => {
  switch (technology) {
    case Technology.SOLAR: return 'gold'
    case Technology.WIND: return 'cyan'
    case Technology.HYDRO: return 'blue'
    case Technology.FORESTRY: return 'green'
    case Technology.CAPTURE: return 'purple'
    case Technology.OTHER: return 'default'
    default: return 'default'
  }
}

const getTechnologyText = (technology: Technology): string => {
  switch (technology) {
    case Technology.SOLAR: return '太阳能'
    case Technology.WIND: return '风能'
    case Technology.HYDRO: return '水能'
    case Technology.FORESTRY: return '林业'
    case Technology.CAPTURE: return '碳捕获'
    case Technology.OTHER: return '其他'
    default: return '未知'
  }
}

// 数据加载
const loadData = async () => {
  try {
    loading.value = true
    
    // 获取碳信用历史记录
    const history = await getCreditHistory(creditId.value)
    creditHistory.value = history
    
    // 获取最新的碳信用记录
    if (history.length > 0) {
      credit.value = history[history.length - 1]
      
      // 获取关联项目信息
      if (credit.value.projectId) {
        try {
          const projectData = await getProjectHistory(credit.value.projectId)
          if (projectData.length > 0) {
            project.value = projectData[projectData.length - 1]
          }
        } catch (error) {
          console.error('获取项目信息失败:', error)
        }
      }
    }
  } catch (error) {
    console.error('加载碳信用详情失败:', error)
    message.error('加载碳信用详情失败')
  } finally {
    loading.value = false
  }
}

const refreshData = () => {
  loadData()
}

// 页面操作
const goBack = () => {
  router.go(-1)
}

const showTransferModal = () => {
  transferModalVisible.value = true
}

const handleTransferSuccess = () => {
  transferModalVisible.value = false
  message.success('转让成功')
  loadData()
}

const showRetireModal = () => {
  retireModalVisible.value = true
}

const handleRetire = async () => {
  if (!credit.value) return
  
  try {
    retireLoading.value = true
    await retireCredits(Number(creditId.value), credit.value.quantity, '碳信用退役')
    message.success('注销成功')
    retireModalVisible.value = false
    loadData()
  } catch (error) {
    console.error('注销失败:', error)
    message.error('注销失败')
  } finally {
    retireLoading.value = false
  }
}

// 生命周期
onMounted(() => {
  loadData()
})

// 监听路由变化
watch(() => route.params.id, () => {
  if (route.params.id) {
    loadData()
  }
})
</script>

<style scoped>
.credit-detail {
  padding: 24px;
  background: #f5f5f5;
  min-height: 100vh;
}

.page-header {
  background: white;
  margin: -24px -24px 24px -24px;
  border-bottom: 1px solid #f0f0f0;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.detail-card {
  margin-bottom: 24px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.quantity-amount {
  font-size: 16px;
  font-weight: 600;
  color: #1890ff;
}

.owner-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.owner-name {
  font-weight: 500;
}

.transfer-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.project-link {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  color: #1890ff;
  text-decoration: none;
}

.project-link:hover {
  color: #40a9ff;
}

.timeline-content {
  padding-left: 8px;
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.event-type {
  font-weight: 600;
  font-size: 16px;
}

.event-time {
  color: #666;
  font-size: 14px;
}

.timeline-details {
  color: #666;
  font-size: 14px;
  line-height: 1.6;
}

.timeline-details p {
  margin: 4px 0;
}

.timeline-details strong {
  color: #333;
}
</style>
