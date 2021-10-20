<template>
    <div>
        <BaseButton class="a" text="Teste Botão" btnlink="true" link="seila"/>
        <BaseInput placeholder="Base Input"/>
        <ExpandableTextArea placeholder="Expandable Text Area"/>
        <StatusBadge class="status-online"/>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
})
</script>

<style lang='scss' scoped>
.a {
    min-width: 13rem;
}
</style>