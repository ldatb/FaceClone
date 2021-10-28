<template>
    <NuxtLink :to="`/profile/${username}`" class="profile-bubble-container">
        <img :src=avatarurl class="profile-image" />
        <span v-if="name.length <= 16" :class="{ 'big-name': name.length>=10 }">{{ name }}</span>
        <span v-else >{{ name.substring(0, 10)+"..." }}</span>
    </NuxtLink>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        username: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
        avatarurl: {
            type: String,
            required: true,
        },
    },
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
.big-name {
    font-size: 11.5px !important;
}
</style>