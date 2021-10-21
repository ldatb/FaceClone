<template>
    <div>
        <input :value="value" :placeholder="placeholder" :type="type" @input="updateValue" />
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
            default: "Placeholder"
        },
        type: {
            type: String,
            default: "text"
        }
    },
    methods: {
        updateValue(event: any) {
            this.$emit("input", event.target.value)
        }
    }
})
</script>

<style lang='scss' scoped>
input {
    outline: none;
    background: color(dark, shade1);
    border-radius: 20px;
    min-height: 2.25rem;
    width: 100%;
    padding: 0 1rem;
}
</style>