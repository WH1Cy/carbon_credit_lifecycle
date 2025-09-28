<template>
  <div class="project-overview">
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="main-title font-display">
            <span class="title-icon">📊</span>
            {{ pageTitle }}
          </h1>
          <p class="subtitle font-body text-lg">{{ pageDescription }}</p>
        </div>
        <div class="header-stats">
          <div class="stat-item">
            <span class="stat-number font-bold text-2xl">{{ filteredProjects.length }}</span>
            <span class="stat-label font-medium text-sm">项目总数</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 操作栏 -->
    <a-card class="action-card card-hover" :bordered="false">
      <div class="action-bar">
        <div class="search-filters">
          <div class="filter-group">
            <a-input-search
              v-model:value="searchText"
              placeholder="搜索项目名称或ID"
              class="search-input"
              @search="handleSearch"
              allow-clear
            >
              <template #prefix>
                <SearchOutlined style="color: #bfbfbf" />
              </template>
            </a-input-search>
          </div>
          <div class="filter-group">
            <a-select
              v-model:value="selectedStatus"
              placeholder="筛选状态"
              class="filter-select"
              allow-clear
              @change="handleStatusFilter"
            >
              <a-select-option :value="0">编辑中</a-select-option>
              <a-select-option :value="1">审核中</a-select-option>
              <a-select-option :value="2">已批准</a-select-option>
              <a-select-option :value="3">已撤销</a-select-option>
            </a-select>
          </div>
          <div class="filter-group">
            <a-select
              v-model:value="selectedTechnology"
              placeholder="筛选技术类型"
              class="filter-select"
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
        </div>
        
        <!-- 监管者可以撤销项目 -->
        <div v-if="isRegulator">
          <a-button 
            type="primary" 
            danger
            :disabled="!selectedRowKeys.length"
            @click="showBatchRevokeModal"
          >
            <DeleteOutlined />
            批量撤销
          </a-button>
        </div>
      </div>
    </a-card>

    <!-- 项目列表 -->
    <a-card class="project-list-card card-hover" :bordered="false">
      <template #title>
        <div class="card-title">
          <div class="title-icon-wrapper">
            <ProjectOutlined class="title-icon" />
          </div>
          <span>项目列表</span>
        </div>
      </template>
      <template #extra>
        <div class="list-actions">
          <div class="total-count">
            <span class="count-number">{{ filteredProjects.length }}</span>
            <span class="count-label">个项目</span>
          </div>
          <a-button 
            @click="fetchProjects" 
            size="small" 
            :loading="loading"
            class="btn-hover"
          >
            <ReloadOutlined />
            刷新
          </a-button>
        </div>
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
          <a-button type="primary" @click="fetchProjects">
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
        :scroll="{ x: 1400 }"
        :row-selection="isRegulator ? rowSelection : undefined"
        row-key="id"
        class="modern-table"
        :row-class-name="getRowClassName"
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

          <!-- 项目所有者 -->
          <template v-else-if="column.key === 'owner'">
            <div class="owner-cell">
              <DefaultAvatar :name="getOwnerName(record.owner)" :size="32" />
              <div class="owner-info">
                <div class="owner-name">{{ getOwnerName(record.owner) }}</div>
                <div class="owner-address">{{ formatAddress(record.owner) }}</div>
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
            <a-tag :color="getProjectStatusColor(record.status)">
              {{ getProjectStatusText(record.status) }}
            </a-tag>
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
              <!-- 监管者可以编辑和撤销任何项目 -->
              <template v-if="isRegulator">
                <a-button size="small" @click="editProject(record)">
                  编辑信息
                </a-button>
                <a-button 
                  size="small" 
                  danger
                  @click="revokeProject(record)"
                  :disabled="record.status === 3"
                >
                  撤销项目
                </a-button>
              </template>
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

    <!-- 编辑项目弹窗 -->
    <ProjectEditModal
      v-model:visible="showEditModal"
      :project="selectedProject"
      @success="handleProjectUpdated"
    />

    <!-- 撤销项目确认弹窗 -->
    <EditReasonModal
      v-model:visible="showRevokeModal"
      title="撤销项目"
      :description="`确定要撤销项目「${selectedProject?.name}」吗？`"
      @confirm="handleProjectRevoke"
    />

    <!-- 批量撤销确认弹窗 -->
    <EditReasonModal
      v-model:visible="showBatchRevokeModalFlag"
      title="批量撤销项目"
      :description="`确定要撤销选中的 ${selectedRowKeys.length} 个项目吗？`"
      @confirm="handleBatchRevoke"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  ReloadOutlined,
  DeleteOutlined
} from '@ant-design/icons-vue'
import { message, Empty } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { getAllProjects } from '@/services/api'
import { banProject } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import type { ProjectRecord } from '@/types'
import { ProjectStatus, Technology } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProjectDetailView from '@/components/ProjectDetailView.vue'
import ProjectEditModal from '@/components/ProjectEditModal.vue'
import EditReasonModal from '@/components/EditReasonModal.vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const projects = ref<ProjectRecord[]>([])
const searchText = ref('')
const selectedStatus = ref<number | undefined>()
const selectedTechnology = ref<number | undefined>()
const showDetailModal = ref(false)
const showEditModal = ref(false)
const showRevokeModal = ref(false)
const showBatchRevokeModalFlag = ref(false)
const selectedProject = ref<ProjectRecord | null>(null)
const selectedRowKeys = ref<number[]>([])

// 权限检查
const isRegulator = computed(() => PermissionService.isRegulator(currentUser.value))

// 页面标题
const pageTitle = computed(() => isRegulator.value ? '项目管理' : '项目总览')
const pageDescription = computed(() => 
  isRegulator.value ? '管理平台所有项目信息' : '查看平台所有项目信息'
)

// 表格列配置
const columns = computed(() => {
  const baseColumns = [
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
      width: 100
    },
    {
      title: '创建时间',
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

  return baseColumns
})

// 行选择配置（仅监管者）
const rowSelection = computed(() => ({
  selectedRowKeys: selectedRowKeys.value,
  onChange: (keys: number[]) => {
    selectedRowKeys.value = keys
  },
  getCheckboxProps: (record: ProjectRecord) => ({
    disabled: record.status === ProjectStatus.REVOKED
  })
}))

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

  // 按技术类型筛选
  if (selectedTechnology.value !== undefined) {
    result = result.filter(project => 
      project && Array.isArray(project.technologies) && 
      project.technologies.includes(selectedTechnology.value!)
    )
  }

  // URL查询参数筛选（来自用户管理页面的跳转）
  if (route.query.user) {
    const userAddress = route.query.user as string
    result = result.filter(project => 
      project && project.owner && 
      project.owner.toLowerCase() === userAddress.toLowerCase()
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

// 获取项目列表
const fetchProjects = async () => {
  try {
    loading.value = true
    const data = await getAllProjects()
    // 确保数据是数组格式
    projects.value = Array.isArray(data) ? data : []
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

// 技术类型筛选处理
const handleTechnologyFilter = () => {
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

// 撤销项目
const revokeProject = (project: ProjectRecord) => {
  selectedProject.value = project
  showRevokeModal.value = true
}

// 显示批量撤销弹窗
const showBatchRevokeModal = () => {
  showBatchRevokeModalFlag.value = true
}

// 项目更新成功处理
const handleProjectUpdated = () => {
  showEditModal.value = false
  fetchProjects()
  message.success('项目信息更新成功')
}

// 项目撤销处理
const handleProjectRevoke = async (reason: string) => {
  if (!selectedProject.value) return

  try {
    await banProject(selectedProject.value.id, reason)
    showRevokeModal.value = false
    fetchProjects()
    message.success('项目撤销成功')
  } catch (error: any) {
    message.error(error.message || '项目撤销失败')
  }
}

// 批量撤销处理
const handleBatchRevoke = async (reason: string) => {
  try {
    await Promise.all(
      selectedRowKeys.value.map(id => banProject(id, reason))
    )
    showBatchRevokeModalFlag.value = false
    selectedRowKeys.value = []
    fetchProjects()
    message.success(`成功撤销 ${selectedRowKeys.value.length} 个项目`)
  } catch (error: any) {
    message.error(error.message || '批量撤销失败')
  }
}

// 获取用户名
const getOwnerName = (address: string) => {
  const user = userStore.getUserByAddress(address)
  return user?.name || '未知用户'
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD')
}

const formatAddress = (address: string) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
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
  fetchProjects()
})
</script>

<style scoped>
.project-overview {
  max-width: 1600px;
  margin: 0 auto;
  padding: 32px 24px;
  min-height: 100vh;
}

/* 页面标题样式 */
.page-header {
  margin-bottom: 32px;
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
  font-size: 32px;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 12px;
  color: var(--color-text);
}

.title-icon {
  font-size: 36px;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
}

.subtitle {
  margin: 8px 0 0 0;
  font-size: 16px;
  opacity: 0.9;
  font-weight: 400;
  color: var(--color-text-secondary);
}

.header-stats {
  display: flex;
  gap: 24px;
}

.stat-item {
  text-align: center;
  padding: 16px 24px;
  background: var(--color-bg-elevated);
  border: 1px solid var(--color-border);
  border-radius: 6px;
}

.stat-number {
  display: block;
  font-size: 24px;
  font-weight: 600;
  line-height: 1;
  color: var(--color-primary);
}

.stat-label {
  display: block;
  font-size: 12px;
  color: var(--color-text-secondary);
  margin-top: 4px;
}

/* 操作栏样式 */
.action-card {
  margin-bottom: 24px;
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  overflow: hidden;
  box-shadow: var(--box-shadow);
}

.action-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 24px;
  flex-wrap: wrap;
}

.search-filters {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
  flex: 1;
}

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.search-input {
  width: 300px;
  min-width: 200px;
}

.filter-select {
  width: 150px;
  min-width: 120px;
}

/* 项目列表样式 */
.project-list-card {
  background: var(--color-bg-container);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  overflow: hidden;
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
  background: #0969da;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 16px;
}

.list-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.total-count {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.count-number {
  font-size: 20px;
  font-weight: 600;
  color: #0969da;
}

.count-label {
  font-size: 14px;
  color: #8c8c8c;
}

/* 表格样式 */
.modern-table {
  margin-top: 16px;
}

.modern-table :deep(.ant-table-thead > tr > th) {
  background: var(--color-bg-elevated);
  border-bottom: 1px solid var(--color-border);
  font-weight: 600;
  color: var(--color-text);
  padding: 12px 16px;
}

.modern-table :deep(.ant-table-tbody > tr > td) {
  padding: 12px 16px;
  border-bottom: 1px solid var(--color-border);
  transition: all 0.1s cubic-bezier(0.4, 0, 0.2, 1);
  color: var(--color-text);
}

.modern-table :deep(.ant-table-tbody > tr:hover > td) {
  background: var(--color-bg-elevated);
}

.modern-table :deep(.ant-table-tbody > tr.ant-table-row-selected > td) {
  background: #dbeafe;
}

/* 项目单元格样式 */
.project-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.project-info {
  flex: 1;
  min-width: 0;
}

.project-name {
  font-size: 14px;
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.project-name a {
  color: inherit;
  text-decoration: none;
  transition: color 0.2s ease;
}

.project-name a:hover {
  color: #0969da;
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

/* 所有者单元格样式 */
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

/* 减排量样式 */
.reduction-amount {
  font-weight: 500;
  color: #1a7f37;
}

/* 状态标签样式 */
.status-tag {
  border-radius: 6px;
  font-weight: 500;
  font-size: 12px;
  padding: 2px 8px;
  border: none;
}

/* 技术标签样式 */
.technology-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.technology-tag {
  font-size: 11px;
  padding: 1px 6px;
  border-radius: 4px;
  background: #f0f0f0;
  color: #666;
  border: none;
}

/* 操作按钮样式 */
.action-buttons {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.action-btn {
  border-radius: 6px;
  font-size: 12px;
  padding: 4px 8px;
  height: auto;
  transition: all 0.2s ease;
}

/* 空状态样式 */
.empty-state {
  padding: 60px 20px;
  text-align: center;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .project-overview {
    padding: 16px;
  }
  
  .header-content {
    flex-direction: column;
    gap: 20px;
    text-align: center;
    padding: 24px;
  }
  
  .main-title {
    font-size: 24px;
  }
  
  .action-bar {
    flex-direction: column;
    align-items: stretch;
  }
  
  .search-filters {
    flex-direction: column;
  }
  
  .search-input,
  .filter-select {
    width: 100%;
  }
  
  .list-actions {
    justify-content: space-between;
  }
}

/* 深度样式覆盖 */
:deep(.ant-card-head) {
  border-bottom: 1px solid #f0f0f0;
  padding: 20px 24px 16px;
}

:deep(.ant-card-body) {
  padding: 24px;
}

:deep(.ant-input-search) {
  border-radius: 8px;
}

:deep(.ant-select) {
  border-radius: 8px;
}

:deep(.ant-btn) {
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.2s ease;
}

:deep(.ant-btn-primary) {
  background: #0969da;
  border-color: #0969da;
  box-shadow: 0 1px 0 rgba(27, 31, 36, 0.04);
}

:deep(.ant-btn-primary:hover) {
  background: #0550ae;
  border-color: #0550ae;
  transform: translateY(-1px);
  box-shadow: 0 1px 0 rgba(27, 31, 36, 0.04);
}

:deep(.ant-pagination) {
  margin-top: 24px;
  text-align: center;
}

:deep(.ant-pagination-item) {
  border-radius: 6px;
}

:deep(.ant-pagination-item-active) {
  background: #0969da;
  border-color: #0969da;
}

:deep(.ant-pagination-item-active a) {
  color: white;
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

:deep(.ant-table-tbody .ant-table-cell) {
  vertical-align: top;
  padding: 16px;
}

.empty-state {
  padding: 40px 0;
  text-align: center;
}
</style>
