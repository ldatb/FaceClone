<template>
    <Aside>
        <div class="content">
            <h3 class="aside-title">Contacts</h3>

            <div class="contacts-wrapper">
                <span class="status-title">Online</span>
                <ul class="contacts">
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Lucas de Ataides" status="online" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Gabriel Rocha Melo" status="busy" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Stephanny Motta" status="online" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Alexandre Mendes" status="busy" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Giovana dos Santos" status="online" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Ronaldo Schimth" status="online" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Heloise Vasconcelos" status="online" />
                </ul>
                
                <span class="status-title">Offline</span>
                <ul class="contacts">
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Deborah Gomes" status="offline" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Jéssica Bareoni" status="offline" />
                    <Contact class="contact" imageurl="profile-pic.jpg" username="Chris Nathaniel" status="offline" />
                </ul>
            </div>
        </div>
    </Aside>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
        
})
</script>

<style lang="scss" scoped>
.aside {
    @include screen('small') {
        display: none;
    }
}
.content {
    display: grid;
    gap: 1.6rem;
    padding: 3rem 0;

    .aside-title {
        color: color(white);
    }
}
.contacts-wrapper {
    display: grid;
    gap: 1rem;

    .status-title {
        color: color(gray, shade1);
        font-size: 0.875rem;
    }
}
.contacts {
    display: grid;
    gap: 0.4rem;

    .contact {
        border-radius: 27px;
        padding: 3px 0;
        padding-left: 3px;
    }
}
</style>