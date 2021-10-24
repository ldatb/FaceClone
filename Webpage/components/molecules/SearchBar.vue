<template>
    <div class="search-bar">
        <img src="@/assets/img/facebook-icon.svg" class="logo" />

        <input v-model="searchcontent" type="text" placeholder="Search on facebook">

        <NuxtLink :to="`/search?query=${searchcontent}`" class="search-button">
            <img src="@/assets/img/search-icon.svg" class="search-button-image">
        </NuxtLink>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    data() {
        return {
            searchcontent: '',
        }
    },
})
</script>

<style scoped>
.search-bar {
    display: grid;
    grid-template-columns: 3.2rem 1fr 2rem;
    justify-content: center;
    align-content: center;
    padding: 0.4rem;
    width: 30rem;
    height: 2.6rem;
    background: #152330;
    border-radius: 2rem;
}
.logo {
    width: 2.1rem;
}
input {
    background: none;
    padding-left: 0.2rem;
}
.search-button {
    background: none;
    padding-top: 0.45rem;
}
.search-button-image {
    width: 1.26rem;
}
</style>