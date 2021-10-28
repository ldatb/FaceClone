import { Plugin } from '@nuxt/types'

const axiosPlugin: Plugin = ({ app, redirect }) => {
    app.$axios.onRequest((config) => {
        const token = app.$cookies.get('access_token')
        if (token) {
            config.headers.access_token = token
        }
    })

    app.$axios.onError((error) => {
        if (error.response?.status === 401) {
            return redirect('/login')
        }
    })
}

export default axiosPlugin