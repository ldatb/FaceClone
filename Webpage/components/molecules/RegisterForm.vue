<template>
    <form @submit.prevent>
        <div class="form-field">
            <BaseInput v-model="firstname" type="text" placeholder="First Name" class="form-input" />
            <BaseInput v-model="lastname" type="text" placeholder="Last Name" class="form-input" />
            <BaseInput v-model="email" type="email" placeholder="Email" class="form-input" />
            <PasswordInput v-model="password" placeholder="Password" class="form-input" :class="{ 'input-error': passMatch }" />
            <PasswordInput v-model="confirm" placeholder="Password" class="form-input" :class="{ 'input-error': passMatch }" />
        </div>

        <BaseButton text="Register" class="register-button" @click="submitRegister" />
    </form>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    data() {
        return {
            firstname: '',
            lastname: '',
            email: '',
            password: '',
            confirm: '',
            passMatch: false,
            passtype: 'password',
            passimageurl: 'show.png',
        }
    },
    methods: {
        async submitRegister(): Promise<void> {
            // Check if there isn't a empty answer
            if (this.firstname === '' || this.lastname === '' || this.email === '' || this.password === '' || this.confirm === '') {
                return this.$notify({type: 'error', text: "Please fill in all fields", duration: 5000})
            }

            // Password and confirmation don't match
            if (this.password !== this.confirm) {
                this.passMatch = true
                return this.$notify({type: 'error', text: "Passwords doesn't not match", duration: 5000})
            } else {
                this.passMatch = false
            }

            // Send to api
            const response = await this.$axios.$post('/users/register', {
                    name: this.firstname,
                    lastname: this.lastname,
                    email: this.email,
                    password: this.password,
            }).catch(error => {
                if (error.response.data.error === "user already exist") {
                    return this.$notify({type: 'error', text: "There's already an user with this email"})
                } else {
                    return this.$notify({type: 'error', text: 'Oops.. something went wrong'})
                }
            })

            // All good
            if (response) {
                this.$notify({type: 'success', text: 'All good! You can log in now.'})
                this.$router.push({path: '/login'})
            }
        },
    }
})
</script>

<style lang="scss" scoped>
form {
    display: grid;
    gap: 2rem;

    .form-field {
        display: grid;
        gap: 0.6rem;
    }

    .register-button {
        width: 100%;
        font-weight: 700;
    }

    a {
        justify-self: end;
        font-size: 14px;
        color: color(white);
    }

    .form-input {
        width: 100% !important;
        padding: 0.1rem 1.2rem;
        background: color(dark, shade1) !important;
    }

    .input-error {
        border: 1px solid red !important;
    }
}
</style>