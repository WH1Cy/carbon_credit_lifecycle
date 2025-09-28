<template>
  <div class="audit-market">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>审核市场</h1>
      <p>接受和处理待审核的项目</p>
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
        <a-col :xs="24" :sm="8">
          <a-card>
            <a-statistic
              title="待审核项目"
              :value="pendingProjects.length"
              :value-style="{ color: '#fa8c16' }"
            >
              <template #prefix>
                <ClockCircleOutlined />
              </template>
            </a-statistic>
          </a-card>
        </a-col>
        
        <a-col :xs="24" :sm="8">
          <a-card>
            <a-statistic
              title="我的审核任务"
              :value="myReviewTasks.length"
              :value-style="{ color: '#1890ff' }"
            >
              <template #prefix>
                <AuditOutlined />
              </template>
            </a-statistic>
          </a-card>
        </a-col>
        
        <a-col :xs="24" :sm="8">
          <a-card>
            <a-statistic
              title="本月审核数"
              :value="monthlyReviewCount"
              :value-style="{ color: '#52c41a' }"
            >
              <template #prefix>
                <CheckCircleOutlined />
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
              v-model:value="selectedTechnology"
              placeholder="筛选技术类型"
              style="width: 150px"
              allow-clear
              @change="handleTechnologyFilter"
            >
              <a-select-option :value="0">太阳能</a-select-option>
              <a-select-option :value="1">风能</a-select-option>
              <a-select-option :value="2">水电</a-select-option>
              <a-select-option :value="3">林业</a-select-option>
              <a-select-option :value="4">碳捕获</a-select-option>
              <a-select-option :value="5">其他</a-select-option>
            </a-select>
          </div>
          
          <a-button @click="fetchPendingProjects">
            <ReloadOutlined />
            刷新
          </a-button>
        </div>
      </a-card>

      <!-- 待审核项目列表 -->
      <a-card title="待审核项目" :loading="loading">
        <template #extra>
          <span class="total-count">共 {{ filteredProjects.length }} 个待审核项目</span>
        </template>

        <div v-if="filteredProjects.length === 0" class="empty-state">
          <a-empty description="暂无待审核项目" />
        </div>

        <div v-else class="project-grid">
          <div
            v-for="project in paginatedProjects"
            :key="project.id"
            class="project-card"
          >
            <a-card
              :title="project.name"
              size="small"
              :actions="[
                { key: 'detail', icon: 'EyeOutlined', title: '查看详情' },
                { key: 'accept', icon: 'CheckOutlined', title: '接受审核' }
              ]"
            >
              <template #extra>
                <a-tag color="orange">待审核</a-tag>
              </template>

              <div class="project-content">
                <div class="project-info">
                  <div class="info-item">
                    <span class="label">项目ID：</span>
                    <span class="value">#{{ project.id }}</span>
                  </div>
                  <div class="info-item">
                    <span class="label">位置：</span>
                    <span class="value">{{ project.location }}</span>
                  </div>
                  <div class="info-item">
                    <span class="label">减排量：</span>
                    <span class="value reduction">{{ project.reduction.toLocaleString() }} 吨CO₂e</span>
                  </div>
                  <div class="info-item">
                    <span class="label">项目方：</span>
                    <div class="owner-info">
                      <DefaultAvatar :name="getOwnerName(project.owner)" :size="20" />
                      <span class="owner-name">{{ getOwnerName(project.owner) || formatAddress(project.owner) }}</span>
                    </div>
                  </div>
                </div>

                <div class="project-tech">
                  <div class="label">技术类型：</div>
                  <a-space wrap size="small">
                    <a-tag
                      v-for="tech in project.technologies"
                      :key="tech"
                      :color="getTechnologyColor(tech)"
                      size="small"
                    >
                      {{ getTechnologyText(tech) }}
                    </a-tag>
                  </a-space>
                </div>

                <div class="project-time">
                  <span class="submit-time">提交时间：{{ formatDate(project.timestamp) }}</span>
                </div>
              </div>

              <template #actions>
                <a-button 
                  type="text" 
                  size="small"
                  @click="viewProjectDetail(project)"
                >
                  <EyeOutlined />
                  查看详情
                </a-button>
                <a-button 
                  type="text" 
                  size="small"
                  @click="acceptReviewTask(project)"
                  :loading="acceptingTasks[project.id]"
                >
                  <CheckOutlined />
                  接受审核
                </a-button>
              </template>
            </a-card>
          </div>
        </div>

        <!-- 分页 -->
        <div v-if="filteredProjects.length > 0" class="pagination-wrapper">
          <a-pagination
            v-model:current="pagination.current"
            v-model:page-size="pagination.pageSize"
            :total="pagination.total"
            :show-total="(total, range) => `第 ${range[0]}-${range[1]} 条，共 ${total} 条`"
            show-size-changer
            show-quick-jumper
          />
        </div>
      </a-card>

      <!-- 项目详情弹窗 -->
      <a-modal
        v-model:open="showDetailModal"
        title="项目详情"
        width="800px"
        :footer="null"
      >
        <ProjectDetailView v-if="selectedProject" :project="selectedProject" />
      </a-modal>

      <!-- 接受审核任务确认弹窗 -->
      <a-modal
        v-model:open="showAcceptModal"
        title="接受审核任务"
        width="500px"
        @ok="handleAcceptTask"
        @cancel="showAcceptModal = false"
        :confirm-loading="acceptLoading"
      >
        <div class="accept-confirm">
          <a-alert
            message="审核任务确认"
            description="接受审核任务后，您将成为该项目的指定审核员，需要在规定时间内完成审核工作。"
            type="info"
            show-icon
            style="margin-bottom: 16px"
          />
          <p>确定要接受项目 <strong>{{ selectedProject?.name }}</strong> 的审核任务吗？</p>
        </div>
      </a-modal>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  ClockCircleOutlined,
  AuditOutlined,
  CheckCircleOutlined,
  ReloadOutlined,
  EyeOutlined,
  CheckOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { getAllProjects } from '@/services/api'
import { acceptVerification } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import { walletInfo } from '@/services/web3'
import type { ProjectRecord } from '@/types'
import { ProjectStatus, Technology } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProjectDetailView from '@/components/ProjectDetailView.vue'

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const acceptLoading = ref(false)
const pendingProjects = ref<ProjectRecord[]>([])
const myReviewTasks = ref<ProjectRecord[]>([])
const monthlyReviewCount = ref(0)
const searchText = ref('')
const selectedTechnology = ref<number | undefined>()
const showDetailModal = ref(false)
const showAcceptModal = ref(false)
const selectedProject = ref<ProjectRecord | null>(null)
const acceptingTasks = ref<Record<number, boolean>>({})

// 过滤后的项目列表
const filteredProjects = computed(() => {
  let result = pendingProjects.value

  // 按文本搜索
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    result = result.filter(project =>
      project.name.toLowerCase().includes(search) ||
      project.id.toString().includes(search)
    )
  }

  // 按技术类型筛选
  if (selectedTechnology.value !== undefined) {
    result = result.filter(project => 
      project.technologies.includes(selectedTechnology.value!)
    )
  }

  return result
})

// 分页配置
const pagination = ref({
  current: 1,
  pageSize: 12,
  total: 0
})

// 分页后的项目列表
const paginatedProjects = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize
  const end = start + pagination.value.pageSize
  pagination.value.total = filteredProjects.value.length
  return filteredProjects.value.slice(start, end)
})

// 获取待审核项目列表
const fetchPendingProjects = async () => {
  try {
    loading.value = true
    const allProjects = await getAllProjects()
    // 筛选状态为"提交送审"的项目
    pendingProjects.value = allProjects.filter(project => 
      project.status === ProjectStatus.REVIEWING
    )
  } catch (error) {
    console.error('获取待审核项目失败:', error)
    message.error('获取待审核项目失败')
  } finally {
    loading.value = false
  }
}

// 获取我的审核任务
const fetchMyReviewTasks = async () => {
  if (!walletInfo.address) return

  try {
    myReviewTasks.value = await getMyReviewTasks(walletInfo.address)
    
    // 计算本月审核数
    const currentMonth = dayjs().month()
    const currentYear = dayjs().year()
    monthlyReviewCount.value = myReviewTasks.value.filter(task => {
      const taskDate = dayjs(task.timestamp * 1000)
      return taskDate.month() === currentMonth && taskDate.year() === currentYear
    }).length
  } catch (error) {
    console.error('获取我的审核任务失败:', error)
  }
}

// 搜索处理
const handleSearch = () => {
  pagination.value.current = 1
}

// 技术类型筛选处理
const handleTechnologyFilter = () => {
  pagination.value.current = 1
}

// 查看项目详情
const viewProjectDetail = (project: ProjectRecord) => {
  selectedProject.value = project
  showDetailModal.value = true
}

// 接受审核任务
const acceptReviewTask = (project: ProjectRecord) => {
  selectedProject.value = project
  showAcceptModal.value = true
}

// 处理接受任务
const handleAcceptTask = async () => {
  if (!selectedProject.value) return

  try {
    acceptLoading.value = true
    await acceptVerification(selectedProject.value.id)
    showAcceptModal.value = false
    
    // 刷新数据
    await Promise.all([
      fetchPendingProjects(),
      fetchMyReviewTasks()
    ])
    
    message.success('审核任务接受成功')
  } catch (error: any) {
    message.error(error.message || '接受任务失败')
  } finally {
    acceptLoading.value = false
  }
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

const getTechnologyColor = (tech: Technology) => {
  const colors = {
    [Technology.SOLAR]: 'orange',
    [Technology.WIND]: 'cyan',
    [Technology.HYDRO]: 'blue',
    [Technology.FORESTRY]: 'green',
    [Technology.CAPTURE]: 'purple',
    [Technology.OTHER]: 'default'
  }
  return colors[tech] || 'default'
}

const getTechnologyText = (tech: Technology) => {
  const texts = {
    [Technology.SOLAR]: '太阳能',
    [Technology.WIND]: '风能',
    [Technology.HYDRO]: '水电',
    [Technology.FORESTRY]: '林业',
    [Technology.CAPTURE]: '碳捕获',
    [Technology.OTHER]: '其他'
  }
  return texts[tech] || '未知'
}

onMounted(() => {
  Promise.all([
    fetchPendingProjects(),
    fetchMyReviewTasks()
  ])
})
</script>

<style scoped>
.audit-market {
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

.empty-state {
  padding: 60px 0;
  text-align: center;
}

.project-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.project-card :deep(.ant-card-body) {
  padding: 16px;
}

.project-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.project-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.label {
  font-size: 12px;
  color: #666;
  min-width: 60px;
}

.value {
  font-size: 14px;
  color: #262626;
}

.value.reduction {
  font-weight: 500;
  color: #52c41a;
}

.owner-info {
  display: flex;
  align-items: center;
  gap: 6px;
}

.owner-name {
  font-size: 13px;
  color: #262626;
}

.project-tech {
  padding-top: 8px;
  border-top: 1px solid #f0f0f0;
}

.project-tech .label {
  display: block;
  margin-bottom: 6px;
  font-size: 12px;
  color: #666;
}

.project-time {
  padding-top: 8px;
  border-top: 1px solid #f0f0f0;
}

.submit-time {
  font-size: 12px;
  color: #8c8c8c;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 24px;
}

.accept-confirm p {
  margin: 0;
  font-size: 14px;
}

:deep(.ant-card-actions) {
  background: #fafafa;
}

:deep(.ant-card-actions > li) {
  margin: 8px 0;
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