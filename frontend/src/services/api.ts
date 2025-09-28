import axios from 'axios'
import type {
  UserRecord,
  ProjectRecord,
  CreditRecord,
  CertificateRecord,
  ReviewRecord,
  PlatformSummary,
  ApiResponse
} from '@/types'
import { PermissionService } from './permission'

// 创建axios实例
const api = axios.create({
  baseURL: '/api',
  timeout: 15000 // 减少超时时间
})

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    console.error('API Error:', error)
    return Promise.reject(error)
  }
)

// 缓存管理
const cache = new Map<string, { data: any; timestamp: number; ttl: number }>()
const pendingRequests = new Map<string, Promise<any>>()

// 缓存配置
const CACHE_TTL = {
  SUMMARY: 30000,      // 30秒
  USERS: 60000,        // 1分钟
  PROJECTS: 60000,     // 1分钟
  CREDITS: 60000,      // 1分钟
  CERTIFICATES: 60000, // 1分钟
  HISTORY: 120000      // 2分钟
}

// 缓存工具函数
const getCacheKey = (endpoint: string, params?: any): string => {
  return params ? `${endpoint}_${JSON.stringify(params)}` : endpoint
}

const getCachedData = (key: string): any | null => {
  const cached = cache.get(key)
  if (cached && (Date.now() - cached.timestamp) < cached.ttl) {
    return cached.data
  }
  cache.delete(key)
  return null
}

const setCachedData = (key: string, data: any, ttl: number): void => {
  cache.set(key, {
    data,
    timestamp: Date.now(),
    ttl
  })
}

// 请求去重函数
const dedupeRequest = async <T>(key: string, requestFn: () => Promise<T>): Promise<T> => {
  if (pendingRequests.has(key)) {
    return pendingRequests.get(key)!
  }

  const promise = requestFn()
  pendingRequests.set(key, promise)

  try {
    const result = await promise
    return result
  } finally {
    pendingRequests.delete(key)
  }
}

// =================================================================================
// 只读操作API - 用于从后端获取区块链数据
// 注意：所有写操作（用户注册、项目创建、转账等）都应该通过 web3.ts 直接与智能合约交互
// =================================================================================

// 平台概览API
export const getSummary = (): Promise<PlatformSummary> => {
  const cacheKey = getCacheKey('/summary')
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get('/summary')
    setCachedData(cacheKey, data, CACHE_TTL.SUMMARY)
    return data
  })
}

// 用户管理API（只读）
export const getAllUsers = async (): Promise<UserRecord[]> => {
  const cacheKey = getCacheKey('/users')
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    console.log('[API] 正在获取所有用户列表')
    try {
      const response = await api.get('/users')
      console.log('[API] 所有用户列表响应:', response)
      const data = response as unknown as UserRecord[]
      setCachedData(cacheKey, data, CACHE_TTL.USERS)
      return data
    } catch (error) {
      console.error('[API] 获取用户列表失败:', error)
      throw error
    }
  })
}

export const getUserHistory = async (address: string): Promise<UserRecord[]> => {
  const cacheKey = getCacheKey(`/users/${address}`)
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    console.log('[API] 正在获取用户历史记录:', address)
    try {
      const response = await api.get(`/users/${address}`)
      console.log('[API] 用户历史记录响应:', response)
      const data = response as unknown as UserRecord[]
      setCachedData(cacheKey, data, CACHE_TTL.HISTORY)
      return data
    } catch (error) {
      console.error('[API] 获取用户历史记录失败:', error)
      throw error
    }
  })
}

// 项目管理API（只读）
export const getAllProjects = (): Promise<ProjectRecord[]> => {
  const cacheKey = getCacheKey('/projects')
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get('/projects')
    setCachedData(cacheKey, data, CACHE_TTL.PROJECTS)
    return data
  })
}

export const getProjectHistory = (id: number): Promise<ProjectRecord[]> => {
  const cacheKey = getCacheKey(`/projects/${id}`)
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get(`/projects/${id}`)
    setCachedData(cacheKey, data, CACHE_TTL.HISTORY)
    return data
  })
}

export const getProjectReviews = (id: number): Promise<ReviewRecord[]> => {
  const cacheKey = getCacheKey(`/projects/${id}/reviews`)
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get(`/projects/${id}/reviews`)
    setCachedData(cacheKey, data, CACHE_TTL.HISTORY)
    return data
  })
}

// 碳信用管理API（只读）
export const getAllCredits = (): Promise<CreditRecord[]> => {
  const cacheKey = getCacheKey('/credits')
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get('/credits')
    setCachedData(cacheKey, data, CACHE_TTL.CREDITS)
    return data
  })
}

export const getCreditHistory = (id: number): Promise<CreditRecord[]> => {
  const cacheKey = getCacheKey(`/credits/${id}`)
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get(`/credits/${id}`)
    setCachedData(cacheKey, data, CACHE_TTL.HISTORY)
    return data
  })
}

// 证书管理API（只读）
export const getAllCertificates = (): Promise<CertificateRecord[]> => {
  const cacheKey = getCacheKey('/certificates')
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get('/certificates')
    setCachedData(cacheKey, data, CACHE_TTL.CERTIFICATES)
    return data
  })
}

export const getCertificateHistory = (id: number): Promise<CertificateRecord[]> => {
  const cacheKey = getCacheKey(`/certificates/${id}`)
  const cached = getCachedData(cacheKey)
  
  if (cached) {
    return Promise.resolve(cached)
  }

  return dedupeRequest(cacheKey, async () => {
    const data = await api.get(`/certificates/${id}`)
    setCachedData(cacheKey, data, CACHE_TTL.HISTORY)
    return data
  })
}

// 基于用户权限的数据获取API（只读）
export const getUserProjects = async (currentUser: UserRecord | null, currentUserAddress: string): Promise<ProjectRecord[]> => {
  const allProjects = await getAllProjects()
  return PermissionService.filterUserProjects(allProjects, currentUser, currentUserAddress)
}

export const getUserCredits = async (currentUser: UserRecord | null, currentUserAddress: string): Promise<CreditRecord[]> => {
  const allCredits = await getAllCredits()
  return PermissionService.filterUserCredits(allCredits, currentUser, currentUserAddress)
}

export const getUserCertificates = async (currentUser: UserRecord | null, currentUserAddress: string): Promise<CertificateRecord[]> => {
  const allCertificates = await getAllCertificates()
  return PermissionService.filterUserCertificates(allCertificates, currentUser, currentUserAddress)
}

// 缓存管理函数
export const clearCache = (pattern?: string) => {
  if (pattern) {
    // 清除匹配模式的缓存
    for (const key of cache.keys()) {
      if (key.includes(pattern)) {
        cache.delete(key)
      }
    }
  } else {
    // 清除所有缓存
    cache.clear()
  }
}

export const getCacheStats = () => {
  return {
    size: cache.size,
    keys: Array.from(cache.keys())
  }
}

// 定期清理过期缓存
setInterval(() => {
  const now = Date.now()
  for (const [key, value] of cache.entries()) {
    if (now - value.timestamp > value.ttl) {
      cache.delete(key)
    }
  }
}, 60000) // 每分钟清理一次

export default api
