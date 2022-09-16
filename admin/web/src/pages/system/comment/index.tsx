import type { BadgeProps, SelectOptionProps, TableColumnProps } from '@arco-design/web-react'
import { Badge, Card, Form, Input, Select, Space, Table, Typography } from '@arco-design/web-react'
import { url } from '~/api'
import { useTable } from '~/hooks'

const { useForm } = Form

const StatusOptions: SelectOptionProps[] = [
  {
    title: '通过',
    value: 1,
  },
  {
    title: '未通过',
    value: 2,
  },
  {
    title: '待审核',
    value: 3,
  },
]

function CommentList() {
  const [form] = useForm()
  const { tableProps, search } = useTable(url.listComment, {
    form,
  })

  const columns: TableColumnProps[] = [
    {
      title: '编号',
      dataIndex: 'id',
      render: value => <Typography.Text copyable>{value}</Typography.Text>,
    },
    {
      title: 'IP归属地',
      dataIndex: 'ip_loc',
    },
    {
      title: '审核状态',
      dataIndex: 'status',
      render: (value: number) => {
        const status: BadgeProps['status'][] = ['success', 'default', 'warning']
        return <Badge status={status[value - 1]} text={StatusOptions[value - 1].title}></Badge>
      },
    },
    {
      title: '评论内容',
      dataIndex: 'content',
    },
    {
      title: '回复数',
      dataIndex: 'reply_count',
    },
    // {
    //   title: '操作',
    //   dataIndex: 'operations',
    //   headerCellStyle: { paddingLeft: '15px' },
    //   render: (_, record) => (
    //     <Button
    //       type="text"
    //       size="small"
    //     >
    //       查看
    //     </Button>
    //   ),
    // },
  ]

  return (
    <Card
      title="评论列表"
      extra={(
        <Space>
          <Form form={form} layout="inline">
            <Space>
              <Form.Item field="status" noStyle>
                <Select placeholder="审核状态" onChange={search.submit} style={{ width: 100 }}>
                  {StatusOptions.map((option, idx) => (
                    <Select.Option key={idx} value={option.value}>
                      {option.title}
                    </Select.Option>
                  ))}
                </Select>
              </Form.Item>
              <Form.Item field="search" noStyle>
                <Input.Search placeholder="搜索" onSearch={search.submit}/>
              </Form.Item>
            </Space>
          </Form>
        </Space>
      )}
      headerStyle={{
        paddingTop: 20,
        paddingBottom: 6,
      }}
    >
      <Table
        {...tableProps}
        columns={columns}
      />
    </Card>
  )
}

export default CommentList
