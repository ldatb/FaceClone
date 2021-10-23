import { Plugin } from '@nuxt/types'

const axiosPlugin: Plugin = ({ app, redirect }) => {
    app.$axios.onRequest((config) => {
        const token = app.$cookies.get('token')

        if (token) {
            config.headers.Access_token = token
        }
    })

    app.$axios.onError((error) => {
        if (error.response?.data.error === "invalid token") {
            app.$cookies.remove('token')

            return redirect('/login')
        }
    })
}

export default axiosPlugin