import {createRouter, createWebHistory,} from 'vue-router'

const routes = [
    { path: '/', component:()=>import('@/components/home.vue') },
    {path:'/admin', component:()=>import('@/components/admin.vue'),children:[
        { path: 'new', component: ()=>import('@/components/adminNew.vue')},
        { path: 'createProblem', component: ()=>import('@/components/adminCreateProblem.vue')},
        { path: 'user', component: ()=>import('@/components/adminUser.vue')},
        { path: 'dashboard', component: ()=>import('@/components/adminDashboard.vue')},
        { path: 'problem', component: ()=>import('@/components/adminProblem.vue')},
    ]},
    { path: '/contests', component: ()=>import('@/components/contest.vue')},
    { path: '/class', component: ()=>import('@/components/class.vue')},
    { path: '/more', component: ()=>import('@/components/more.vue')},
    { path: '/login', component: ()=>import('@/components/login.vue')},
    { path: '/register', component: ()=>import('@/components/register.vue')},
    { path: '/problem', component: ()=>import('@/components/problem.vue')},
    { path: '/problemItem', component: ()=>import('@/components/problemItem.vue')},
    { path: '/rank', component: ()=>import('@/components/rank.vue')},
    { path: '/status', component: ()=>import('@/components/status.vue')},
    // { path: '/user', component: ()=>import('@/components/User/AppUser.vue'),children:[
    //     { path: 'profile', component: ()=>import('@/components/User/profile.vue')},
    //     {path: 'security', component: ()=>import('@/components/User/security.vue')},
    //     {path: 'account', component: ()=>import('@/components/User/account.vue')},
    //     {path: 'message', component: ()=>import('@/components/User/message.vue')},
    // ]},
    // { path: '/group/:id', component: ()=>import('@/components/Group/GroupItem.vue') },

    // { path: '/:pathMatch(.*)*', component: ()=>import('@/components/NotFound.vue') },
    // { path: '/test', component: ()=>import('@/components/Problem/ProblemItem.vue')},
]
const router = createRouter(
  {
    history: createWebHistory(),
    routes,
  }
)
router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')
    if (to.path === '/login' || to.path === '/register') {
        token?next('/'):next()
    }else{
        next()
    }
})
export default router
