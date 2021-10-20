<template>
    <div class="expand-wrapper">
        <textarea rows="1" :placeholder="placeholder" :value="value" @focus="$emit('focused')" @input="handleExpansion($event)"></textarea>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        placeholder: {
            type: String,
            default: "Placeholder"
        },
        value: {
            type: [String, Number],
            default: null
        }
    },
    methods: {
        handleExpansion(event: any) {
            event.target.parentNode.dataset.replicatedValue = event.target.value
            this.$emit('input', event.target.value)
        }
    }
})
</script>

<style lang="scss" scoped>
.expand-wrapper {
    display: grid;
    height: max-content;
}

.expand-wrapper::after {
  content: attr(data-replicated-value) ' ';
  white-space: pre-wrap;
  visibility: hidden;
}
.expand-wrapper > textarea {
  resize: none;
  overflow: hidden;
  background: none;
  color: color(white);
}
.expand-wrapper > textarea,
.expand-wrapper::after {
  grid-area: 1 / 1 / 2 / 2;
}
</style>