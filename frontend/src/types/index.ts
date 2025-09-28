// 角色枚举
export enum Role {
  NONE = 0,
  REGULATOR = 1,
  PROJECT_OWNER = 2,
  VERIFIER = 3,
  BUYER = 4
}

// 技术类型枚举
export enum Technology {
  SOLAR = 0,
  WIND = 1,
  HYDRO = 2,
  FORESTRY = 3,
  CAPTURE = 4,
  OTHER = 5
}

// 项目状态枚举
export enum ProjectStatus {
  EDITING = 0,
  REVIEWING = 1,
  APPROVED = 2,
  REVOKED = 3
}

// 信用事件类型枚举
export enum CreditEventType {
  MINT = 0,
  TRANSFER = 1,
  RETIRE = 2,
  REVOKE = 3
}

// 证书状态枚举
export enum CertificateStatus {
  ACTIVE = 0,
  REVOKED = 1
}

// 审核意见枚举
export enum ReviewOpinion {
  APPROVE = 0,
  REJECT = 1,
  REQUEST_CHANGES = 2
}

// 用户记录
export interface UserRecord {
  timestamp: number
  addr: string
  name: string
  bio: string
  credentialsHash: string[]
  roles: number[]  // 改为number[]以匹配后端返回的uint8[]
  editor: string
  editorRecordIndex: number
  editReason: string
}

// 项目记录
export interface ProjectRecord {
  timestamp: number
  id: number
  owner: string
  ownerRecordIndex: number
  name: string
  location: string
  technologies: Technology[]
  description: string
  reduction: number
  status: ProjectStatus
  documentsHash: string[]
  editor: string
  editorRecordIndex: number
  editReason: string
}

// 审核记录
export interface ReviewRecord {
  id: number
  timestamp: number
  reviewer: string
  verifierRecordIndex: number
  opinion: ReviewOpinion
  comment: string
  issuedBatchId?: number
}

// 碳信用记录
export interface CreditRecord {
  timestamp: number
  eventType: CreditEventType
  projectId: number
  vintageYear: number
  quantity: number
  currentOwner: string
  ownerRecordIndex: number
  from: string
  fromRecordIndex: number
  to: string
  toRecordIndex: number
  parentBatchId: number
  childBatchId: number
  reason: string
}

// 证书记录
export interface CertificateRecord {
  timestamp: number
  id: number
  batchId: number // 证书批次ID
  projectId: number // 关联项目ID
  quantity: number
  status: CertificateStatus
  owner: string
  ownerRecordIndex: number
  issuedAt: number // 核发时间
  revokedAt?: number // 撤销时间（可选）
  reason?: string // 操作原因（可选）
  editor: string
  editorRecordIndex: number
  editReason: string
}

// 平台摘要
export interface PlatformSummary {
  totalUsers: number
  totalProjects: number
  totalCreditBatches: number
  totalCertificates: number
}

// API响应包装
export interface ApiResponse<T = any> {
  code?: number
  message?: string
  data?: T
}


// Web3钱包信息
export interface WalletInfo {
  address: string
  isConnected: boolean
  chainId: number
}
