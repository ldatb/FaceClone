import { Middleware } from "@nuxt/types";

const auth: Middleware = ({ app, redirect }) => {
    const token = app.$cookies.get('access_token')

    if (!token) {
        return redirect('/login')
    }
}

export default auth