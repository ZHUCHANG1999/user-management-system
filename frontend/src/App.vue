<template>
  <div id="app">
    <component :is="layoutComponent">
      <router-view />
    </component>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import Layout from '@/components/Layout.vue'

const route = useRoute()

// 登录/注册页面不使用 Layout
const layoutComponent = computed(() => {
  if (['/login', '/register'].includes(route.path)) {
    return 'div'
  }
  return Layout
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  min-height: 100vh;
}
</style>
