<template>
    <div class="header-container">
        <NuxtLink to="/" class="left-header">
            <img src="@/assets/img/home-icon.svg" />
        </NuxtLink>

        <SearchBar class="center-header" />

        <div class="right-header">
            <ProfileBubbleHeader />

            <div class="header-links">
                <NuxtLink to="/messenger">
                    <img src="@/assets/img/messenger-icon.svg" class="link-image">
                </NuxtLink>
                <NuxtLink to="/notifications">
                    <img src="@/assets/img/notification-icon.svg" class="link-image">
                </NuxtLink>
                <NuxtLink to="/profile/change">
                    <img src="@/assets/img/settings-icon.svg" class="link-image">
                </NuxtLink>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
        
})
</script>

<style lang="scss" scoped>
.header-container {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    height: 4rem;
    background: color(dark, shade2);
    box-shadow: 0 1rem 1rem rgba(0, 0, 0, 0.25);
}
.left-header {
    margin-top: 4px;
    margin-right: 9rem;
    align-self: center;
    justify-self: center;
}
.center-header {
    align-self: center;
    justify-self: center;
}
.right-header {
    margin-right: 7rem;
    display: grid;
    grid-template-columns: auto auto;
    align-content: center;
    justify-content: center;
    gap: 3rem;
}
.header-links {
    display: grid;
    height: fit-content;
    align-self: center;
    margin-top: 4.5px;
    grid-template-columns: auto auto auto;
    gap: 1.2rem;

    .link-image {
        width: 1.4rem;
    }
}
</style>