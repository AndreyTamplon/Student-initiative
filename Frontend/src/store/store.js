import { createStore } from 'vuex'

const store = createStore({
    state () {
        return {

        }
    },
    getters: {
        isLoggedIn() {
            return localStorage.getItem('name') !== null
        },
        getName() {
            return localStorage.getItem('name')
        },
        getEmail(){
            return localStorage.getItem('email')
        },
        getToken(){
            return localStorage.getItem('token')
        },
    },
    mutations: {
    },
})

export default store;