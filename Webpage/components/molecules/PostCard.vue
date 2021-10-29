<template>
    <div class="post-card">
        <div class="post-card-info">
            <NuxtLink to="/profile" class="profile-avatar">
                <img :src=avatarurl />

                <div class="profile-spam">
                    <p>{{ name }}</p>
                    <span>{{ datetime }}</span>
                </div>
            </NuxtLink>

            <div>
                <p class="post-text">
                    {{ description }}
                </p>
            </div>
        </div>

        <div :v-if="mediaurl != ''" class="card-image">
            <img :src=mediaurl />
        </div>

        <div class="reactions-counter">
            <div v-if="numberlikes !== 0" :key="numberlikes" class="counter-box">
                <img src="@/assets/img/reactions/like.svg" alt="" />
                <span>{{ numberlikes }}</span>
            </div>

            <div v-if="numberhearts !== 0" :key="numberhearts" class="counter-box">
                <img src="@/assets/img/reactions/heart.svg" alt="" />
                <span>{{ numberhearts }}</span>
            </div>

            <div v-if="numberlaughs !== 0" :key="numberlaughs" class="counter-box">
                <img src="@/assets/img/reactions/laugh.svg" alt="" />
                <span>{{ numberlaughs }}</span>
            </div>

            <div v-if="numbersads !== 0" :key="numbersads" class="counter-box">
                <img src="@/assets/img/reactions/sad.svg" alt="" />
                <span>{{ numbersads }}</span>
            </div>

            <div v-if="numberangries !== 0" :key="numberangries" class="counter-box">
                <img src="@/assets/img/reactions/angry.svg" alt="" />
                <span>{{ numberangries }}</span>
            </div>
        </div>

        <div class="post-actions">
            <div class="button-reactions">
                <ul class="reactions">
                    <li class="reaction-like" @click="reactLike">
                        <img src="@/assets/img/reactions/like.svg" />
                    </li>
                    <li class="reaction-heart" @click="reactHeart">
                        <img src="@/assets/img/reactions/heart.svg" />
                    </li>
                    <li class="reaction-laugh" @click="reactLaugh">
                        <img src="@/assets/img/reactions/laugh.svg" />
                    </li>
                    <li class="reaction-sad" @click="reactSad">
                        <img src="@/assets/img/reactions/sad.svg" />
                    </li>
                    <li class="reaction-angry" @click="reactAngry">
                        <img src="@/assets/img/reactions/angry.svg" />
                    </li>
                </ul>

                <img src="@/assets/img/thumb-icon.svg" />
                <span>React</span>
            </div>

            <div class="button-open-comments" @click="openComments = !openComments">
                <img src="@/assets/img/comment-icon.svg" />
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

            <!--<div class="comments">
                <div class="comment">
                    <img src="@/assets/img/profile-pic.jpg" alt="" />
                    <div class="comment-content">
                        <span>Deborah Gomes</span>
                        <p>
                            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Non
                            culpa temporibus unde doloremque eos dolore.
                        </p>
                    </div>
                </div>-->

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
import axios from 'axios'
export default Vue.extend({
    props: {
        username: {
            type: String,
            required: true,
        },
        avatarurl: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
        postid: {
            type: Number,
            required: true,
        },
        time: {
            type: String,
            required: true,
        },
        description: {
            type: String,
            required: true,
        },
        mediaurl: {
            type: String,
            required: true,
        },
        likes: {
            type: Number,
            required: true,
        },
        hearts: {
            type: Number,
            required: true,
        },
        laughs: {
            type: Number,
            required: true,
        },
        sads: {
            type: Number,
            required: true,
        },
        angries: {
            type: Number,
            required: true,
        },
    },
    data(): any {
        let timeTarget = "Just now"

        const timeNow = Date.now()
        const timePostInt = Date.parse(this.time)
        let timeFromPost = timeNow - timePostInt
        if (timeFromPost > 60000) {
            timeFromPost = Math.floor(timeFromPost / 60000)
            timeTarget = timeFromPost + " minutes ago"

            if (timeFromPost > 60) {
                timeFromPost = Math.floor(timeFromPost / 60)
                timeTarget = timeFromPost + " hours ago"

                if (timeFromPost > 24) {
                    timeFromPost = Math.floor(timeFromPost / 24)
                    timeTarget = timeFromPost + " days ago"

                    if (timeFromPost > 30) {
                        const dateTarget = new Date(timePostInt)
                        timeTarget = dateTarget.getDay().toString() + " / " + dateTarget.getMonth().toString() + " / " + dateTarget.getFullYear().toString()
                    }
                }
            }
        }

        return {
            openComments: false,
            text: "",
            datetime: timeTarget,
            numberlikes: this.likes,
            numberhearts: this.hearts,
            numberlaughs: this.laughs,
            numbersads: this.sads,
            numberangries: this.angries,
        }
    },
    methods: {
        async reactLike() {
            await axios.post("http://localhost:3000/posts/react", {
                "postid": this.postid.toString(),
                "reaction": "like"
            })

            this.numberlikes += 1
            this.numberhearts = this.hearts
            this.numberlaughs = this.laughs
            this.numbersads = this.sads
            this.numberangries = this.angries
        },
        async reactHeart() {
            await axios.post("http://localhost:3000/posts/react", {
                "postid": this.postid.toString(),
                "reaction": "heart"
            })

            this.numberlikes = this.likes
            this.numberhearts += 1
            this.numberlaughs = this.laughs
            this.numbersads = this.sads
            this.numberangries = this.angries
        },
        async reactLaugh() {
            await axios.post("http://localhost:3000/posts/react", {
                "postid": this.postid.toString(),
                "reaction": "laugh"
            })

            this.numberlikes = this.likes
            this.numberhearts = this.hearts
            this.numberlaughs += 1
            this.numbersads = this.sads
            this.numberangries = this.angries
        },
        async reactSad() {
            await axios.post("http://localhost:3000/posts/react", {
                "postid": this.postid.toString(),
                "reaction": "sad"
            })

            this.numberlikes = this.likes
            this.numberhearts = this.hearts
            this.numberlaughs = this.laughs
            this.numbersads += 1
            this.numberangries = this.angries
        },
        async reactAngry() {
            await axios.post("http://localhost:3000/posts/react", {
                "postid": this.postid.toString(),
                "reaction": "angry"
            })

            this.numberlikes = this.likes
            this.numberhearts = this.hearts
            this.numberlaughs = this.laughs
            this.numbersads += this.sads
            this.numberangries += 1
        },
    }
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