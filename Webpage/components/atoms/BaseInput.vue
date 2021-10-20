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
    background: color(dark, shade1);
    border-radius: 1.5rem;
    color: color(white);
    width: max-content;
    min-height: 2.3rem;
    padding-left: 1rem;
    padding-right: 1rem;
}
</style>