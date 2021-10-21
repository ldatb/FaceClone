<template>
    <FullscreenContainer class="fullscreen-container">
        <HomeLeftAside class="home-aside"/>
        <HomeFeed class="home-feed" />
        <HomeRightAside class="home-aside" />
    </FullscreenContainer>
</template>

<style lang="scss" scoped>
.fullscreen-container {
    display: grid;
    grid-template-columns: 19rem 1fr 19rem;
    @include screen('large', 'medium') {
        grid-template-columns: auto 17rem;
        padding-left: 2.4rem;
        grid-gap: 3rem;
    }
    @include screen('small') {
        grid-template-columns: auto;
        padding-left: 0;
    }
}
.home-feed {
    padding-bottom: 4rem;
}
.home-aside {
    @include screen('infinity') {
        display: block !important;
    }
}
</style>