<template>
  <div class="certificate-overview">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>{{ pageTitle }}</h1>
      <p>{{ pageDescription }}</p>
    </div>

    <!-- 统计信息卡片 -->
    <a-row :gutter="[16, 16]" class="stats-cards">
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="证书总数"
            :value="totalCertificates"
            :value-style="{ color: '#1890ff' }"
          >
            <template #prefix>
              <FileTextOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="已核发证书"
            :value="issuedCertificates"
            :value-style="{ color: '#52c41a' }"
          >
            <template #prefix>
              <CheckCircleOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="已撤销证书"
            :value="revokedCertificates"
            :value-style="{ color: '#f5222d' }"
          >
            <template #prefix>
              <StopOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="本月新增"
            :value="monthlyNewCertificates"
            :value-style="{ color: '#fa8c16' }"
          >
            <template #prefix>
              <PlusCircleOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
    </a-row>

    <!-- 操作栏 -->
    <a-card class="action-card">
      <div class="action-bar">
        <div class="search-filters">
          <a-input-search
            v-model:value="searchText"
            placeholder="搜索证书编号、项目名称或所有者"
            style="width: 300px"
            @search="handleSearch"
            allow-clear
          />
          <a-select
            v-model:value="selectedStatus"
            placeholder="筛选状态"
            style="width: 120px"
            allow-clear
            @change="handleStatusFilter"
          >
            <a-select-option :value="0">已核发</a-select-option>
            <a-select-option :value="1">已撤销</a-select-option>
          </a-select>
          <a-select
            v-model:value="selectedYear"
            placeholder="筛选年份"
            style="width: 120px"
            allow-clear
            @change="handleYearFilter"
          >
            <a-select-option
              v-for="year in availableYears"
              :key="year"
              :value="year"
            >
              {{ year }}年
            </a-select-option>
          </a-select>
        </div>
        
        <a-space>
          <a-button @click="fetchCertificates">
            <ReloadOutlined />
            刷新
          </a-button>
        </a-space>
      </div>
    </a-card>

    <!-- 证书列表 -->
    <a-card :title="listTitle" :loading="loading">
      <template #extra>
        <a-space>
          <span class="total-count">共 {{ filteredCertificates.length }} 条记录</span>
          <a-button @click="fetchCertificates" size="small" :loading="loading">
            <ReloadOutlined />
            刷新
          </a-button>
        </a-space>
      </template>

      <!-- 错误状态 -->
      <div v-if="!loading && certificates.length === 0" class="empty-state">
        <a-empty
          description="暂无证书数据"
          :image="Empty.PRESENTED_IMAGE_SIMPLE"
        >
          <template #description>
            <span>暂无证书数据，请检查网络连接或联系管理员</span>
          </template>
          <a-button type="primary" @click="fetchCertificates">
            重新加载
          </a-button>
        </a-empty>
      </div>

      <!-- 证书表格 -->
      <a-table
        v-else
        :columns="columns"
        :data-source="paginatedCertificates"
        :pagination="pagination"
        :scroll="{ x: 1300 }"
        row-key="id"
      >
        <template #bodyCell="{ column, record }: { column: any, record: any }">
          <!-- 证书信息 -->
          <template v-if="column.key === 'certificate'">
            <div class="certificate-cell">
              <div class="certificate-info">
                <div class="certificate-id">证书 #{{ record.id }}</div>
                <div class="batch-info">批次 #{{ record.batchId }}</div>
              </div>
            </div>
          </template>

          <!-- 关联项目 -->
          <template v-else-if="column.key === 'project'">
            <div class="project-cell">
              <div class="project-info">
                <div class="project-name">
                  <router-link :to="`/project/${record.projectId}`">
                    {{ getProjectName(record.projectId) }}
                  </router-link>
                </div>
                <div class="project-id">项目 #{{ record.projectId }}</div>
              </div>
            </div>
          </template>

          <!-- 所有者 -->
          <template v-else-if="column.key === 'owner'">
            <div class="owner-cell">
              <DefaultAvatar :name="getOwnerName(record.owner)" :size="32" />
              <div class="owner-info">
                <div class="owner-name">{{ getOwnerName(record.owner) }}</div>
                <div class="owner-address">{{ formatAddress(record.owner) }}</div>
              </div>
            </div>
          </template>

          <!-- 数量 -->
          <template v-else-if="column.key === 'quantity'">
            <span class="quantity-amount">{{ record.quantity.toLocaleString() }} 吨CO₂e</span>
          </template>

          <!-- 状态 -->
          <template v-else-if="column.key === 'status'">
            <a-tag :color="getCertificateStatusColor(record.status)">
              {{ getCertificateStatusText(record.status) }}
            </a-tag>
          </template>

          <!-- 核发时间 -->
          <template v-else-if="column.key === 'issuedAt'">
            {{ formatDate(record.timestamp) }}
          </template>

          <!-- 操作按钮 -->
          <template v-else-if="column.key === 'actions'">
            <a-space>
              <a-button size="small" @click="viewCertificateDetail(record)">
                详情
              </a-button>
              <a-button 
                size="small" 
                @click="viewRelatedProject(record)"
              >
                关联项目
              </a-button>
              <!-- 只有REGULATOR可以撤销证书 -->
              <a-button 
                v-if="canRevokeCertificate(record)"
                size="small"
                type="primary"
                danger
                @click="revokeCertificate(record)"
              >
                撤销证书
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 证书详情弹窗 -->
    <a-modal
      v-model:open="showDetailModal"
      title="证书详情"
      width="700px"
      :footer="null"
    >
      <CertificateDetailView v-if="selectedCertificate" :certificate="selectedCertificate" />
    </a-modal>

    <!-- 撤销证书确认弹窗 -->
    <EditReasonModal
      v-model:visible="showRevokeModal"
      title="撤销证书"
      :description="`确定要撤销证书 #${selectedCertificate?.id} 吗？撤销后无法恢复。`"
      placeholder="请说明撤销原因"
      @confirm="handleCertificateRevoke"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  FileTextOutlined,
  CheckCircleOutlined,
  StopOutlined,
  PlusCircleOutlined,
  ReloadOutlined
} from '@ant-design/icons-vue'
import { message, Empty } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { getAllCertificates, getAllProjects } from '@/services/api'
import { revokeCertificate } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import type { CertificateRecord, ProjectRecord } from '@/types'
import { CertificateStatus } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import CertificateDetailView from '@/components/CertificateDetailView.vue'
import EditReasonModal from '@/components/EditReasonModal.vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const certificates = ref<CertificateRecord[]>([])
const projects = ref<ProjectRecord[]>([])
const searchText = ref('')
const selectedStatus = ref<number | undefined>()
const selectedYear = ref<number | undefined>()
const showDetailModal = ref(false)
const showRevokeModal = ref(false)
const selectedCertificate = ref<CertificateRecord | null>(null)

// 页面标题和描述
const pageTitle = computed(() => {
  return PermissionService.isRegulator(currentUser.value) ? '证书管理' : '证书总览'
})

const pageDescription = computed(() => {
  return PermissionService.isRegulator(currentUser.value) 
    ? '管理平台上的所有证书' 
    : '查看平台上的所有证书'
})

const listTitle = computed(() => {
  return PermissionService.isRegulator(currentUser.value) ? '证书管理列表' : '证书总览列表'
})

// 统计数据
const totalCertificates = computed(() => certificates.value.length)

const issuedCertificates = computed(() => {
  return certificates.value.filter(cert => cert.status === 0).length // ISSUED = 0
})

const revokedCertificates = computed(() => {
  return certificates.value.filter(cert => cert.status === 1).length // REVOKED = 1
})

const monthlyNewCertificates = computed(() => {
  const currentMonth = dayjs().month()
  const currentYear = dayjs().year()
  return certificates.value.filter(cert => {
    const certDate = dayjs(cert.timestamp * 1000) // 使用timestamp而不是issuedAt
    return certDate.month() === currentMonth && certDate.year() === currentYear
  }).length
})

// 可用年份
const availableYears = computed(() => {
  const years = new Set(certificates.value.map(cert => dayjs(cert.timestamp * 1000).year()))
  return Array.from(years).sort((a, b) => b - a)
})

// 表格列配置
const columns = computed(() => {
  const baseColumns = [
    {
      title: '证书信息',
      key: 'certificate',
      width: 150,
      fixed: 'left' as const
    },
    {
      title: '关联项目',
      key: 'project',
      width: 200
    },
    {
      title: '所有者',
      key: 'owner',
      width: 180
    },
    {
      title: '数量',
      key: 'quantity',
      width: 120
    },
    {
      title: '状态',
      key: 'status',
      width: 100
    },
    {
      title: '核发时间',
      key: 'issuedAt',
      width: 150
    }
  ]

  // 如果是REGULATOR，添加操作列
  if (PermissionService.isRegulator(currentUser.value)) {
    baseColumns.push({
      title: '操作',
      key: 'actions',
      width: 200,
      fixed: 'right'
    })
  } else {
    baseColumns.push({
      title: '操作',
      key: 'actions',
      width: 150
    })
  }

  return baseColumns
})

// 过滤后的证书列表
const filteredCertificates = computed(() => {
  // 确保certificates.value是数组
  if (!Array.isArray(certificates.value)) {
    return []
  }
  
  let result = certificates.value

  // 按文本搜索
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    result = result.filter(cert => {
      if (!cert) return false
      try {
        return (cert.id && cert.id.toString().includes(search)) ||
               (cert.batchId && cert.batchId.toString().includes(search)) ||
               getProjectName(cert.projectId).toLowerCase().includes(search) ||
               getOwnerName(cert.owner).toLowerCase().includes(search) ||
               (cert.owner && cert.owner.toLowerCase().includes(search))
      } catch (error) {
        return false
      }
    })
  }

  // 按状态筛选
  if (selectedStatus.value !== undefined) {
    result = result.filter(cert => 
      cert && cert.status === selectedStatus.value
    )
  }

  // 按年份筛选
  if (selectedYear.value !== undefined) {
    result = result.filter(cert => {
      if (!cert || !cert.timestamp) return false
      try {
        return dayjs(cert.timestamp * 1000).year() === selectedYear.value
      } catch (error) {
        return false
      }
    })
  }

  // URL查询参数筛选（来自用户管理页面的跳转）
  if (route.query.user) {
    const userAddress = route.query.user as string
    result = result.filter(cert => 
      cert && cert.owner && 
      cert.owner.toLowerCase() === userAddress.toLowerCase()
    )
  }

  return result.sort((a, b) => {
    if (!a || !b) return 0
    return b.timestamp - a.timestamp
  })
})

// 分页配置
const pagination = ref({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showQuickJumper: true,
  showTotal: (total: number, range: [number, number]) =>
    `第 ${range[0]}-${range[1]} 条，共 ${total} 条`
})

// 分页后的证书列表
const paginatedCertificates = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize
  const end = start + pagination.value.pageSize
  pagination.value.total = filteredCertificates.value.length
  return filteredCertificates.value.slice(start, end)
})

// 获取证书列表
const fetchCertificates = async () => {
  try {
    loading.value = true
    const [certificateData, projectData] = await Promise.all([
      getAllCertificates(),
      getAllProjects()
    ])
    
    // 确保数据是数组格式
    certificates.value = Array.isArray(certificateData) ? certificateData : []
    projects.value = Array.isArray(projectData) ? projectData : []
  } catch (error: any) {
    console.error('获取证书列表失败:', error)
    // 设置空数组避免页面崩溃
    certificates.value = []
    projects.value = []
    
    // 根据错误类型显示不同消息
    if (error.code === 'NETWORK_ERROR' || error.message?.includes('Network Error')) {
      message.error('网络连接失败，请检查网络后重试')
    } else if (error.response?.status === 404) {
      message.error('后端服务未启动，请启动后端服务')
    } else {
      message.error('获取证书列表失败，请稍后重试')
    }
  } finally {
    loading.value = false
  }
}

// 搜索处理
const handleSearch = () => {
  pagination.value.current = 1
}

// 状态筛选处理
const handleStatusFilter = () => {
  pagination.value.current = 1
}

// 年份筛选处理
const handleYearFilter = () => {
  pagination.value.current = 1
}

// 查看证书详情
const viewCertificateDetail = (certificate: CertificateRecord) => {
  selectedCertificate.value = certificate
  showDetailModal.value = true
}

// 查看关联项目
const viewRelatedProject = (certificate: CertificateRecord) => {
  router.push(`/project/${certificate.projectId}`)
}

// 撤销证书
const revokeCertificate = (certificate: CertificateRecord) => {
  selectedCertificate.value = certificate
  showRevokeModal.value = true
}

// 处理证书撤销
const handleCertificateRevoke = async (reason: string) => {
  if (!selectedCertificate.value) return

  try {
    await revokeCertificate(selectedCertificate.value.id, reason)
    showRevokeModal.value = false
    fetchCertificates()
    message.success('证书撤销成功')
  } catch (error: any) {
    message.error(error.message || '证书撤销失败')
  }
}

// 检查是否可以撤销证书
const canRevokeCertificate = (certificate: CertificateRecord) => {
  return PermissionService.isRegulator(currentUser.value) && 
         certificate.status === 0 // ISSUED = 0
}

// 获取项目名称
const getProjectName = (projectId: number) => {
  const project = projects.value.find(p => p.id === projectId)
  return project?.name || `项目 #${projectId}`
}

// 获取用户名
const getOwnerName = (address: string) => {
  const user = userStore.getUserByAddress(address)
  return user?.name || ''
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm')
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

onMounted(() => {
  fetchCertificates()
})
</script>

<style scoped>
.certificate-overview {
  max-width: 1600px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 24px;
}

.page-header h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.page-header p {
  margin: 8px 0 0 0;
  color: #666;
}

.stats-cards {
  margin-bottom: 24px;
}

.action-card {
  margin-bottom: 16px;
}

.action-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-filters {
  display: flex;
  gap: 16px;
  align-items: center;
}

.total-count {
  color: #666;
  font-size: 14px;
}

.certificate-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.certificate-info {
  min-width: 0;
  flex: 1;
}

.certificate-id {
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.batch-info {
  font-size: 12px;
  color: #8c8c8c;
}

.project-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.project-info {
  min-width: 0;
  flex: 1;
}

.project-name {
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.project-name a {
  color: inherit;
  text-decoration: none;
}

.project-name a:hover {
  color: #1890ff;
}

.project-id {
  font-size: 12px;
  color: #8c8c8c;
}

.owner-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.owner-info {
  min-width: 0;
  flex: 1;
}

.owner-name {
  font-size: 14px;
  color: #262626;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.owner-address {
  font-size: 12px;
  color: #8c8c8c;
  font-family: monospace;
}

.quantity-amount {
  font-weight: 500;
  color: #52c41a;
}

:deep(.ant-table-tbody .ant-table-cell) {
  vertical-align: top;
  padding: 16px;
}

.empty-state {
  padding: 40px 0;
  text-align: center;
}

:deep(.ant-statistic-title) {
  font-size: 14px;
  color: #666;
}

:deep(.ant-statistic-content) {
  font-size: 24px;
  font-weight: 600;
}
</style>
