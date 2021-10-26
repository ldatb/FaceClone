<template>
    <div class="follow-container">
        <NuxtLink to="/profile" class="profile-picture">
            <img :src="require(`@/assets/img/${avatarurl}`)" />
        </NuxtLink>

        <div class="info-container">
            <NuxtLink to="/profile" class="user-info">
                <span class="user-name">{{ name }}</span>
                <span class="user-username">{{ username }}</span>
            </NuxtLink>

             <button @click="followAction">
                <span>{{ actiontext }}</span>
            </button>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        avatarurl: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
        username: {
            type: String,
            required: true,
        },
        actiontext: {
            type: String,
            default: "Follow / Unfollow",
        },
    },
    methods: {
        followAction() {
            return null
        }
    }
})
</script>

<style lang="scss" scoped>
.follow-container {
    display: grid;
    grid-auto-flow: row;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding-top: 1rem;
    z-index: -1 !important;

    .profile-picture {
        justify-self: center;

        img {
            width: 7rem;
            border: 4px solid color(dark, shade2);
            border-radius: 50%;
            position: relative;
            top: 4.3rem;
            transition: 300ms all ease;
        }
    }

    .info-container {
        width: 19rem;
        height: 13rem;
        border-radius: 2rem;
        background: color(dark, shade2);
        padding-top: 8rem;
        padding-bottom: 4rem;
        display: grid;
        align-content: center;
        justify-content: center;
        gap: 2rem;
        transition: 200ms all ease;

        .user-info {
            display: grid;
            gap: 0.3rem;
            grid-auto-flow: row;
            align-content: center;
            justify-content: center;
            text-align: center;

            .user-name {
                color: color(white);
                font-size: 1.3rem;
                font-weight: 700;
            }

            .user-username {
                color: color(gray, shade1);
                font-size: 1rem;
                font-weight: 500;
            }
        }

        button {
            justify-self: center;
            background: none;
            color: color(gray);
            width: fit-content;
            height: fit-content;
            cursor: pointer;

            &:hover {
                text-decoration: underline;
            }
        }
    }

    &:hover {
        .profile-picture {
            img {
                border: 7px solid color(dark, shade2);
            }
        }

        .info-container {
            box-shadow: 1px 4px 10px rgba(0, 0, 0, 0.5);
        }
    }
}
</style>