<template>
  <div class="project-detail-view">
    <div class="project-header">
      <h3>{{ project.name }}</h3>
      <a-tag :color="getProjectStatusColor(project.status)">
        {{ getProjectStatusText(project.status) }}
      </a-tag>
    </div>

    <a-divider />

    <a-descriptions :column="2" size="middle">
      <a-descriptions-item label="项目ID">
        #{{ project.id }}
      </a-descriptions-item>
      
      <a-descriptions-item label="项目所有者">
        <div class="owner-info">
          <DefaultAvatar :name="getOwnerName(project.owner)" :size="24" />
          <span class="owner-name">{{ getOwnerName(project.owner) || formatAddress(project.owner) }}</span>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="项目位置">
        {{ project.location }}
      </a-descriptions-item>
      
      <a-descriptions-item label="预期减排量">
        <span class="reduction-amount">{{ project.reduction.toLocaleString() }} 吨CO₂e</span>
      </a-descriptions-item>
      
      <a-descriptions-item label="技术类型" :span="2">
        <a-space wrap>
          <a-tag
            v-for="tech in project.technologies"
            :key="tech"
            :color="getTechnologyColor(tech)"
          >
            {{ getTechnologyText(tech) }}
          </a-tag>
        </a-space>
      </a-descriptions-item>
      
      <a-descriptions-item label="项目描述" :span="2">
        <div class="project-description">
          {{ project.description || '暂无描述' }}
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item label="相关文档" :span="2">
        <div v-if="project.documentsHash && project.documentsHash.length > 0">
          <div
            v-for="(hash, index) in project.documentsHash"
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
      
      <a-descriptions-item label="创建时间">
        {{ formatDate(project.timestamp) }}
      </a-descriptions-item>
      
      <a-descriptions-item label="编辑者">
        <div class="editor-info">
          <span>{{ getOwnerName(project.editor) || formatAddress(project.editor) }}</span>
          <a-tag size="small" color="blue">记录索引: {{ project.editorRecordIndex }}</a-tag>
        </div>
      </a-descriptions-item>
      
      <a-descriptions-item v-if="project.editReason" label="编辑原因" :span="2">
        <div class="edit-reason">
          {{ project.editReason }}
        </div>
      </a-descriptions-item>
    </a-descriptions>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import type { ProjectRecord } from '@/types'
import { ProjectStatus, Technology } from '@/types'
import { useUserStore } from '@/stores/user'
import DefaultAvatar from './DefaultAvatar.vue'

interface Props {
  project: ProjectRecord
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
</script>

<style scoped>
.project-detail-view {
  padding: 16px 0;
}

.project-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.project-header h3 {
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

.reduction-amount {
  font-weight: 500;
  color: #52c41a;
}

.project-description {
  max-height: 100px;
  overflow-y: auto;
  line-height: 1.6;
  color: #666;
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
