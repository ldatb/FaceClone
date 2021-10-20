<template>
    <NuxtLink v-if="btnLink" class="base-button" :to="link">{{ text }}></NuxtLink>
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
            required: true
        },
        link: {
            type: String,
            default: "error"
        },
        btnLink: {
            type: Boolean,
            default: false
        }
    }
})
</script>

<style lang='scss' scoped>
button, .base-button {
    background: color(blue);
    color: color(white);
    padding: 0.6rem 0;
    min-height: 2.3rem;
    display: grid;
    align-items: center;
    justify-content: center;
    border-radius: 1.5rem;
    font-size: 0.97rem;
    transition: all 300ms ease;
    cursor: pointer;
    &:hover {
        box-shadow: 0px 2px 11px -2px rgba(69, 137, 247, 0.8);
    }
    &:active {
        background: color(blue, shade2);
    }
}
button {
    outline: none;
}
</style>