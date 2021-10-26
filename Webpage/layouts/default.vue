<template>
  <div class="default-layout">
    <clientOnly>
      <notifications position="bottom center" classes="notifications" :max="1" />
    </clientOnly>
    <Header class="sticky-header"/>
    <Nuxt class="page-content"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mobile } from '@/store'
export default Vue.extend({
  computed: {
    $isMenuActive() {
      return mobile.$isMenuActive
    }
  }
})
</script>

<style lang="scss">
body {
  background: color(dark);
}
.default-layout {
  min-height: 100%;
}
.sticky-header {
  position: sticky;
  top: 0;
  width: 100%;
  z-index: 0;
}
.page-content {
  z-index: -1;
}
</style>