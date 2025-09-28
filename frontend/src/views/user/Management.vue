<template>
  <div class="user-management">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>用户管理</h1>
      <p>管理平台所有用户信息</p>
    </div>

    <!-- 权限检查 -->
    <template v-if="!PermissionService.isRegulator(currentUser)">
      <a-result
        status="403"
        title="无权限访问"
        sub-title="您需要监管者权限才能访问此页面"
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
              placeholder="搜索用户名或地址"
              style="width: 300px"
              @search="handleSearch"
              allow-clear
            />
            <a-select
              v-model:value="selectedRole"
              placeholder="筛选角色"
              style="width: 150px"
              allow-clear
              @change="handleRoleFilter"
            >
              <a-select-option :value="1">监管者</a-select-option>
              <a-select-option :value="2">项目方</a-select-option>
              <a-select-option :value="3">审核员</a-select-option>
              <a-select-option :value="4">购买方</a-select-option>
            </a-select>
          </div>
          
          <a-button type="primary" @click="showRegisterModal = true">
            <UserAddOutlined />
            注册新用户
          </a-button>
        </div>
      </a-card>

      <!-- 用户列表 -->
      <a-card title="用户列表" :loading="loading">
        <template #extra>
          <a-space>
            <span class="total-count">共 {{ filteredUsers.length }} 位用户</span>
            <a-button @click="fetchUsers" size="small">
              <ReloadOutlined />
              刷新
            </a-button>
          </a-space>
        </template>

        <a-table
          :columns="columns"
          :data-source="paginatedUsers"
          :pagination="pagination"
          :scroll="{ x: 1200 }"
          row-key="addr"
        >
          <template #bodyCell="{ column, record }: { column: any, record: UserRecord }">
            <!-- 用户头像和基本信息 -->
            <template v-if="column.key === 'user'">
              <div class="user-cell">
                <DefaultAvatar :name="record.name || '用户'" :size="40" />
                <div class="user-info">
                  <div class="user-name">{{ record.name || '未设置用户名' }}</div>
                  <div class="user-address">{{ formatAddress(record.addr) }}</div>
                </div>
              </div>
            </template>

            <!-- 角色标签 -->
            <template v-else-if="column.key === 'roles'">
              <a-space wrap>
                <a-tag
                  v-for="role in record.roles"
                  :key="role"
                  :color="getRoleColor(role)"
                >
                  {{ PermissionService.getRoleDisplayName(role) }}
                </a-tag>
              </a-space>
            </template>

            <!-- 个人简介 -->
            <template v-else-if="column.key === 'bio'">
              <div class="bio-cell">
                {{ record.bio || '暂无简介' }}
              </div>
            </template>

            <!-- 注册时间 -->
            <template v-else-if="column.key === 'timestamp'">
              {{ formatDate(record.timestamp) }}
            </template>

            <!-- 操作按钮 -->
            <template v-else-if="column.key === 'actions'">
              <a-space>
                <a-button size="small" @click="showUserDetail(record)">
                  详情
                </a-button>
                <a-button size="small" @click="editUser(record)">
                  编辑资料
                </a-button>
                <a-dropdown>
                  <template #overlay>
                    <a-menu>
                      <a-menu-item key="roles" @click="manageUserRoles(record)">
                        <UserOutlined />
                        权限管理
                      </a-menu-item>
                      <a-menu-divider />
                      <a-sub-menu key="records" title="记录查询">
                        <a-menu-item @click="queryUserProjects(record)">
                          用户的项目查询
                        </a-menu-item>
                        <a-menu-item @click="queryUserCredits(record)">
                          用户持有的碳信用查询
                        </a-menu-item>
                        <a-menu-item @click="queryUserCertificates(record)">
                          用户关联的证书查询
                        </a-menu-item>
                      </a-sub-menu>
                    </a-menu>
                  </template>
                  <a-button size="small">
                    更多
                    <DownOutlined />
                  </a-button>
                </a-dropdown>
              </a-space>
            </template>
          </template>
        </a-table>
      </a-card>

      <!-- 用户详情弹窗 -->
      <a-modal
        v-model:open="showDetailModal"
        title="用户详情"
        width="600px"
        :footer="null"
      >
        <UserDetailView v-if="selectedUser" :user="selectedUser" />
      </a-modal>

      <!-- 编辑用户资料弹窗 -->
      <ProfileModal 
        v-model:visible="showEditModal" 
        :user="selectedUser"
        @success="handleUserUpdated"
      />

      <!-- 注册新用户弹窗 -->
      <ProfileModal 
        v-model:visible="showRegisterModal" 
        :user="null"
        @success="handleUserRegistered"
      />

      <!-- 权限管理弹窗 -->
      <UserRoleModal
        v-model:visible="showRoleModal"
        :user="selectedUser"
        @success="handleRoleUpdated"
      />
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  UserAddOutlined,
  ReloadOutlined,
  UserOutlined,
  DownOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

import { useUserStore } from '@/stores/user'
import { getAllUsers } from '@/services/api'
import { PermissionService } from '@/services/permission'
import type { UserRecord } from '@/types'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProfileModal from '@/components/ProfileModal.vue'
import UserDetailView from '@/components/UserDetailView.vue'
import UserRoleModal from '@/components/UserRoleModal.vue'

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 状态
const loading = ref(false)
const users = ref<UserRecord[]>([])
const searchText = ref('')
const selectedRole = ref<number | undefined>()
const showDetailModal = ref(false)
const showEditModal = ref(false)
const showRegisterModal = ref(false)
const showRoleModal = ref(false)
const selectedUser = ref<UserRecord | null>(null)

// 表格列配置
const columns = [
  {
    title: '用户信息',
    key: 'user',
    width: 250,
    fixed: 'left' as const
  },
  {
    title: '角色',
    key: 'roles',
    width: 200
  },
  {
    title: '个人简介',
    key: 'bio',
    width: 300,
    ellipsis: true
  },
  {
    title: '注册时间',
    key: 'timestamp',
    width: 180
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    fixed: 'right' as const
  }
]

// 过滤后的用户列表
const filteredUsers = computed(() => {
  let result = users.value

  // 按文本搜索
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    result = result.filter(user =>
      (user.name && user.name.toLowerCase().includes(search)) ||
      user.addr.toLowerCase().includes(search)
    )
  }

  // 按角色筛选
  if (selectedRole.value !== undefined) {
    result = result.filter(user =>
      user.roles && user.roles.includes(selectedRole.value!)
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

// 分页后的用户列表
const paginatedUsers = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize
  const end = start + pagination.value.pageSize
  pagination.value.total = filteredUsers.value.length
  return filteredUsers.value.slice(start, end)
})

// 获取用户列表
const fetchUsers = async () => {
  try {
    loading.value = true
    users.value = await getAllUsers()
  } catch (error) {
    console.error('获取用户列表失败:', error)
    message.error('获取用户列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索处理
const handleSearch = () => {
  pagination.value.current = 1
}

// 角色筛选处理
const handleRoleFilter = () => {
  pagination.value.current = 1
}

// 显示用户详情
const showUserDetail = (user: UserRecord) => {
  selectedUser.value = user
  showDetailModal.value = true
}

// 编辑用户
const editUser = (user: UserRecord) => {
  selectedUser.value = user
  showEditModal.value = true
}

// 管理用户角色
const manageUserRoles = (user: UserRecord) => {
  selectedUser.value = user
  showRoleModal.value = true
}

// 查询用户项目
const queryUserProjects = (user: UserRecord) => {
  router.push({
    path: '/project/overview',
    query: { user: user.addr }
  })
}

// 查询用户碳信用
const queryUserCredits = (user: UserRecord) => {
  router.push({
    path: '/credit/market',
    query: { user: user.addr }
  })
}

// 查询用户证书
const queryUserCertificates = (user: UserRecord) => {
  router.push({
    path: '/certificate/overview',
    query: { user: user.addr }
  })
}

// 用户更新成功处理
const handleUserUpdated = () => {
  showEditModal.value = false
  fetchUsers()
  message.success('用户资料更新成功')
}

// 用户注册成功处理
const handleUserRegistered = () => {
  showRegisterModal.value = false
  fetchUsers()
  message.success('新用户注册成功')
}

// 角色更新成功处理
const handleRoleUpdated = () => {
  showRoleModal.value = false
  fetchUsers()
  message.success('用户权限更新成功')
}

// 工具函数
const formatDate = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm')
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
  fetchUsers()
})
</script>

<style scoped>
.user-management {
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

.user-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-info {
  min-width: 0;
  flex: 1;
}

.user-name {
  font-weight: 500;
  color: #262626;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-address {
  font-size: 12px;
  color: #8c8c8c;
  font-family: monospace;
}

.bio-cell {
  max-width: 300px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #666;
}

:deep(.ant-table-tbody .ant-table-cell) {
  vertical-align: top;
  padding: 16px;
}

:deep(.ant-dropdown-menu-submenu-title) {
  padding: 5px 12px;
}
</style>
