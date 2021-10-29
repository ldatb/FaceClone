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
        :posts=posts
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
            posts: [],
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

        // Get user posts
        await this.$axios.get(`/posts/user-posts/${this.username}`).then(response => {
            if (response.data.posts != null) {
                this.posts = response.data.posts
                this.postsquantity = this.posts.length
            }
        })

        // Get user in private route to check if it's the same as the requested user
        await this.$axios.get('/private/user').then(response => {
            if (response) {
                if (response.data.user.username === data.user.username) {
                    this.owner = true
                }
            }
        })
    },
})
</script>
