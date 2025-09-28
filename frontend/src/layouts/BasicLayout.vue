<template>
  <a-layout class="layout">
    <!-- 侧边导航 -->
    <a-layout-sider
      v-model:collapsed="collapsed"
      :trigger="null"
      collapsible
      width="256"
      theme="light"
      class="sider"
      breakpoint="lg"
      :collapsed-width="0"
    >
      <!-- Logo -->
      <div class="logo">
        <img src="/logo.svg" alt="Logo" v-if="!collapsed" />
        <img src="/logo-small.svg" alt="Logo" v-else />
        <h1 v-if="!collapsed">碳信通</h1>
      </div>

      <!-- 菜单 -->
      <a-menu
        v-model:selectedKeys="selectedKeys"
        v-model:openKeys="openKeys"
        mode="inline"
        :items="menuItems"
        @click="handleMenuClick"
        class="main-menu"
      />
    </a-layout-sider>

    <!-- 主要内容区域 -->
    <a-layout>
      <!-- 顶部导航栏 -->
      <a-layout-header class="header">
        <div class="header-left">
          <MenuUnfoldOutlined
            v-if="collapsed"
            class="trigger"
            @click="() => (collapsed = !collapsed)"
          />
          <MenuFoldOutlined
            v-else
            class="trigger"
            @click="() => (collapsed = !collapsed)"
          />
          
          <!-- 面包屑导航 -->
          <a-breadcrumb class="breadcrumb">
            <a-breadcrumb-item>
              <router-link to="/">首页</router-link>
            </a-breadcrumb-item>
            <a-breadcrumb-item v-if="currentRoute.meta?.title">
              {{ currentRoute.meta.title }}
            </a-breadcrumb-item>
          </a-breadcrumb>
        </div>

        <div class="header-right">
          <!-- 钱包连接状态和用户信息 -->
          <a-space>
            <!-- 主题切换按钮 -->
            <ThemeToggle />
            <!-- 未连接钱包的情况 -->
            <a-button
              v-if="!walletInfo.isConnected"
              type="primary"
              size="small"
              :loading="isConnecting"
              @click="handleConnectWallet"
            >
              <WalletOutlined />
              连接钱包
            </a-button>

            <!-- 用户信息区域 -->
            <template v-else>
              <!-- 已连接钱包但用户未注册 -->
              <div v-if="!currentUser" class="user-status">
                <a-dropdown :trigger="['hover']" placement="bottomRight">
                  <div class="user-info">
                    <DefaultAvatar :name="'新用户'" />
                    <div class="user-details">
                      <div class="user-name">新用户</div>
                      <div class="wallet-address">{{ formatAddress(walletInfo.address) }}</div>
                    </div>
                    <DownOutlined />
                  </div>
                  <template #overlay>
                    <a-menu>
                      <a-menu-item key="complete-profile" @click="showProfileModal = true">
                        <UserOutlined />
                        完善个人资料
                      </a-menu-item>
                      <a-menu-item key="switch-address">
                        <WalletOutlined />
                        切换地址
                      </a-menu-item>
                      <a-menu-divider />
                      <a-menu-item key="disconnect" @click="handleDisconnect" class="disconnect-item">
                        <LogoutOutlined />
                        退出钱包
                      </a-menu-item>
                    </a-menu>
                  </template>
                </a-dropdown>
              </div>

              <!-- 已注册用户 -->
              <div v-else class="user-status">
                <a-dropdown :trigger="['hover']" placement="bottomRight">
                  <div class="user-info">
                    <DefaultAvatar :name="currentUser.name || '用户'" />
                    <div class="user-details">
                      <div class="user-name">{{ currentUser.name || '未设置用户名' }}</div>
                      <div class="user-roles" v-if="currentUser.roles && currentUser.roles.length > 0">
                        <a-tag
                          v-for="role in currentUser.roles"
                          :key="role"
                          size="small"
                          :color="getRoleColor(role)"
                        >
                          {{ PermissionService.getRoleDisplayName(role) }}
                        </a-tag>
                      </div>
                      <div v-else class="no-role">暂无权限</div>
                    </div>
                    <DownOutlined />
                  </div>
                  <template #overlay>
                    <a-menu>
                      <a-menu-item key="profile" @click="goToProfile">
                        <UserOutlined />
                        个人资料
                      </a-menu-item>
                      <a-menu-item key="switch-address">
                        <WalletOutlined />
                        切换地址
                      </a-menu-item>
                      <a-menu-divider />
                      <a-menu-item key="disconnect" @click="handleDisconnect" class="disconnect-item">
                        <LogoutOutlined />
                        退出钱包
                      </a-menu-item>
                    </a-menu>
                  </template>
                </a-dropdown>
              </div>
            </template>
          </a-space>
        </div>
      </a-layout-header>

      <!-- 内容区域 -->
      <a-layout-content class="content">
        <div class="content-wrapper">
          <RouterView />
        </div>
      </a-layout-content>
    </a-layout>

    <!-- 个人资料完善弹窗 -->
    <ProfileModal 
      v-model:visible="showProfileModal" 
      :user="currentUser"
      @success="handleProfileUpdated"
    />
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, h } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  DashboardOutlined,
  UserOutlined,
  ProjectOutlined,
  GoldOutlined,
  SafetyCertificateOutlined,
  WalletOutlined,
  LogoutOutlined,
  DownOutlined,
  CheckOutlined,
  AuditOutlined,
  TeamOutlined,
  FileSearchOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'
import { connectWallet, disconnectWallet, walletInfo, isConnecting } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import DefaultAvatar from '@/components/DefaultAvatar.vue'
import ProfileModal from '@/components/ProfileModal.vue'
import ThemeToggle from '@/components/ThemeToggle.vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

// 布局状态
const collapsed = ref(false)
const selectedKeys = ref<string[]>([])
const openKeys = ref<string[]>([])
const showProfileModal = ref(false)

// 当前路由
const currentRoute = computed(() => route)

// 动态菜单项配置
const menuItems = computed(() => {
  const items = []

  // 概览 - 所有用户都可以访问（无论是否连接钱包）
  items.push({
    key: '/dashboard',
    icon: () => h(DashboardOutlined),
    label: '概览',
    title: '概览'
  })

  // 未连接钱包：只显示概览
  if (!walletInfo.isConnected) {
    return items
  }

  // 钱包已连接但用户未注册：只显示概览
  if (!currentUser.value) {
    return items
  }

  // 用户已注册，根据权限显示二级菜单
  const user = currentUser.value

  // 用户菜单 - 根据权限显示不同子菜单
  const userSubMenus = []
  
  // 个人资料 - 所有有用户信息的用户可访问
  userSubMenus.push({
    key: '/user/profile',
    label: '个人资料',
    icon: () => h(UserOutlined)
  })

  // 用户管理 - 仅监管者可访问
  if (PermissionService.isRegulator(user)) {
    userSubMenus.push({
      key: '/user/management',
      label: '用户管理',
      icon: () => h(TeamOutlined)
    })
  }

  if (userSubMenus.length > 0) {
    items.push({
      key: 'user',
      icon: () => h(UserOutlined),
      label: '用户',
      children: userSubMenus
    })
  }

  // 项目菜单
  const projectSubMenus = []
  
  // 项目总览/项目管理 - 所有用户都可访问
  const projectOverviewLabel = PermissionService.isRegulator(user) ? '项目管理' : '项目总览'
  projectSubMenus.push({
    key: '/project/overview',
    label: projectOverviewLabel,
    icon: () => h(ProjectOutlined)
  })

  // 我的项目 - 项目方可访问
  if (PermissionService.isProjectOwner(user)) {
    projectSubMenus.push({
      key: '/project/my',
      label: '我的项目',
      icon: () => h(FileSearchOutlined)
    })
  }

  items.push({
    key: 'project',
    icon: () => h(ProjectOutlined),
    label: '项目',
    children: projectSubMenus
  })

  // 审核菜单 - 仅审核员可访问
  if (PermissionService.isVerifier(user)) {
    const auditSubMenus = [
      {
        key: '/audit/market',
        label: '审核市场',
        icon: () => h(AuditOutlined)
      },
      {
        key: '/audit/my',
        label: '我的审核',
        icon: () => h(FileSearchOutlined)
      }
    ]

    items.push({
      key: 'audit',
      icon: () => h(AuditOutlined),
      label: '审核',
      children: auditSubMenus
    })
  }

  // 碳信用菜单
  const creditSubMenus = []
  
  // 碳信用交易市场 - 所有用户都可访问
  creditSubMenus.push({
    key: '/credit/market',
    label: '碳信用交易市场',
    icon: () => h(GoldOutlined)
  })

  // 我的碳信用 - 项目方和购买方可访问
  if (PermissionService.isProjectOwner(user) || PermissionService.isBuyer(user)) {
    creditSubMenus.push({
      key: '/credit/my',
      label: '我的碳信用',
      icon: () => h(FileSearchOutlined)
    })
  }

  items.push({
    key: 'credit',
    icon: () => h(GoldOutlined),
    label: '碳信用',
    children: creditSubMenus
  })

  // 证书菜单
  const certificateSubMenus = []
  
  // 证书总览/证书管理 - 所有用户都可访问
  const certificateOverviewLabel = PermissionService.isRegulator(user) ? '证书管理' : '证书总览'
  certificateSubMenus.push({
    key: '/certificate/overview',
    label: certificateOverviewLabel,
    icon: () => h(SafetyCertificateOutlined)
  })

  // 我的证书 - 如果有相关逻辑可以添加
  // if (PermissionService.isBuyer(user)) {
  //   certificateSubMenus.push({
  //     key: '/certificate/my',
  //     label: '我的证书',
  //     icon: () => h(FileSearchOutlined)
  //   })
  // }

  items.push({
    key: 'certificate',
    icon: () => h(SafetyCertificateOutlined),
    label: '证书',
    children: certificateSubMenus
  })

  return items
})

// 菜单点击处理
const handleMenuClick = ({ key }: { key: string | number }) => {
  router.push(String(key))
}

// 格式化地址显示
const formatAddress = (address: string) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

// 连接钱包
const handleConnectWallet = async () => {
  try {
    await connectWallet()
    await userStore.fetchCurrentUser()
    message.success('钱包连接成功')
  } catch (error: any) {
    message.error(error.message || '钱包连接失败')
  }
}

// 跳转到个人资料页面
const goToProfile = () => {
  router.push('/user/profile')
}

// 个人资料更新成功处理
const handleProfileUpdated = () => {
  showProfileModal.value = false
  userStore.fetchCurrentUser()
  message.success('个人资料更新成功')
}

// 断开钱包连接
const handleDisconnect = () => {
  disconnectWallet()
  userStore.currentUser = null
  router.push('/')
  message.info('钱包已断开连接')
}

// 获取角色颜色
const getRoleColor = (role: number) => {
  const colorMap: Record<number, string> = {
    1: 'purple',   // 监管者 - 使用紫色更突出
    2: 'blue',     // 项目方 - 蓝色
    3: 'green',    // 审核员 - 绿色
    4: 'orange'    // 购买方 - 橙色
  }
  return colorMap[role] || 'default'
}

// 监听路由变化更新选中菜单
watch(
  () => route.path,
  (newPath) => {
    selectedKeys.value = [newPath]
    
    // 自动展开对应的菜单
    if (newPath.startsWith('/user/')) {
      openKeys.value = ['user']
    } else if (newPath.startsWith('/project/')) {
      openKeys.value = ['project']
    } else if (newPath.startsWith('/audit/')) {
      openKeys.value = ['audit']
    } else if (newPath.startsWith('/credit/')) {
      openKeys.value = ['credit']
    } else if (newPath.startsWith('/certificate/')) {
      openKeys.value = ['certificate']
    }
  },
  { immediate: true }
)

// 监听钱包连接状态变化
watch(
  () => walletInfo.isConnected,
  (connected) => {
    console.log('[BasicLayout] 钱包连接状态变化:', connected)
    if (connected) {
      userStore.fetchCurrentUser()
    }
  }
)

// 监听钱包连接和用户状态变化，自动弹出个人资料完善窗口
watch(
  [() => walletInfo.isConnected, () => currentUser.value],
  ([connected, user]) => {
    if (connected && !user && route.path === '/dashboard') {
      // 钱包已连接但用户未注册，且在概览页面时自动弹出
      showProfileModal.value = true
    }
  },
  { immediate: true }
)

// 监听currentUser变化
watch(
  () => currentUser.value,
  (user) => {
    console.log('[BasicLayout] currentUser变化:', user)
    if (user) {
      console.log('[BasicLayout] 用户名:', user.name)
      console.log('[BasicLayout] 用户角色原始数据:', user.roles)
      console.log('[BasicLayout] 用户角色类型:', typeof user.roles, Array.isArray(user.roles))
      
      // 检查每个角色的显示名称
      if (user.roles && Array.isArray(user.roles)) {
        user.roles.forEach((role, index) => {
          console.log(`[BasicLayout] 角色[${index}]: ${role} (类型: ${typeof role}) -> ${PermissionService.getRoleDisplayName(role)}`)
        })
      }
      
      console.log('[BasicLayout] 角色检查:')
      console.log('  - isRegulator:', PermissionService.isRegulator(user))
      console.log('  - isProjectOwner:', PermissionService.isProjectOwner(user))
      console.log('  - isVerifier:', PermissionService.isVerifier(user))
      console.log('  - isBuyer:', PermissionService.isBuyer(user))
      console.log('  - hasAnyRole:', PermissionService.hasAnyRole(user))
    }
  },
  { immediate: true }
)

onMounted(() => {
  // 初始化用户数据
  userStore.fetchUsers()
})
</script>

<style scoped>
.layout {
  min-height: 100vh;
}

.sider {
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  z-index: 100;
  border-right: 1px solid #f0f0f0;
}

.logo {
  height: 64px;
  display: flex;
  align-items: center;
  padding: 0 16px;
  border-bottom: 1px solid #f0f0f0;
}

.logo img {
  height: 32px;
  margin-right: 8px;
}

.logo h1 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #1890ff;
}

.header {
  position: fixed;
  top: 0;
  right: 0;
  left: 256px;
  z-index: 99;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  border-bottom: 1px solid #f0f0f0;
  padding: 0 24px;
  transition: left 0.2s;
}

.layout :deep(.ant-layout-sider-collapsed) + .ant-layout .header {
  left: 80px;
}

.header-left {
  display: flex;
  align-items: center;
}

.trigger {
  font-size: 18px;
  padding: 0 16px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger:hover {
  color: #1890ff;
}

.breadcrumb {
  margin-left: 16px;
}

.header-right {
  display: flex;
  align-items: center;
}

.content {
  margin-left: 256px;
  margin-top: 64px;
  transition: margin-left 0.2s;
}

.layout :deep(.ant-layout-sider-collapsed) + .ant-layout .content {
  margin-left: 80px;
}

.content-wrapper {
  padding: 24px;
  min-height: calc(100vh - 64px);
  background: #f0f2f5;
}

.user-status {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 6px;
  transition: all 0.3s;
  border: 1px solid var(--color-border);
  background: var(--color-bg-container);
  color: var(--color-text);
}

.user-info:hover {
  background-color: var(--color-bg-elevated);
  border-color: var(--color-primary);
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 2px;
  min-width: 120px;
}

.user-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
  line-height: 1.2;
}

.wallet-address {
  font-size: 12px;
  color: var(--color-text-secondary);
  font-family: monospace;
}

.user-roles {
  display: flex;
  gap: 4px;
  flex-wrap: wrap;
}

.no-role {
  font-size: 12px;
  color: var(--color-text-secondary);
  font-style: italic;
}

.main-menu {
  height: calc(100vh - 64px);
  border-right: none;
}

/* 钱包相关样式 */
.wallet-section {
  position: relative;
}

.wallet-tag {
  cursor: pointer;
  padding: 4px 8px;
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  border-radius: 4px;
  transition: all 0.3s;
}

.wallet-tag:hover {
  background-color: #52c41a !important;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(82, 196, 26, 0.3);
}

:deep(.wallet-menu) {
  min-width: 280px;
  padding: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border-radius: 8px;
}

.wallet-menu-header {
  padding: 12px 16px;
  background: #f8f9fa;
  border-bottom: 1px solid #e8e8e8;
}

.current-network {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  font-weight: 500;
  color: #666;
}

.network-dot {
  width: 8px;
  height: 8px;
  background: #52c41a;
  border-radius: 50%;
}

.account-list {
  padding: 8px 0;
  max-height: 240px;
  overflow-y: auto;
}

.account-list-title {
  padding: 8px 16px;
  font-size: 12px;
  font-weight: 500;
  color: #666;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.account-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
  position: relative;
}

.account-item:hover {
  background: #f0f7ff;
}

.account-item.active {
  background: #e6f7ff;
  border-left: 3px solid #1890ff;
}

.account-info {
  flex: 1;
  min-width: 0;
}

.account-address {
  font-weight: 500;
  font-size: 14px;
  color: #262626;
}

.account-full-address {
  font-size: 11px;
  color: #8c8c8c;
  font-family: 'Monaco', 'Menlo', monospace;
  margin-top: 2px;
  word-break: break-all;
}

.check-icon {
  color: #1890ff;
  font-size: 14px;
}

:deep(.disconnect-item) {
  color: #ff4d4f !important;
}

:deep(.disconnect-item:hover) {
  background: #fff2f0 !important;
  color: #ff4d4f !important;
}
</style>
