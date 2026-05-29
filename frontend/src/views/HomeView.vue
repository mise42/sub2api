<template>
  <div class="min-h-screen">
    <iframe
      v-if="iframeSrc"
      :src="iframeSrc"
      class="h-screen w-full border-0"
      allowfullscreen
    />
    <div v-else v-html="homeContent"></div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAppStore } from '@/stores'

const appStore = useAppStore()

const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')
const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

const iframeSrc = computed(() => {
  if (homeContent.value) {
    return isHomeContentUrl.value ? homeContent.value.trim() : ''
  }
  return '/sub2api-promo.html'
})

onMounted(() => {
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>
