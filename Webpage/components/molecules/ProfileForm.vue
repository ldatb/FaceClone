<template>
    <form>
        <div class="form-container">
            <div class="form-field">
                <label>Name</label>
                <BaseInput class="input-container" />
            </div>

            <div class="form-field">
                <label>Username</label>
                <BaseInput class="input-container" />
            </div>
        </div>

        <div class="form-container">
            <div class="form-field">
                <label>Email</label>
                <BaseInput class="input-container" />
            </div>


            <div class="form-field">
                <label>New password</label>
                <BaseInput type="password" class="input-container" />
            </div>


            <div class="form-field">
                <label>Repeat new password</label>
                <BaseInput type="password" class="input-container" />
            </div>
        </div>

        <div class="form-actions">
            <BaseButton text="Send"/>
        </div>
    </form>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
        
})
</script>

<style lang="scss" scoped>
form {
    display: grid;
    gap: 2rem;
    @include screen('medium', 'small') {
        gap: 4rem;
    }
}
.form-container {
        display: grid;
        gap: 0.6rem;
}
.form-field {
    display: grid;
    grid-template-columns: 0.5fr 1fr;
    gap: 1rem;
    align-items: center;
    text-align: start;
    @include screen('medium', 'small') {
      grid-template-columns: 1fr;
      gap: 0.6rem;
    }

    label {
        text-align: right;
        color: color(white);
        font-weight: 700;
        @include screen('medium', 'small') {
            text-align: left;
        }
    }

    .input-container {
        background: color(dark);
        width: 100% !important;
    }
}
.form-actions {
    display: grid;
    grid-template-columns: 0.5fr 1fr;
    grid-template-areas: "_ Button";
    justify-self: center;
    margin-left: -1rem;
    gap: 1rem;
    margin-top: 1.3rem;
    @include screen('medium', 'small') {
        grid-template-columns: 1fr;
        grid-template-areas: 'Button';
    }

    button {
        width: max-content;
        grid-area: Button;
    }
}
</style>