<template>
    <NuxtLink class="profile-tab" :to="link">
        <span>{{ number }}</span>
        <b>{{ text }}</b>
    </NuxtLink>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        link: {
            type: String,
            required: true
        },
        number: {
            type: Number,
            required: true
        },
        text: {
            type: String,
            required: true
        }
    }
})
</script>

<style lang="scss" scoped>
.profile-tab {
    display: grid;
    grid-auto-flow: row;
    justify-content: center;
    justify-self: center;
    text-align: center;
    gap: 0.4rem;
    width: 7rem;
    color: color(white);
    transition: 200ms all ease;

    span {
        color: color(gray, shade1);
    }

    &:hover {
        border-radius: 4px;
        background: rgba(0, 0, 0, 0.1);
        color: color(blue);

        span {
            color: color(gray);
        }
    }
}
</style>