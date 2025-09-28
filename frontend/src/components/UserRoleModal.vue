<template>
  <a-modal
    v-model:open="isVisible"
    title="权限管理"
    width="500px"
    :mask-closable="false"
    @ok="handleSubmit"
    @cancel="handleCancel"
    :confirm-loading="loading"
  >
    <template #footer>
      <a-space>
        <a-button @click="handleCancel" :disabled="loading">
          取消
        </a-button>
        <a-button 
          type="primary" 
          @click="handleSubmit" 
          :loading="loading"
          :disabled="!isFormValid"
        >
          更新权限
        </a-button>
      </a-space>
    </template>

    <div v-if="user" class="role-management">
      <!-- 用户信息 -->
      <div class="user-info">
        <DefaultAvatar :name="user.name || '用户'" :size="50" />
        <div class="user-details">
          <div class="user-name">{{ user.name || '未设置用户名' }}</div>
          <div class="user-address">{{ formatAddress(user.addr) }}</div>
        </div>
      </div>

      <a-divider />

      <a-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        layout="vertical"
        @finish="onFinish"
      >
        <a-form-item label="当前权限">
          <a-space wrap>
            <a-tag
              v-for="role in user.roles"
              :key="role"
              :color="getRoleColor(role)"
            >
              {{ PermissionService.getRoleDisplayName(role) }}
            </a-tag>
          </a-space>
        </a-form-item>

        <a-form-item
          name="roles"
          label="新权限"
          :rules="rules.roles"
        >
          <a-checkbox-group v-model:value="formData.roles">
            <div class="role-options">
              <a-checkbox :value="1" class="role-option">
                <div class="role-item">
                  <a-tag color="red">监管者</a-tag>
                  <span class="role-desc">拥有平台最高管理权限</span>
                </div>
              </a-checkbox>
              
              <a-checkbox :value="2" class="role-option">
                <div class="role-item">
                  <a-tag color="blue">项目方</a-tag>
                  <span class="role-desc">可以创建和管理碳减排项目</span>
                </div>
              </a-checkbox>
              
              <a-checkbox :value="3" class="role-option">
                <div class="role-item">
                  <a-tag color="green">审核员</a-tag>
                  <span class="role-desc">负责审核项目和核发碳信用</span>
                </div>
              </a-checkbox>
              
              <a-checkbox :value="4" class="role-option">
                <div class="role-item">
                  <a-tag color="orange">购买方</a-tag>
                  <span class="role-desc">可以购买和退役碳信用</span>
                </div>
              </a-checkbox>
            </div>
          </a-checkbox-group>
        </a-form-item>

        <a-form-item
          name="editReason"
          label="编辑原因"
          :rules="rules.editReason"
        >
          <a-textarea
            v-model:value="formData.editReason"
            placeholder="请说明本次权限变更的原因"
            :rows="3"
            :maxlength="200"
            show-count
          />
        </a-form-item>
      </a-form>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import type { UserRecord } from '@/types'
import { adminSetUserRoles } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import DefaultAvatar from './DefaultAvatar.vue'

interface Props {
  visible: boolean
  user?: UserRecord | null
}

interface Emits {
  (e: 'update:visible', value: boolean): void
  (e: 'success'): void
}

const props = withDefaults(defineProps<Props>(), {
  user: null
})

const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const loading = ref(false)

// 表单数据
const formData = ref({
  roles: [] as number[],
  editReason: ''
})

// 表单验证规则
const rules = {
  roles: [
    { required: true, message: '请选择至少一个权限', trigger: 'change', type: 'array', min: 1 }
  ],
  editReason: [
    { required: true, message: '请说明编辑原因', trigger: 'blur' },
    { min: 5, max: 200, message: '编辑原因应在5-200字符之间', trigger: 'blur' }
  ]
}

// 控制弹窗显示
const isVisible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value)
})

// 表单验证状态
const isFormValid = computed(() => {
  return formData.value.roles.length > 0 && formData.value.editReason.trim().length >= 5
})

// 监听用户数据变化，初始化表单
watch(
  () => props.user,
  (user) => {
    if (user) {
      formData.value = {
        roles: [...(user.roles || [])],
        editReason: ''
      }
    }
  },
  { immediate: true }
)

// 监听弹窗显示状态
watch(
  () => props.visible,
  (visible) => {
    if (visible && props.user) {
      formData.value = {
        roles: [...(props.user.roles || [])],
        editReason: ''
      }
    }
  }
)

// 处理表单提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    await onFinish()
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}

// 表单提交处理
const onFinish = async () => {
  if (!props.user) {
    message.error('用户信息缺失')
    return
  }

  try {
    loading.value = true
    
    await adminSetUserRoles(
      props.user.addr,
      formData.value.roles,
      formData.value.editReason.trim()
    )

    message.success('用户权限更新成功')
    emit('success')
  } catch (error: any) {
    console.error('更新用户权限失败:', error)
    message.error(error.message || '更新失败，请重试')
  } finally {
    loading.value = false
  }
}

// 取消处理
const handleCancel = () => {
  isVisible.value = false
}

// 工具函数
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
</script>

<style scoped>
.role-management {
  padding: 8px 0;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.user-details {
  flex: 1;
}

.user-name {
  font-size: 16px;
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.user-address {
  font-size: 12px;
  color: #8c8c8c;
  font-family: monospace;
}

.role-options {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.role-option {
  width: 100%;
}

.role-option :deep(.ant-checkbox) {
  align-self: flex-start;
  margin-top: 4px;
}

.role-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-left: 8px;
}

.role-desc {
  font-size: 12px;
  color: #666;
  line-height: 1.4;
}

:deep(.ant-checkbox-group) {
  width: 100%;
}

:deep(.ant-checkbox-wrapper) {
  display: flex;
  align-items: flex-start;
  width: 100%;
  padding: 12px;
  border: 1px solid #f0f0f0;
  border-radius: 6px;
  transition: all 0.3s;
}

:deep(.ant-checkbox-wrapper:hover) {
  border-color: #40a9ff;
  background-color: #f6ffed;
}

:deep(.ant-checkbox-wrapper-checked) {
  border-color: #1890ff;
  background-color: #e6f7ff;
}
</style>
