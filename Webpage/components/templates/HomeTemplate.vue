<template>
    <FullscreenContainer class="fullscreen-container">
        <HomeLeftAside class="home-aside" :username=username :avatarurl=avatarurl :name=name />
        <HomeFeed class="home-feed" :avatarurl=avatarurl />
        <HomeRightAside class="home-aside" />
    </FullscreenContainer>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data() {
    return {
      username: '',
      avatarurl: '',
      name: '',
    }
  },
  async fetch() {
    const { data } = await this.$axios.get('/private/user')

    if (data) {
      this.username = data.user.username
      this.avatarurl = data.user.avatar_url
      this.name = data.user.fullname
    }
  },
})
</script>

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