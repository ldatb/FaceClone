<template>
    <NuxtLink :to="link" class="aside-link" @click="$emit('click')">
        <img :src="require(`@/assets/img/aside/${imageurl}`)" alt="Events">
        <p>{{ text }}</p>
    </NuxtLink>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    props: {
        imageurl: {
            type: String,
            required: true,
        },
        text: {
            type: String,
            required: true,
        },
        link: {
            type: String,
            default: '/',
        },
    }
})
</script>

<style lang="scss" scoped>
.aside-link {
    display: grid;
    padding: 1rem 0;
    gap: 0.8rem;
    justify-items: center;
    background: color(dark, shade2);
    border-radius: 0.938rem;
    transition: all 300ms ease;
    cursor: pointer;
    &:hover {
        transform: scale(1.04);
        box-shadow: 1px 2px 10px rgba(0, 0, 0, 0.25);
    }
}
img {
    height: 2.25rem;
    filter: drop-shadow(2px 2px 5px rgba(0, 0, 0, 0.3));
}
p {
    color: color(white);
}
</style>