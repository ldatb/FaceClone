<template>
    <Container class="feed-container">
        <NewPostForm />

        <div class="posts-container">
            <PostCard v-for="n in 3" :key="n" />
        </div>
    </Container>
</template>

<style lang="scss" scoped>
.feed-container {
    display: grid;
    gap: 1.3rem;
}
.posts-container {
    display: grid;
    gap: 1.5rem;
}
</style>