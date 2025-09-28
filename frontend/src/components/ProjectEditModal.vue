<template>
  <a-modal
    :open="visible"
    @update:open="(v:boolean)=>emit('update:visible', v)"
    :title="isEdit ? '编辑项目' : '创建项目'"
    width="800px"
    @ok="handleSubmit"
    @cancel="handleCancel"
    :confirm-loading="loading"
    :mask-closable="false"
  >
    <a-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      layout="vertical"
    >
      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item label="项目名称" name="name">
            <a-input
              v-model:value="formData.name"
              placeholder="请输入项目名称"
              :disabled="isEdit && !canEdit"
            />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="项目地址" name="location">
            <a-input
              v-model:value="formData.location"
              placeholder="请输入项目地址"
              :disabled="isEdit && !canEdit"
            />
          </a-form-item>
        </a-col>
      </a-row>

      <a-form-item label="项目描述" name="description">
        <a-textarea
          v-model:value="formData.description"
          placeholder="请输入项目描述"
          :rows="4"
          :disabled="isEdit && !canEdit"
        />
      </a-form-item>

      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item label="技术类型" name="technologies">
            <a-select
              v-model:value="formData.technologies"
              mode="multiple"
              placeholder="请选择技术类型"
              :disabled="isEdit && !canEdit"
            >
              <a-select-option :value="0">太阳能</a-select-option>
              <a-select-option :value="1">风能</a-select-option>
              <a-select-option :value="2">水电</a-select-option>
              <a-select-option :value="3">林业</a-select-option>
              <a-select-option :value="4">碳捕获</a-select-option>
              <a-select-option :value="5">其他</a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="年减排量 (吨CO₂e)" name="annualReduction">
            <a-input-number
              v-model:value="formData.annualReduction"
              placeholder="请输入年减排量"
              style="width: 100%"
              :min="0"
              :precision="2"
              :disabled="isEdit && !canEdit"
            />
          </a-form-item>
        </a-col>
      </a-row>


      <a-form-item label="项目文档" name="documents">
        <a-upload
          v-model:file-list="formData.documents"
          :before-upload="beforeUpload"
          multiple
          :disabled="isEdit && !canEdit"
        >
          <a-button>
            <upload-outlined />
            上传文档
          </a-button>
        </a-upload>
      </a-form-item>
    </a-form>
    <EditReasonModal
      v-model:visible="showReasonModal"
      :title="isEdit ? '确认更新项目信息' : '确认创建项目'"
      :description="isEdit ? '将对当前项目进行信息更新，请填写更新原因。' : '将在链上创建新项目，请填写创建原因。'"
      @confirm="handleReasonConfirm"
    />
  </a-modal>
</template>

<script setup lang="ts">
import { ref, computed, watch, nextTick } from 'vue'
import { message } from 'ant-design-vue'
import { UploadOutlined } from '@ant-design/icons-vue'
import dayjs, { type Dayjs } from 'dayjs'
import type { FormInstance, Rule } from 'ant-design-vue/es/form'
import type { UploadFile } from 'ant-design-vue/es/upload/interface'
import type { ProjectRecord } from '@/types'
import { registerProject, updateProjectDetails } from '@/services/web3'
import { PermissionService } from '@/services/permission'
import { useUserStore } from '@/stores/user'
import { storeToRefs } from 'pinia'
import EditReasonModal from '@/components/EditReasonModal.vue'

interface Props {
  visible: boolean
  project?: ProjectRecord | null
}

interface Emits {
  (e: 'update:visible', value: boolean): void
  (e: 'success'): void
}

const props = withDefaults(defineProps<Props>(), {
  project: null
})

const emit = defineEmits<Emits>()

// 表单引用
const formRef = ref<FormInstance>()
const loading = ref(false)
const showReasonModal = ref(false)
let pendingProjectData: {
  name: string
  description: string
  location: string
  technologies: number[]
  reduction: number
  documentsHash: string[]
} | null = null

// 计算属性
const isEdit = computed(() => !!props.project)
const canEdit = computed(() => {
  if (!props.project) return true
  // 只有编辑中的项目可以编辑
  return props.project.status === 0
})

// 表单数据
const formData = ref({
  name: '',
  description: '',
  location: '',
  technologies: [] as number[],
  annualReduction: null as number | null,
  startDate: null as Dayjs | null,
  endDate: null as Dayjs | null,
  documents: [] as UploadFile[]
})

// 表单验证规则
const formRules: Record<string, Rule[]> = {
  name: [
    { required: true, message: '请输入项目名称', trigger: 'blur' },
    { min: 2, max: 100, message: '项目名称长度应在2-100字符之间', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入项目描述', trigger: 'blur' },
    { min: 10, max: 1000, message: '项目描述长度应在10-1000字符之间', trigger: 'blur' }
  ],
  location: [
    { required: true, message: '请输入项目地址', trigger: 'blur' }
  ],
  technologies: [
    { required: true, message: '请选择技术类型', trigger: 'change' }
  ],
  annualReduction: [
    { required: true, message: '请输入年减排量', trigger: 'blur' },
    { type: 'number', min: 0.01, message: '年减排量必须大于0', trigger: 'blur' }
  ]
}

// 监听弹窗显示状态
watch(() => props.visible, (newValue) => {
  if (newValue) {
    initForm()
  } else {
    resetForm()
  }
})

// 监听项目数据变化
watch(() => props.project, () => {
  if (props.visible) {
    initForm()
  }
})

// 初始化表单
const initForm = () => {
  if (props.project) {
    formData.value = {
      name: props.project.name,
      description: props.project.description,
      location: props.project.location,
      technologies: [...props.project.technologies],
      annualReduction: props.project.reduction,
      startDate: null,
      endDate: null,
      documents: []
    }
  } else {
    resetForm()
  }
  
  // 清除验证状态
  nextTick(() => {
    formRef.value?.clearValidate()
  })
}

// 重置表单
const resetForm = () => {
  formData.value = {
    name: '',
    description: '',
    location: '',
    technologies: [],
    annualReduction: null,
    startDate: null,
    endDate: null,
    documents: []
  }
  formRef.value?.resetFields()
}

// 文件上传前处理
const beforeUpload = (file: UploadFile) => {
  const isValidType = file.type === 'application/pdf' || 
    file.type === 'application/msword' || 
    file.type === 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
  
  if (!isValidType) {
    message.error('只能上传 PDF 或 Word 文档!')
    return false
  }
  
  const isLt10M = file.size! / 1024 / 1024 < 10
  if (!isLt10M) {
    message.error('文件大小不能超过 10MB!')
    return false
  }
  
  return false // 阻止自动上传
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    
    // 验证必要字段
    if (!formData.value.annualReduction || formData.value.annualReduction <= 0) {
      message.error('请输入有效的年减排量')
      return
    }
    
    if (!formData.value.technologies || formData.value.technologies.length === 0) {
      message.error('请选择至少一种技术类型')
      return
    }
    
    const projectData = {
      name: formData.value.name.trim(),
      description: formData.value.description.trim(),
      location: formData.value.location.trim(),
      technologies: formData.value.technologies,
      reduction: Math.floor(Number(formData.value.annualReduction) * 100) / 100, // 确保是有效的数字，保留2位小数
      documentsHash: formData.value.documents.map(file => file.name || '').filter(name => name)
    }
    
    console.log('准备提交项目数据:', projectData)
    pendingProjectData = projectData
    showReasonModal.value = true
  } catch (error: any) {
    console.error('表单验证失败:', error)
    if (error.errorFields) {
      message.error('请检查表单填写是否正确')
    } else {
      message.error('提交项目失败: ' + (error.message || '未知错误'))
    }
  }
}

// 理由确认回调
const handleReasonConfirm = async (reason: string) => {
  if (!pendingProjectData) return
  
  // 检查用户权限
  if (!isEdit.value) {
    const userStore = useUserStore()
    const { currentUser } = storeToRefs(userStore)
    if (!PermissionService.isProjectOwner(currentUser.value)) {
      message.error('您没有项目方权限，无法创建项目')
      return
    }
  }
  
  try {
    loading.value = true
    console.log('开始提交项目，数据:', pendingProjectData)
    
    if (isEdit.value) {
      await updateProjectDetails(
        props.project!.id,
        pendingProjectData.name,
        pendingProjectData.location,
        pendingProjectData.technologies,
        pendingProjectData.description,
        pendingProjectData.reduction,
        pendingProjectData.documentsHash,
        reason
      )
      message.success('项目更新成功')
    } else {
      await registerProject(
        pendingProjectData.name,
        pendingProjectData.location,
        pendingProjectData.technologies,
        pendingProjectData.description,
        pendingProjectData.reduction,
        pendingProjectData.documentsHash,
        reason
      )
      message.success('项目创建成功')
    }
    emit('success')
    emit('update:visible', false)
  } catch (error: any) {
    console.error('提交项目失败:', error)
    
    // 显示具体的错误信息
    let errorMessage = '提交项目失败'
    if (error.message) {
      if (error.message.includes('用户取消了交易')) {
        errorMessage = '用户取消了交易'
      } else if (error.message.includes('交易执行失败')) {
        errorMessage = '交易执行失败，请检查参数或余额'
      } else if (error.message.includes('请先连接钱包')) {
        errorMessage = '请先连接钱包'
      } else if (error.message.includes('权限不足')) {
        errorMessage = '权限不足，请确保您有项目方权限'
      } else {
        errorMessage = error.message
      }
    }
    
    message.error(errorMessage)
  } finally {
    loading.value = false
    pendingProjectData = null
    showReasonModal.value = false
  }
}

// 取消操作
const handleCancel = () => {
  emit('update:visible', false)
}
</script>

<style scoped>
.ant-form-item {
  margin-bottom: 16px;
}

.ant-upload {
  width: 100%;
}
</style>
