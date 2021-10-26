<template>
    <div class="messenger-template">
        <MessengerAsideLeft class="aside-chats-left"/>
        <MessengerActiveChat class="open-chat-container" />
        <MessengerAsideRight avatarurl="profile-pic.jpg" name="Joana D'Arc" class="aside-open-chat-profile" />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.messenger-template {
    display: grid;
    grid-auto-flow: column;
    align-content: center;
    grid-template-areas: auto 1fr auto;
}
.aside-chats-left {
    justify-self: start;
    width: 100%;
    height: 100%;
    max-height: 100vw;
}
.open-chat-container {
    justify-self: center;
}
.aside-open-chat-profile {
    justify-self: end;
}
</style>