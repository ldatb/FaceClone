import { NuxtAxiosInstance } from '@nuxtjs/axios'
import { NuxtCookies } from 'cookie-universal-nuxt'
import { Auth } from '@nuxtjs/auth-next'

/* eslint-disable import/no-mutable-exports */
let $axios: NuxtAxiosInstance
let $cookies: NuxtCookies
let $auth: Auth

export const initializeAxios = (axiosInstance: NuxtAxiosInstance) => {
  $axios = axiosInstance
}

export const initializeCookies = (cookiesInstance: NuxtCookies) => {
  $cookies = cookiesInstance
}


export const initializeAuth = (authInstance: Auth) => {
  $auth = authInstance
}

export { $axios, $cookies, $auth }