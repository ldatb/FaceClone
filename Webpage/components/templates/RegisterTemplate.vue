<template>
  <div class="register-template-container">
      <div class="background-image"></div>

      <div v-if="true" class="auth-box">
          <img src="@/assets/img/facebook-logo.svg" alt="Facebook Logo" class="logo">

          <div class="welcome">
              <h2>Register now</h2>
              <p>It's easy and free!</p>
          </div>

          <RegisterForm />

          <AlreadyHaveAccount text="Already have an account?" />
      </div>

      <div v-else class="auth-box">
          <EmailSent />
      </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.register-template-container {
    display: grid;
    grid-template-columns: 1fr 500px;
    height: 100%;
    @include screen('medium', 'small') {
        grid-template-columns: 1fr;
    }

    .background-image {
        background-image: url('@/assets/img/capa-login.png');
        background-position: center;
        background-size: cover;
        @include screen('medium', 'small') {
        display: none;
        }
    }
}
.auth-box {
    display: grid;
    grid-auto-rows: max-content;
    gap: 3rem;
    padding: 5.4rem;
    align-content: center;
    @include screen('small') {
        padding: 3rem;
    }

    .logo {
        width: 11rem;
    }

    .welcome {
        line-height: 1.4;

        h2 {
            font-size: 22px;
            font-weight: 700;
            color: color(white);
        }

        p {
            font-size: 19px;
            color: color(gray, shade1);
        }
    }
}
</style>