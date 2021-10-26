<template>
    <div class="post-card">
        <div class="post-card-info">
            <NuxtLink to="/profile" class="profile-avatar">
                <img src="@/assets/img/profile-pic.jpg" alt="" />

                <div class="profile-spam">
                    <p>User name</p>
                    <span>18:37</span>
                </div>
            </NuxtLink>

            <div>
                <p class="post-text">
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Laborum non
                    labore sit nesciunt fuga fugit ut, voluptas iusto.
                </p>
            </div>
        </div>

        <div class="card-image">
            <img src="@/assets/img/post1.png" alt="" />
        </div>

        <div class="reactions-counter">
            <div class="counter-box">
                <img src="@/assets/img/like-icon.svg" alt="" />
                <span>13</span>
            </div>

            <div class="counter-box">
                <img src="@/assets/img/heart-icon.svg" alt="" />
                <span>08</span>
            </div>
        </div>

        <div class="post-actions">
            <div class="button-reactions">
                <ul class="reactions">
                    <li class="reaction-like">
                        <img src="@/assets/img/reactions/like.svg" alt="" />
                    </li>
                    <li class="reaction-heart">
                        <img src="@/assets/img/reactions/heart.svg" alt="" />
                    </li>
                    <li class="reaction-laugh">
                        <img src="@/assets/img/reactions/laugh.svg" alt="" />
                    </li>
                    <li class="reaction-sad">
                        <img src="@/assets/img/reactions/sad.svg" alt="" />
                    </li>
                    <li class="reaction-angry">
                        <img src="@/assets/img/reactions/angry.svg" alt="" />
                    </li>
                </ul>

                <img src="@/assets/img/thumb-icon.svg" alt="" />
                <span>Reagir</span>
            </div>

            <div class="button-open-comments" @click="openComments = !openComments">
                <img src="@/assets/img/comment-icon.svg" alt="" />
                <span>Comment</span>
            </div>
        </div>

        <div class="post-comments" :class="{ 'open-comments': openComments }">
            <div class="comment-form">
                <form>
                    <img src="@/assets/img/profile-pic.jpg" alt="" />
                    <ExpandableTextArea v-model="text" class="grow-wrap" placeholder="Write your comment"/>
                </form>
            </div>

            <div class="comments">
                <div class="comment">
                    <img src="@/assets/img/profile-pic.jpg" alt="" />
                    <div class="comment-content">
                        <span>Deborah Gomes</span>
                        <p>
                            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Non
                            culpa temporibus unde doloremque eos dolore.
                        </p>
                    </div>
                </div>

                <div class="comment">
                    <img src="@/assets/img/profile-pic.jpg" alt="" />
                    <div class="comment-content">
                        <span>Deborah Gomes</span>
                        <p>
                            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Non
                            culpa temporibus unde doloremque eos dolore.
                        </p>
                    </div>
                </div>
            </div>

            <div class="commments-actions">
                <button>
                    Show more
                </button>
                <button @click="openComments = !openComments">
                    Hide comments
                </button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    data(): any {
        return {
            openComments: false,
            text: "",
        }
    },
})
</script>

<style lang="scss" scoped>
.post-card {
    background: color(dark, shade1);
    border-radius: 1.25rem;
    box-shadow: 1px 3px 9px rgba(1, 1, 1, 0.2);
    padding-bottom: 1.5rem;
}
.post-card-info {
    display: grid;
    gap: 1rem;
    padding: 1.2rem;

    .profile-avatar {
        display: grid;
        grid-template-columns: repeat(2, auto);
        gap: 1.2rem;
        width: fit-content;
        justify-content: start;
        align-items: center;

        .profile-spam {
            display: grid;
            gap: 0.3rem;

            p {
                color: color(white);
                font-weight: 700;
            }

            span {
                color: color(gray, shade1);
                font-size: 0.875rem;
            }
        }

        img {
            width: 3.5rem;
            border-radius: 50%;
        }
    }
}
.post-text {
    color: color(white);
}
.card-image {
    img {
        width: 100%;
    }
}
.reactions-counter {
    display: grid;
    grid-auto-flow: column;
    gap: 1rem;
    justify-content: start;
    padding-left: 1.7rem;
    position: relative;
    top: -21px;

    .counter-box {
        display: grid;
        gap: 0.6rem;
        position: relative;
        color: #98a9bc;
        align-items: center;
        justify-items: center;

        img {
            width: 1.875rem;
        }
    }
}
.post-actions {
    display: grid;
    grid-auto-flow: column;
    grid-template-columns: repeat(2, 1fr);
    border-top: 1px solid #23354b;
    border-bottom: 1px solid #23354b;

    > div {
        cursor: pointer;
        display: grid;
        gap: 0.6rem;
        grid-auto-flow: column;
        color: #f4f6f8;
        font-weight: 700;
        align-items: center;
        justify-content: center;
        padding: 1rem 0;
    }

    .button-reactions {
        position: relative;

        .reactions {
            visibility: hidden;
            opacity: 0;
            width: auto;
            position: absolute;
            display: grid;
            grid-auto-flow: column;
            gap: 0.6rem;
            top: -3.8rem;
            left: 0.5rem;
            background: #1f354d;
            padding: 0.3rem 0.4rem;
            border-radius: 3.5rem;
            box-shadow: 1px -6px 8px rgba(0, 0, 0, 0.25);
            transition: all 300ms ease;
            @include screen('medium', 'small') {
                top: -3.2rem;
            }

            li {
                display: grid;
                align-items: center;
                position: relative;
                transform: scale(0.2);
                opacity: 0;
                transition: all 300ms ease;

                img {
                    width: 3.2rem;
                    transition: all 300ms ease;
                    @include screen('medium', 'small') {
                    width: 2.5rem;
                    }
                    &:hover {
                        transform: scale(1.2);
                    }
                }
            }
        }

        &:hover .reactions {
            visibility: visible;
            opacity: 1;

            li {
                transform: scale(1);
                opacity: 1;
            }

            .reaction-like {
                transition-delay: 0.04s;
            }

            .reaction-heart {
                transition-delay: 0.08s;
            }

            .reaction-laugh {
                transition-delay: 0.12s;
            }

            .reaction-sad {
                transition-delay: 0.16s;
            }

            .reaction-angry {
                transition-delay: 0.20s;
            }
        }
    }
}
.post-comments {
    display: grid;
    gap: 1.4rem;
    padding: 1.2rem;
    padding-bottom: 0;
    height: 0;
    opacity: 0;
    overflow: hidden;
    transition: all 300ms ease;

    .comment-form {
        form {
            display: grid;
            grid-template-columns: auto 1fr;
            gap: 1rem;

            img {
                width: 2.6rem;
                border-radius: 50%;
            }

            .grow-wrap {
                background: #112331;
                padding: 0.9rem 1.375rem;
                border-radius: 2rem;
            }
        }
    }

    .comments {
        display: grid;
        gap: 0.7rem;

        .comment {
            display: grid;
            gap: 1rem;
            grid-template-columns: auto 1fr;
            align-items: center;

            .comment-content {
                display: grid;
                gap: 0.5rem;
                justify-content: start;
                border-radius: 1.25rem;
                padding: 0.9rem 1.375rem;
                background: color(dark, shade2);
                color: color(white);

                span {
                    font-size: 0.875rem;
                    font-weight: 700;
                }
            }

            img {
                width: 2.7rem;
                border-radius: 50%;
            }
        }
    }

    .commments-actions {
        display: grid;
        grid-auto-flow: column;
        justify-content: space-between;

        button {
            outline: none;
            background: none;
            color: color(gray, shade1);
            cursor: pointer;
        }
    }
}
.open-comments {
    height: auto;
    opacity: 1;
}
</style>