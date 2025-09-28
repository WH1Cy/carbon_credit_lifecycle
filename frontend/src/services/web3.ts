import Web3 from 'web3'
import detectEthereumProvider from '@metamask/detect-provider'
import { ref, reactive } from 'vue'
import type { WalletInfo } from '@/types'

// 合约配置
const CONTRACT_ADDRESS = '0x71f9cbd442304a86ba905968fea9d250d9f08b04'
const SEPOLIA_CHAIN_ID = '0xaa36a7' // 11155111 in hex

// 状态管理
export const walletInfo = reactive<WalletInfo>({
  address: '',
  isConnected: false,
  chainId: 0
})

export const isConnecting = ref(false)
export const availableAccounts = ref<string[]>([])

let web3: Web3 | null = null
let contract: any = null

// 合约ABI（使用完整ABI）
const CONTRACT_ABI = [
  {
    "inputs": [],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "certificateId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "recordIndex",
        "type": "uint256"
      }
    ],
    "name": "CertificateHistoryUpdated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "batchId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "recordIndex",
        "type": "uint256"
      }
    ],
    "name": "CreditHistoryUpdated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "projectId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "recordIndex",
        "type": "uint256"
      }
    ],
    "name": "ProjectHistoryUpdated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "projectId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "recordIndex",
        "type": "uint256"
      }
    ],
    "name": "ReviewHistoryUpdated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "userAddress",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "recordIndex",
        "type": "uint256"
      }
    ],
    "name": "UserHistoryUpdated",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_projectId",
        "type": "uint256"
      }
    ],
    "name": "acceptVerification",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_user",
        "type": "address"
      },
      {
        "internalType": "enum CarbonTrustConnect.Role[]",
        "name": "_newRoles",
        "type": "uint8[]"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "adminSetUserRoles",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_projectId",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_location",
        "type": "string"
      },
      {
        "internalType": "enum CarbonTrustConnect.Technology[]",
        "name": "_technologies",
        "type": "uint8[]"
      },
      {
        "internalType": "string",
        "name": "_description",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "_reduction",
        "type": "uint256"
      },
      {
        "internalType": "string[]",
        "name": "_documentsHash",
        "type": "string[]"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "adminUpdateProjectDetails",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_user",
        "type": "address"
      },
      {
        "internalType": "string",
        "name": "_name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_bio",
        "type": "string"
      },
      {
        "internalType": "string[]",
        "name": "_credentialsHash",
        "type": "string[]"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "adminUpdateUserProfile",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_projectId",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "banProject",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_projectId",
        "type": "uint256"
      },
      {
        "internalType": "bool",
        "name": "_isApproved",
        "type": "bool"
      },
      {
        "internalType": "string",
        "name": "_comment",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "_creditQuantityToIssue",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_vintageYear",
        "type": "uint256"
      }
    ],
    "name": "finalizeVerification",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_buyer",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_quantity",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_vintageYear",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "mintCreditsForBuyer",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "_name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_location",
        "type": "string"
      },
      {
        "internalType": "enum CarbonTrustConnect.Technology[]",
        "name": "_technologies",
        "type": "uint8[]"
      },
      {
        "internalType": "string",
        "name": "_description",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "_reduction",
        "type": "uint256"
      },
      {
        "internalType": "string[]",
        "name": "_documentsHash",
        "type": "string[]"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "registerProject",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_batchId",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_quantity",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_retirementReason",
        "type": "string"
      }
    ],
    "name": "retireCredits",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_certificateId",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "revokeCertificate",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_batchId",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "revokeCredits",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_batchId",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_quantity",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "_to",
        "type": "address"
      }
    ],
    "name": "splitAndTransfer",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_projectId",
        "type": "uint256"
      }
    ],
    "name": "submitForVerification",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_projectId",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_location",
        "type": "string"
      },
      {
        "internalType": "enum CarbonTrustConnect.Technology[]",
        "name": "_technologies",
        "type": "uint8[]"
      },
      {
        "internalType": "string",
        "name": "_description",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "_reduction",
        "type": "uint256"
      },
      {
        "internalType": "string[]",
        "name": "_documentsHash",
        "type": "string[]"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "updateProjectDetails",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "_name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_bio",
        "type": "string"
      },
      {
        "internalType": "string[]",
        "name": "_credentialsHash",
        "type": "string[]"
      },
      {
        "internalType": "string",
        "name": "_reason",
        "type": "string"
      }
    ],
    "name": "updateUserProfile",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]

// 初始化Web3连接
export const initWeb3 = async () => {
  try {
    const provider = await detectEthereumProvider()
    if (!provider) {
      throw new Error('请安装MetaMask钱包')
    }

    web3 = new Web3(provider as any)
    contract = new web3.eth.Contract(CONTRACT_ABI as any, CONTRACT_ADDRESS)

    // 监听账户变化
    provider.on('accountsChanged', (accounts: string[]) => {
      if (accounts.length === 0) {
        disconnectWallet()
      } else {
        walletInfo.address = accounts[0]
        availableAccounts.value = accounts
      }
    })

    // 监听网络变化
    provider.on('chainChanged', (chainId: string) => {
      walletInfo.chainId = parseInt(chainId, 16)
      if (chainId !== SEPOLIA_CHAIN_ID) {
        console.warn('请切换到Sepolia测试网')
      }
    })

    return true
  } catch (error) {
    console.error('Web3初始化失败:', error)
    return false
  }
}

// 连接钱包
export const connectWallet = async () => {
  if (!web3) {
    const initialized = await initWeb3()
    if (!initialized) {
      throw new Error('Web3初始化失败')
    }
  }

  try {
    isConnecting.value = true
    
    // 请求连接账户
    const accounts = await window.ethereum.request({
      method: 'eth_requestAccounts'
    })

    if (accounts.length === 0) {
      throw new Error('未获取到账户')
    }

    // 检查网络
    const chainId = await window.ethereum.request({
      method: 'eth_chainId'
    })

    if (chainId !== SEPOLIA_CHAIN_ID) {
      // 尝试切换到Sepolia网络
      try {
        await window.ethereum.request({
          method: 'wallet_switchEthereumChain',
          params: [{ chainId: SEPOLIA_CHAIN_ID }]
        })
      } catch (switchError: any) {
        if (switchError.code === 4902) {
          // 网络不存在，添加网络
          await window.ethereum.request({
            method: 'wallet_addEthereumChain',
            params: [{
              chainId: SEPOLIA_CHAIN_ID,
              chainName: 'Sepolia Test Network',
              rpcUrls: ['https://sepolia.infura.io/v3/22c489e1df7a458181cc8ddc54e70f1b'],
              nativeCurrency: {
                name: 'SepoliaETH',
                symbol: 'SepoliaETH',
                decimals: 18
              },
              blockExplorerUrls: ['https://sepolia.etherscan.io']
            }]
          })
        } else {
          throw switchError
        }
      }
    }

    // 更新状态
    walletInfo.address = accounts[0]
    walletInfo.isConnected = true
    walletInfo.chainId = parseInt(chainId, 16)
    availableAccounts.value = accounts

    return accounts[0]
  } catch (error: any) {
    console.error('连接钱包失败:', error)
    if (error.code === 4001) {
      throw new Error('用户拒绝了连接请求')
    } else {
      throw new Error(error.message || '连接钱包失败')
    }
  } finally {
    isConnecting.value = false
  }
}

// 断开钱包连接
export const disconnectWallet = () => {
  walletInfo.address = ''
  walletInfo.isConnected = false
  walletInfo.chainId = 0
  availableAccounts.value = []
}

// 交易状态管理
const pendingTransactions = new Map<string, Promise<any>>()

// 签名并发送交易的通用函数（带防重复提交）
export const signAndSendTransaction = async (
  functionName: string,
  params: any[],
  value: string = '0'
) => {
  if (!web3 || !contract || !walletInfo.isConnected) {
    throw new Error('请先连接钱包')
  }

  // 创建交易唯一标识
  const txKey = `${functionName}_${JSON.stringify(params)}_${walletInfo.address}`
  
  // 检查是否有相同的交易正在进行
  if (pendingTransactions.has(txKey)) {
    throw new Error('相同交易正在进行中，请稍后再试')
  }

  const transactionPromise = executeTransaction(functionName, params, value)
  pendingTransactions.set(txKey, transactionPromise)

  try {
    const result = await transactionPromise
    return result
  } finally {
    pendingTransactions.delete(txKey)
  }
}

// 执行交易的核心函数
const executeTransaction = async (
  functionName: string,
  params: any[],
  value: string = '0'
) => {
  try {
    // 估算gas（添加超时）
    const gasEstimateRaw = await Promise.race([
      contract.methods[functionName](...params).estimateGas({
        from: walletInfo.address,
        value: web3!.utils.toWei(value, 'ether')
      }),
      new Promise((_, reject) => 
        setTimeout(() => reject(new Error('Gas估算超时')), 10000)
      )
    ])

    // 确保gasEstimate是number类型
    const gasEstimate = typeof gasEstimateRaw === 'bigint' 
      ? Number(gasEstimateRaw) 
      : Number(gasEstimateRaw)

    // 获取gas价格（添加缓存）
    const gasPrice = await getCachedGasPrice()

    // 构建交易对象
    const txObject = {
      from: walletInfo.address,
      to: CONTRACT_ADDRESS,
      data: contract!.methods[functionName](...params).encodeABI(),
      gas: web3!.utils.toHex(Math.floor(gasEstimate * 1.2)), // 增加20%的gas余量
      gasPrice: web3!.utils.toHex(gasPrice),
      value: web3!.utils.toHex(web3!.utils.toWei(value, 'ether'))
    }

    console.log('发送交易:', { functionName, params, txObject })

    // 直接通过MetaMask发送交易
    const txHash = await window.ethereum.request({
      method: 'eth_sendTransaction',
      params: [txObject]
    })

    console.log('交易已发送:', txHash)
    
    // 等待交易确认（优化等待逻辑）
    const receipt = await waitForTransactionReceipt(txHash)

    console.log('交易确认成功:', receipt)
    return {
      message: '交易成功',
      transaction_hash: txHash,
      receipt
    }
  } catch (error: any) {
    console.error('交易失败:', error)
    if (error.code === 4001) {
      throw new Error('用户取消了交易')
    } else if (error.code === -32603) {
      throw new Error('交易执行失败，请检查参数或余额')
    } else {
      throw new Error(error.message || '交易失败')
    }
  }
}

// Gas价格缓存
let gasPriceCache: { price: string; timestamp: number } | null = null
const GAS_PRICE_CACHE_DURATION = 30000 // 30秒缓存

const getCachedGasPrice = async (): Promise<string> => {
  const now = Date.now()
  
  if (gasPriceCache && (now - gasPriceCache.timestamp) < GAS_PRICE_CACHE_DURATION) {
    return gasPriceCache.price
  }

  const price = await web3!.eth.getGasPrice()
  const priceString = price.toString()
  gasPriceCache = { price: priceString, timestamp: now }
  return priceString
}

// 优化的交易确认等待
const waitForTransactionReceipt = async (txHash: string): Promise<any> => {
  const maxAttempts = 30 // 减少等待次数
  const interval = 3000 // 增加等待间隔到3秒

  for (let attempt = 0; attempt < maxAttempts; attempt++) {
    try {
      const receipt = await web3!.eth.getTransactionReceipt(txHash)
      if (receipt) {
        if (!receipt.status) {
          throw new Error('交易执行失败')
        }
        return receipt
      }
    } catch (error) {
      // 交易可能还未上链，继续等待
    }
    
    await new Promise(resolve => setTimeout(resolve, interval))
  }

  throw new Error('交易确认超时，请检查区块链浏览器')
}

// 业务函数封装
export const updateUserProfile = (name: string, bio: string, credentials: string[], reason: string) => {
  return signAndSendTransaction('updateUserProfile', [name, bio, credentials, reason])
}

export const registerProject = (
  name: string,
  location: string,
  technologies: number[],
  description: string,
  reduction: number,
  documents: string[],
  reason: string
) => {
  console.log('调用registerProject，参数:', {
    name,
    location,
    technologies,
    description,
    reduction,
    documents,
    reason
  })
  
  // 确保reduction是有效的数字
  const validReduction = Number(reduction)
  if (isNaN(validReduction) || validReduction <= 0) {
    throw new Error('年减排量必须是大于0的数字')
  }
  
  return signAndSendTransaction('registerProject', [
    name, 
    location, 
    technologies, 
    description, 
    validReduction, 
    documents, 
    reason
  ])
}

export const submitProjectForVerification = (projectId: number) => {
  return signAndSendTransaction('submitForVerification', [projectId])
}

export const acceptVerification = (projectId: number) => {
  return signAndSendTransaction('acceptVerification', [projectId])
}

export const finalizeVerification = (
  projectId: number,
  isApproved: boolean,
  comment: string,
  creditQuantity: number,
  vintageYear: number
) => {
  return signAndSendTransaction('finalizeVerification', [
    projectId, isApproved, comment, creditQuantity, vintageYear
  ])
}

export const transferCredits = (batchId: number, quantity: number, toAddress: string) => {
  return signAndSendTransaction('splitAndTransfer', [batchId, quantity, toAddress])
}

export const retireCredits = (batchId: number, quantity: number, reason: string) => {
  return signAndSendTransaction('retireCredits', [batchId, quantity, reason])
}

// 管理员功能
export const adminSetUserRoles = (userAddress: string, roles: number[], reason: string) => {
  return signAndSendTransaction('adminSetUserRoles', [userAddress, roles, reason])
}

export const adminUpdateUserProfile = (
  userAddress: string,
  name: string,
  bio: string,
  credentials: string[],
  reason: string
) => {
  return signAndSendTransaction('adminUpdateUserProfile', [userAddress, name, bio, credentials, reason])
}

export const adminUpdateProjectDetails = (
  projectId: number,
  name: string,
  location: string,
  technologies: number[],
  description: string,
  reduction: number,
  documents: string[],
  reason: string
) => {
  return signAndSendTransaction('adminUpdateProjectDetails', [
    projectId, name, location, technologies, description, reduction, documents, reason
  ])
}

export const banProject = (projectId: number, reason: string) => {
  return signAndSendTransaction('banProject', [projectId, reason])
}

export const mintCreditsForBuyer = (buyerAddress: string, quantity: number, vintageYear: number, reason: string) => {
  return signAndSendTransaction('mintCreditsForBuyer', [buyerAddress, quantity, vintageYear, reason])
}

export const revokeCredits = (batchId: number, reason: string) => {
  return signAndSendTransaction('revokeCredits', [batchId, reason])
}

// 证书管理功能
export const revokeCertificate = (certificateId: number, reason: string) => {
  return signAndSendTransaction('revokeCertificate', [certificateId, reason])
}

// 项目更新功能
export const updateProjectDetails = (
  projectId: number,
  name: string,
  location: string,
  technologies: number[],
  description: string,
  reduction: number,
  documents: string[],
  reason: string
) => {
  return signAndSendTransaction('updateProjectDetails', [
    projectId, name, location, technologies, description, reduction, documents, reason
  ])
}