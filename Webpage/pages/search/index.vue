<template>
    <div>
        <b>
            {{ $route.query.query }}
        </b>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    layout: 'blank'
})
</script>

<style lang="scss" scoped>

</style>