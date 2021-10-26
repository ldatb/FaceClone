<template>
    <button class="side-chat-container" :class="{'selected-chat': selected}" @click="openChat">
        <div class="side-chat">
            <img :src="require(`@/assets/img/${avatarurl}`)">

            <div class="side-chat-info">
                <b v-if="name.length<20">{{ name }}</b>
                <b v-else>{{ name.substring(0,19)+"..." }}</b>

                <span v-if="lastmessage.length<30">{{ lastmessage }}</span>
                <span v-else>{{ lastmessage.substring(0,29)+" ..." }}</span>
            </div>
        </div>

        <div :class="{'unread-message': unread}"></div>
    </button>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        name: {
            type: String,
            required: true,
        },
        lastmessage: {
            type: String,
            default: '',
        },
        avatarurl: {
            type: String,
            required: true,        
        },
        unread: {
            type: Boolean,
            default: false,
        },
        selected: {
            type: Boolean,
            default: false,
        }
    },
    methods: {
        openChat() {
            return null
        },
    },
})
</script>

<style lang="scss" scoped>
.side-chat-container {
    display: grid;
    grid-auto-flow: column;
    align-content: center;
    justify-content: start !important;
    background: none;
    width: 27rem;
    height: 4.5rem;
    cursor: pointer;

    .side-chat {
        padding-left: 5rem;
        display: grid;
        grid-auto-flow: column;
        gap: 1rem;


        img {
            width: 3rem;
            border-radius: 50%;
        }

        .side-chat-info {
            display: grid;
            grid-auto-flow: row;
            gap: 0.1rem;
            align-self: center;
            color: color(white);

            b {
                justify-self: start;
                font-size: 1.05rem;
            }

            span {
                justify-self: start;
                font-size: 0.84rem;
            }
        }
    }

    .unread-message {
        position: absolute;
        left: 24rem;
        width: 0.9rem;
        height: 0.9rem;
        border-radius: 50%;
        background: color(red);
        align-self: center;
        justify-self: center;
    }

    &:hover {
        background: color(dark);
    }
}
.selected-chat {
    background: color(dark, shade2);

    &:hover {
        background: color(dark, shade2) !important;
    }
}
</style>