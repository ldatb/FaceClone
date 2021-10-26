<template>
    <button class="notification-container" @click="gotoNotification">
        <img src="@/assets/img/profile-pic.jpg">
        <span class="notification-message">{{ text }}</span>
        <span class="notification-time">{{ time }}</span>
    </button>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        text: {
            type: String,
            default: ''
        },
        time: {
            type: String,
            default: ''
        },
    },
    methods: {
        gotoNotification() {
            return null
        }
    }
})
</script>

<style lang="scss" scoped>
.notification-container {
    width: 27rem;
    height: 4rem;
    border-radius: 3rem;
    background: color(dark, shade2);
    display: grid;
    gap: 1rem;
    grid-auto-flow: column;
    align-content: center;
    padding: 0.2rem 1rem 0.2rem 0.5rem;
    cursor: pointer;
    transition: 200ms all ease;

    img {
        justify-self: start;
        align-self: center;
        width: 3rem;
        border-radius: 50%;
    }

    span {
        color: color(white);
    }

    .notification-message {
        align-self: center;
    }

    .notification-time {
        font-size: 0.7rem;
    }

    &:hover {
        background: color(blue, shade2);
        box-shadow: 4px 7px 10px rgba(0, 0, 0, 0.4);
    }
}
</style>