/**
 * 权限控制服务
 * 基于智能合约中的角色权限系统
 */

import type { UserRecord, ProjectRecord, CreditRecord, CertificateRecord } from '@/types'

export enum Role {
  NONE = 0,
  REGULATOR = 1,
  PROJECT_OWNER = 2,
  VERIFIER = 3,
  BUYER = 4
}

export class PermissionService {
  /**
   * 检查用户是否有指定角色
   */
  static hasRole(user: UserRecord | null, role: Role): boolean {
    if (!user || !user.roles) return false
    return user.roles.includes(role as number)  // 确保类型匹配
  }

  /**
   * 检查用户是否是监管者
   */
  static isRegulator(user: UserRecord | null): boolean {
    return this.hasRole(user, Role.REGULATOR)
  }

  /**
   * 检查用户是否是项目方
   */
  static isProjectOwner(user: UserRecord | null): boolean {
    return this.hasRole(user, Role.PROJECT_OWNER)
  }

  /**
   * 检查用户是否是审核员
   */
  static isVerifier(user: UserRecord | null): boolean {
    return this.hasRole(user, Role.VERIFIER)
  }

  /**
   * 检查用户是否是购买方
   */
  static isBuyer(user: UserRecord | null): boolean {
    return this.hasRole(user, Role.BUYER)
  }

  /**
   * 检查用户是否有任何有效角色
   */
  static hasAnyRole(user: UserRecord | null): boolean {
    if (!user || !user.roles) return false
    return user.roles.some(role => role !== Role.NONE)
  }

  /**
   * 过滤用户可以查看的项目
   */
  static filterUserProjects(projects: ProjectRecord[], currentUser: UserRecord | null, currentUserAddress: string): ProjectRecord[] {
    if (!currentUser) return []

    // 监管者可以查看所有项目
    if (this.isRegulator(currentUser)) {
      return projects
    }

    // 审核员可以查看待审且未被锁定的项目 + 自己拥有的项目
    if (this.isVerifier(currentUser)) {
      return projects.filter(project => {
        // 自己拥有的项目
        if (project.owner.toLowerCase() === currentUserAddress.toLowerCase()) {
          return true
        }
        // 待审核的项目（状态为REVIEWING=1）
        return project.status === 1
      })
    }

    // 项目方和购买方只能查看自己相关的项目
    return projects.filter(project => 
      project.owner.toLowerCase() === currentUserAddress.toLowerCase()
    )
  }

  /**
   * 过滤用户可以查看的碳信用
   */
  static filterUserCredits(credits: CreditRecord[], currentUser: UserRecord | null, currentUserAddress: string): CreditRecord[] {
    if (!currentUser) return []

    // 监管者可以查看所有碳信用
    if (this.isRegulator(currentUser)) {
      return credits
    }

    // 其他角色只能查看自己拥有的碳信用
    return credits.filter(credit => 
      credit.currentOwner.toLowerCase() === currentUserAddress.toLowerCase()
    )
  }

  /**
   * 过滤用户可以查看的证书
   */
  static filterUserCertificates(certificates: CertificateRecord[], currentUser: UserRecord | null, currentUserAddress: string): CertificateRecord[] {
    if (!currentUser) return []

    // 监管者可以查看所有证书
    if (this.isRegulator(currentUser)) {
      return certificates
    }

    // 其他角色只能查看自己拥有的证书
    return certificates.filter(certificate => 
      certificate.owner.toLowerCase() === currentUserAddress.toLowerCase()
    )
  }

  /**
   * 检查用户是否可以编辑项目
   */
  static canEditProject(project: ProjectRecord, currentUser: UserRecord | null, currentUserAddress: string): boolean {
    if (!currentUser) return false

    // 监管者可以编辑任何项目
    if (this.isRegulator(currentUser)) {
      return true
    }

    // 项目所有者可以编辑自己的项目（仅在EDITING状态）
    if (project.owner.toLowerCase() === currentUserAddress.toLowerCase()) {
      return project.status === 0 // EDITING状态
    }

    return false
  }

  /**
   * 检查用户是否可以提交项目审核
   */
  static canSubmitForVerification(project: ProjectRecord, currentUser: UserRecord | null, currentUserAddress: string): boolean {
    if (!currentUser) return false

    // 只有项目所有者可以提交自己的项目进行审核
    return project.owner.toLowerCase() === currentUserAddress.toLowerCase() && 
           project.status === 0 && // EDITING状态
           this.isProjectOwner(currentUser)
  }

  /**
   * 检查用户是否可以审核项目
   */
  static canVerifyProject(project: ProjectRecord, currentUser: UserRecord | null): boolean {
    if (!currentUser) return false

    // 只有审核员可以审核项目，且项目必须处于待审核状态
    return this.isVerifier(currentUser) && project.status === 1 // REVIEWING状态
  }

  /**
   * 检查用户是否可以转移碳信用
   */
  static canTransferCredit(credit: CreditRecord, currentUser: UserRecord | null, currentUserAddress: string): boolean {
    if (!currentUser) return false

    // 只有碳信用的当前所有者可以转移，且碳信用未被撤销或退役
    return credit.currentOwner.toLowerCase() === currentUserAddress.toLowerCase() &&
           credit.eventType !== 2 && // 非RETIRE
           credit.eventType !== 3    // 非REVOKE
  }

  /**
   * 检查用户是否可以退役碳信用
   */
  static canRetireCredit(credit: CreditRecord, currentUser: UserRecord | null, currentUserAddress: string): boolean {
    if (!currentUser) return false

    // 只有碳信用的当前所有者可以退役
    return credit.currentOwner.toLowerCase() === currentUserAddress.toLowerCase() &&
           credit.eventType !== 2 && // 非RETIRE
           credit.eventType !== 3 && // 非REVOKE
           (this.isBuyer(currentUser) || this.isProjectOwner(currentUser))
  }

  /**
   * 检查用户是否可以撤销碳信用
   */
  static canRevokeCredit(currentUser: UserRecord | null): boolean {
    if (!currentUser) return false
    
    // 只有监管者可以撤销碳信用
    return this.isRegulator(currentUser)
  }

  /**
   * 检查用户是否可以撤销证书
   */
  static canRevokeCertificate(currentUser: UserRecord | null): boolean {
    if (!currentUser) return false
    
    // 只有监管者可以撤销证书
    return this.isRegulator(currentUser)
  }

  /**
   * 检查用户是否可以设置其他用户的角色
   */
  static canManageUserRoles(currentUser: UserRecord | null): boolean {
    if (!currentUser) return false
    
    // 只有监管者可以管理用户角色
    return this.isRegulator(currentUser)
  }

  /**
   * 获取用户可用的菜单项
   */
  static getAvailableMenuItems(currentUser: UserRecord | null) {
    const menuItems = []

    // 概览页面 - 所有已认证用户都可以访问
    if (this.hasAnyRole(currentUser)) {
      menuItems.push({
        key: '/dashboard',
        label: '概览',
        icon: 'DashboardOutlined'
      })
    }

    // 用户管理 - 只有监管者可以访问
    if (this.isRegulator(currentUser)) {
      menuItems.push({
        key: '/users',
        label: '用户管理',
        icon: 'UserOutlined'
      })
    }

    // 项目管理 - 所有角色都可以访问（但看到的内容不同）
    if (this.hasAnyRole(currentUser)) {
      menuItems.push({
        key: '/projects',
        label: '项目管理',
        icon: 'ProjectOutlined'
      })
    }

    // 碳信用管理 - 所有角色都可以访问（但看到的内容不同）
    if (this.hasAnyRole(currentUser)) {
      menuItems.push({
        key: '/credits',
        label: '碳信用管理',
        icon: 'GoldOutlined'
      })
    }

    // 证书管理 - 所有角色都可以访问（但看到的内容不同）
    if (this.hasAnyRole(currentUser)) {
      menuItems.push({
        key: '/certificates',
        label: '证书管理',
        icon: 'SafetyCertificateOutlined'
      })
    }

    return menuItems
  }

  /**
   * 获取角色显示名称
   */
  static getRoleDisplayName(role: number): string {
    const roleNames = {
      0: '无',
      1: '监管者',
      2: '项目方',
      3: '审核员',
      4: '购买方'
    }
    return roleNames[role as keyof typeof roleNames] || '未知'
  }

  /**
   * 获取用户所有角色的显示名称
   */
  static getUserRoleNames(user: UserRecord | null): string {
    if (!user || !user.roles || !Array.isArray(user.roles)) return '无角色'
    
    console.log('[PermissionService] getUserRoleNames - 用户角色:', user.roles)
    const roleNames = user.roles.map(role => {
      const displayName = this.getRoleDisplayName(role)
      console.log(`[PermissionService] 角色 ${role} -> ${displayName}`)
      return displayName
    })
    
    const result = roleNames.join(', ')
    console.log('[PermissionService] 最终角色名称:', result)
    return result
  }
}
