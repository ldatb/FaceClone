<template>
    <div class="profile-info">
        <div class="profile-cover">
            <img src="@/assets/img/capa.jpg" class="cover-image" />
        </div>
        <div class="profile-avatar">
            <img :src=avatarurl class="avatar-image" />
            <span class="profile-name">{{ name }}</span>
        </div>
    </div>
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