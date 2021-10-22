<template>
    <form @submit.prevent="onSubmit">
        <div class="form-field">
            <BaseInput v-model="firstname" type="text" placeholder="First Name" class="input" />
            <BaseInput v-model="lastname" type="text" placeholder="Last Name" class="input" />
            <BaseInput v-model="email" type="email" placeholder="Email" class="input" />
            <BaseInput v-model="password" type="password" placeholder="Password" class="input" />
        </div>

        <BaseButton text="Register" class="button" />
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
        }
    },
    methods: {
        async onSubmit() {
            try {
                await this.$axios.$post('/users/register', {
                    name: this.firstname,
                    lastname: this.lastname,
                    email: this.email,
                    password: this.password,
                }).catch(error => {
                    this.$notify({type: 'error', text: error.toString()})
                })

                this.$notify({type: 'success', text: 'All good! Check your email for the verification code.'})
            } catch {
                this.$notify({type: 'error', text: 'Oops... Something went wrong'})
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

    .button {
        width: 100%;
        font-weight: 700;
    }

    a {
        justify-self: end;
        font-size: 14px;
        color: color(white);
    }

    .input {
        width: 100% !important;
        padding: 0.1rem 1.2rem;
        background: color(dark, shade1) !important;
    }
}
</style>