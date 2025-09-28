<template>
  <div class="simple-layout">
    <div class="header">
      <h1>碳信通平台</h1>
      <div class="header-actions">
        <ThemeToggle />
        <a-button 
          v-if="!walletInfo.isConnected" 
          type="primary" 
          @click="handleConnectWallet"
          :loading="isConnecting"
        >
          连接钱包
        </a-button>
      </div>
    </div>
    
    <div class="content">
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'
import { walletInfo, isConnecting, connectWallet } from '@/services/web3'
import ThemeToggle from '@/components/ThemeToggle.vue'

const router = useRouter()
const userStore = useUserStore()
const { currentUser } = storeToRefs(userStore)

const handleConnectWallet = async () => {
  try {
    await connectWallet()
    await userStore.fetchCurrentUser()
  } catch (error: any) {
    console.error('钱包连接失败:', error)
  }
}
</script>

<style scoped>
.simple-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background: #ffffff;
  border-bottom: 1px solid #d0d7de;
  padding: 16px 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header h1 {
  margin: 0;
  color: #0969da;
  font-size: 20px;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.content {
  flex: 1;
  padding: 24px;
  background: #f6f8fa;
}
</style>
