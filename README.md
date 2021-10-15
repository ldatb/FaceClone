# FaceClone
This is not a clone per-se, but a simple API library and a web page based on Facebook.

## 1. API (Back-End)
The API for this project was done in [Go - Fiber](https://docs.gofiber.io/), a Golang library for API building.

### 1.1 API Installation
1. Download and install Go at https://golang.org/ (The version used in this project is 1.17.1, you can use anyone above 1.14)
    1.1 Alternatively, you can use `docker-compose up`
2. Install the dependencies using `go get -u -v ./...`
3. Create a [MailTrap](https://mailtrap.io/) account, which is the email service provider.
4. Create your .env file, use .env.example to guide you.
5. Run with `go run main.go`

## 2 Webpage (Front-end)
The webpage for this project was done using [Vue-Nuxt](https://nuxtjs.org/) and the [Atomic Design principles by Brad Frost](https://atomicdesign.bradfrost.com/)

### 1.2 Webpage installation
For specific instructions read the README.md file inside the Webpage folder
1. Install dependencies using `npm install`
2. Serve with host at port 3000 with `npm run dev`
3. Run with `npm run build` and `npm run start`
4. You can also generate a static page (not recommended since it was made with the intention of being a server-side dynamic webpage) with `npm run generate`
