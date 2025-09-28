<template>
  <div class="user-profile">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>个人资料</h1>
      <p>查看和管理您的个人信息</p>
    </div>

    <!-- 权限检查 -->
    <template v-if="!currentUser">
      <a-result
        status="403"
        title="无权限访问"
        sub-title="请先完善个人资料"
      >
        <template #extra>
          <a-button type="primary" @click="router.push('/dashboard')">
            返回首页
          </a-button>
        </template>
      </a-result>
    </template>

    <template v-else>
      <!-- 用户信息卡片 -->
      <a-card title="基本信息" class="profile-info-card" :loading="loading">
        <template #extra>
          <a-button type="primary" @click="showEditModal = true">
            编辑资料
          </a-button>
        </template>

        <div class="profile-content">
          <div class="profile-avatar">
            <DefaultAvatar :name="currentUser.name || '用户'" :size="120" />
          </div>
          <div class="profile-details">
            <a-descriptions :column="2" size="middle">
              <a-descriptions-item label="钱包地址" :span="2">
                <a-typography-text copyable>{{ walletInfo.address }}</a-typography-text>
              </a-descriptions-item>
              <a-descriptions-item label="用户名">
                {{ currentUser.name || '未设置' }}
              </a-descriptions-item>
              <a-descriptions-item label="角色">
                <a-space>
                  <a-tag
                    v-for="role in currentUser.roles"
                    :key="role"
                    :color="getRoleColor(role)"
                    size="large"
                  >
                    {{ PermissionService.getRoleDisplayName(role) }}
                  </a-tag>
                </a-space>
              </a-descriptions-item>
              <a-descriptions-item label="个人简介" :span="2">
                {{ currentUser.bio || '暂无个人简介' }}
              </a-descriptions-item>
              <a-descriptions-item label="资质文件" :span="2">
                <div v-if="currentUser.credentialsHash && currentUser.credentialsHash.length > 0">
                  <a-tag
                    v-for="(hash, index) in currentUser.credentialsHash"
                    :key="index"
                    class="credential-tag"
                  >
                    {{ hash.slice(0, 8) }}...{{ hash.slice(-8) }}
                  </a-tag>
                </div>
                <span v-else class="no-data">暂无资质文件</span>
              </a-descriptions-item>
              <a-descriptions-item label="注册时间">
                {{ formatDate(currentUser.timestamp) }}
              </a-descriptions-item>
              <a-descriptions-item label="最后更新">
                {{ formatDate(currentUser.timestamp) }}
              </a-descriptions-item>
            </a-descriptions>
          </div>
        </div>
      </a-card>

      <!-- 变更历史 -->
      <a-card title="个人资料变更历史" class="history-card" :loading="historyLoading">
        <template #extra>
          <a-button @click="fetchUserHistory" size="small">
            <ReloadOutlined />
            刷新
          </a-button>
        </template>

        <a-timeline mode="left" class="history-timeline">
          <a-timeline-item
            v-for="(record, index) in userHistory"
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
                  {{ index === 0 ? '当前版本' : `历史版本 ${userHistory.length - index}` }}
                </a-tag>
                <span class="editor-info">
                  {{ record.editor === walletInfo.address ? '本人' : `由 ${formatAddress(record.editor)} 编辑` }}
                </span>
              </div>
              
              <div class="timeline-details">
                <a-descriptions :column="1" size="small">
                  <a-descriptions-item label="用户名">
                    {{ record.name || '未设置' }}
                  </a-descriptions-item>
                  <a-descriptions-item label="个人简介">
                    {{ record.bio || '暂无' }}
                  </a-descriptions-item>
                  <a-descriptions-item label="角色">
                    <a-space>
                      <a-tag
                        v-for="role in record.roles"
                        :key="role"
                        :color="getRoleColor(role)"
                        size="small"
                      >
                        {{ PermissionService.getRoleDisplayName(role) }}
                      </a-tag>
                    </a-space>
                  </a-descriptions-item>
                  <a-descriptions-item v-if="record.editReason" label="编辑原因">
                    {{ record.editReason }}
                  </a-descriptions-item>
                </a-descriptions>
              </div>
            </div>
          </a-timeline-item>
        </a-timeline>
      </a-card>

      <!-- 编辑个人资料弹窗 -->
      <ProfileModal 
        v-model:visible="showEditModal" 
        :user="currentUser"
        @success="handleProfileUpdated"
      />
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ReloadOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { walletInfo } from '@/services/web3'
import { getUserHistory } from '@/services/api'
import { PermissionService } from '@/services/permission'
import type { UserRecord } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProfileModal from '@/components/ProfileModal.vue'

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const historyLoading = ref(false)
const showEditModal = ref(false)
const userHistory = ref<UserRecord[]>([])

// 获取用户历史记录
const fetchUserHistory = async () => {
  if (!walletInfo.address) return
  
  try {
    historyLoading.value = true
    userHistory.value = await getUserHistory(walletInfo.address)
  } catch (error) {
    console.error('获取用户历史记录失败:', error)
    message.error('获取用户历史记录失败')
  } finally {
    historyLoading.value = false
  }
}

// 个人资料更新成功处理
const handleProfileUpdated = () => {
  showEditModal.value = false
  userStore.fetchCurrentUser()
  fetchUserHistory()
  message.success('个人资料更新成功')
}

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

onMounted(() => {
  fetchUserHistory()
})
</script>

<style scoped>
.user-profile {
  max-width: 1000px;
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

.profile-info-card {
  margin-bottom: 24px;
}

.profile-content {
  display: flex;
  align-items: flex-start;
  gap: 32px;
}

.profile-avatar {
  flex-shrink: 0;
}

.profile-details {
  flex: 1;
}

.credential-tag {
  margin-bottom: 4px;
  font-family: monospace;
}

.no-data {
  color: #999;
  font-style: italic;
}

.history-card {
  margin-bottom: 24px;
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
</style>