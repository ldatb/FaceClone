<template>
    <div class="feed">
        <NewPostForm :avatarurl=avatarurl />

        <div v-for="post in posts" :key="post.id" class="posts">
            <PostCard
            :username=username
            :avatarurl=avatarurl
            :name=name
            :postid=post.id
            :time=post.time
            :description=post.description
            :mediaurl=post.media_url
            :likes=post.likes
            :hearts=post.hearts
            :laughs=post.laughs
            :sads=post.sads
            :angries=post.angries
            />
        </div>
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
.feed {
    display: grid;
    grid-auto-rows: max-content;
    gap: 1.3rem;
    padding-top: 3rem;
    justify-self: center;
    @include screen('infinity', 'large') {
        max-width: 650px;
    }
    @include screen('small') {
        width: 90%;
    }
}
.posts {
    display: grid;
    gap: 1.5rem;
}
</style>