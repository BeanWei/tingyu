export interface IRoute {
  name: string
  key: string
  // 当前页是否展示面包屑
  breadcrumb?: boolean
  children?: IRoute[]
  // 当前路由是否渲染菜单项，为 true 的话不会在菜单中显示，但可通过路由地址访问。
  ignore?: boolean
}

const routes: IRoute[] = [
  {
    name: '仪表盘',
    key: 'dashboard',
    children: [
      {
        name: '工作台',
        key: 'dashboard/workplace',
      },
    ],
  },
  {
    name: '系统管理',
    key: 'system',
    children: [
      {
        name: '分类管理',
        key: 'system/category',
      },
      {
        name: '话题管理',
        key: 'system/topic',
      },
      {
        name: '帖子管理',
        key: 'system/post',
      },
      {
        name: '评论管理',
        key: 'system/comment',
      },
      {
        name: '回复管理',
        key: 'system/reply',
      },
      {
        name: '用户管理',
        key: 'system/user',
      },
    ],
  },
]

export default routes
