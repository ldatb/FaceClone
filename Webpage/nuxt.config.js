export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'FaceClone',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ['normalize.css/normalize.css', '@/assets/scss/base.scss'],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: ['@/plugins/accessor', '@/plugins/notifications.client'],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: [{ path: '@/components/', pathPrefix: false }],

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    '@nuxtjs/style-resources',
    '@nuxtjs/fontawesome'
  ],
  
  styleResources: {
    scss: ['@/components/bosons/*.scss']
  },
  fontawesome: {
    component: 'fa',
    icons: {
      solid: true,
      brands: true
    }
  },

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/auth-next'
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    baseURL: 'http://localhost:3000',
  },

  // Auth module: https://auth.nuxtjs.org/
  auth: {
    strategies: {
      local: {
        token: {
          property: 'access_token',
          global: true,
          name: 'access_token',
          type: '',
          maxAge: 60 * 60 * 24 * 30 // 30 days
        },
        endpoints: {
          login: { url: '/users/jwtlogin', method: 'post' },
          logout: { url: '/users/logout', method: 'delete' },
          user: { url: '/users/jwtuser', method: 'get' }
        }
      }
    }
  },
/*
  // This checks if there is a token set, and redirects the user either to the homepage or the login page
  router: {
    middleware: ['auth']
  },
*/
  server: {
    port: 8000,
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },
}
