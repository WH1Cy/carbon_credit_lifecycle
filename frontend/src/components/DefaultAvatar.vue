<template>
  <a-avatar 
    :size="size" 
    :style="{ 
      backgroundColor: backgroundColor, 
      color: '#fff',
      fontWeight: '600'
    }"
  >
    {{ initials }}
  </a-avatar>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  name?: string
  size?: number | 'small' | 'default' | 'large'
}

const props = withDefaults(defineProps<Props>(), {
  name: '用户',
  size: 'default'
})

// 生成用户名首字母
const initials = computed(() => {
  if (!props.name) return 'U'
  
  // 处理中文和英文
  const trimmedName = props.name.trim()
  if (!trimmedName) return 'U'
  
  // 如果是中文，取最后一个字符（通常是名字）
  if (/[\u4e00-\u9fa5]/.test(trimmedName)) {
    return trimmedName.slice(-1)
  }
  
  // 如果是英文，取首字母
  const words = trimmedName.split(' ')
  if (words.length >= 2) {
    return (words[0][0] + words[1][0]).toUpperCase()
  }
  return words[0][0].toUpperCase()
})

// 根据用户名生成背景色
const backgroundColor = computed(() => {
  if (!props.name) return '#1890ff'
  
  const colors = [
    '#f56a00', '#7265e6', '#ffbf00', '#00a2ae',
    '#1890ff', '#722ed1', '#eb2f96', '#52c41a',
    '#fa541c', '#faad14', '#13c2c2', '#2f54eb'
  ]
  
  let hash = 0
  for (let i = 0; i < props.name.length; i++) {
    hash = props.name.charCodeAt(i) + ((hash << 5) - hash)
  }
  
  return colors[Math.abs(hash) % colors.length]
})
</script>
