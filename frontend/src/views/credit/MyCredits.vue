<template>
  <div class="my-credits">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>我的碳信用</h1>
      <p>管理您持有的所有碳信用</p>
    </div>

    <!-- 权限检查 -->
    <template v-if="!canAccessMyCredits">
      <a-result
        status="403"
        title="无权限访问"
        sub-title="您需要项目方或购买方权限才能访问此页面"
      >
        <template #extra>
          <a-button type="primary" @click="router.push('/dashboard')">
            返回首页
          </a-button>
        </template>
      </a-result>
    </template>

    <template v-else>
      <!-- 统计信息卡片 -->
      <a-row :gutter="[16, 16]" class="stats-cards">
        <a-col :xs="24" :sm="6">
          <a-card>
            <a-statistic
              title="持有总量"
              :value="totalHoldings"
              suffix="吨CO₂e"
              :value-style="{ color: '#52c41a' }"
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
              title="可用数量"
              :value="availableHoldings"
              suffix="吨CO₂e"
              :value-style="{ color: '#1890ff' }"
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
              :value="retiredHoldings"
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
              title="持有批次数"
              :value="uniqueBatches"
              :value-style="{ color: '#fa8c16' }"
            >
              <template #prefix>
                <AppstoreOutlined />
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
              v-model:value="selectedStatus"
              placeholder="筛选状态"
              style="width: 120px"
              allow-clear
              @change="handleStatusFilter"
            >
              <a-select-option value="available">可用</a-select-option>
              <a-select-option value="retired">已退役</a-select-option>
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
            <a-button type="primary" @click="showTransferModal = true">
              <SwapOutlined />
              转让碳信用
            </a-button>
            <a-button @click="fetchMyCredits">
              <ReloadOutlined />
              刷新
            </a-button>
          </a-space>
        </div>
      </a-card>

      <!-- 我的碳信用列表 -->
      <a-card title="我的碳信用列表" :loading="loading">
        <template #extra>
          <a-space>
            <span class="total-count">共 {{ filteredCredits.length }} 条记录</span>
            <a-button @click="fetchMyCredits" size="small" :loading="loading">
              <ReloadOutlined />
              刷新
            </a-button>
          </a-space>
        </template>

        <!-- 错误状态 -->
        <div v-if="!loading && myCredits.length === 0" class="empty-state">
          <a-empty
            description="暂无碳信用数据"
            :image="Empty.PRESENTED_IMAGE_SIMPLE"
          >
            <template #description>
              <span>暂无碳信用数据，请检查网络连接或联系管理员</span>
            </template>
            <a-button type="primary" @click="fetchMyCredits">
              重新加载
            </a-button>
          </a-empty>
        </div>

        <!-- 过滤后无数据状态 -->
        <div v-else-if="!loading && filteredCredits.length === 0 && myCredits.length > 0" class="empty-state">
          <a-empty
            description="没有符合条件的碳信用记录"
            :image="Empty.PRESENTED_IMAGE_SIMPLE"
          >
            <template #description>
              <span>请尝试调整筛选条件</span>
            </template>
          </a-empty>
        </div>

        <!-- 碳信用表格 -->
        <a-table
          v-else
          :columns="columns"
          :data-source="paginatedCredits"
          :pagination="pagination"
          :scroll="{ x: 1200 }"
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

            <!-- 数量 -->
            <template v-else-if="column.key === 'quantity'">
              <span class="quantity-amount">{{ record.quantity.toLocaleString() }} 吨CO₂e</span>
            </template>

            <!-- 状态 -->
            <template v-else-if="column.key === 'status'">
              <a-tag :color="getCreditStatusColor(record)">
                {{ getCreditStatusText(record) }}
              </a-tag>
            </template>

            <!-- 获得方式 -->
            <template v-else-if="column.key === 'eventType'">
              <a-tag :color="getCreditEventColor(record.eventType)">
                {{ getCreditEventText(record.eventType) }}
              </a-tag>
            </template>

            <!-- 获得时间 -->
            <template v-else-if="column.key === 'timestamp'">
              {{ formatDate(record.timestamp) }}
            </template>

            <!-- 操作按钮 -->
            <template v-else-if="column.key === 'actions'">
              <a-space>
                <a-button size="small" @click="viewCreditDetail(record)">
                  详情
                </a-button>
                <a-button 
                  size="small" 
                  @click="viewRelatedProject(record)"
                >
                  关联项目
                </a-button>
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
        :available-credits="availableCreditsForTransfer"
        @success="handleTransferSuccess"
      />

      <!-- 退役碳信用确认弹窗 -->
      <EditReasonModal
        v-model:visible="showRetireModal"
        title="退役碳信用"
        :description="`确定要退役 ${selectedCredit?.quantity} 吨CO₂e 的碳信用吗？退役后无法撤销。`"
        placeholder="请说明退役原因（如用于抵消排放等）"
        @confirm="handleCreditRetire"
      />
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  GoldOutlined,
  ShoppingCartOutlined,
  DeleteOutlined,
  AppstoreOutlined,
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
import CreditDetailView from '@/components/CreditDetailView.vue'
import CreditTransferModal from '@/components/CreditTransferModal.vue'
import EditReasonModal from '@/components/EditReasonModal.vue'

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const myCredits = ref<CreditRecord[]>([])
const projects = ref<ProjectRecord[]>([])
const searchText = ref('')
const selectedStatus = ref<string | undefined>()
const selectedYear = ref<number | undefined>()
const showDetailModal = ref(false)
const showTransferModal = ref(false)
const showRetireModal = ref(false)
const selectedCredit = ref<CreditRecord | null>(null)

// 权限检查
const canAccessMyCredits = computed(() => 
  PermissionService.isProjectOwner(currentUser.value) || 
  PermissionService.isBuyer(currentUser.value)
)

// 统计数据
const totalHoldings = computed(() => {
  return myCredits.value.reduce((sum, credit) => sum + credit.quantity, 0)
})

const availableHoldings = computed(() => {
  return myCredits.value
    .filter(credit => credit.eventType !== CreditEventType.RETIRE && credit.eventType !== CreditEventType.REVOKE)
    .reduce((sum, credit) => sum + credit.quantity, 0)
})

const retiredHoldings = computed(() => {
  return myCredits.value
    .filter(credit => credit.eventType === CreditEventType.RETIRE)
    .reduce((sum, credit) => sum + credit.quantity, 0)
})

const uniqueBatches = computed(() => {
  const batchIds = new Set(myCredits.value.map(credit => getBatchId(credit)))
  return batchIds.size
})

// 可用年份
const availableYears = computed(() => {
  const years = new Set(myCredits.value.map(credit => credit.vintageYear))
  return Array.from(years).sort((a, b) => b - a)
})

// 可用于转让的碳信用
const availableCreditsForTransfer = computed(() => {
  return myCredits.value.filter(credit => 
    credit.eventType !== CreditEventType.RETIRE && 
    credit.eventType !== CreditEventType.REVOKE
  )
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
    title: '获得方式',
    key: 'eventType',
    width: 100
  },
  {
    title: '获得时间',
    key: 'timestamp',
    width: 150
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    fixed: 'right' as const
  }
]

// 过滤后的碳信用列表
const filteredCredits = computed(() => {
  // 确保myCredits.value是数组
  if (!Array.isArray(myCredits.value)) {
    return []
  }
  
  let result = myCredits.value

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

  // 按状态筛选
  if (selectedStatus.value) {
    if (selectedStatus.value === 'available') {
      result = result.filter(credit => 
        credit && credit.eventType !== CreditEventType.RETIRE && credit.eventType !== CreditEventType.REVOKE
      )
    } else if (selectedStatus.value === 'retired') {
      result = result.filter(credit => 
        credit && credit.eventType === CreditEventType.RETIRE
      )
    }
  }

  // 按年份筛选
  if (selectedYear.value !== undefined) {
    result = result.filter(credit => 
      credit && credit.vintageYear === selectedYear.value
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

// 获取我的碳信用列表
const fetchMyCredits = async () => {
  if (!walletInfo.address) {
    myCredits.value = []
    projects.value = []
    return
  }

  try {
    loading.value = true
    const [creditData, projectData] = await Promise.all([
      getAllCredits(),
      getAllProjects()
    ])
    
    // 确保数据是数组格式
    const allCredits = Array.isArray(creditData) ? creditData : []
    const allProjects = Array.isArray(projectData) ? projectData : []
    
    // 筛选当前用户持有的碳信用，添加安全检查
    myCredits.value = allCredits.filter(credit => {
      if (!credit || !credit.currentOwner || !walletInfo.address) return false
      return credit.currentOwner.toLowerCase() === walletInfo.address.toLowerCase()
    })
    projects.value = allProjects
  } catch (error: any) {
    console.error('获取我的碳信用失败:', error)
    // 设置空数组避免页面崩溃
    myCredits.value = []
    projects.value = []
    
    // 根据错误类型显示不同消息
    if (error.code === 'NETWORK_ERROR' || error.message?.includes('Network Error')) {
      message.error('网络连接失败，请检查网络后重试')
    } else if (error.response?.status === 404) {
      message.error('后端服务未启动，请启动后端服务')
    } else {
      message.error('获取我的碳信用失败，请稍后重试')
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

// 查看碳信用详情
const viewCreditDetail = (credit: CreditRecord) => {
  selectedCredit.value = credit
  showDetailModal.value = true
}

// 查看关联项目
const viewRelatedProject = (credit: CreditRecord) => {
  router.push(`/project/${credit.projectId}`)
}

// 退役碳信用
const retireCredit = (credit: CreditRecord) => {
  selectedCredit.value = credit
  showRetireModal.value = true
}

// 转让成功处理
const handleTransferSuccess = () => {
  showTransferModal.value = false
  fetchMyCredits()
  message.success('碳信用转让成功')
}

// 处理碳信用退役
const handleCreditRetire = async (reason: string) => {
  if (!selectedCredit.value) return

  try {
    await retireCredits(getBatchId(selectedCredit.value), selectedCredit.value.quantity, reason)
    showRetireModal.value = false
    fetchMyCredits()
    message.success('碳信用退役成功')
  } catch (error: any) {
    message.error(error.message || '碳信用退役失败')
  }
}

// 检查是否可以退役碳信用
const canRetireCredit = (credit: CreditRecord) => {
  return credit.eventType !== CreditEventType.RETIRE && 
         credit.eventType !== CreditEventType.REVOKE
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

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm')
}

const getCreditStatusColor = (credit: CreditRecord) => {
  if (credit.eventType === CreditEventType.RETIRE) return 'purple'
  if (credit.eventType === CreditEventType.REVOKE) return 'red'
  return 'green'
}

const getCreditStatusText = (credit: CreditRecord) => {
  if (credit.eventType === CreditEventType.RETIRE) return '已退役'
  if (credit.eventType === CreditEventType.REVOKE) return '已撤销'
  return '可用'
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
    [CreditEventType.MINT]: '铸造获得',
    [CreditEventType.TRANSFER]: '转让获得',
    [CreditEventType.RETIRE]: '已退役',
    [CreditEventType.REVOKE]: '已撤销'
  }
  return texts[eventType] || '未知'
}

onMounted(() => {
  fetchMyCredits()
})
</script>

<style scoped>
.my-credits {
  max-width: 1400px;
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

.empty-state {
  padding: 60px 0;
  text-align: center;
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

:deep(.ant-table-tbody .ant-table-cell) {
  vertical-align: top;
  padding: 16px;
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