<template>
    <Aside :class="{ 'is-menu-active': isMenuActive }">
        <div class="content">
            <div class="profile-avatar">
                <img src="@/assets/img/profile-pic.jpg" alt="Profile Picture">

                <p>Teste de Perfil</p>

                <BaseButton btnlink text="Profile" link="/profile" @click.native="toggleMenuActive" />
            </div>

            <div class="aside-links">
                <AsideLink imageurl="messenger-link.svg" text="Messenger" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="followers-link.svg" text="Followers" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="groups-link.svg" text="Groups" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="pages-link.svg" text="Pages" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="events-link.svg" text="Events" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="foundations-link.svg" text="Foundations" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="memories-link.svg" text="Memories" @click.native="toggleMenuActive"/>
                <AsideLink imageurl="videos-link.svg" text="Videos" @click.native="toggleMenuActive"/>
            </div>
        </div>

        <button class="button-close" @click="toggleMenuActive">
            <fa :icon="['fas', 'times']" class="closed" />
        </button>
    </Aside>
</template>

<script lang="ts">
import Vue from 'vue'
import { mobile } from '@/store'
export default Vue.extend({
    computed: {
        isMenuActive() {
            return mobile.$isMenuActive
        },
    },
    methods: {
        toggleMenuActive() {
            const body = document.querySelector('body') as HTMLElement
            const html = document.querySelector('html') as HTMLElement
            const width = document.body.clientWidth

            if (width > 1200) return null

            body.classList.toggle('overflow-hidden')
            html.classList.toggle('overflow-hidden')
            mobile.toggle()
        },
    }
})
</script>

<style lang="scss" scoped> 
.aside {
    background: color(dark, shade1);
    @include screen('large', 'medium', 'small') {
        position: fixed;
        top: 0;
        left: 0;
        z-index: 99999;
        height: 100vh;
        width: 100%;
        transform: translate3d(-100%, 0, 0);
        transition: all 300ms ease;
    }
    @include screen('infinity') {
        display: none;
    }
    &.is-menu-active {
        transform: translate3d(0, 0, 0);
        ::v-deep, .wrapper {
            position: relative;
            top: 0;
            overflow: auto;
            max-height: unset;
            height: inherit;
            overflow-x: hidden;
        }
    }
}
.content {
    padding: 3rem 1.4rem;
    display: grid;
    gap: 2rem;

    .profile-avatar {
        display: grid;
        gap: 1rem;
        justify-items: center;
        img {
            width: 6.5rem;
            border-radius: 50%;
        }
        p {
            font-weight: 700;
            color: color(white);
            font-size: 1.25rem;
        }
    }

    .aside-links {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 1.2rem;
    }
}
.button-close {
    position: absolute;
    top: 21px;
    right: 20px;
    background: none;
    outline: none;
    cursor: pointer;
    @include screen('large', 'infinity') {
        display: none;
    }

    .closed {
        font-size: 26px;
        color: white;
    }
}
</style>