<template>
  <div class="my-audits">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>我的审核</h1>
      <p>管理您的所有审核任务</p>
    </div>

    <!-- 权限检查 -->
    <template v-if="!PermissionService.isVerifier(currentUser)">
      <a-result
        status="403"
        title="无权限访问"
        sub-title="您需要审核员权限才能访问此页面"
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
              title="待处理任务"
              :value="pendingTasks.length"
              :value-style="{ color: '#fa8c16' }"
            >
              <template #prefix>
                <ClockCircleOutlined />
              </template>
            </a-statistic>
          </a-card>
        </a-col>
        
        <a-col :xs="24" :sm="6">
          <a-card>
            <a-statistic
              title="已完成任务"
              :value="completedTasks.length"
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
              title="通过率"
              :value="approvalRate"
              suffix="%"
              :value-style="{ color: '#1890ff' }"
            >
              <template #prefix>
                <TrophyOutlined />
              </template>
            </a-statistic>
          </a-card>
        </a-col>
        
        <a-col :xs="24" :sm="6">
          <a-card>
            <a-statistic
              title="本月审核数"
              :value="monthlyCount"
              :value-style="{ color: '#722ed1' }"
            >
              <template #prefix>
                <CalendarOutlined />
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
              placeholder="搜索项目名称或ID"
              style="width: 300px"
              @search="handleSearch"
              allow-clear
            />
            <a-select
              v-model:value="selectedStatus"
              placeholder="筛选状态"
              style="width: 150px"
              allow-clear
              @change="handleStatusFilter"
            >
              <a-select-option value="pending">待处理</a-select-option>
              <a-select-option value="completed">已完成</a-select-option>
            </a-select>
          </div>
          
          <a-button @click="fetchMyTasks">
            <ReloadOutlined />
            刷新
          </a-button>
        </div>
      </a-card>

      <!-- 审核任务列表 -->
      <a-card title="审核任务列表" :loading="loading">
        <template #extra>
          <span class="total-count">共 {{ filteredTasks.length }} 个任务</span>
        </template>

        <a-table
          :columns="columns"
          :data-source="paginatedTasks"
          :pagination="pagination"
          :scroll="{ x: 1200 }"
          row-key="id"
        >
          <template #bodyCell="{ column, record }: { column: any, record: any }">
            <!-- 项目信息 -->
            <template v-if="column.key === 'project'">
              <div class="project-cell">
                <div class="project-info">
                  <div class="project-name">
                    <router-link :to="`/project/${record.project.id}`">
                      {{ record.project.name }}
                    </router-link>
                  </div>
                  <div class="project-id">#{{ record.project.id }}</div>
                  <div class="project-location">{{ record.project.location }}</div>
                </div>
              </div>
            </template>

            <!-- 项目所有者 -->
            <template v-else-if="column.key === 'owner'">
              <div class="owner-cell">
                <DefaultAvatar :name="getOwnerName(record.project.owner)" :size="32" />
                <div class="owner-info">
                  <div class="owner-name">{{ getOwnerName(record.project.owner) }}</div>
                  <div class="owner-address">{{ formatAddress(record.project.owner) }}</div>
                </div>
              </div>
            </template>

            <!-- 减排量 -->
            <template v-else-if="column.key === 'reduction'">
              <span class="reduction-amount">{{ record.project.reduction.toLocaleString() }} 吨CO₂e</span>
            </template>

            <!-- 任务状态 -->
            <template v-else-if="column.key === 'status'">
              <a-tag :color="getTaskStatusColor(record.status)">
                {{ getTaskStatusText(record.status) }}
              </a-tag>
            </template>

            <!-- 审核结果 -->
            <template v-else-if="column.key === 'result'">
              <div v-if="record.review">
                <a-tag :color="record.review.approved ? 'success' : 'error'">
                  {{ record.review.approved ? '通过' : '不通过' }}
                </a-tag>
              </div>
              <span v-else class="pending-review">-</span>
            </template>

            <!-- 接受时间 -->
            <template v-else-if="column.key === 'acceptTime'">
              {{ formatDate(record.acceptTime) }}
            </template>

            <!-- 操作按钮 -->
            <template v-else-if="column.key === 'actions'">
              <a-space>
                <a-button size="small" @click="viewTaskDetail(record)">
                  详情
                </a-button>
                <a-button 
                  v-if="record.status === 'pending'"
                  type="primary"
                  size="small"
                  @click="submitReviewOpinion(record)"
                >
                  提交审核意见
                </a-button>
                <a-button 
                  v-else-if="record.review"
                  size="small"
                  @click="viewReviewDetail(record)"
                >
                  查看审核详情
                </a-button>
              </a-space>
            </template>
          </template>
        </a-table>
      </a-card>

      <!-- 任务详情弹窗 -->
      <a-modal
        v-model:open="showDetailModal"
        title="审核任务详情"
        width="800px"
        :footer="null"
      >
        <ProjectDetailView v-if="selectedTask" :project="selectedTask.project" />
      </a-modal>

      <!-- 提交审核意见弹窗 -->
      <ReviewOpinionModal
        v-model:visible="showReviewModal"
        :task="selectedTask"
        @success="handleReviewSubmitted"
      />

      <!-- 审核详情弹窗 -->
      <a-modal
        v-model:open="showReviewDetailModal"
        title="审核详情"
        width="600px"
        :footer="null"
      >
        <ReviewDetailView v-if="selectedTask" :review="selectedTask.review" />
      </a-modal>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  ClockCircleOutlined,
  CheckCircleOutlined,
  TrophyOutlined,
  CalendarOutlined,
  ReloadOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
// 注意：getMyReviewTasks 已从 API 中移除，因为审核任务应该通过监听区块链事件获取
// 这里可能需要重新设计审核任务获取逻辑
import { getAllProjects } from '@/services/api'
import { PermissionService } from '@/services/permission'
import { walletInfo } from '@/services/web3'
import type { ProjectRecord, ReviewRecord } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProjectDetailView from '@/components/ProjectDetailView.vue'
import ReviewOpinionModal from '@/components/ReviewOpinionModal.vue'
import ReviewDetailView from '@/components/ReviewDetailView.vue'

// 审核任务接口
interface ReviewTask {
  id: string
  project: ProjectRecord
  status: 'pending' | 'completed'
  acceptTime: number
  review?: ReviewRecord
}

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const tasks = ref<ReviewTask[]>([])
const searchText = ref('')
const selectedStatus = ref<string | undefined>()
const showDetailModal = ref(false)
const showReviewModal = ref(false)
const showReviewDetailModal = ref(false)
const selectedTask = ref<ReviewTask | null>(null)

// 计算属性
const pendingTasks = computed(() => tasks.value.filter(task => task.status === 'pending'))
const completedTasks = computed(() => tasks.value.filter(task => task.status === 'completed'))

const approvalRate = computed(() => {
  const completed = completedTasks.value
  if (completed.length === 0) return 0
  const approved = completed.filter(task => task.review?.approved).length
  return Math.round((approved / completed.length) * 100)
})

const monthlyCount = computed(() => {
  const currentMonth = dayjs().month()
  const currentYear = dayjs().year()
  return tasks.value.filter(task => {
    const taskDate = dayjs(task.acceptTime * 1000)
    return taskDate.month() === currentMonth && taskDate.year() === currentYear
  }).length
})

// 表格列配置
const columns = [
  {
    title: '项目信息',
    key: 'project',
    width: 250,
    fixed: 'left' as const
  },
  {
    title: '项目所有者',
    key: 'owner',
    width: 180
  },
  {
    title: '减排量',
    key: 'reduction',
    width: 150
  },
  {
    title: '任务状态',
    key: 'status',
    width: 100
  },
  {
    title: '审核结果',
    key: 'result',
    width: 100
  },
  {
    title: '接受时间',
    key: 'acceptTime',
    width: 150
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    fixed: 'right' as const
  }
]

// 过滤后的任务列表
const filteredTasks = computed(() => {
  let result = tasks.value

  // 按文本搜索
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    result = result.filter(task =>
      task.project.name.toLowerCase().includes(search) ||
      task.project.id.toString().includes(search)
    )
  }

  // 按状态筛选
  if (selectedStatus.value) {
    result = result.filter(task => task.status === selectedStatus.value)
  }

  return result
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

// 分页后的任务列表
const paginatedTasks = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize
  const end = start + pagination.value.pageSize
  pagination.value.total = filteredTasks.value.length
  return filteredTasks.value.slice(start, end)
})

// 获取我的审核任务
const fetchMyTasks = async () => {
  if (!walletInfo.address) return

  try {
    loading.value = true
    // 这里应该调用具体的API获取审核任务
    // 暂时使用模拟数据
    const mockTasks: ReviewTask[] = []
    tasks.value = mockTasks
  } catch (error) {
    console.error('获取审核任务失败:', error)
    message.error('获取审核任务失败')
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

// 查看任务详情
const viewTaskDetail = (task: ReviewTask) => {
  selectedTask.value = task
  showDetailModal.value = true
}

// 提交审核意见
const submitReviewOpinion = (task: ReviewTask) => {
  selectedTask.value = task
  showReviewModal.value = true
}

// 查看审核详情
const viewReviewDetail = (task: ReviewTask) => {
  selectedTask.value = task
  showReviewDetailModal.value = true
}

// 审核提交成功处理
const handleReviewSubmitted = () => {
  showReviewModal.value = false
  fetchMyTasks()
  message.success('审核意见提交成功')
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

const getTaskStatusColor = (status: string) => {
  const colors = {
    pending: 'orange',
    completed: 'green'
  }
  return colors[status as keyof typeof colors] || 'default'
}

const getTaskStatusText = (status: string) => {
  const texts = {
    pending: '待处理',
    completed: '已完成'
  }
  return texts[status as keyof typeof texts] || '未知'
}

onMounted(() => {
  fetchMyTasks()
})
</script>

<style scoped>
.my-audits {
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
  font-family: monospace;
  margin-bottom: 2px;
}

.project-location {
  font-size: 12px;
  color: #999;
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

.reduction-amount {
  font-weight: 500;
  color: #52c41a;
}

.pending-review {
  color: #999;
  font-style: italic;
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
