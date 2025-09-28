<template>
  <div class="user-detail-view">
    <div class="user-header">
      <DefaultAvatar :name="user.name || '用户'" :size="80" />
      <div class="user-basic-info">
        <h3>{{ user.name || '未设置用户名' }}</h3>
        <a-typography-text copyable>{{ user.addr }}</a-typography-text>
      </div>
    </div>

    <a-divider />

    <a-descriptions :column="1" size="middle">
      <a-descriptions-item label="用户名">
        {{ user.name || '未设置' }}
      </a-descriptions-item>
      
      <a-descriptions-item label="个人简介">
        <div class="bio-content">
          {{ user.bio || '暂无个人简介' }}
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="角色权限">
        <a-space wrap>
          <a-tag
            v-for="role in user.roles"
            :key="role"
            :color="getRoleColor(role)"
            size="large"
          >
            {{ PermissionService.getRoleDisplayName(role) }}
          </a-tag>
        </a-space>
      </a-descriptions-item>
      
      <a-descriptions-item label="资质文件">
        <div v-if="user.credentialsHash && user.credentialsHash.length > 0">
          <div
            v-for="(hash, index) in user.credentialsHash"
            :key="index"
            class="credential-item"
          >
            <a-tag class="credential-tag">
              {{ hash }}
            </a-tag>
          </div>
        </div>
        <span v-else class="no-data">暂无资质文件</span>
      </a-descriptions-item>
      
      <a-descriptions-item label="注册时间">
        {{ formatDate(user.timestamp) }}
      </a-descriptions-item>
      
      <a-descriptions-item label="编辑者">
        <div class="editor-info">
          <span>{{ formatAddress(user.editor) }}</span>
          <a-tag size="small" color="blue">记录索引: {{ user.editorRecordIndex }}</a-tag>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="user.editReason" label="编辑原因">
        <div class="edit-reason">
          {{ user.editReason }}
        </div>
      </a-descriptions-item>
    </a-descriptions>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import type { UserRecord } from '@/types'
import { PermissionService } from '@/services/permission'
import DefaultAvatar from './DefaultAvatar.vue'

interface Props {
  user: UserRecord
}

defineProps<Props>()

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm:ss')
}

const formatAddress = (address: string) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const getRoleColor = (role: number) => {
  const colors: Record<number, string> = {
    1: 'red',      // 监管者
    2: 'blue',     // 项目方
    3: 'green',    // 审核员
    4: 'orange'    // 购买方
  }
  return colors[role] || 'default'
}
</script>

<style scoped>
.user-detail-view {
  padding: 16px 0;
}

.user-header {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 20px;
}

.user-basic-info h3 {
  margin: 0 0 8px 0;
  font-size: 18px;
  font-weight: 600;
}

.bio-content {
  max-height: 100px;
  overflow-y: auto;
  line-height: 1.6;
  color: #666;
}

.credential-item {
  margin-bottom: 8px;
}

.credential-tag {
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

.editor-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.edit-reason {
  background: #fafafa;
  padding: 12px;
  border-radius: 6px;
  border-left: 3px solid #1890ff;
  color: #666;
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
