<template>
  <a-modal
    v-model:open="isVisible"
    title="转让碳信用"
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
          确认转让
        </a-button>
      </a-space>
    </template>

    <div class="credit-transfer-modal">
      <a-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        layout="vertical"
        @finish="onFinish"
      >
        <a-form-item
          name="selectedCredit"
          label="选择要转让的碳信用"
          :rules="rules.selectedCredit"
        >
          <a-select
            v-model:value="formData.selectedCredit"
            placeholder="请选择要转让的碳信用批次"
            @change="handleCreditSelect"
          >
            <a-select-option
              v-for="credit in availableCredits"
              :key="getBatchId(credit)"
              :value="getBatchId(credit)"
            >
              <div class="credit-option">
                <div class="credit-basic">
                  <span class="batch-id">批次 #{{ getBatchId(credit) }}</span>
                  <span class="quantity">{{ credit.quantity.toLocaleString() }} 吨CO₂e</span>
                </div>
                <div class="credit-project">
                  项目 #{{ credit.projectId }} • {{ credit.vintageYear }}年
                </div>
              </div>
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item
          v-if="selectedCreditInfo"
          label="所选碳信用信息"
        >
          <a-card size="small">
            <a-descriptions :column="2" size="small">
              <a-descriptions-item label="批次ID">
                #{{ selectedCreditInfo.batchId }}
              </a-descriptions-item>
              <a-descriptions-item label="可用数量">
                {{ selectedCreditInfo.quantity.toLocaleString() }} 吨CO₂e
              </a-descriptions-item>
              <a-descriptions-item label="关联项目">
                项目 #{{ selectedCreditInfo.projectId }}
              </a-descriptions-item>
              <a-descriptions-item label="年份">
                {{ selectedCreditInfo.vintageYear }}年
              </a-descriptions-item>
            </a-descriptions>
          </a-card>
        </a-form-item>

        <a-form-item
          name="toAddress"
          label="转让目标地址"
          :rules="rules.toAddress"
        >
          <a-input
            v-model:value="formData.toAddress"
            placeholder="请输入接收方的钱包地址"
          />
        </a-form-item>

        <a-form-item
          name="quantity"
          label="转让数量"
          :rules="rules.quantity"
        >
          <a-input-number
            v-model:value="formData.quantity"
            placeholder="请输入转让数量"
            :min="1"
            :max="selectedCreditInfo?.quantity || 0"
            :precision="0"
            style="width: 100%"
            addon-after="吨CO₂e"
          />
          <div v-if="selectedCreditInfo" class="quantity-hint">
            可转让数量：{{ selectedCreditInfo.quantity.toLocaleString() }} 吨CO₂e
          </div>
        </a-form-item>

        <a-form-item
          name="newBatchId"
          label="新批次ID"
          :rules="rules.newBatchId"
          extra="拆分转让时需要为新批次指定ID"
        >
          <a-input-number
            v-model:value="formData.newBatchId"
            placeholder="请输入新批次ID"
            :min="1"
            style="width: 100%"
          />
        </a-form-item>

        <a-form-item
          name="reason"
          label="转让原因"
          :rules="rules.reason"
        >
          <a-textarea
            v-model:value="formData.reason"
            placeholder="请说明转让原因"
            :rows="3"
            :maxlength="200"
            show-count
          />
        </a-form-item>

        <!-- 转让说明 -->
        <a-form-item label="转让说明">
          <a-alert
            message="转让须知"
            type="info"
            show-icon
          >
            <template #description>
              <ul class="transfer-notes">
                <li>转让后碳信用的所有权将转移给接收方</li>
                <li>如果转让数量小于持有数量，将进行批次拆分</li>
                <li>转让操作一旦确认无法撤销</li>
                <li>请确认接收方地址正确无误</li>
              </ul>
            </template>
          </a-alert>
        </a-form-item>
      </a-form>
    </div>
  </a-modal>
  <EditReasonModal
    v-model:visible="showReasonModal"
    title="确认碳信用转让"
    description="将发起链上转让交易，请填写转让原因。"
    @confirm="handleReasonConfirm"
  />
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import { transferCredits } from '@/services/web3'
import type { CreditRecord } from '@/types'
import EditReasonModal from '@/components/EditReasonModal.vue'

interface Props {
  visible: boolean
  availableCredits?: CreditRecord[]
}

interface Emits {
  (e: 'update:visible', value: boolean): void
  (e: 'success'): void
}

const props = withDefaults(defineProps<Props>(), {
  availableCredits: () => []
})

const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const showReasonModal = ref(false)
let pendingPayload: { batchId: number; quantity: number; to: string } | null = null

// 表单数据
const formData = ref({
  selectedCredit: undefined as number | undefined,
  toAddress: '',
  quantity: undefined as number | undefined,
  newBatchId: undefined as number | undefined,
  reason: ''
})

// 所选碳信用信息
const selectedCreditInfo = computed(() => {
  if (!formData.value.selectedCredit) return null
  
  const credit = props.availableCredits.find(c => getBatchId(c) === formData.value.selectedCredit)
  if (!credit) return null
  
  return {
    batchId: getBatchId(credit),
    quantity: credit.quantity,
    projectId: credit.projectId,
    vintageYear: credit.vintageYear
  }
})

// 表单验证规则
const rules = {
  selectedCredit: [
    { required: true, message: '请选择要转让的碳信用', trigger: 'change' }
  ],
  toAddress: [
    { required: true, message: '请输入接收方地址', trigger: 'blur' },
    { 
      pattern: /^0x[a-fA-F0-9]{40}$/,
      message: '请输入有效的以太坊地址',
      trigger: 'blur'
    }
  ],
  quantity: [
    { required: true, message: '请输入转让数量', trigger: 'blur' },
    { 
      validator: (rule: any, value: number) => {
        if (!value || value <= 0) {
          return Promise.reject('转让数量必须大于0')
        }
        if (selectedCreditInfo.value && value > selectedCreditInfo.value.quantity) {
          return Promise.reject('转让数量不能超过可用数量')
        }
        return Promise.resolve()
      },
      trigger: 'blur'
    }
  ],
  newBatchId: [
    { required: true, message: '请输入新批次ID', trigger: 'blur' }
  ],
  reason: [
    { required: true, message: '请说明转让原因', trigger: 'blur' },
    { min: 5, max: 200, message: '转让原因应在5-200字符之间', trigger: 'blur' }
  ]
}

// 控制弹窗显示
const isVisible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value)
})

// 表单验证状态
const isFormValid = computed(() => {
  return formData.value.selectedCredit !== undefined &&
         formData.value.toAddress.trim() !== '' &&
         formData.value.quantity !== undefined &&
         formData.value.quantity > 0 &&
         formData.value.newBatchId !== undefined &&
         formData.value.reason.trim().length >= 5
})

// 监听弹窗显示状态，重置表单
watch(
  () => props.visible,
  (visible) => {
    if (visible) {
      formData.value = {
        selectedCredit: undefined,
        toAddress: '',
        quantity: undefined,
        newBatchId: undefined,
        reason: ''
      }
    }
  }
)

// 处理碳信用选择
const handleCreditSelect = () => {
  // 重置数量和新批次ID
  formData.value.quantity = undefined
  formData.value.newBatchId = undefined
}

// 处理表单提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    pendingPayload = {
      batchId: formData.value.selectedCredit!,
      quantity: formData.value.quantity!,
      to: formData.value.toAddress.trim()
    }
    showReasonModal.value = true
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}

// 理由确认
const handleReasonConfirm = async (reason: string) => {
  // 表单提交处理
  if (!selectedCreditInfo.value) {
    message.error('请选择要转让的碳信用')
    return
  }

  try {
    loading.value = true
    
    await transferCredits(
      pendingPayload!.batchId,
      pendingPayload!.quantity,
      pendingPayload!.to
    )

    message.success('碳信用转让成功')
    emit('success')
  } catch (error: any) {
    console.error('转让碳信用失败:', error)
    message.error(error.message || '转让失败，请重试')
  } finally {
    loading.value = false
    showReasonModal.value = false
    pendingPayload = null
  }
}

// 取消处理
const handleCancel = () => {
  isVisible.value = false
}

// 获取批次ID
const getBatchId = (credit: CreditRecord) => {
  return credit.childBatchId || credit.parentBatchId || 1
}
</script>

<style scoped>
.credit-transfer-modal {
  padding: 8px 0;
}

.credit-option {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.credit-basic {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.batch-id {
  font-weight: 500;
  color: #262626;
}

.quantity {
  color: #52c41a;
  font-weight: 500;
}

.credit-project {
  font-size: 12px;
  color: #8c8c8c;
}

.quantity-hint {
  margin-top: 4px;
  font-size: 12px;
  color: #666;
}

.transfer-notes {
  margin: 0;
  padding-left: 16px;
}

.transfer-notes li {
  margin-bottom: 4px;
  line-height: 1.5;
}

:deep(.ant-select-item-option-content) {
  width: 100%;
}

:deep(.ant-form-item-label) {
  font-weight: 500;
}
</style>
