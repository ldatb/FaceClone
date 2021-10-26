<template>
    <div class="open-chat-container">
        <div class="chat-messages">
            <MessageBox 
                message="0 Lorem ipsum dolor sit amet."
            sender
            class="message-box-container" />
            <MessageBox 
                message=" 1Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            class="message-box-container" />
            <MessageBox 
                message="2 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            sender
            class="message-box-container" />
            <MessageBox 
                message="3 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            class="message-box-container" />
            <MessageBox 
                message="4 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            sender
            class="message-box-container" />
            <MessageBox 
                message="5 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            sender
            class="message-box-container" />
            <MessageBox 
                message="6 Lorem ipsum dolor sit amet."
            sender
            class="message-box-container" />
            <MessageBox 
                message="7 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            class="message-box-container" />
            <MessageBox 
                message="8 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            sender
            class="message-box-container" />
            <MessageBox 
                message="9 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            class="message-box-container" />
            <MessageBox 
                message="10 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            sender
            class="message-box-container" />
            <MessageBox 
                message="11 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque libero diam, feugiat quis nulla in, luctus malesuada lorem. Duis sollicitudin scelerisque mi ut aliquet."
            sender
            class="message-box-container" />
        </div>

        <div class="chat-actions">
            <ExpandableTextArea v-model="message" placeholder="Hi! Type a message" type="text" class="message-input" />
            <button class="message-send-button" @click="sendMessage">
                <img src="@/assets/img/send.svg" />
            </button>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    data() {
        return {
            message: '',
        }
    },
    methods: {
        sendMessage() {
            return null
        },
    },
})
</script>

<style lang="scss" scoped>
.open-chat-container {
    align-self: end;
    padding-bottom: 1.2rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 55vw;
    height: 93vh;
}
.chat-messages {
    flex: 1;
    display: grid;
    grid-auto-flow: row;
    gap: 1rem;
    padding: 0 1rem 0 1rem;
    height: auto;
    overflow: auto;
}
.chat-actions {
    display: grid;
    grid-auto-flow: column;
    justify-content: center;
    gap: 1rem;

    .message-input {
        align-self: center;
        justify-self: center;
        background: color(dark, shade2);
        width: 56rem;
        border-radius: 20px;
        padding: 0.8rem 1rem 0.8rem 1.5rem;
    }

    .message-send-button {
        display: grid;
        justify-content: center;
        align-content: center;
        background: color(blue);
        border-radius: 50%;
        width: 2.8rem;
        height: 2.8rem;
        cursor: pointer;

        img {
            width: 1.2rem;
        }
    }
}
</style>