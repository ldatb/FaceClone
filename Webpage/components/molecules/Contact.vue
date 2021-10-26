<template>
    <li class="contact">
        <NuxtLink to="/messenger">
            <div class="contact-image">
                <img :src="require(`@/assets/img/${imageurl}`)" alt="" />
                <StatusBadge class="status-badge" :status="getStatus" />
            </div>

            <div class="contact-info">
                <p>{{ username }}</p>
                <span class="last-message">
                    {{ lastmessage }}
                </span>
            </div>
        </NuxtLink>
    </li>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: {
        imageurl: {
            type: String,
            required: true
        },
        status: {
            type: String,
            required: true
        },
        username: {
            type: String,
            required: true
        },
        lastmessage: {
            type: String,
            default: '',
        },
    },
    computed: {
        getStatus(): string {
            return this.status
        },
    },
})
</script>

<style lang="scss" scoped>
.contact {
    transition: all 300ms ease;
    cursor: pointer;

    a {
        display: grid;
        grid-template-columns: auto 1fr;
        gap: 1rem;
        align-items: center;

        .contact-image {
            width: 2.6rem;
            display: grid;
            align-items: center;
            position: relative;

            .status-badge {
                width: 0.75rem;
                height: 0.75rem;
                border-radius: 50%;
                position: absolute;
                bottom: 0.063rem;
                right: -0.063rem;
                border: 2px solid #18232f;
            }

            img {
                width: 100%;
                border-radius: 50%;
            }
        }

        p {
            color: color(white);
            font-weight: bold;
            margin-top: 6px;
        }
    }

    .contact-info {
        display: grid;
        gap: 0.3rem;

        span {
            color: color(gray, shade1);
        }
    }

    &:hover {
        background: color(dark, shade2);
    }
}
</style>