import { Middleware } from "@nuxt/types";

const guest: Middleware = ({ app, redirect }) => {
    const token = app.$cookies.get('access_token')

    if (token) {
        return redirect('/')
    }
}

export default guest