<template>
  <div class="full-screen-container">
    <slot />
  </div>
</template>

<style lang="scss" scoped>
.full-screen-container {
  max-width: 1460px;
  height: 100%;
  margin: 0 auto;
  display: grid;
}
</style>