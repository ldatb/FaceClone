<template>
    <div class="profile-template">
        <Profile :avatarurl=avatarurl :name=name />

        <ProfileTabs :username=username :timeline=postsquantity :followers=followers :following=following activetimeline />

        <ProfileFeed class="profile-feed" :username=username :name=name :avatarurl=avatarurl :owner=owner :posts=posts />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        username: {
            type: String,
            required: true,
        },
        avatarurl: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
        postsquantity: {
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
        owner: {
            type: Boolean,
            required: true,
        },
        posts: {
            type: Array,
            required: true,
        },
    }
})
</script>

<style lang="scss" scoped>
.profile-template {
    display: grid;
    padding-bottom: 4rem;
}
</style>