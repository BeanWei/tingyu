export interface IRoute {
  name: string
  key: string
  // 当前页是否展示面包屑
  breadcrumb?: boolean
  children?: IRoute[]
  // 当前路由是否渲染菜单项，为 true 的话不会在菜单中显示，但可通过路由地址访问。
  ignore?: boolean
}

export const routes: IRoute[] = [
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
]
