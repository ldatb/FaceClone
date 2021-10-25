<template>
    <div class="profile-tabs">
        <ProfileTab link="/profile" text="Timeline" :number=timeline :class="{ 'active-tab': activetimeline }" />
        <ProfileTab link="/profile/followers" text="Followers" :number=followers :class="{ 'active-tab': activefollowers }" />
        <ProfileTab link="/profile/following" text="Following" :number=following :class="{ 'active-tab': activefollowing }" />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        timeline: {
            type: Number,
            required: true,
        },
        followers: {
            type: Number,
            required: true,
        },
        following: {
            type: Number,
            required: true,
        },
        activetimeline: {
            type: Boolean,
            default: false,
        },
        activefollowers: {
            type: Boolean,
            default: false,
        },
        activefollowing: {
            type: Boolean,
            default: false,
        },
    }
})
</script>

<style lang="scss" scoped>
.profile-tabs {
    display: grid;
    grid-auto-flow: column;
    justify-content: center;
    gap: 4rem;
    padding-top: 18rem;

    .active-tab {
        border-bottom: 2px solid color(blue);
    }
}
</style>