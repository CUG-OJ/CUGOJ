import {createRouter, createWebHashHistory, createWebHistory,} from 'vue-router'

const routes = [
    { path: '/', component:()=>import('../components/Home.vue') },
    {path:'/admin', component:()=>import('@/components/Admin/Admin.vue')},
    { path: '/user', component: ()=>import('@/components/User/AppUser.vue'),children:[
        { path: 'profile', component: ()=>import('@/components/User/profile.vue')},
        {path: 'security', component: ()=>import('@/components/User/security.vue')},
        {path: 'account', component: ()=>import('@/components/User/account.vue')},
        {path: 'message', component: ()=>import('@/components/User/message.vue')},
    ]},
    { path: '/group/:id', component: ()=>import('@/components/Group/GroupItem.vue') },
    { path: '/contests', component: ()=>import('../components/Contest/Contests.vue')},
    { path: '/group', component: ()=>import('@/components/Group/Group.vue')},
    { path: '/help', component: ()=>import('../components/Help.vue')},
    { path: '/problem', component: ()=>import('@/components/Problem/Problem.vue')},
    { path: '/rank', component: ()=>import('../components/Rank.vue')},
    { path: '/status', component: ()=>import('../components/Status.vue')},
    { path: '/:pathMatch(.*)*', component: ()=>import('../components/NotFound.vue') },
    { path: '/test', component: ()=>import('@/components/Problem/ProblemItem.vue')},
]
const router = createRouter(
  {
    history: createWebHistory(),
    routes,
  }
)

export default router
