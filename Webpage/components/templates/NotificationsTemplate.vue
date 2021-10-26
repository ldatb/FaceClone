<template>
    <div class="notifications-template-container">
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
        <Notification text="Notification test text :)" time="13:29" />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.notifications-template-container {
    display: grid;
    grid-auto-flow: row;
    align-content: center;
    justify-content: center;
    gap: 1rem;
    padding-top: 3rem;
}
</style>