<template>
    <div class="profile-followers-template">
        <Profile coverurl="capa.jpg" avatarurl="profile-pic.jpg" username="Teste user" />

        <ProfileTabs timeline=300 followers=400 following=500 activefollowers />

        <div class="followers-container">
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow back"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow back"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow back"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow back"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow back"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow back"/>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.profile-followers-template {
    padding-bottom: 4rem;
}
.followers-container {
    display: grid;
    justify-content: center;
    align-content: center;
    grid-template-columns: auto auto;
    column-gap: 4rem;
}
</style>