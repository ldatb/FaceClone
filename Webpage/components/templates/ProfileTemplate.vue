<template>
    <div class="profile-template">
        <div class="profile-info">
            <div class="profile-cover">
                <img src="@/assets/img/capa.jpg" class="cover-image" />
            </div>
            <div class="profile-avatar">
                <img src="@/assets/img/profile-pic.jpg" class="avatar-image" />
                <span class="profile-name">Lucas Barreto</span>
            </div>
        </div>

        <div class="profile-tabs">
            <ProfileTab link="/profile" text="Timeline" number=300 class="active-profile-tab" />
            <ProfileTab link="/profile/followers" text="Followers" number=1000 />
            <ProfileTab link="/profile/following" text="Following" number=20 />
        </div>

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
.profile-info {
    display: grid;
    grid-auto-flow: row;

    .profile-cover {
        height: 25rem;
        width: 100%;

        .cover-image {
            height: 100%;
            width: 100%;
            object-fit: cover !important;
        }
    }

    .profile-avatar {
        position: absolute;
        top: 21.5rem;
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
.profile-tabs {
    display: grid;
    grid-auto-flow: column;
    justify-content: center;
    gap: 4rem;
    padding-top: 18rem;

    .active-profile-tab {
        border-bottom: 2px solid color(blue);
    }
}
.profile-feed {
    padding-top: 4rem;
}
</style>