<template>
    <div class="profile-info">
        <div class="profile-cover">
            <img :src="require(`@/assets/img/${coverurl}`)" class="cover-image" />
        </div>
        <div class="profile-avatar">
            <img :src="require(`@/assets/img/${avatarurl}`)" class="avatar-image" />
            <span class="profile-name">{{ username }}</span>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        coverurl: {
            type: String,
            required: true,
        },
        avatarurl: {
            type: String,
            required: true,
        },
        username: {
            type: String,
            required: true,
        },
    }
})
</script>

<style lang="scss" scoped>
.profile-info {
    display: grid;
    grid-auto-flow: row;

    .profile-cover {
        height: 25rem;
        width: 100%;
        z-index: -1 !important;

        .cover-image {
            height: 100%;
            width: 100%;
            object-fit: cover !important;
        }
    }

    .profile-avatar {
        position: relative;
        z-index: -1 !important;
        bottom: 8rem;
        display: grid;
        gap: 2.5rem;
        grid-auto-flow: row;
        justify-self: center;

        .avatar-image {
            width: 16rem;
            border: 0.5rem solid color(dark);
            border-radius: 50%;
            justify-self: center;
        }

        .profile-name {
            justify-self: center;
            font-weight: 700;
            font-size: 1.6rem;
            color: color(white);
        }
    }
}
</style>