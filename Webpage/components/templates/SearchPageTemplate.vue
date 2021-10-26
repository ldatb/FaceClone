<template>
    <div class="search-page-template">
        <div class="search-container">
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow"/>
            <ProfileFollow avatarurl="profile-pic.jpg" name="Lucas de Ataides" username="teste123" actiontext="Follow"/>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.search-page-template {
    padding-bottom: 4rem;
}
.search-container {
    display: grid;
    justify-content: center;
    align-content: center;
    grid-template-columns: auto auto;
    column-gap: 4rem;
}
</style>