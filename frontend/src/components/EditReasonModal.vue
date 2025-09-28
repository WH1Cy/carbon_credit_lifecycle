<template>
  <a-modal
    v-model:open="isVisible"
    :title="title"
    width="500px"
    :mask-closable="false"
    @ok="handleConfirm"
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
          @click="handleConfirm" 
          :loading="loading"
          :disabled="!isFormValid"
          danger
        >
          确认
        </a-button>
      </a-space>
    </template>

    <div class="edit-reason-modal">
      <!-- 描述信息 -->
      <div v-if="description" class="description">
        <a-alert
          :message="description"
          type="warning"
          show-icon
          style="margin-bottom: 16px"
        />
      </div>

      <a-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        layout="vertical"
        @finish="onFinish"
      >
        <a-form-item
          name="reason"
          label="操作原因"
          :rules="rules.reason"
        >
          <a-textarea
            v-model:value="formData.reason"
            :placeholder="placeholder"
            :rows="4"
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

interface Props {
  visible: boolean
  title?: string
  description?: string
  placeholder?: string
  loading?: boolean
}

interface Emits {
  (e: 'update:visible', value: boolean): void
  (e: 'confirm', reason: string): void
}

const props = withDefaults(defineProps<Props>(), {
  title: '确认操作',
  description: '',
  placeholder: '请说明本次操作的原因',
  loading: false
})

const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()

// 表单数据
const formData = ref({
  reason: ''
})

// 表单验证规则
const rules = {
  reason: [
    { required: true, message: '请说明操作原因', trigger: 'blur' },
    { min: 5, max: 200, message: '操作原因应在5-200字符之间', trigger: 'blur' }
  ]
}

// 控制弹窗显示
const isVisible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value)
})

// 表单验证状态
const isFormValid = computed(() => {
  return formData.value.reason.trim().length >= 5
})

// 监听弹窗显示状态，重置表单
watch(
  () => props.visible,
  (visible) => {
    if (visible) {
      formData.value = {
        reason: ''
      }
    }
  }
)

// 处理确认
const handleConfirm = async () => {
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
  const reason = formData.value.reason.trim()
  if (!reason) {
    message.error('请说明操作原因')
    return
  }

  emit('confirm', reason)
}

// 取消处理
const handleCancel = () => {
  isVisible.value = false
}
</script>

<style scoped>
.edit-reason-modal {
  padding: 8px 0;
}

.description {
  margin-bottom: 16px;
}

:deep(.ant-form-item-label) {
  font-weight: 500;
}
</style>
