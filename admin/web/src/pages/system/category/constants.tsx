import type { TableColumnProps } from '@arco-design/web-react'
import { Button, Typography } from '@arco-design/web-react'

export const columns: TableColumnProps[] = [
  {
    title: '编号',
    dataIndex: 'id',
    render: value => <Typography.Text copyable>{value}</Typography.Text>,
  },
  {
    title: '分类名称',
    dataIndex: 'name',
  },
  {
    title: '排序值',
    dataIndex: 'rank',
    sorter: (a, b) => a.count - b.count,
    render(x) {
      return Number(x).toLocaleString()
    },
  },
  {
    title: '操作',
    dataIndex: 'operations',
    headerCellStyle: { paddingLeft: '15px' },
    render: (_, _record) => (
      <Button
        type="text"
        size="small"
      >
        查看
      </Button>
    ),
  },
]
