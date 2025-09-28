<template>
  <a-modal
    v-model:open="isVisible"
    title="提交审核意见"
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
          提交审核意见
        </a-button>
      </a-space>
    </template>

    <div v-if="task" class="review-opinion-modal">
      <!-- 项目信息 -->
      <div class="project-summary">
        <h4>审核项目信息</h4>
        <a-descriptions :column="2" size="small" bordered>
          <a-descriptions-item label="项目名称">
            {{ task.project.name }}
          </a-descriptions-item>
          <a-descriptions-item label="项目ID">
            #{{ task.project.id }}
          </a-descriptions-item>
          <a-descriptions-item label="项目位置">
            {{ task.project.location }}
          </a-descriptions-item>
          <a-descriptions-item label="预期减排量">
            {{ task.project.reduction.toLocaleString() }} 吨CO₂e
          </a-descriptions-item>
          <a-descriptions-item label="技术类型" :span="2">
            <a-space wrap>
              <a-tag
                v-for="tech in task.project.technologies"
                :key="tech"
                :color="getTechnologyColor(tech)"
                size="small"
              >
                {{ getTechnologyText(tech) }}
              </a-tag>
            </a-space>
          </a-descriptions-item>
        </a-descriptions>
      </div>

      <a-divider />

      <a-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        layout="vertical"
        @finish="onFinish"
      >
        <a-form-item
          name="approved"
          label="审核结果"
          :rules="rules.approved"
        >
          <a-radio-group v-model:value="formData.approved">
            <a-radio :value="true">
              <a-tag color="success">通过</a-tag>
              <span style="margin-left: 8px;">项目符合要求，建议批准</span>
            </a-radio>
            <a-radio :value="false">
              <a-tag color="error">不通过</a-tag>
              <span style="margin-left: 8px;">项目存在问题，建议拒绝</span>
            </a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item
          v-if="formData.approved"
          name="issuedBatchId"
          label="核发批次ID"
          extra="审核通过时需要指定核发的碳信用批次ID"
        >
          <a-input-number
            v-model:value="formData.issuedBatchId"
            placeholder="请输入批次ID"
            :min="1"
            style="width: 200px"
          />
        </a-form-item>

        <a-form-item
          name="comment"
          label="审核意见"
          :rules="rules.comment"
        >
          <a-textarea
            v-model:value="formData.comment"
            placeholder="请详细说明审核意见和理由"
            :rows="4"
            :maxlength="500"
            show-count
          />
        </a-form-item>

        <!-- 审核要点提醒 -->
        <a-form-item label="审核要点提醒">
          <a-alert
            message="请注意以下审核要点"
            type="info"
            show-icon
          >
            <template #description>
              <ul class="review-checklist">
                <li>项目技术方案的可行性和科学性</li>
                <li>减排量计算的准确性和合理性</li>
                <li>项目文档的完整性和真实性</li>
                <li>项目实施的可操作性和可监测性</li>
                <li>是否符合相关标准和法规要求</li>
              </ul>
            </template>
          </a-alert>
        </a-form-item>
      </a-form>
    </div>
  </a-modal>
  <EditReasonModal
    v-model:visible="showReasonModal"
    title="确认提交审核意见"
    description="将更新项目审核记录并可能核发碳信用，请填写理由。"
    @confirm="handleReasonConfirm"
  />
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import { finalizeVerification } from '@/services/web3'
import { Technology } from '@/types'
import EditReasonModal from '@/components/EditReasonModal.vue'

// 审核任务接口
interface ReviewTask {
  id: string
  project: {
    id: number
    name: string
    location: string
    reduction: number
    technologies: Technology[]
  }
  status: string
  acceptTime: number
}

interface Props {
  visible: boolean
  task?: ReviewTask | null
}

interface Emits {
  (e: 'update:visible', value: boolean): void
  (e: 'success'): void
}

const props = withDefaults(defineProps<Props>(), {
  task: null
})

const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const showReasonModal = ref(false)

// 表单数据
const formData = ref({
  approved: undefined as boolean | undefined,
  comment: '',
  issuedBatchId: undefined as number | undefined
})

// 表单验证规则
const rules = {
  approved: [
    { required: true, message: '请选择审核结果', trigger: 'change' }
  ],
  comment: [
    { required: true, message: '请填写审核意见', trigger: 'blur' },
    { min: 10, max: 500, message: '审核意见应在10-500字符之间', trigger: 'blur' }
  ],
  issuedBatchId: [
    { 
      validator: (rule: any, value: number) => {
        if (formData.value.approved && !value) {
          return Promise.reject('审核通过时必须指定批次ID')
        }
        return Promise.resolve()
      },
      trigger: 'blur'
    }
  ]
}

// 控制弹窗显示
const isVisible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value)
})

// 表单验证状态
const isFormValid = computed(() => {
  if (formData.value.approved === undefined) return false
  if (!formData.value.comment.trim()) return false
  if (formData.value.approved && !formData.value.issuedBatchId) return false
  return true
})

// 监听弹窗显示状态，重置表单
watch(
  () => props.visible,
  (visible) => {
    if (visible) {
      formData.value = {
        approved: undefined,
        comment: '',
        issuedBatchId: undefined
      }
    }
  }
)

// 处理表单提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    showReasonModal.value = true
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}

// 理由确认
const handleReasonConfirm = async (reason: string) => {
  if (!props.task) {
    message.error('任务信息缺失')
    return
  }

  try {
    loading.value = true
    
    await finalizeVerification(
      props.task.project.id,
      formData.value.approved!,
      `${formData.value.comment.trim()}\n理由: ${reason}`,
      formData.value.issuedBatchId || 0,
      new Date().getFullYear() // 使用当前年份作为vintageYear
    )

    message.success('审核意见提交成功')
    emit('success')
  } catch (error: any) {
    console.error('提交审核意见失败:', error)
    message.error(error.message || '提交失败，请重试')
  } finally {
    loading.value = false
    showReasonModal.value = false
  }
}

// 取消处理
const handleCancel = () => {
  isVisible.value = false
}

// 工具函数
const getTechnologyColor = (tech: Technology) => {
  const colors = {
    [Technology.SOLAR]: 'orange',
    [Technology.WIND]: 'cyan',
    [Technology.HYDRO]: 'blue',
    [Technology.FORESTRY]: 'green',
    [Technology.CAPTURE]: 'purple',
    [Technology.OTHER]: 'default'
  }
  return colors[tech] || 'default'
}

const getTechnologyText = (tech: Technology) => {
  const texts = {
    [Technology.SOLAR]: '太阳能',
    [Technology.WIND]: '风能',
    [Technology.HYDRO]: '水电',
    [Technology.FORESTRY]: '林业',
    [Technology.CAPTURE]: '碳捕获',
    [Technology.OTHER]: '其他'
  }
  return texts[tech] || '未知'
}
</script>

<style scoped>
.review-opinion-modal {
  padding: 8px 0;
}

.project-summary {
  margin-bottom: 16px;
}

.project-summary h4 {
  margin: 0 0 12px 0;
  font-size: 16px;
  font-weight: 600;
}

.review-checklist {
  margin: 0;
  padding-left: 16px;
}

.review-checklist li {
  margin-bottom: 4px;
  line-height: 1.5;
}

:deep(.ant-radio-wrapper) {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
  padding: 8px;
  border: 1px solid #f0f0f0;
  border-radius: 6px;
  transition: all 0.3s;
}

:deep(.ant-radio-wrapper:hover) {
  border-color: #40a9ff;
  background-color: #f6ffed;
}

:deep(.ant-radio-wrapper-checked) {
  border-color: #1890ff;
  background-color: #e6f7ff;
}

:deep(.ant-form-item-label) {
  font-weight: 500;
}
</style>
