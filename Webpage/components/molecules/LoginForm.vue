<template>
    <form @submit.prevent>
        <div class="inputs-and-password">
            <BaseInput v-model="email" type="email" placeholder="Email" class="form-input" />
            <PasswordInput v-model="password" placeholder="Password" class="form-input" />

            <NuxtLink class="recover-button" to="/recover-password">
                Forgot your password?
            </NuxtLink>
        </div>

        <BaseButton class="form-button" text="Login" @click="submitLogin" />
    </form>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    data() {
        return {
            email: '',
            password: '',
        }
    },
    methods: {
        async submitLogin(): Promise<void> {
            // Check if there isn't a empty answer
            if (this.email === '' || this.password === '') {
                return this.$notify({type: 'error', text: "Please fill in all fields", duration: 5000})
            }

            // Send to api
            const response = await this.$axios.$post('/users/login', {
                    email: this.email,
                    password: this.password,
            }).catch(error => {
                window.console.log(error.response.data.error)
                if (error.response.data.error === "invalid password") {
                    return this.$notify({type: 'error', text: "Incorrect password"})
                } if (error.response.data.error === "invalid user") {
                    return this.$notify({type: 'error', text: "This user does not exist"})
                } else {
                    return this.$notify({type: 'error', text: 'Oops.. something went wrong'})
                }
            })

            // All good
            if (response) {
                // Save JWT token
                

                window.console.log(response.token)
                this.$router.push({path: '/login'})
            }
        },
    }
})
</script>

<style lang="scss" scoped>
form {
    display: grid;
    gap: 1.5rem;
}
.inputs-and-password {
    display: inherit;
    gap: 0.7rem;
}
.form-input {
    width: 100% !important;
    padding: 0.1rem 0.8rem;
    background: color(dark, shade1) !important;
}
.recover-button {
    justify-self: end;
    font-size: 14px;
    color: color(white);
}
.form-button {
    width: 100%;
}
</style>