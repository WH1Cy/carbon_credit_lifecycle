<template>
  <div class="credit-market">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>碳信用交易市场</h1>
      <p>查看和交易平台上的所有碳信用</p>
    </div>

    <!-- 统计信息卡片 -->
    <a-row :gutter="[16, 16]" class="stats-cards">
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="总碳信用批次"
            :value="totalBatches"
            :value-style="{ color: '#1890ff' }"
          >
            <template #prefix>
              <GoldOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="可交易数量"
            :value="availableCredits"
            suffix="吨CO₂e"
            :value-style="{ color: '#52c41a' }"
          >
            <template #prefix>
              <ShoppingCartOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="已退役数量"
            :value="retiredCredits"
            suffix="吨CO₂e"
            :value-style="{ color: '#722ed1' }"
          >
            <template #prefix>
              <DeleteOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      
      <a-col :xs="24" :sm="6">
        <a-card>
          <a-statistic
            title="本月交易量"
            :value="monthlyVolume"
            suffix="吨CO₂e"
            :value-style="{ color: '#fa8c16' }"
          >
            <template #prefix>
              <BarChartOutlined />
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
            placeholder="搜索项目名称或批次ID"
            style="width: 300px"
            @search="handleSearch"
            allow-clear
          />
          <a-select
            v-model:value="selectedEventType"
            placeholder="筛选事件类型"
            style="width: 150px"
            allow-clear
            @change="handleEventTypeFilter"
          >
            <a-select-option :value="0">铸造</a-select-option>
            <a-select-option :value="1">转让</a-select-option>
            <a-select-option :value="2">退役</a-select-option>
            <a-select-option :value="3">撤销</a-select-option>
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
        
        <div v-if="canTransfer">
          <a-button type="primary" @click="showTransferModal = true">
            <SwapOutlined />
            转让碳信用
          </a-button>
        </div>
      </div>
    </a-card>

    <!-- 碳信用交易列表 -->
    <a-card title="碳信用交易记录" :loading="loading">
      <template #extra>
        <a-space>
          <span class="total-count">共 {{ filteredCredits.length }} 条记录</span>
          <a-button @click="fetchCredits" size="small" :loading="loading">
            <ReloadOutlined />
            刷新
          </a-button>
        </a-space>
      </template>

      <!-- 错误状态 -->
      <div v-if="!loading && credits.length === 0" class="empty-state">
        <a-empty
          description="暂无碳信用数据"
          :image="Empty.PRESENTED_IMAGE_SIMPLE"
        >
          <template #description>
            <span>暂无碳信用数据，请检查网络连接或联系管理员</span>
          </template>
          <a-button type="primary" @click="fetchCredits">
            重新加载
          </a-button>
        </a-empty>
      </div>

      <!-- 碳信用表格 -->
      <a-table
        v-else
        :columns="columns"
        :data-source="paginatedCredits"
        :pagination="pagination"
        :scroll="{ x: 1400 }"
        row-key="id"
      >
        <template #bodyCell="{ column, record }: { column: any, record: CreditRecord }">
          <!-- 项目信息 -->
          <template v-if="column.key === 'project'">
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

          <!-- 批次信息 -->
          <template v-else-if="column.key === 'batch'">
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
              <DefaultAvatar :name="getOwnerName(record.currentOwner)" :size="32" />
              <div class="owner-info">
                <div class="owner-name">{{ getOwnerName(record.currentOwner) }}</div>
                <div class="owner-address">{{ formatAddress(record.currentOwner) }}</div>
              </div>
            </div>
          </template>

          <!-- 交易信息 -->
          <template v-else-if="column.key === 'transfer'">
            <div v-if="record.eventType === CreditEventType.TRANSFER" class="transfer-info">
              <div class="transfer-from">
                <span class="label">转出方：</span>
                <span class="address">{{ formatAddress(record.from) }}</span>
              </div>
              <div class="transfer-to">
                <span class="label">转入方：</span>
                <span class="address">{{ formatAddress(record.to) }}</span>
              </div>
            </div>
            <span v-else class="no-transfer">-</span>
          </template>

          <!-- 时间 -->
          <template v-else-if="column.key === 'timestamp'">
            {{ formatDate(record.timestamp) }}
          </template>

          <!-- 操作按钮 -->
          <template v-else-if="column.key === 'actions'">
            <a-space>
              <a-button size="small" @click="viewCreditDetail(record)">
                详情
              </a-button>
              <!-- 只有当前所有者且是PROJECT_OWNER或BUYER才能退役 -->
              <a-button 
                v-if="canRetireCredit(record)"
                size="small"
                type="primary"
                danger
                @click="retireCredit(record)"
              >
                退役
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 碳信用详情弹窗 -->
    <a-modal
      v-model:open="showDetailModal"
      title="碳信用详情"
      width="700px"
      :footer="null"
    >
      <CreditDetailView v-if="selectedCredit" :credit="selectedCredit" />
    </a-modal>

    <!-- 转让碳信用弹窗 -->
    <CreditTransferModal
      v-model:visible="showTransferModal"
      @success="handleTransferSuccess"
    />

    <!-- 退役碳信用确认弹窗 -->
    <EditReasonModal
      v-model:visible="showRetireModal"
      title="退役碳信用"
      :description="`确定要退役 ${selectedCredit?.quantity} 吨CO₂e 的碳信用吗？退役后无法撤销。`"
      placeholder="请说明退役原因"
      @confirm="handleCreditRetire"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  GoldOutlined,
  ShoppingCartOutlined,
  DeleteOutlined,
  BarChartOutlined,
  SwapOutlined,
  ReloadOutlined
} from '@ant-design/icons-vue'
import { message, Empty } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { getAllCredits, getAllProjects } from '@/services/api'
import { retireCredits } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import { walletInfo } from '@/services/web3'
import type { CreditRecord, ProjectRecord } from '@/types'
import { CreditEventType } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import CreditDetailView from '@/components/CreditDetailView.vue'
import CreditTransferModal from '@/components/CreditTransferModal.vue'
import EditReasonModal from '@/components/EditReasonModal.vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const credits = ref<CreditRecord[]>([])
const projects = ref<ProjectRecord[]>([])
const searchText = ref('')
const selectedEventType = ref<number | undefined>()
const selectedYear = ref<number | undefined>()
const showDetailModal = ref(false)
const showTransferModal = ref(false)
const showRetireModal = ref(false)
const selectedCredit = ref<CreditRecord | null>(null)

// 权限检查
const canTransfer = computed(() => 
  PermissionService.isProjectOwner(currentUser.value) || 
  PermissionService.isBuyer(currentUser.value)
)

// 统计数据
const totalBatches = computed(() => {
  const batchIds = new Set(credits.value.map(credit => getBatchId(credit)))
  return batchIds.size
})

const availableCredits = computed(() => {
  return credits.value
    .filter(credit => credit.eventType !== CreditEventType.RETIRE && credit.eventType !== CreditEventType.REVOKE)
    .reduce((sum, credit) => sum + credit.quantity, 0)
})

const retiredCredits = computed(() => {
  return credits.value
    .filter(credit => credit.eventType === CreditEventType.RETIRE)
    .reduce((sum, credit) => sum + credit.quantity, 0)
})

const monthlyVolume = computed(() => {
  const currentMonth = dayjs().month()
  const currentYear = dayjs().year()
  return credits.value
    .filter(credit => {
      const creditDate = dayjs(credit.timestamp * 1000)
      return creditDate.month() === currentMonth && 
             creditDate.year() === currentYear &&
             credit.eventType === CreditEventType.TRANSFER
    })
    .reduce((sum, credit) => sum + credit.quantity, 0)
})

// 可用年份
const availableYears = computed(() => {
  const years = new Set(credits.value.map(credit => credit.vintageYear))
  return Array.from(years).sort((a, b) => b - a)
})

// 表格列配置
const columns = [
  {
    title: '关联项目',
    key: 'project',
    width: 200,
    fixed: 'left' as const
  },
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
    title: '交易信息',
    key: 'transfer',
    width: 200
  },
  {
    title: '时间',
    key: 'timestamp',
    width: 150
  },
  {
    title: '操作',
    key: 'actions',
    width: 150,
    fixed: 'right' as const
  }
]

// 过滤后的碳信用列表
const filteredCredits = computed(() => {
  // 确保credits.value是数组
  if (!Array.isArray(credits.value)) {
    return []
  }
  
  let result = credits.value

  // 按文本搜索
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    result = result.filter(credit => {
      if (!credit) return false
      try {
        return getProjectName(credit.projectId).toLowerCase().includes(search) ||
               (credit.projectId && credit.projectId.toString().includes(search)) ||
               getBatchId(credit).toString().includes(search)
      } catch (error) {
        return false
      }
    })
  }

  // 按事件类型筛选
  if (selectedEventType.value !== undefined) {
    result = result.filter(credit => 
      credit && credit.eventType === selectedEventType.value
    )
  }

  // 按年份筛选
  if (selectedYear.value !== undefined) {
    result = result.filter(credit => 
      credit && credit.vintageYear === selectedYear.value
    )
  }

  // URL查询参数筛选（来自用户管理页面的跳转）
  if (route.query.user) {
    const userAddress = route.query.user as string
    result = result.filter(credit => 
      credit && credit.currentOwner && 
      credit.currentOwner.toLowerCase() === userAddress.toLowerCase()
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

// 分页后的碳信用列表
const paginatedCredits = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize
  const end = start + pagination.value.pageSize
  pagination.value.total = filteredCredits.value.length
  return filteredCredits.value.slice(start, end)
})

// 获取碳信用列表
const fetchCredits = async () => {
  try {
    loading.value = true
    const [creditData, projectData] = await Promise.all([
      getAllCredits(),
      getAllProjects()
    ])
    
    // 确保数据是数组格式
    credits.value = Array.isArray(creditData) ? creditData : []
    projects.value = Array.isArray(projectData) ? projectData : []
  } catch (error: any) {
    console.error('获取碳信用列表失败:', error)
    // 设置空数组避免页面崩溃
    credits.value = []
    projects.value = []
    
    // 根据错误类型显示不同消息
    if (error.code === 'NETWORK_ERROR' || error.message?.includes('Network Error')) {
      message.error('网络连接失败，请检查网络后重试')
    } else if (error.response?.status === 404) {
      message.error('后端服务未启动，请启动后端服务')
    } else {
      message.error('获取碳信用列表失败，请稍后重试')
    }
  } finally {
    loading.value = false
  }
}

// 搜索处理
const handleSearch = () => {
  pagination.value.current = 1
}

// 事件类型筛选处理
const handleEventTypeFilter = () => {
  pagination.value.current = 1
}

// 年份筛选处理
const handleYearFilter = () => {
  pagination.value.current = 1
}

// 查看碳信用详情
const viewCreditDetail = (credit: CreditRecord) => {
  selectedCredit.value = credit
  showDetailModal.value = true
}

// 退役碳信用
const retireCredit = (credit: CreditRecord) => {
  selectedCredit.value = credit
  showRetireModal.value = true
}

// 转让成功处理
const handleTransferSuccess = () => {
  showTransferModal.value = false
  fetchCredits()
  message.success('碳信用转让成功')
}

// 处理碳信用退役
const handleCreditRetire = async (reason: string) => {
  if (!selectedCredit.value) return

  try {
    await retireCredits(getBatchId(selectedCredit.value), selectedCredit.value.quantity, reason)
    showRetireModal.value = false
    fetchCredits()
    message.success('碳信用退役成功')
  } catch (error: any) {
    message.error(error.message || '碳信用退役失败')
  }
}

// 检查是否可以退役碳信用
const canRetireCredit = (credit: CreditRecord) => {
  if (!walletInfo.address || !currentUser.value) return false
  
  // 必须是当前所有者
  if (credit.currentOwner.toLowerCase() !== walletInfo.address.toLowerCase()) return false
  
  // 必须是PROJECT_OWNER或BUYER
  if (!PermissionService.isProjectOwner(currentUser.value) && !PermissionService.isBuyer(currentUser.value)) return false
  
  // 不能是已退役或已撤销的
  if (credit.eventType === CreditEventType.RETIRE || credit.eventType === CreditEventType.REVOKE) return false
  
  return true
}

// 获取项目名称
const getProjectName = (projectId: number) => {
  const project = projects.value.find(p => p.id === projectId)
  return project?.name || `项目 #${projectId}`
}

// 获取批次ID
const getBatchId = (credit: CreditRecord) => {
  return credit.childBatchId || credit.parentBatchId || 1
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
  fetchCredits()
})
</script>

<style scoped>
.credit-market {
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

.quantity-amount {
  font-weight: 500;
  color: #52c41a;
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

.transfer-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.transfer-from,
.transfer-to {
  font-size: 12px;
}

.label {
  color: #666;
  margin-right: 4px;
}

.address {
  font-family: monospace;
  color: #262626;
}

.no-transfer {
  color: #999;
  font-style: italic;
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
