import { Plugin } from '@nuxt/types'
import { initializeAxios, initializeCookies, initializeAuth } from '@/utils/nuxt-instance'

const accessor: Plugin = ({ app }) => {
  initializeAxios(app.$axios)
  initializeCookies(app.$cookies)
  initializeAuth(app.$auth)
}

export default accessor
