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
            default: ""
        },
        type: {
            type: String,
            default: ""
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
div {
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
    
    &::placeholder {
        text-justify: center;
    }
}
</style>