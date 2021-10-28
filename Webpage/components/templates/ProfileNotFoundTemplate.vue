<template>
    <div class="profile-template">
        <Profile avatarurl="NOT FOUND" name="NOT FOUND" />
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        avatarurl: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
        postsquantity: {
            type: String,
            required: true,
        },
        followers: {
            type: String,
            required: true,
        },
        following: {
            type: String,
            required: true,
        },
    }
})
</script>

<style lang="scss" scoped>
.profile-template {
    display: grid;
    padding-bottom: 4rem;
}
</style>