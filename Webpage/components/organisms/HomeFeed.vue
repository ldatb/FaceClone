<template>
    <div class="feed">
        <NewPostForm />

        <div class="posts">
            <PostCard />
        </div>
    </div>
</template>

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