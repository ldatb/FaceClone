<template>
  <div class="email-sent">
      <fa :icon="['fas', 'check-circle']" class="check-circle" />
      <p>
          We have sent a message to your <b>email address</b> with the necessary instructions
          to complete your registration! Don't forget to check the <b>Promotions box</b>, <b>Garbage</b>
          and <b>Spam</b>.
      </p>

      <AlreadyHaveAccount text="Already confirmed your email?" />
  </div>
</template>

<style lang="scss" scoped>
.email-sent {
    display: grid;
    gap: 1rem;
    justify-items: center;

    svg {
        color: white;
        font-size: 50px;
    }

    p {
        color: color(gray, shade1);
        font-weight: 300;
        line-height: 1.7;
        text-align: center;

        b {
            color: white;
        }
    }
}
</style>