<template>
    <div class="profile-container">
        <div class="profile-avatar">
            <div class="avatar-image">
                <img src="@/assets/img/profile-pic.jpg">
            </div>

            <div class="avatar-info">
                <p>Nome do Usuário</p>
                <label for="avatar">Change profile picture</label>
                <input id="avatar" type="file">
            </div>
        </div>

        <ProfileForm class="profile-info" />

        <div class="disable-account">
            <button>Deactivate my account temporarily</button>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
        
})
</script>

<style lang="scss" scoped>
.profile-container {
    display: grid;
    gap: 2rem;
    background: color(dark, shade1);
    border-radius: 24px;
    padding: 1.4rem;
    margin-top: 3rem;

}
.profile-avatar {
    display: grid;
    grid-template-columns: 0.5fr 1fr;
    gap: 1rem;
    align-items: center;
    justify-self: center;

    .avatar-image {
        text-align: right;

        img {
            width: 5rem;
            border-radius: 50%;
        }
    }

    .avatar-info {
        p {
            color: color(white);
        }

        label {
            background: none;
            outline: none;
            color: color(blue);
            font-size: 14px;
            cursor: pointer;
        }

        input {
            display: none;
        }
    }
}
.profile-info {
    justify-self: center;
    width: 100% !important;
    margin: 0 1rem 0 -1rem;
}
.disable-account {
    display: grid;
    justify-content: center;

    button {
        outline: none;
        background: none;
        color: color(blue);
        font-size: 15px;
        font-weight: 700;
        cursor: pointer;
    }
}
</style>