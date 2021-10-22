<template>
    <div class="login-redirect">
        <p>
            <span>{{ text }}</span>
            <NuxtLink to="/login">
                Login
            </NuxtLink>
        </p>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        text: {
            type: String,
            required: true,
        },
    },
})
</script>

<style lang="scss" scoped>
.login-redirect {
    text-align: center;

    p {
        font-size: 14px;

        span {
            color: color(gray, shade2);
        }

        a {
            color: color(white);
            font-weight: 700;
        }
    }
}
</style>