<template>
    <Aside class="aside-open-chat-container">
        <NuxtLink to="/profile" class="profile-info">
            <img :src="require(`@/assets/img/${avatarurl}`)" />
            <b>{{ name }}</b>
        </NuxtLink>
    </Aside>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        avatarurl: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
    }
})
</script>

<style lang="scss" scoped>
.aside-open-chat-container {
    display: grid;
    width: 27rem;
    background: color(dark, shade1);

    .profile-info {
        padding-top: 4rem;
        display: grid;
        grid-auto-flow: row;
        gap: 2.5rem;
        justify-content: center;
        text-align: center;

        img {
            width: 15rem;
            border-radius: 50%;
        }

        b {
            color: color(white);
            font-size: 1.5rem;
        }
    }
}
</style>