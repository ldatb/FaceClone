<template>
    <form>
        <div class="inputs-and-password">
            <div class="form-field">
                <BaseInput class="form-input" type="email" placeholder="Email" />
            </div>
            <div class="form-field">
                <BaseInput class="form-input" type="password" placeholder="Password" />
            </div>

            <NuxtLink class="recover-button" to="/recover-password">
                Forgot your password?
            </NuxtLink>
        </div>

        <BaseButton class="form-button" text="Login" />
    </form>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({})
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
.recover-button {
    justify-self: end;
    font-size: 14px;
    color: color(white);
}
.form-button {
    width: 100%;
}
</style>