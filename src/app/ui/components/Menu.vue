<template>
  <nav class="navbar is-black" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <nuxt-link class="navbar-item" to="/" data-cy="home-link">
        <strong>govueapp</strong>
      </nuxt-link>

      <button class="button navbar-burger">
        <span />
        <span />
        <span />
      </button>
    </div>

    <div class="navbar-end">
      <div class="navbar-item has-dropdown is-hoverable">
        <a class="navbar-link">
          Menu
        </a>

        <div class="navbar-dropdown is-right">
          <nuxt-link class="navbar-item" to="/login">
            Login
          </nuxt-link>
          <nuxt-link class="navbar-item" to="/about">
            About
          </nuxt-link>
          <hr class="navbar-divider" />
          <a v-if="isAuthenticated" class="dropdown-item" @click="logout"
            >Logout</a
          >
          <div class="navbar-item">
            v0.5.0
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
import { mapGetters } from 'vuex'
const Cookie = process.client ? require('js-cookie') : undefined

export default {
  name: 'Menu',
  computed: {
    ...mapGetters(['isAuthenticated'])
  },
  methods: {
    logout() {
      Cookie.remove('auth')
      this.$store.commit('setAuth', null)
      this.$router.push('/login')
    }
  }
}
</script>
