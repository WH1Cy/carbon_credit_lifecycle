<template>
  <div class="credit-detail-view">
    <div class="credit-header">
      <div class="credit-basic">
        <h3>批次 #{{ getBatchId(credit) }}</h3>
        <a-tag :color="getCreditEventColor(credit.eventType)" size="large">
          {{ getCreditEventText(credit.eventType) }}
        </a-tag>
      </div>
    </div>

    <a-divider />

    <a-descriptions :column="2" size="middle">
      <a-descriptions-item label="关联项目">
        <router-link :to="`/project/${credit.projectId}`">
          项目 #{{ credit.projectId }}
        </router-link>
      </a-descriptions-item>
      
      <a-descriptions-item label="年份">
        {{ credit.vintageYear }}年
      </a-descriptions-item>
      
      <a-descriptions-item label="数量">
        <span class="quantity-amount">{{ credit.quantity.toLocaleString() }} 吨CO₂e</span>
      </a-descriptions-item>
      
      <a-descriptions-item label="当前所有者">
        <div class="owner-info">
          <DefaultAvatar :name="getOwnerName(credit.currentOwner)" :size="24" />
          <span class="owner-name">{{ getOwnerName(credit.currentOwner) || formatAddress(credit.currentOwner) }}</span>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="所有者记录索引">
        <a-tag size="small" color="blue">#{{ credit.ownerRecordIndex }}</a-tag>
      </a-descriptions-item>
      
      <a-descriptions-item label="事件时间">
        {{ formatDate(credit.timestamp) }}
      </a-descriptions-item>
      
      <a-descriptions-item v-if="credit.eventType === CreditEventType.TRANSFER" label="转出方" :span="2">
        <div class="transfer-info">
          <div class="transfer-from">
            <span class="label">转出方：</span>
            <span class="name">{{ getOwnerName(credit.from) || formatAddress(credit.from) }}</span>
            <a-tag size="small" color="orange">记录索引: {{ credit.fromRecordIndex }}</a-tag>
          </div>
          <div class="transfer-to">
            <span class="label">转入方：</span>
            <span class="name">{{ getOwnerName(credit.to) || formatAddress(credit.to) }}</span>
            <a-tag size="small" color="green">记录索引: {{ credit.toRecordIndex }}</a-tag>
          </div>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="credit.parentBatchId" label="父批次ID">
        <a-tag color="purple">#{{ credit.parentBatchId }}</a-tag>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="credit.childBatchId" label="子批次ID">
        <a-tag color="cyan">#{{ credit.childBatchId }}</a-tag>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="credit.reason" label="操作原因" :span="2">
        <div class="reason-content">
          {{ credit.reason }}
        </div>
      </a-descriptions-item>
    </a-descriptions>

    <!-- 碳信用状态说明 -->
    <div class="credit-status-info">
      <a-alert
        :type="getAlertType(credit.eventType)"
        :message="getStatusMessage(credit.eventType)"
        show-icon
      >
        <template #description>
          <div v-if="credit.eventType === CreditEventType.MINT">
            该碳信用是从批准的项目中新铸造产生的，可以用于交易或退役。
          </div>
          <div v-else-if="credit.eventType === CreditEventType.TRANSFER">
            该碳信用已从 {{ getOwnerName(credit.from) || formatAddress(credit.from) }} 转让给 {{ getOwnerName(credit.to) || formatAddress(credit.to) }}。
          </div>
          <div v-else-if="credit.eventType === CreditEventType.RETIRE">
            该碳信用已被永久退役，不能再次交易，通常用于抵消碳排放。
          </div>
          <div v-else-if="credit.eventType === CreditEventType.REVOKE">
            该碳信用已被监管机构撤销，通常是由于项目存在问题。
          </div>
        </template>
      </a-alert>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import type { CreditRecord } from '@/types'
import { CreditEventType } from '@/types'
import { useUserStore } from '@/stores/user'
import DefaultAvatar from './DefaultAvatar.vue'

interface Props {
  credit: CreditRecord
}

defineProps<Props>()

const userStore = useUserStore()

// 获取用户名
const getOwnerName = (address: string) => {
  const user = userStore.getUserByAddress(address)
  return user?.name || ''
}

// 获取批次ID
const getBatchId = (credit: CreditRecord) => {
  return credit.childBatchId || credit.parentBatchId || 1
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm:ss')
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

const getAlertType = (eventType: CreditEventType) => {
  const types = {
    [CreditEventType.MINT]: 'success',
    [CreditEventType.TRANSFER]: 'info',
    [CreditEventType.RETIRE]: 'warning',
    [CreditEventType.REVOKE]: 'error'
  }
  return types[eventType] || 'info'
}

const getStatusMessage = (eventType: CreditEventType) => {
  const messages = {
    [CreditEventType.MINT]: '碳信用已铸造',
    [CreditEventType.TRANSFER]: '碳信用已转让',
    [CreditEventType.RETIRE]: '碳信用已退役',
    [CreditEventType.REVOKE]: '碳信用已撤销'
  }
  return messages[eventType] || '未知状态'
}
</script>

<style scoped>
.credit-detail-view {
  padding: 16px 0;
}

.credit-header {
  margin-bottom: 20px;
}

.credit-basic {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.credit-basic h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
}

.owner-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.owner-name {
  font-weight: 500;
}

.quantity-amount {
  font-weight: 500;
  color: #52c41a;
  font-size: 16px;
}

.transfer-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.transfer-from,
.transfer-to {
  display: flex;
  align-items: center;
  gap: 8px;
}

.label {
  color: #666;
  font-size: 14px;
  min-width: 60px;
}

.name {
  font-weight: 500;
  color: #262626;
}

.reason-content {
  background: #fafafa;
  padding: 12px;
  border-radius: 6px;
  border-left: 3px solid #1890ff;
  color: #666;
  line-height: 1.6;
}

.credit-status-info {
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
