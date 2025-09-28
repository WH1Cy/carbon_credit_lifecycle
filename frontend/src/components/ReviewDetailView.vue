<template>
  <div class="review-detail-view">
    <div class="review-header">
      <div class="review-result">
        <a-tag :color="review.approved ? 'success' : 'error'" size="large">
          {{ review.approved ? '审核通过' : '审核不通过' }}
        </a-tag>
      </div>
      <div class="review-time">
        {{ formatDate(review.timestamp) }}
      </div>
    </div>

    <a-divider />

    <a-descriptions :column="1" size="middle">
      <a-descriptions-item label="审核员">
        <div class="reviewer-info">
          <DefaultAvatar :name="getReviewerName(review.handlingVerifier)" :size="32" />
          <div class="reviewer-details">
            <div class="reviewer-name">{{ getReviewerName(review.handlingVerifier) || '未知审核员' }}</div>
            <div class="reviewer-address">{{ formatAddress(review.handlingVerifier) }}</div>
          </div>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="审核结果">
        <a-tag :color="review.approved ? 'success' : 'error'">
          {{ review.approved ? '通过' : '不通过' }}
        </a-tag>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="review.approved && review.issuedBatchId" label="核发批次">
        <a-tag color="green">#{{ review.issuedBatchId }}</a-tag>
      </a-descriptions-item>
      
      <a-descriptions-item label="审核意见">
        <div class="review-comment">
          {{ review.comment || '无审核意见' }}
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="审核时间">
        {{ formatDate(review.timestamp) }}
      </a-descriptions-item>
      
      <a-descriptions-item label="审核员记录索引">
        <a-tag size="small" color="blue">#{{ review.verifierRecordIndex }}</a-tag>
      </a-descriptions-item>
    </a-descriptions>

    <!-- 审核状态说明 -->
    <div class="review-status-info">
      <a-alert
        :type="review.approved ? 'success' : 'error'"
        :message="review.approved ? '项目审核通过' : '项目审核不通过'"
        show-icon
      >
        <template #description>
          <div v-if="review.approved">
            该项目已通过审核，可以进行碳信用的核发。
            <span v-if="review.issuedBatchId">已核发批次ID：#{{ review.issuedBatchId }}</span>
          </div>
          <div v-else>
            该项目未通过审核，请根据审核意见进行修改后重新提交。
          </div>
        </template>
      </a-alert>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import type { ReviewRecord } from '@/types'
import { useUserStore } from '@/stores/user'
import DefaultAvatar from './DefaultAvatar.vue'

interface Props {
  review: ReviewRecord
}

defineProps<Props>()

const userStore = useUserStore()

// 获取审核员名称
const getReviewerName = (address: string) => {
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
</script>

<style scoped>
.review-detail-view {
  padding: 16px 0;
}

.review-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.review-result {
  /* 样式 */
}

.review-time {
  font-size: 14px;
  color: #666;
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

.reviewer-address {
  font-size: 12px;
  color: #8c8c8c;
  font-family: monospace;
}

.review-comment {
  background: #fafafa;
  padding: 12px;
  border-radius: 6px;
  border-left: 3px solid #1890ff;
  color: #666;
  line-height: 1.6;
  min-height: 40px;
}

.review-status-info {
  margin-top: 20px;
}

:deep(.ant-descriptions-item-label) {
  font-weight: 500;
  color: #333;
}

:deep(.ant-descriptions-item-content) {
  color: #666;
}
</style>
