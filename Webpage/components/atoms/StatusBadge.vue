<template>
    <span class="status-badge" :class="handleStatus"></span>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        // eslint-disable-next-line vue/require-default-prop
        status: {
            type: String
        }
    },
    computed: {
        handleStatus(): object {
            const classes = {
                'status-online': this.status === 'online',
                'status-offline': this.status === 'offline',
                'status-busy': this.status === 'busy',
            }
            return classes
        },
    },
})
</script>

<style lang="scss" scoped>
.status-badge {
    position: absolute;
    width: 0.75rem;
    height: 0.75rem;
    bottom: 1px;
    right: -1px;
    border: 2px solid #18232f;
    border-radius: 50%;
}
.status-online {
  background: color(green);
}
.status-busy {
  background: color(red);
}
.status-offline {
  background: color(gray, shade2);
}
</style>