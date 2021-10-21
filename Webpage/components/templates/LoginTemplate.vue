<template>
    <div class="login-template-container">
        <div class="background-image"></div>

        <div class="form-box">
            <img class="logo" src="@/assets/img/facebook-logo.svg" alt="Facebook logo">
            
            <div class="login-welcome-text">
                <p>Welcome to FaceClone!</p>
                <h2>Login to your account</h2>
            </div>

            <div class="login-register-box">
                <LoginForm class="login-form"/>

                <div class="register-box">
                    <p>
                        <span>Don't have an account?</span>
                        <NuxtLink to="/register">Create one!</NuxtLink>
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({})
</script>

<style lang="scss" scoped>
.login-template-container {
    display: grid;
    grid-template-columns: 1fr 40rem;
    height: 100%;
    @include screen('medium', 'small') {
        grid-template-columns: 1fr;
    }
}
.background-image {
    background-image: url("@/assets/img/capa-login.png");
    background-position: center;
    background-size: cover;
    @include screen('medium', 'small') {
        display: none;
    }
}
.form-box {
    display: grid;
    grid-auto-rows: max-content;
    gap: 3.7rem;
    justify-content: center;
    align-content: center;

    .logo {
        width: 14rem;
    }

    .login-welcome-text {
        display: grid;
        gap: 0.6rem;
        p {
            font-size: 17px;
            color: color(gray, shade1);
        }
        h2 {
            font-size: 20px;
            font-weight: 700;
            color: color(white);
        }
    }

    .login-register-box {
        display: grid;
        gap: 2.6rem;
    }

    .register-box {
        text-align: center;

        p {
            font-size: 16px;

            span {
                color: color(gray, shade1);
            }

            a {
                color: color(white);
                font-weight: 700;
            }
        }
    }
}
</style>