<template>
    <div class="profile-following-template">
        <Profile coverurl="capa.jpg" avatarurl="profile-pic.jpg" username="Teste user" />

        <ProfileTabs timeline=300 followers=400 following=500 activefollowing />

        <div class="following-container">
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Unfollow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Unfollow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Unfollow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Unfollow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Unfollow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Unfollow"/>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.profile-following-template {
    padding-bottom: 4rem;
}
.following-container {
    display: grid;
    justify-content: center;
    align-content: center;
    grid-template-columns: auto auto;
    column-gap: 4rem;
}
</style>