<template>
    <div class="profile-template">
        <Profile coverurl="capa.jpg" avatarurl="profile-pic.jpg" username="Teste user" />

        <ProfileTabs timeline=300 followers=400 following=500 activetimeline=true />

        <ProfileFeed class="profile-feed" />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.profile-template {
    display: grid;
    padding-bottom: 4rem;
}
.profile-feed {
    padding-top: 4rem;
}
</style>