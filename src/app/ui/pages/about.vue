<template>
  <div>
    <section class="section">
      <div class="container">
        <h1 class="title">
          {{ title }}
        </h1>
        <h2 class="subtitle">
          This shows you how to build a website using
          <strong>Nuxt</strong>, <strong>Vue</strong>, and
          <strong>Bulma</strong>.
        </h2>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: 'About'
    }
  },
  head() {
    return {
      title: this.title
    }
  }
}
</script>
