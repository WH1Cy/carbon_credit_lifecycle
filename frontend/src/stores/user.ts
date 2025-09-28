import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { UserRecord, Role } from '@/types'
import { getAllUsers, getUserHistory } from '@/services/api'
import { walletInfo } from '@/services/web3'

export const useUserStore = defineStore('user', () => {
  // 状态
  const users = ref<UserRecord[]>([])
  const currentUser = ref<UserRecord | null>(null)
  const loading = ref(false)

  // 计算属性
  const currentUserRoles = computed(() => {
    return currentUser.value?.roles || []
  })

  const isRegulator = computed(() => {
    return currentUserRoles.value.includes(1) // Role.REGULATOR
  })

  const isProjectOwner = computed(() => {
    return currentUserRoles.value.includes(2) // Role.PROJECT_OWNER
  })

  const isVerifier = computed(() => {
    return currentUserRoles.value.includes(3) // Role.VERIFIER
  })

  const isBuyer = computed(() => {
    return currentUserRoles.value.includes(4) // Role.BUYER
  })

  // 操作
  const fetchUsers = async () => {
    try {
      loading.value = true
      users.value = await getAllUsers()
    } catch (error) {
      console.error('获取用户列表失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const fetchCurrentUser = async () => {
    if (!walletInfo.isConnected || !walletInfo.address) {
      console.log('[UserStore] 钱包未连接或地址为空')
      currentUser.value = null
      return
    }

    try {
      console.log('[UserStore] 正在获取用户历史记录:', walletInfo.address)
      const userHistory = await getUserHistory(walletInfo.address)
      console.log('[UserStore] 获取到的用户历史记录:', userHistory)
      
      if (userHistory.length > 0) {
        const latestUser = userHistory[userHistory.length - 1] // 获取最新记录
        console.log('[UserStore] 最新用户记录:', latestUser)
        console.log('[UserStore] 用户角色:', latestUser.roles)
        currentUser.value = latestUser
      } else {
        console.log('[UserStore] 未找到用户历史记录')
        currentUser.value = null
      }
    } catch (error) {
      console.error('[UserStore] 获取当前用户信息失败:', error)
      currentUser.value = null
    }
  }

  const getUserByAddress = (address: string) => {
    return users.value.find(user => user.addr.toLowerCase() === address.toLowerCase())
  }

  const getRoleNames = (roles: number[]) => {
    const roleMap = {
      0: '无',
      1: '监管者',
      2: '项目所有者', 
      3: '验证者',
      4: '购买者'
    }
    return roles.map(role => roleMap[role as keyof typeof roleMap]).join(', ')
  }

  return {
    users,
    currentUser,
    loading,
    currentUserRoles,
    isRegulator,
    isProjectOwner,
    isVerifier,
    isBuyer,
    fetchUsers,
    fetchCurrentUser,
    getUserByAddress,
    getRoleNames
  }
})
