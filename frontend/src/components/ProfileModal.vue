<template>
  <a-modal
    v-model:open="isVisible"
    :title="user ? '编辑个人资料' : '完善个人资料'"
    width="600px"
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
          {{ user ? '更新' : '保存' }}
        </a-button>
      </a-space>
    </template>

    <a-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      layout="vertical"
      @finish="onFinish"
    >
      <a-row :gutter="16">
        <a-col :span="8">
          <div class="avatar-section">
            <div class="avatar-preview">
              <DefaultAvatar :name="formData.name || '用户'" :size="80" />
            </div>
            <div class="avatar-info">
              <p>头像将根据用户名自动生成</p>
            </div>
          </div>
        </a-col>
        
        <a-col :span="16">
          <a-form-item
            name="name"
            label="用户名"
            :rules="rules.name"
          >
            <a-input
              v-model:value="formData.name"
              placeholder="请输入用户名"
              :maxlength="50"
            />
          </a-form-item>

          <a-form-item
            name="bio"
            label="个人简介"
          >
            <a-textarea
              v-model:value="formData.bio"
              placeholder="请输入个人简介（可选）"
              :rows="3"
              :maxlength="200"
              show-count
            />
          </a-form-item>

          <a-form-item
            name="credentials"
            label="资质文件哈希"
          >
            <a-textarea
              v-model:value="credentialsText"
              placeholder="请输入资质文件哈希，每行一个（可选）"
              :rows="3"
              :maxlength="500"
            />
            <div class="field-hint">
              请输入相关资质证明的文件哈希值，每行一个
            </div>
          </a-form-item>
        </a-col>
      </a-row>

      <!-- 编辑原因（仅更新时显示） -->
      <a-form-item
        v-if="user"
        name="editReason"
        label="编辑原因"
        :rules="rules.editReason"
      >
        <a-textarea
          v-model:value="formData.editReason"
          placeholder="请说明本次编辑的原因"
          :rows="2"
          :maxlength="200"
          show-count
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import type { UserRecord } from '@/types'
import { walletInfo, updateUserProfile } from '@/services/web3'
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
  name: '',
  bio: '',
  editReason: ''
})

const credentialsText = ref('')

// 表单验证规则
const rules = {
  name: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 2, max: 50, message: '用户名长度应在2-50字符之间', trigger: 'blur' }
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
  if (!formData.value.name?.trim()) return false
  if (props.user && !formData.value.editReason?.trim()) return false
  return true
})

// 监听用户数据变化，初始化表单
watch(
  () => props.user,
  (user) => {
    if (user) {
      formData.value = {
        name: user.name || '',
        bio: user.bio || '',
        editReason: ''
      }
      credentialsText.value = user.credentialsHash?.join('\n') || ''
    }
  },
  { immediate: true }
)

// 监听弹窗显示状态
watch(
  () => props.visible,
  (visible) => {
    if (visible && !props.user) {
      // 新用户注册，重置表单
      formData.value = {
        name: '',
        bio: '',
        editReason: ''
      }
      credentialsText.value = ''
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
  if (!walletInfo.address) {
    message.error('请先连接钱包')
    return
  }

  try {
    loading.value = true
    
    // 处理资质文件哈希
    const credentialsHash = credentialsText.value
      .split('\n')
      .map(line => line.trim())
      .filter(line => line.length > 0)

    const userData = {
      name: formData.value.name.trim(),
      bio: formData.value.bio.trim(),
      credentialsHash: credentialsHash
    }

    // 无论是新用户还是更新用户，都使用相同的智能合约方法
    await updateUserProfile(
      userData.name,
      userData.bio,
      userData.credentialsHash,
      formData.value.editReason.trim() || (props.user ? '更新个人资料' : '完善个人资料')
    )
    
    if (props.user) {
      message.success('个人资料更新成功')
    } else {
      message.success('个人资料创建成功')
    }

    emit('success')
  } catch (error: any) {
    console.error('保存用户信息失败:', error)
    message.error(error.message || '保存失败，请重试')
  } finally {
    loading.value = false
  }
}

// 取消处理
const handleCancel = () => {
  isVisible.value = false
}
</script>

<style scoped>
.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.avatar-preview {
  margin-bottom: 12px;
}

.avatar-info {
  color: #666;
  font-size: 12px;
  line-height: 1.4;
}

.field-hint {
  margin-top: 4px;
  color: #666;
  font-size: 12px;
}
</style>
