<template>
    <div v-if="sender" class="message-box sender-box">
        <span>{{ message }}</span>
    </div>
    <div v-else class="message-box receiver-box">
        <span>{{ message }}</span>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        sender: {
            type: Boolean,
            default: false,
        },
        message: {
            type: String,
            required: true,
        },
    }
})
</script>

<style lang="scss" scoped>
.message-box {
    width: fit-content;
    height: fit-content;
    max-width: 42rem;
    padding: 1rem 1rem 1rem 1rem;
    color: color(white);
    font-size: 1rem;
}
.sender-box {
    background: color(blue) !important;
    border-radius: 20px 20px 0 20px;
    justify-self: end;
}
.receiver-box {
    background: color(dark, shade2) !important;
    border-radius: 20px 20px 20px 0px;
    justify-self: start;
}
</style>