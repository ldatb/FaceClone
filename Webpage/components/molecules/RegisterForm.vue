<template>
    <form>
        <div class="form-field">
            <BaseInput v-model="email" type="email" placeholder="Email" class="input" />
        </div>

        <BaseButton text="Next step" class="button" />
    </form>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    data() {
        return {
            email: ''
        }
    }
})
</script>

<style lang="scss" scoped>
form {
    display: grid;
    gap: 0.8rem;

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
    }
}
</style>