<template>
  <div class="review-progress">
    <a-steps 
      :current="currentStep" 
      :status="stepStatus"
      direction="vertical"
      size="small"
    >
      <a-step 
        title="项目提交"
        :description="stepDescriptions.submit"
      >
        <template #icon>
          <FileTextOutlined />
        </template>
      </a-step>
      
      <a-step 
        title="审核中"
        :description="stepDescriptions.review"
      >
        <template #icon>
          <AuditOutlined />
        </template>
      </a-step>
      
      <a-step 
        title="审核完成"
        :description="stepDescriptions.complete"
      >
        <template #icon>
          <CheckCircleOutlined v-if="project.status === 2" />
          <CloseCircleOutlined v-else-if="project.status === 3" />
          <LoadingOutlined v-else />
        </template>
      </a-step>
    </a-steps>

    <!-- 审核意见列表 -->
    <div v-if="reviews.length > 0" class="review-opinions">
      <a-divider orientation="left">审核意见</a-divider>
      
      <div v-for="review in reviews" :key="review.id" class="review-item">
        <a-card size="small">
          <div class="review-header">
            <div class="reviewer-info">
              <DefaultAvatar 
                :address="review.reviewer" 
                :size="32"
              />
              <div class="reviewer-detail">
                <div class="reviewer-name">{{ getReviewerName(review.reviewer) }}</div>
                <div class="review-time">{{ formatDate(review.timestamp) }}</div>
              </div>
            </div>
            <a-tag :color="getOpinionColor(review.opinion)">
              {{ getOpinionText(review.opinion) }}
            </a-tag>
          </div>
          
          <div v-if="review.comment" class="review-comment">
            {{ review.comment }}
          </div>
        </a-card>
      </div>
    </div>

    <!-- 当前状态说明 -->
    <div class="status-info">
      <a-divider orientation="left">状态说明</a-divider>
      <a-alert
        :message="statusInfo.title"
        :description="statusInfo.description"
        :type="statusInfo.type"
        show-icon
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import {
  FileTextOutlined,
  AuditOutlined,
  CheckCircleOutlined,
  CloseCircleOutlined,
  LoadingOutlined
} from '@ant-design/icons-vue'
import dayjs from 'dayjs'
import type { ProjectRecord, ReviewRecord } from '@/types'
import { ProjectStatus, ReviewOpinion } from '@/types'
import { getProjectReviews } from '@/services/api'
import DefaultAvatar from '@/components/DefaultAvatar.vue'

interface Props {
  project: ProjectRecord
}

const props = defineProps<Props>()

// 状态
const reviews = ref<ReviewRecord[]>([])
const loading = ref(false)

// 当前步骤
const currentStep = computed(() => {
  switch (props.project.status) {
    case ProjectStatus.EDITING:
      return 0
    case ProjectStatus.UNDER_REVIEW:
      return 1
    case ProjectStatus.APPROVED:
    case ProjectStatus.REVOKED:
      return 2
    default:
      return 0
  }
})

// 步骤状态
const stepStatus = computed(() => {
  if (props.project.status === ProjectStatus.REVOKED) {
    return 'error'
  } else if (props.project.status === ProjectStatus.APPROVED) {
    return 'finish'
  } else {
    return 'process'
  }
})

// 步骤描述
const stepDescriptions = computed(() => ({
  submit: `项目于 ${formatDate(props.project.timestamp)} 提交审核`,
  review: props.project.status === ProjectStatus.UNDER_REVIEW 
    ? `正在等待审核员评估，已收到 ${reviews.value.length} 条审核意见`
    : props.project.status > ProjectStatus.UNDER_REVIEW 
    ? `审核已完成，共收到 ${reviews.value.length} 条审核意见`
    : '等待提交审核',
  complete: getCompleteDescription()
}))

// 完成步骤描述
function getCompleteDescription() {
  switch (props.project.status) {
    case ProjectStatus.APPROVED:
      return `项目于 ${formatDate(getLastReviewTime())} 通过审核`
    case ProjectStatus.REVOKED:
      return `项目于 ${formatDate(getLastReviewTime())} 被撤销`
    default:
      return '等待审核完成'
  }
}

// 状态信息
const statusInfo = computed(() => {
  switch (props.project.status) {
    case ProjectStatus.EDITING:
      return {
        title: '项目编辑中',
        description: '项目正在编辑中，完成后可提交审核。',
        type: 'info' as const
      }
    case ProjectStatus.UNDER_REVIEW:
      return {
        title: '审核进行中',
        description: '项目正在接受审核员评估，请耐心等待审核结果。',
        type: 'warning' as const
      }
    case ProjectStatus.APPROVED:
      return {
        title: '审核通过',
        description: '恭喜！项目已通过审核，可以开始发行碳信用。',
        type: 'success' as const
      }
    case ProjectStatus.REVOKED:
      return {
        title: '项目已撤销',
        description: '项目已被撤销，无法继续进行。',
        type: 'error' as const
      }
    default:
      return {
        title: '未知状态',
        description: '项目状态异常，请联系管理员。',
        type: 'error' as const
      }
  }
})

// 获取审核意见颜色
const getOpinionColor = (opinion: ReviewOpinion) => {
  switch (opinion) {
    case ReviewOpinion.APPROVE:
      return 'green'
    case ReviewOpinion.REJECT:
      return 'red'
    case ReviewOpinion.REQUEST_CHANGES:
      return 'orange'
    default:
      return 'default'
  }
}

// 获取审核意见文本
const getOpinionText = (opinion: ReviewOpinion) => {
  switch (opinion) {
    case ReviewOpinion.APPROVE:
      return '批准'
    case ReviewOpinion.REJECT:
      return '拒绝'
    case ReviewOpinion.REQUEST_CHANGES:
      return '要求修改'
    default:
      return '未知'
  }
}

// 获取审核员名称
const getReviewerName = (address: string) => {
  // 这里可以根据实际情况获取审核员的真实姓名
  // 暂时显示地址的前8位
  return `审核员 ${address.slice(0, 8)}...`
}

// 获取最后审核时间
const getLastReviewTime = () => {
  if (reviews.value.length === 0) return props.project.timestamp
  return Math.max(...reviews.value.map(r => r.timestamp))
}

// 格式化日期
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm')
}

// 获取项目审核记录
const fetchReviews = async () => {
  try {
    loading.value = true
    reviews.value = await getProjectReviews(props.project.id)
  } catch (error) {
    console.error('获取审核记录失败:', error)
  } finally {
    loading.value = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchReviews()
})
</script>

<style scoped>
.review-progress {
  padding: 16px 0;
}

.review-opinions {
  margin-top: 24px;
}

.review-item {
  margin-bottom: 12px;
}

.review-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.reviewer-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.reviewer-detail {
  display: flex;
  flex-direction: column;
}

.reviewer-name {
  font-weight: 500;
  font-size: 14px;
}

.reviewer-time {
  color: #666;
  font-size: 12px;
}

.review-comment {
  padding: 8px 12px;
  background-color: #f5f5f5;
  border-radius: 4px;
  color: #333;
  line-height: 1.5;
}

.status-info {
  margin-top: 24px;
}

.ant-steps {
  max-width: 100%;
}

.ant-alert {
  border-radius: 6px;
}
</style>
