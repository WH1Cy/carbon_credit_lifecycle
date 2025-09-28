<template>
  <div class="certificate-detail-view">
    <div class="certificate-header">
      <div class="certificate-basic">
        <h3>证书 #{{ certificate.id }}</h3>
        <a-tag :color="getCertificateStatusColor(certificate.status)" size="large">
          {{ getCertificateStatusText(certificate.status) }}
        </a-tag>
      </div>
    </div>

    <a-divider />

    <a-descriptions :column="2" size="middle">
      <a-descriptions-item label="证书编号">
        <span class="certificate-id">#{{ certificate.id }}</span>
      </a-descriptions-item>
      
      <a-descriptions-item label="批次编号">
        <a-tag color="blue">#{{ certificate.batchId }}</a-tag>
      </a-descriptions-item>
      
      <a-descriptions-item label="关联项目">
        <router-link :to="`/project/${certificate.projectId}`">
          项目 #{{ certificate.projectId }}
        </router-link>
      </a-descriptions-item>
      
      <a-descriptions-item label="数量">
        <span class="quantity-amount">{{ certificate.quantity.toLocaleString() }} 吨CO₂e</span>
      </a-descriptions-item>
      
      <a-descriptions-item label="所有者">
        <div class="owner-info">
          <DefaultAvatar :name="getOwnerName(certificate.owner)" :size="24" />
          <span class="owner-name">{{ getOwnerName(certificate.owner) || formatAddress(certificate.owner) }}</span>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="核发时间">
        {{ formatDate(certificate.issuedAt) }}
      </a-descriptions-item>
      
      <a-descriptions-item v-if="certificate.revokedAt" label="撤销时间">
        <span class="revoked-time">{{ formatDate(certificate.revokedAt) }}</span>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="certificate.reason" label="操作原因" :span="2">
        <div class="reason-content">
          {{ certificate.reason }}
        </div>
      </a-descriptions-item>
    </a-descriptions>

    <!-- 证书状态说明 -->
    <div class="certificate-status-info">
      <a-alert
        :type="getAlertType(certificate.status)"
        :message="getStatusMessage(certificate.status)"
        show-icon
      >
        <template #description>
          <div v-if="certificate.status === 0">
            该证书已正式核发，证明对应的碳信用已被确认和验证。证书确认了碳减排项目的真实性和有效性。
          </div>
          <div v-else-if="certificate.status === 1">
            该证书已被监管机构撤销，通常是由于项目或碳信用存在问题。撤销后的证书不再具有法律效力。
            <div v-if="certificate.revokedAt" class="revoke-details">
              撤销时间：{{ formatDate(certificate.revokedAt) }}
            </div>
          </div>
        </template>
      </a-alert>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import type { CertificateRecord } from '@/types'
import { CertificateStatus } from '@/types'
import { useUserStore } from '@/stores/user'
import DefaultAvatar from './DefaultAvatar.vue'

interface Props {
  certificate: CertificateRecord
}

defineProps<Props>()

const userStore = useUserStore()

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

const getCertificateStatusColor = (status: CertificateStatus) => {
  const colors = {
    0: 'green', // ISSUED
    1: 'red'    // REVOKED
  }
  return colors[status] || 'default'
}

const getCertificateStatusText = (status: CertificateStatus) => {
  const texts = {
    0: '已核发', // ISSUED
    1: '已撤销'  // REVOKED
  }
  return texts[status] || '未知'
}

const getAlertType = (status: CertificateStatus): 'success' | 'error' | 'warning' | 'info' => {
  const types: Record<number, 'success' | 'error' | 'warning' | 'info'> = {
    0: 'success', // ISSUED
    1: 'error'    // REVOKED
  }
  return types[status] || 'info'
}

const getStatusMessage = (status: CertificateStatus) => {
  const messages = {
    0: '证书已核发', // ISSUED
    1: '证书已撤销'  // REVOKED
  }
  return messages[status] || '未知状态'
}
</script>

<style scoped>
.certificate-detail-view {
  padding: 16px 0;
}

.certificate-header {
  margin-bottom: 20px;
}

.certificate-basic {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.certificate-basic h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
}

.certificate-id {
  font-size: 16px;
  font-weight: 600;
  color: #1890ff;
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

.revoked-time {
  color: #f5222d;
  font-weight: 500;
}

.reason-content {
  background: #fafafa;
  padding: 12px;
  border-radius: 6px;
  border-left: 3px solid #1890ff;
  color: #666;
  line-height: 1.6;
}

.certificate-status-info {
  margin-top: 20px;
}

.revoke-details {
  margin-top: 8px;
  font-size: 12px;
  color: #666;
}

:deep(.ant-descriptions-item-label) {
  font-weight: 500;
  color: #333;
}

:deep(.ant-descriptions-item-content) {
  color: #666;
}
</style>
