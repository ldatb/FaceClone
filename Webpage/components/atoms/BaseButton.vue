<template>
    <NuxtLink v-if="btnlink" class="button-link" :to="link">
        {{ text }}
    </NuxtLink>
    <button v-else @click="$emit('click')">
        {{ text }}
    </button>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        text: {
            type: String,
            required: true,
        },
        link: {
            type: String,
            required: false,
            default: "/",
        },
        btnlink: {
            type: Boolean,
            default: false,
        },
    },
})
</script>

<style lang='scss' scoped>
button, .button-link {
    background: color(blue);
    color: color(white);
    padding: 0.6rem 1rem;
    min-width: 2.3rem;
    display: grid;
    align-items: center;
    justify-content: center;
    border-radius: 1.25rem;
    font-size: 0.97rem;
    transition: all 300ms ease;
    cursor: pointer;
    &:hover {
        box-shadow: 0px 2px 11px -2px rgba(69, 137, 247, 0.45);
    }
    &:active {
        background: color(blue, shade2);
    }
}
button {
    outline: none;
}
</style>