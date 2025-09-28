<template>
  <div class="my-projects">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>我的项目</h1>
      <p>管理您名下的所有项目</p>
    </div>

    <!-- 权限检查 -->
    <template v-if="!PermissionService.isProjectOwner(currentUser)">
      <a-result
        status="403"
        title="无权限访问"
        sub-title="您需要项目方权限才能访问此页面"
      >
        <template #extra>
          <a-button type="primary" @click="router.push('/dashboard')">
            返回首页
          </a-button>
        </template>
      </a-result>
    </template>

    <template v-else>
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
              <a-select-option :value="0">等待送审</a-select-option>
              <a-select-option :value="1">提交送审</a-select-option>
              <a-select-option :value="2">审核通过</a-select-option>
              <a-select-option :value="3">已撤销</a-select-option>
            </a-select>
          </div>
          
          <a-button type="primary" @click="showCreateModal = true">
            <PlusOutlined />
            注册新项目
          </a-button>
        </div>
      </a-card>

      <!-- 项目列表 -->
      <a-card title="我的项目列表" :loading="loading">
        <template #extra>
          <a-space>
            <span class="total-count">共 {{ filteredProjects.length }} 个项目</span>
            <a-button @click="fetchMyProjects" size="small" :loading="loading">
              <ReloadOutlined />
              刷新
            </a-button>
          </a-space>
        </template>

        <!-- 错误状态 -->
        <div v-if="!loading && projects.length === 0" class="empty-state">
          <a-empty
            description="暂无项目数据"
            :image="Empty.PRESENTED_IMAGE_SIMPLE"
          >
            <template #description>
              <span>暂无项目数据，请检查网络连接或联系管理员</span>
            </template>
            <a-button type="primary" @click="fetchMyProjects">
              重新加载
            </a-button>
          </a-empty>
        </div>

        <!-- 项目表格 -->
        <a-table
          v-else
          :columns="columns"
          :data-source="paginatedProjects"
          :pagination="pagination"
          :scroll="{ x: 1200 }"
          row-key="id"
        >
          <template #bodyCell="{ column, record }: { column: any, record: ProjectRecord }">
            <!-- 项目基本信息 -->
            <template v-if="column.key === 'project'">
              <div class="project-cell">
                <div class="project-info">
                  <div class="project-name">
                    <router-link :to="`/project/${record.id}`">
                      {{ record.name }}
                    </router-link>
                  </div>
                  <div class="project-id">#{{ record.id }}</div>
                  <div class="project-location">{{ record.location }}</div>
                </div>
              </div>
            </template>

            <!-- 技术类型 -->
            <template v-else-if="column.key === 'technologies'">
              <a-space wrap>
                <a-tag
                  v-for="tech in record.technologies"
                  :key="tech"
                  :color="getTechnologyColor(tech)"
                >
                  {{ getTechnologyText(tech) }}
                </a-tag>
              </a-space>
            </template>

            <!-- 减排量 -->
            <template v-else-if="column.key === 'reduction'">
              <span class="reduction-amount">{{ record.reduction.toLocaleString() }} 吨CO₂e</span>
            </template>

            <!-- 项目状态 -->
            <template v-else-if="column.key === 'status'">
              <div class="status-cell">
                <a-tag :color="getProjectStatusColor(record.status)">
                  {{ getProjectStatusText(record.status) }}
                </a-tag>
                <!-- 状态说明 -->
                <div class="status-desc">
                  {{ getStatusDescription(record.status) }}
                </div>
              </div>
            </template>

            <!-- 创建时间 -->
            <template v-else-if="column.key === 'timestamp'">
              {{ formatDate(record.timestamp) }}
            </template>

            <!-- 操作按钮 -->
            <template v-else-if="column.key === 'actions'">
              <a-space>
                <a-button size="small" @click="showProjectDetail(record)">
                  详情
                </a-button>
                <!-- 只有"等待送审"状态可以编辑 -->
                <a-button 
                  size="small" 
                  @click="editProject(record)"
                  :disabled="record.status !== ProjectStatus.EDITING"
                >
                  编辑信息
                </a-button>
                <!-- 提交审核按钮 -->
                <a-button 
                  v-if="record.status === ProjectStatus.EDITING"
                  type="primary"
                  size="small"
                  @click="submitForReview(record)"
                >
                  提交审核
                </a-button>
                <!-- 查看审核进度 -->
                <a-button 
                  v-else-if="record.status === ProjectStatus.REVIEWING"
                  size="small"
                  @click="viewReviewProgress(record)"
                >
                  审核进度
                </a-button>
              </a-space>
            </template>
          </template>
        </a-table>
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

      <!-- 创建/编辑项目弹窗 -->
      <ProjectEditModal
        v-model:visible="showEditModal"
        :project="selectedProject"
        @success="handleProjectUpdated"
      />

      <!-- 注册新项目弹窗 -->
      <ProjectEditModal
        v-model:visible="showCreateModal"
        :project="null"
        @success="handleProjectCreated"
      />

      <!-- 提交审核确认弹窗 -->
      <a-modal
        v-model:open="showSubmitModal"
        title="提交审核确认"
        width="500px"
        @ok="handleSubmitForReview"
        @cancel="showSubmitModal = false"
        :confirm-loading="submitLoading"
      >
        <div class="submit-confirm">
          <a-alert
            message="提交审核提醒"
            description="项目提交审核后将无法再次编辑，请确认项目信息已完善。审核员将对项目进行全面评估。"
            type="warning"
            show-icon
            style="margin-bottom: 16px"
          />
          <p>确定要提交项目 <strong>{{ selectedProject?.name }}</strong> 进行审核吗？</p>
        </div>
      </a-modal>

      <!-- 审核进度弹窗 -->
      <a-modal
        v-model:open="showProgressModal"
        title="审核进度"
        width="600px"
        :footer="null"
      >
        <ProjectReviewProgress v-if="selectedProject" :project="selectedProject" />
      </a-modal>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  PlusOutlined,
  ReloadOutlined
} from '@ant-design/icons-vue'
import { message, Empty } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { getAllProjects } from '@/services/api'
import { submitProjectForVerification } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import { walletInfo } from '@/services/web3'
import type { ProjectRecord } from '@/types'
import { ProjectStatus, Technology } from '@/types'
import ProjectDetailView from '@/components/ProjectDetailView.vue'
import ProjectEditModal from '@/components/ProjectEditModal.vue'
import ProjectReviewProgress from '@/components/ProjectReviewProgress.vue'

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const submitLoading = ref(false)
const projects = ref<ProjectRecord[]>([])
const searchText = ref('')
const selectedStatus = ref<number | undefined>()
const showDetailModal = ref(false)
const showEditModal = ref(false)
const showCreateModal = ref(false)
const showSubmitModal = ref(false)
const showProgressModal = ref(false)
const selectedProject = ref<ProjectRecord | null>(null)

// 表格列配置
const columns = [
  {
    title: '项目信息',
    key: 'project',
    width: 250,
    fixed: 'left' as const
  },
  {
    title: '技术类型',
    key: 'technologies',
    width: 200
  },
  {
    title: '减排量',
    key: 'reduction',
    width: 150
  },
  {
    title: '状态',
    key: 'status',
    width: 150
  },
  {
    title: '创建时间',
    key: 'timestamp',
    width: 120
  },
  {
    title: '操作',
    key: 'actions',
    width: 250,
    fixed: 'right' as const
  }
]

// 过滤后的项目列表
const filteredProjects = computed(() => {
  // 确保projects.value是数组
  if (!Array.isArray(projects.value)) {
    return []
  }
  
  let result = projects.value

  // 按文本搜索
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    result = result.filter(project => {
      if (!project || typeof project.name !== 'string') return false
      return project.name.toLowerCase().includes(search) ||
             (project.id && project.id.toString().includes(search))
    })
  }

  // 按状态筛选
  if (selectedStatus.value !== undefined) {
    result = result.filter(project => 
      project && project.status === selectedStatus.value
    )
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

// 分页后的项目列表
const paginatedProjects = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize
  const end = start + pagination.value.pageSize
  pagination.value.total = filteredProjects.value.length
  return filteredProjects.value.slice(start, end)
})

// 获取我的项目列表
const fetchMyProjects = async () => {
  if (!walletInfo.address) {
    projects.value = []
    return
  }

  try {
    loading.value = true
    const data = await getAllProjects()
    // 确保数据是数组格式
    const allProjects = Array.isArray(data) ? data : []
    
    // 只显示当前用户的项目，添加安全检查
    projects.value = allProjects.filter(project => {
      if (!project || !project.owner || !walletInfo.address) return false
      return project.owner.toLowerCase() === walletInfo.address.toLowerCase()
    })
  } catch (error: any) {
    console.error('获取项目列表失败:', error)
    // 设置空数组避免页面崩溃
    projects.value = []
    
    // 根据错误类型显示不同消息
    if (error.code === 'NETWORK_ERROR' || error.message?.includes('Network Error')) {
      message.error('网络连接失败，请检查网络后重试')
    } else if (error.response?.status === 404) {
      message.error('后端服务未启动，请启动后端服务')
    } else {
      message.error('获取项目列表失败，请稍后重试')
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

// 显示项目详情
const showProjectDetail = (project: ProjectRecord) => {
  selectedProject.value = project
  showDetailModal.value = true
}

// 编辑项目
const editProject = (project: ProjectRecord) => {
  selectedProject.value = project
  showEditModal.value = true
}

// 提交审核
const submitForReview = (project: ProjectRecord) => {
  selectedProject.value = project
  showSubmitModal.value = true
}

// 查看审核进度
const viewReviewProgress = (project: ProjectRecord) => {
  selectedProject.value = project
  showProgressModal.value = true
}

// 项目更新成功处理
const handleProjectUpdated = () => {
  showEditModal.value = false
  fetchMyProjects()
  message.success('项目信息更新成功')
}

// 项目创建成功处理
const handleProjectCreated = () => {
  showCreateModal.value = false
  fetchMyProjects()
  message.success('项目注册成功')
}

// 处理提交审核
const handleSubmitForReview = async () => {
  if (!selectedProject.value) return

  try {
    submitLoading.value = true
    await submitProjectForVerification(selectedProject.value.id)
    showSubmitModal.value = false
    fetchMyProjects()
    message.success('项目已提交审核')
  } catch (error: any) {
    message.error(error.message || '提交审核失败')
  } finally {
    submitLoading.value = false
  }
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD')
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
    [ProjectStatus.EDITING]: '等待送审',
    [ProjectStatus.REVIEWING]: '提交送审',
    [ProjectStatus.APPROVED]: '审核通过',
    [ProjectStatus.REVOKED]: '已撤销'
  }
  return texts[status] || '未知'
}

const getStatusDescription = (status: ProjectStatus) => {
  const descriptions = {
    [ProjectStatus.EDITING]: '可以编辑项目信息',
    [ProjectStatus.REVIEWING]: '等待审核员处理',
    [ProjectStatus.APPROVED]: '可以申请碳信用',
    [ProjectStatus.REVOKED]: '项目已被撤销'
  }
  return descriptions[status] || ''
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
  fetchMyProjects()
})
</script>

<style scoped>
.my-projects {
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

.status-cell {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.status-desc {
  font-size: 12px;
  color: #666;
  line-height: 1.2;
}

.reduction-amount {
  font-weight: 500;
  color: #52c41a;
}

.submit-confirm p {
  margin: 0;
  font-size: 14px;
}

:deep(.ant-table-tbody .ant-table-cell) {
  vertical-align: top;
  padding: 16px;
}

.empty-state {
  padding: 40px 0;
  text-align: center;
}
</style>

