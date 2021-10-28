<template>
    <div>
        <ProfileTemplate
        :username=username
        :avatarurl=avatarurl
        :name=name 
        :postsquantity=postsquantity 
        :followers=followers 
        :following=following
        :owner=owner
        />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    layout: 'default',
    middleware: 'auth',
    data() {
        return {
            username: '',
            avatarurl: '',
            name: '',
            postsquantity: 0,
            followers: 0,
            following: 0,
            owner: false,
        }
    },
    async fetch() {
        const { data } = await this.$axios.get(`/users/user?keyword=${this.$route.params.slug}`)

        if (data) {
            this.username = data.user.username
            this.avatarurl = data.user.avatar_url
            this.name = data.user.fullname
            this.followers = data.user.followers
            this.following = data.user.following
        }

        // Get user in private route to check if it's the same as the requested user
        const privateResponse = await this.$axios.get('/private/user')
        if (privateResponse) {
            if (privateResponse.data.user.username === data.user.username) {
                this.owner = true
            }
        }
    }
})
</script>
