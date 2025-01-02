// https://nuxt.com/docs/api/configuration/nuxt-config

import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
    compatibilityDate: '2024-11-01',
    devtools: { enabled: true },
    modules: ['@nuxtjs/apollo'],
    apollo: {
        clients: {
            default: {
                httpEndpoint: 'http://localhost:8080/graphql'
            }
        },
    },
})
