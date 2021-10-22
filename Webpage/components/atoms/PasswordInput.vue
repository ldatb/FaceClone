<template>
    <div>
        <input :value="value" :placeholder="placeholder" :type="type" @input="updateValue" />
        <button class="button-show" @click="showPassword">
            <img :src="require(`@/assets/img/${imageURL}`)" />
        </button>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        value: {
            type: [String, Number],
            default: null
        },
        placeholder: {
            type: String,
            default: ""
        },
    },
    data() {
        return {
            imageURL: 'show_icon.svg',
            type: "password",
        };
    },
    methods: {
        updateValue(event: any) {
            this.$emit("input", event.target.value)
        },
        showPassword() {
            if (this.type === "password") {
                this.type = 'text'
                this.imageURL = 'hide_icon.svg'
            } else {
                this.type = 'password'
                this.imageURL = 'show_icon.svg'
            }
        }
    }
})
</script>

<style lang='scss' scoped>
div {
    display: grid;
    grid-template-columns: auto auto;
    background: color(white);
    border-radius: 20px;
}
input {
    outline: none;
    background: transparent;
    border-radius: 20px;
    min-height: 2.25rem;
    width: 100%;
    padding: 0 1rem;
}
button {
    display: grid;
    background: none;
    align-content: center;
    justify-content: center;

    img {
        height: 2rem;
        width: auto;
    }
}
</style>