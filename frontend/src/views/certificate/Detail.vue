<template>
  <div class="certificate-detail">
    <!-- 页面标题和导航 -->
    <div class="page-header">
      <a-page-header
        :title="`证书详情 #${certificateId}`"
        @back="goBack"
      >
        <template #extra>
          <a-space>
            <a-button @click="refreshData">
              <ReloadOutlined />
              刷新
            </a-button>
          </a-space>
        </template>
      </a-page-header>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <a-spin size="large">
        <template #tip>正在加载证书信息...</template>
      </a-spin>
    </div>

    <!-- 证书不存在 -->
    <template v-else-if="!certificate">
      <a-result
        status="404"
        title="证书不存在"
        sub-title="抱歉，您访问的证书不存在或已被删除"
      >
        <template #extra>
          <a-button type="primary" @click="goBack">
            返回
          </a-button>
        </template>
      </a-result>
    </template>

    <!-- 证书详情内容 -->
    <template v-else>
      <!-- 证书基本信息 -->
      <a-card title="证书基本信息" class="detail-card">
        <template #extra>
          <a-tag :color="getCertificateStatusColor(certificate.status)" size="large">
            {{ getCertificateStatusText(certificate.status) }}
          </a-tag>
        </template>

        <a-descriptions :column="2" size="middle">
          <a-descriptions-item label="证书编号">
            <span class="certificate-id">#{{ certificate.id }}</span>
          </a-descriptions-item>
          
          <a-descriptions-item label="批次编号">
            <a-tag color="blue">#{{ certificate.batchId }}</a-tag>
          </a-descriptions-item>
          
          <a-descriptions-item label="关联项目">
            <router-link :to="`/project/${certificate.projectId}`">
              {{ projectName }} (#{{ certificate.projectId }})
            </router-link>
          </a-descriptions-item>
          
          <a-descriptions-item label="数量">
            <span class="quantity-amount">{{ certificate.quantity.toLocaleString() }} 吨CO₂e</span>
          </a-descriptions-item>
          
          <a-descriptions-item label="所有者">
            <div class="owner-info">
              <DefaultAvatar :name="ownerName" :size="32" />
              <div class="owner-detail">
                <div class="owner-name">{{ ownerName || '未知用户' }}</div>
                <div class="owner-address">{{ formatAddress(certificate.owner) }}</div>
              </div>
            </div>
          </a-descriptions-item>
          
          <a-descriptions-item label="核发时间">
            {{ formatDate(certificate.timestamp) }}
          </a-descriptions-item>
          
          <a-descriptions-item v-if="certificate.revokedAt" label="撤销时间">
            <span class="revoked-time">{{ formatDate(certificate.revokedAt) }}</span>
          </a-descriptions-item>
          
          <a-descriptions-item v-if="certificate.reason" label="操作原因" :span="2">
            <div class="reason-content">
              {{ certificate.reason }}
            </div>
          </a-descriptions-item>
        </a-descriptions>

        <!-- 证书状态说明 -->
        <div class="certificate-status-info">
          <a-alert
            :type="getAlertType(certificate.status)"
            :message="getStatusMessage(certificate.status)"
            show-icon
          >
            <template #description>
              <div v-if="certificate.status === 0">
                该证书已正式核发，证明对应的碳信用已被确认和验证。
              </div>
              <div v-else-if="certificate.status === 1">
                该证书已被监管机构撤销，通常是由于项目或碳信用存在问题。
                <div v-if="certificate.revokedAt" class="revoke-details">
                  撤销时间：{{ formatDate(certificate.revokedAt) }}
                </div>
              </div>
            </template>
          </a-alert>
        </div>

        <!-- 管理操作 -->
        <div v-if="canRevokeCertificate" class="admin-actions">
          <a-divider>管理操作</a-divider>
          <a-space>
            <a-button 
              type="primary"
              danger
              @click="revokeCertificate"
            >
              <StopOutlined />
              撤销证书
            </a-button>
          </a-space>
        </div>
      </a-card>

      <!-- 关联项目信息 -->
      <a-card v-if="project" title="关联项目信息" class="detail-card">
        <ProjectDetailView :project="project" :show-actions="false" />
      </a-card>

      <!-- 相关碳信用记录 -->
      <a-card title="相关碳信用记录" class="detail-card">
        <template #extra>
          <a-button size="small" @click="fetchRelatedCredits">
            <ReloadOutlined />
            刷新
          </a-button>
        </template>

        <div v-if="relatedCredits.length === 0" class="empty-state">
          <a-empty description="暂无相关碳信用记录" />
        </div>

        <a-table
          v-else
          :columns="creditColumns"
          :data-source="relatedCredits"
          :pagination="false"
          size="small"
          row-key="id"
        >
          <template #bodyCell="{ column, record }: { column: any, record: any }">
            <!-- 批次信息 -->
            <template v-if="column.key === 'batch'">
              <div class="batch-cell">
                <div class="batch-id">批次 #{{ getBatchId(record) }}</div>
                <div class="vintage-year">{{ record.vintageYear }}年</div>
              </div>
            </template>

            <!-- 事件类型 -->
            <template v-else-if="column.key === 'eventType'">
              <a-tag :color="getCreditEventColor(record.eventType)">
                {{ getCreditEventText(record.eventType) }}
              </a-tag>
            </template>

            <!-- 数量 -->
            <template v-else-if="column.key === 'quantity'">
              <span class="quantity-amount">{{ record.quantity.toLocaleString() }} 吨CO₂e</span>
            </template>

            <!-- 当前所有者 -->
            <template v-else-if="column.key === 'currentOwner'">
              <div class="owner-cell">
                <DefaultAvatar :name="getOwnerName(record.currentOwner)" :size="24" />
                <div class="owner-info">
                  <div class="owner-name">{{ getOwnerName(record.currentOwner) }}</div>
                  <div class="owner-address">{{ formatAddress(record.currentOwner) }}</div>
                </div>
              </div>
            </template>

            <!-- 时间 -->
            <template v-else-if="column.key === 'timestamp'">
              {{ formatDate(record.timestamp) }}
            </template>

            <!-- 操作 -->
            <template v-else-if="column.key === 'actions'">
              <a-button size="small" @click="viewCreditDetail(record)">
                详情
              </a-button>
            </template>
          </template>
        </a-table>
      </a-card>
    </template>

    <!-- 撤销证书确认弹窗 -->
    <EditReasonModal
      v-model:visible="showRevokeModal"
      title="撤销证书"
      :description="`确定要撤销证书 #${certificate?.id} 吗？撤销后无法恢复。`"
      placeholder="请说明撤销原因"
      @confirm="handleCertificateRevoke"
    />

    <!-- 碳信用详情弹窗 -->
    <a-modal
      v-model:open="showCreditDetailModal"
      title="碳信用详情"
      width="700px"
      :footer="null"
    >
      <CreditDetailView v-if="selectedCredit" :credit="selectedCredit" />
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  ReloadOutlined,
  StopOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { 
  getAllCertificates, 
  getAllProjects, 
  getAllCredits,
  revokeCertificate as apiRevokeCertificate 
} from '@/services/api'
import { PermissionService } from '@/services/permission'
import type { CertificateRecord, ProjectRecord, CreditRecord } from '@/types'
import { CertificateStatus, CreditEventType } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProjectDetailView from '@/components/ProjectDetailView.vue'
import CreditDetailView from '@/components/CreditDetailView.vue'
import EditReasonModal from '@/components/EditReasonModal.vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 路由参数
const certificateId = computed(() => Number(route.params.id))

// 状态
const loading = ref(false)
const certificate = ref<CertificateRecord | null>(null)
const project = ref<ProjectRecord | null>(null)
const relatedCredits = ref<CreditRecord[]>([])
const showRevokeModal = ref(false)
const showCreditDetailModal = ref(false)
const selectedCredit = ref<CreditRecord | null>(null)

// 计算属性
const projectName = computed(() => {
  return project.value?.name || `项目 #${certificate.value?.projectId}`
})

const ownerName = computed(() => {
  if (!certificate.value) return ''
  const user = userStore.getUserByAddress(certificate.value.owner)
  return user?.name || ''
})

const canRevokeCertificate = computed(() => {
  return PermissionService.isRegulator(currentUser.value) && 
         certificate.value?.status === 0 // ISSUED = 0
})

// 碳信用表格列
const creditColumns = [
  {
    title: '批次信息',
    key: 'batch',
    width: 120
  },
  {
    title: '事件类型',
    key: 'eventType',
    width: 100
  },
  {
    title: '数量',
    key: 'quantity',
    width: 120
  },
  {
    title: '当前所有者',
    key: 'currentOwner',
    width: 180
  },
  {
    title: '时间',
    key: 'timestamp',
    width: 150
  },
  {
    title: '操作',
    key: 'actions',
    width: 80
  }
]

// 获取证书详情
const fetchCertificateDetail = async () => {
  try {
    loading.value = true
    const [certificateData, projectData] = await Promise.all([
      getAllCertificates(),
      getAllProjects()
    ])
    
    // 确保数据是数组格式
    const certificates = Array.isArray(certificateData) ? certificateData : []
    const projects = Array.isArray(projectData) ? projectData : []
    
    certificate.value = certificates.find(cert => cert && cert.id === certificateId.value) || null
    
    if (certificate.value) {
      project.value = projects.find(proj => proj && proj.id === certificate.value!.projectId) || null
      await fetchRelatedCredits()
    }
  } catch (error: any) {
    console.error('获取证书详情失败:', error)
    // 设置空值避免页面崩溃
    certificate.value = null
    project.value = null
    relatedCredits.value = []
    
    // 根据错误类型显示不同消息
    if (error.code === 'NETWORK_ERROR' || error.message?.includes('Network Error')) {
      message.error('网络连接失败，请检查网络后重试')
    } else if (error.response?.status === 404) {
      message.error('后端服务未启动，请启动后端服务')
    } else {
      message.error('获取证书详情失败，请稍后重试')
    }
  } finally {
    loading.value = false
  }
}

// 获取相关碳信用记录
const fetchRelatedCredits = async () => {
  if (!certificate.value) {
    relatedCredits.value = []
    return
  }

  try {
    const creditData = await getAllCredits()
    
    // 确保数据是数组格式
    const credits = Array.isArray(creditData) ? creditData : []
    
    // 找到与此证书批次相关的碳信用记录，添加安全检查
    relatedCredits.value = credits.filter(credit => {
      if (!credit || !certificate.value) return false
      return credit.parentBatchId === certificate.value!.batchId ||
             credit.childBatchId === certificate.value!.batchId
    }).sort((a, b) => {
      if (!a || !b) return 0
      return b.timestamp - a.timestamp
    })
  } catch (error: any) {
    console.error('获取相关碳信用记录失败:', error)
    // 设置空数组避免页面崩溃
    relatedCredits.value = []
    
    // 根据错误类型显示不同消息
    if (error.code === 'NETWORK_ERROR' || error.message?.includes('Network Error')) {
      message.error('网络连接失败，请检查网络后重试')
    } else if (error.response?.status === 404) {
      message.error('后端服务未启动，请启动后端服务')
    } else {
      message.error('获取相关碳信用记录失败，请稍后重试')
    }
  }
}

// 刷新数据
const refreshData = () => {
  fetchCertificateDetail()
}

// 返回上一页
const goBack = () => {
  router.back()
}

// 撤销证书
const revokeCertificate = () => {
  showRevokeModal.value = true
}

// 处理证书撤销
const handleCertificateRevoke = async (reason: string) => {
  if (!certificate.value) return

  try {
    await apiRevokeCertificate(certificate.value.id, reason)
    showRevokeModal.value = false
    fetchCertificateDetail()
    message.success('证书撤销成功')
  } catch (error: any) {
    message.error(error.message || '证书撤销失败')
  }
}

// 查看碳信用详情
const viewCreditDetail = (credit: CreditRecord) => {
  selectedCredit.value = credit
  showCreditDetailModal.value = true
}

// 获取用户名
const getOwnerName = (address: string) => {
  const user = userStore.getUserByAddress(address)
  return user?.name || ''
}

// 获取批次ID
const getBatchId = (credit: CreditRecord) => {
  return credit.childBatchId || credit.parentBatchId || 1
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm:ss')
}

const formatAddress = (address: string) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const getCertificateStatusColor = (status: CertificateStatus) => {
  const colors = {
    0: 'green', // ISSUED
    1: 'red'    // REVOKED
  }
  return colors[status] || 'default'
}

const getCertificateStatusText = (status: CertificateStatus) => {
  const texts = {
    0: '已核发', // ISSUED
    1: '已撤销'  // REVOKED
  }
  return texts[status] || '未知'
}

const getAlertType = (status: CertificateStatus): 'success' | 'error' | 'warning' | 'info' => {
  const types: Record<number, 'success' | 'error' | 'warning' | 'info'> = {
    0: 'success', // ISSUED
    1: 'error'    // REVOKED
  }
  return types[status] || 'info'
}

const getStatusMessage = (status: CertificateStatus) => {
  const messages = {
    0: '证书已核发', // ISSUED
    1: '证书已撤销'  // REVOKED
  }
  return messages[status] || '未知状态'
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

onMounted(() => {
  fetchCertificateDetail()
})
</script>

<style scoped>
.certificate-detail {
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 16px;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}

.detail-card {
  margin-bottom: 16px;
}

.certificate-id {
  font-size: 18px;
  font-weight: 600;
  color: #1890ff;
}

.quantity-amount {
  font-weight: 500;
  color: #52c41a;
  font-size: 16px;
}

.owner-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.owner-detail {
  min-width: 0;
  flex: 1;
}

.owner-name {
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.owner-address {
  font-size: 12px;
  color: #8c8c8c;
  font-family: monospace;
}

.revoked-time {
  color: #f5222d;
  font-weight: 500;
}

.reason-content {
  background: #fafafa;
  padding: 12px;
  border-radius: 6px;
  border-left: 3px solid #1890ff;
  color: #666;
  line-height: 1.6;
}

.certificate-status-info {
  margin-top: 20px;
}

.revoke-details {
  margin-top: 8px;
  font-size: 12px;
  color: #666;
}

.admin-actions {
  margin-top: 20px;
}

.empty-state {
  padding: 40px 0;
  text-align: center;
}

.batch-cell {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.batch-id {
  font-weight: 500;
  color: #262626;
  font-size: 14px;
}

.vintage-year {
  font-size: 12px;
  color: #8c8c8c;
}

.owner-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.owner-cell .owner-info {
  min-width: 0;
  flex: 1;
}

.owner-cell .owner-name {
  font-size: 14px;
  color: #262626;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-bottom: 2px;
}

.owner-cell .owner-address {
  font-size: 12px;
  color: #8c8c8c;
  font-family: monospace;
}

:deep(.ant-descriptions-item-label) {
  font-weight: 500;
  color: #333;
}

:deep(.ant-descriptions-item-content) {
  color: #666;
}

:deep(.ant-page-header-heading-title) {
  font-size: 20px;
  font-weight: 600;
}
</style>
