<template>
    <NuxtLink to="/profile" class="profile-bubble-container">
        <img src="@/assets/img/profile-pic.jpg" class="profile-image">
        <span>User name</span>
    </NuxtLink>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
        
})
</script>

<style lang="scss" scoped>
.profile-bubble-container {
    display: grid;
    grid-template-columns: auto auto;
    align-content: center;
    width: 9.4rem;
    height: fit-content;
    padding: 0.2rem 0.3rem;
    background: #152330;
    border-radius: 2rem;
    color: white;

    span {
        align-self: center;
        justify-self: start !important;
        font-size: 14.5px;
        font-weight: 700;
    }

    .profile-image {
        width: 2.1rem;
        border-radius: 50%;
    }
}
</style>