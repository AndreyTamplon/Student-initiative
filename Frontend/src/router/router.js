import {createRouter, createWebHistory} from 'vue-router'
import AboutPage from "@/pages/AboutPage";
import Home from "@/pages/Home";
import SignUpPage from "@/pages/SignUpPage";
import SignInPage from "@/pages/SignInPage";
import CreatePetitionPage from "@/pages/CreatePetitionPage";
import store from "@/store/store";
import PetitionPage from "@/pages/PetitionPage";

const routes = [
    {
        path: '/',
        component: Home,
        name: 'home'
    },
    {
        path: '/about',
        component: AboutPage,
        name: 'about'
    },
    {
        path: '/sign-up',
        component: SignUpPage,
        name: 'sign-up'
    },
    {
        path: '/sign-in',
        component: SignInPage,
        name: 'sign-in'
    },
    {
        path: '/create-petition',
        component: CreatePetitionPage,
        name: 'create-petition',
        beforeEnter: (to, from, next) => {
            if(store.getters.isLoggedIn){
                next()
            } else {
                next({
                    name: "sign-in"
                });
            }
        }
    },
    {
        path: '/petition/:id',
        component: PetitionPage,
        name: 'petition'
    },
]


const history = createWebHistory();

const router = createRouter({
    history,
    routes,
})

const DEFAULT_TITLE = 'Студенческая инициатива';
router.beforeEach((to)=> {
    document.title = to.meta.title || DEFAULT_TITLE;
})

export default router