import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

let RouteList = [
  {
    path: '/',
    component: resolve => require(['~/views/Layout/App.vue'], resolve),
    meta: {
      title: '首页',
      keepAlive: false,
    },
    children: [
      {
        path: '/',
        name: 'Dashboard',
        meta: {
          title: '首页',
          keepAlive: false
        },
        component: resolve => require(['~/views/Home/Index.vue'], resolve),
      },
      {
        path: '/deploy_repos',
        name: 'DeployRepo',
        meta: {
          title: '部署仓库列表',
          keepAlive: false
        },
        component: resolve => require(['~/views/DeployRepo/index.vue'], resolve),
      },
      {
        path: '/deploy_repos/:repo_id/servers',
        name: 'DeployServer',
        meta: {
          title: '部署分支列表',
          keepAlive: false
        },
        component: resolve => require(['~/views/DeployServer/index.vue'], resolve),
        props: true,
      },
      {
        path: '/deploy_servers/:server_id/records',
        name: 'DeployRecord',
        meta: {
          title: '部署记录列表',
          keepAlive: false
        },
        component: resolve => require(['~/views/DeployRecord/index.vue'], resolve),
        props: true,
      },
    ]
  },
  {
    path: '/login',
    name: 'Login',
    meta: {
      title: '后台登录',
      keepAlive: false
    },
    components: {
      blank: resolve => require(['~/views/Login/Login.vue'], resolve),
    }
  },
]

export default new Router({routes: RouteList})
