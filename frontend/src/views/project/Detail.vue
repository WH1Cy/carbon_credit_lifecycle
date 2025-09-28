<template>
  <div class="project-detail">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <a-spin size="large" />
    </div>

    <!-- 项目不存在 -->
    <div v-else-if="!currentProject" class="not-found">
      <a-result
        status="404"
        title="项目不存在"
        sub-title="未找到指定的项目信息"
      >
        <template #extra>
          <a-button type="primary" @click="router.push('/project/overview')">
            返回项目列表
          </a-button>
        </template>
      </a-result>
    </div>

    <!-- 项目详情内容 -->
    <div v-else class="project-content">
      <!-- 项目基本信息 -->
      <a-card title="项目信息" class="project-info-card">
        <template #extra>
          <a-space>
            <a-button @click="router.back()" size="small">
              <ArrowLeftOutlined />
              返回
            </a-button>
            <a-button @click="fetchProjectHistory" size="small">
              <ReloadOutlined />
              刷新
            </a-button>
          </a-space>
        </template>

        <div class="project-header">
          <div class="project-basic">
            <h2>{{ currentProject.name }}</h2>
            <div class="project-meta">
              <a-tag>项目 ID: {{ currentProject.id }}</a-tag>
              <a-tag :color="getProjectStatusColor(currentProject.status)">
                {{ getProjectStatusText(currentProject.status) }}
              </a-tag>
            </div>
          </div>
        </div>

        <a-row :gutter="[24, 16]">
          <a-col :xs="24" :md="12">
            <a-descriptions :column="1" size="middle">
              <a-descriptions-item label="项目所有者">
                <div class="owner-info">
                  <DefaultAvatar :name="getOwnerName(currentProject.owner)" :size="32" />
                  <div class="owner-details">
                    <div class="owner-name">{{ getOwnerName(currentProject.owner) }}</div>
                    <div class="owner-address">{{ formatAddress(currentProject.owner) }}</div>
                  </div>
                </div>
              </a-descriptions-item>
              
              <a-descriptions-item label="项目位置">
                {{ currentProject.location }}
              </a-descriptions-item>
              
              <a-descriptions-item label="预期减排量">
                <span class="reduction-amount">{{ currentProject.reduction.toLocaleString() }} 吨CO₂e</span>
              </a-descriptions-item>
              
              <a-descriptions-item label="创建时间">
                {{ formatDate(currentProject.timestamp) }}
              </a-descriptions-item>
            </a-descriptions>
          </a-col>
          
          <a-col :xs="24" :md="12">
            <a-descriptions :column="1" size="middle">
              <a-descriptions-item label="技术类型">
                <a-space wrap>
                  <a-tag
                    v-for="tech in currentProject.technologies"
                    :key="tech"
                    :color="getTechnologyColor(tech)"
                    size="large"
                  >
                    {{ getTechnologyText(tech) }}
                  </a-tag>
                </a-space>
              </a-descriptions-item>
              
              <a-descriptions-item label="项目描述">
                <div class="project-description">
                  {{ currentProject.description || '暂无描述' }}
                </div>
              </a-descriptions-item>
              
              <a-descriptions-item label="相关文档">
                <div v-if="currentProject.documentsHash && currentProject.documentsHash.length > 0">
                  <div
                    v-for="(hash, index) in currentProject.documentsHash"
                    :key="index"
                    class="document-item"
                  >
                    <a-tag class="document-tag">
                      文档 {{ index + 1 }}: {{ hash.slice(0, 12) }}...{{ hash.slice(-8) }}
                    </a-tag>
                  </div>
                </div>
                <span v-else class="no-data">暂无相关文档</span>
              </a-descriptions-item>
            </a-descriptions>
          </a-col>
        </a-row>
      </a-card>

      <!-- 项目变更历史 -->
      <a-card title="项目变更历史" class="history-card">
        <template #extra>
          <span class="history-count">共 {{ projectHistory.length }} 条记录</span>
        </template>

        <a-timeline mode="left" class="history-timeline">
          <a-timeline-item
            v-for="(record, index) in projectHistory"
            :key="index"
            :color="index === 0 ? 'green' : 'blue'"
          >
            <template #label>
              <div class="timeline-label">
                {{ formatDate(record.timestamp) }}
              </div>
            </template>
            
            <div class="timeline-content">
              <div class="timeline-header">
                <a-tag :color="index === 0 ? 'green' : 'blue'">
                  {{ index === 0 ? '当前版本' : `历史版本 ${projectHistory.length - index}` }}
                </a-tag>
                <span class="editor-info">
                  由 {{ getOwnerName(record.editor) || formatAddress(record.editor) }} 编辑
                </span>
              </div>
              
              <div class="timeline-details">
                <a-descriptions :column="2" size="small">
                  <a-descriptions-item label="项目名称">
                    {{ record.name }}
                  </a-descriptions-item>
                  <a-descriptions-item label="项目位置">
                    {{ record.location }}
                  </a-descriptions-item>
                  <a-descriptions-item label="预期减排量">
                    {{ record.reduction.toLocaleString() }} 吨CO₂e
                  </a-descriptions-item>
                  <a-descriptions-item label="项目状态">
                    <a-tag :color="getProjectStatusColor(record.status)">
                      {{ getProjectStatusText(record.status) }}
                    </a-tag>
                  </a-descriptions-item>
                  <a-descriptions-item label="技术类型" :span="2">
                    <a-space wrap>
                      <a-tag
                        v-for="tech in record.technologies"
                        :key="tech"
                        :color="getTechnologyColor(tech)"
                        size="small"
                      >
                        {{ getTechnologyText(tech) }}
                      </a-tag>
                    </a-space>
                  </a-descriptions-item>
                  <a-descriptions-item v-if="record.editReason" label="编辑原因" :span="2">
                    <div class="edit-reason">
                      {{ record.editReason }}
                    </div>
                  </a-descriptions-item>
                </a-descriptions>
              </div>
            </div>
          </a-timeline-item>
        </a-timeline>
      </a-card>

      <!-- 审核记录 -->
      <a-card title="审核记录" class="review-card" :loading="reviewLoading">
        <template #extra>
          <a-button @click="fetchReviewRecords" size="small">
            <ReloadOutlined />
            刷新
          </a-button>
        </template>

        <div v-if="reviewRecords.length === 0" class="no-reviews">
          <a-empty description="暂无审核记录" />
        </div>

        <div v-else class="review-list">
          <div
            v-for="(review, index) in reviewRecords"
            :key="index"
            class="review-item"
          >
            <div class="review-header">
              <div class="reviewer-info">
                <DefaultAvatar :name="getOwnerName(review.handlingVerifier)" :size="40" />
                <div class="reviewer-details">
                  <div class="reviewer-name">{{ getOwnerName(review.handlingVerifier) }}</div>
                  <div class="review-time">{{ formatDate(review.timestamp) }}</div>
                </div>
              </div>
              <a-tag :color="review.approved ? 'success' : 'error'">
                {{ review.approved ? '审核通过' : '审核不通过' }}
              </a-tag>
            </div>
            
            <div class="review-content">
              <div class="review-comment">
                <strong>审核意见：</strong>
                {{ review.comment || '无审核意见' }}
              </div>
              
              <div v-if="review.approved && review.issuedBatchId" class="issued-batch">
                <strong>核发批次：</strong>
                <a-tag color="green">#{{ review.issuedBatchId }}</a-tag>
              </div>
            </div>
          </div>
        </div>
      </a-card>

      <!-- 关联的碳信用 -->
      <a-card title="关联的碳信用" class="credits-card" :loading="creditsLoading">
        <template #extra>
          <a-button @click="fetchRelatedCredits" size="small">
            <ReloadOutlined />
            刷新
          </a-button>
        </template>

        <div v-if="relatedCredits.length === 0" class="no-credits">
          <a-empty description="暂无关联的碳信用" />
        </div>

        <a-table
          v-else
          :columns="creditColumns"
          :data-source="relatedCredits"
          :pagination="false"
          size="small"
        >
          <template #bodyCell="{ column, record }: { column: any, record: any }">
            <template v-if="column.key === 'eventType'">
              <a-tag :color="getCreditEventColor(record.eventType)">
                {{ getCreditEventText(record.eventType) }}
              </a-tag>
            </template>
            <template v-else-if="column.key === 'quantity'">
              {{ record.quantity.toLocaleString() }} 吨CO₂e
            </template>
            <template v-else-if="column.key === 'timestamp'">
              {{ formatDate(record.timestamp) }}
            </template>
          </template>
        </a-table>
      </a-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  ArrowLeftOutlined,
  ReloadOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { 
  getProjectHistory, 
  getProjectReviews, 
  getAllCredits 
} from '@/services/api'
import type { ProjectRecord, ReviewRecord, CreditRecord } from '@/types'
import { ProjectStatus, Technology, CreditEventType } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

// 状态
const loading = ref(false)
const reviewLoading = ref(false)
const creditsLoading = ref(false)
const projectHistory = ref<ProjectRecord[]>([])
const reviewRecords = ref<ReviewRecord[]>([])
const relatedCredits = ref<CreditRecord[]>([])

// 项目ID
const projectId = computed(() => {
  const id = route.params.id
  return Array.isArray(id) ? parseInt(id[0]) : parseInt(id as string)
})

// 当前项目（最新版本）
const currentProject = computed(() => {
  return projectHistory.value.length > 0 ? projectHistory.value[0] : null
})

// 碳信用表格列配置
const creditColumns = [
  {
    title: '事件类型',
    key: 'eventType',
    width: 120
  },
  {
    title: '数量',
    key: 'quantity',
    width: 150
  },
  {
    title: '年份',
    dataIndex: 'vintageYear',
    width: 100
  },
  {
    title: '当前所有者',
    dataIndex: 'currentOwner',
    ellipsis: true,
    customRender: ({ text }: { text: string }) => formatAddress(text)
  },
  {
    title: '时间',
    key: 'timestamp',
    width: 150
  }
]

// 获取项目历史记录
const fetchProjectHistory = async () => {
  try {
    loading.value = true
    projectHistory.value = await getProjectHistory(projectId.value)
  } catch (error) {
    console.error('获取项目历史记录失败:', error)
    message.error('获取项目历史记录失败')
  } finally {
    loading.value = false
  }
}

// 获取审核记录
const fetchReviewRecords = async () => {
  try {
    reviewLoading.value = true
    reviewRecords.value = await getProjectReviews(projectId.value)
  } catch (error) {
    console.error('获取审核记录失败:', error)
    message.error('获取审核记录失败')
  } finally {
    reviewLoading.value = false
  }
}

// 获取关联的碳信用
const fetchRelatedCredits = async () => {
  try {
    creditsLoading.value = true
    const allCredits = await getAllCredits()
    // 筛选与当前项目相关的碳信用
    relatedCredits.value = allCredits.filter(credit => 
      credit.projectId === projectId.value
    )
  } catch (error) {
    console.error('获取关联碳信用失败:', error)
    message.error('获取关联碳信用失败')
  } finally {
    creditsLoading.value = false
  }
}

// 获取用户名
const getOwnerName = (address: string) => {
  const user = userStore.getUserByAddress(address)
  return user?.name || ''
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm:ss')
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
  fetchProjectHistory()
  fetchReviewRecords()
  fetchRelatedCredits()
})
</script>

<style scoped>
.project-detail {
  max-width: 1200px;
  margin: 0 auto;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.not-found {
  margin-top: 100px;
}

.project-content {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.project-info-card {
  /* 样式 */
}

.project-header {
  margin-bottom: 24px;
}

.project-basic h2 {
  margin: 0 0 12px 0;
  font-size: 24px;
  font-weight: 600;
}

.project-meta {
  display: flex;
  gap: 8px;
  align-items: center;
}

.owner-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.owner-details {
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

.reduction-amount {
  font-weight: 500;
  color: #52c41a;
}

.project-description {
  line-height: 1.6;
  color: #666;
  max-height: 100px;
  overflow-y: auto;
}

.document-item {
  margin-bottom: 8px;
}

.document-tag {
  font-family: monospace;
  font-size: 12px;
  padding: 4px 8px;
  word-break: break-all;
  display: inline-block;
  max-width: 100%;
}

.no-data {
  color: #999;
  font-style: italic;
}

.history-card,
.review-card,
.credits-card {
  /* 样式 */
}

.history-count {
  color: #666;
  font-size: 14px;
}

.history-timeline {
  margin-top: 16px;
}

.timeline-label {
  font-size: 12px;
  color: #666;
  text-align: right;
  padding-right: 16px;
}

.timeline-content {
  padding: 12px 16px;
  background: #fafafa;
  border-radius: 6px;
}

.timeline-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
}

.editor-info {
  font-size: 12px;
  color: #666;
}

.timeline-details :deep(.ant-descriptions-item-label) {
  font-weight: 500;
  color: #333;
}

.timeline-details :deep(.ant-descriptions-item-content) {
  color: #666;
}

.edit-reason {
  background: #f0f7ff;
  padding: 8px 12px;
  border-radius: 4px;
  border-left: 3px solid #1890ff;
  color: #666;
  line-height: 1.5;
}

.no-reviews,
.no-credits {
  padding: 40px 0;
  text-align: center;
}

.review-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.review-item {
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  padding: 16px;
  background: #fafafa;
}

.review-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.reviewer-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.reviewer-details {
  flex: 1;
}

.reviewer-name {
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.review-time {
  font-size: 12px;
  color: #8c8c8c;
}

.review-content {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.review-comment,
.issued-batch {
  font-size: 14px;
  line-height: 1.5;
}

:deep(.ant-descriptions-item-label) {
  font-weight: 500;
  color: #333;
}

:deep(.ant-descriptions-item-content) {
  color: #666;
}
</style>
