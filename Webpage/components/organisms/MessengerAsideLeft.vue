<template>
    <Aside class="aside-chats-container">
        <div class="title-container">
            <b>Chat Messenger</b>
        </div>

        <div class="aside-chats">
            <AsideChat name="Rafaela Lopes" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" selected class="side-chat" />
            <AsideChat name="Alexandre Mendes" lastmessage="Iae men kk Iae men kk Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Gabriel" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Nataniel Gomes" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Júlio César" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Juscelino Kubitscheck" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Augusto Pereira" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Rafael Hilário" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Heloise Vasconçelos" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Paulo Lucas" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Rafael Hilário" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Heloise Vasconçelos" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Paulo Lucas" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Rafael Hilário" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
            <AsideChat name="Heloise Vasconçelos" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" class="side-chat" />
            <AsideChat name="Paulo Lucas" lastmessage="Iae men kk" avatarurl="profile-pic.jpg" unread class="side-chat" />
        </div>
    </Aside>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.aside-chats-container {
    display: grid;
    grid-auto-flow: row;
    width: fit-content;
    background: color(dark, shade1);

    .title-container {
        padding-left: 5rem;
        padding-top: 3rem;
        padding-bottom: 3rem;

        b {
            color: white;
            font-size: 1.45rem;
            
        }
    }
}

.aside-chats {
    display: grid;
    grid-auto-flow: row;
}
</style>