<template>
    <form ref="postform" class="post-form" :class="{'active-form': isWritingNewPost, 'has-text': hasText}" @submit.prevent="sendPost()">
        <div class="new-post-box">
            <img :src=avatarurl alt="Profile picture">
            <ExpandableTextArea v-model="text" class="grow-wrap" placeholder="What are you thinking about?" @focused="isWritingNewPost = true"/>
        </div>

        <div v-show="false" class="uploaded-image-preview">
            <img src="@/assets/img/post1.png" alt="" />
        </div>

        <div v-show="isWritingNewPost || hasText" class="form-action">
            <button class="upload-image-button">
                <img src="@/assets/img/camera-icon.svg" alt="Upload image">
            </button>

            <div class="post-commands">
                <button class="cancel-post-button" @click="cancelPost()">Cancel</button>
                <button class="send-post-button" >Send</button>
            </div>
        </div>
    </form>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    props: {
        avatarurl: {
            type: String,
            required: true,
        },
    },
    data(): any {
        return {
            text: '',
            isWritingNewPost: false,
        }
    },
    computed: {
        hasText(): boolean {
            return this.text.trim().length > 0
        },
    },
    mounted() {
        window.addEventListener('click', (e) => {
            if (this.$refs.postform.contains(e.target)) {
                return null
            } else {
                this.isWritingNewPost = false
            }
        })
    },
    methods: {
        sendPost() {
            // eslint-disable-next-line no-console
            console.log(this.text)
        },
        cancelPost() {
            this.text = ''
            this.isWritingNewPost = false
        },
    },
})
</script>

<style lang="scss" scoped>
.post-form {
    display: grid;
    gap: 1.5rem;
    align-content: start;
    width: 100%;
    min-height: 2.5rem;
    background: color(dark, shade1);
    border-radius: 1.25rem;
    overflow: hidden;
    transition: all 300ms ease;
    padding-top: 0.25rem;
}
.new-post-box {
    display: grid;
    grid-template-columns: auto 1fr;
    gap: 1rem;
    align-items: center;
    height: 100%;

    img {
        width: 2rem;
        border-radius: 50%;
        margin-left: 0.2rem;
    }
}
.uploaded-image-preview {
    width: 100%;
    padding: 0 1rem;

    img {
        width: 100%;
        max-height: 51vh;
        object-fit: cover;
        box-shadow: 1px 4.5px 9px rgba(0, 0, 0, 0.4);
    }
}
.form-action {
    display: grid;
    grid-auto-flow: column;
    justify-content: space-between;
    align-self: end;
    border-top: 1px solid #23354b;
    padding: 0.4rem 1rem;

    > div {
        display: grid;
        grid-auto-flow: column;
        gap: 0.5rem;
    }

    button {
        background: transparent;
        outline: one;
        color: color(blue);
        cursor: pointer;
        font-size: 0.9rem;
        padding: 0.3rem 1.2rem;
        border-radius: 1.8;
        transition: 300ms all ease;

        &:hover {
            color: #98a9bc;
            background: #112331;
            border-radius: 20px;
        }
    }

    .send-post-button {
        border-radius: 20px;
        background: color(blue, shade2);
        color: white;

        &:hover {
            background: color(blue);
            color: color(dark, shade2);
        }
    }
}
.upload-image-button {
    padding: 0;
    border-radius: none;
    &:hover {
        color: unset;
        background: unset;
    }
    img {
        width: 1.25rem;
    }
}
.active-form {
    align-content: unset;
    min-height: 4.5rem;
    padding-top: 0;
    transition: 300ms all ease;

    .new-post-box {
        align-items: start;
        padding-top: 0.2rem;
        min-height: inherit;

        .grow-wrap {
            padding-top: 0.44rem;
            min-height: inherit;
        }
    }
}
.has-text {
    height: auto;

    .new-post-box {
        align-items: start;

        .grow-wrap {
            padding-top: 0.44rem;
        }
    }
}

[placeholder]:empty::before {
    content: attr(placeholder);
    color: #667c95;
}
[placeholder]:empty:focus::before {
  content: '';
}
</style>