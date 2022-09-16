import type { BadgeProps, SelectOptionProps, TableColumnProps } from '@arco-design/web-react'
import { Badge, Card, Form, Input, Select, Space, Table, Typography } from '@arco-design/web-react'
import { url } from '~/api'
import { useTable } from '~/hooks'

const { useForm } = Form

const StatusOptions: SelectOptionProps[] = [
  {
    title: '正常',
    value: 1,
  },
  {
    title: '停用',
    value: 2,
  },
]

function UserList() {
  const [form] = useForm()
  const { tableProps, search } = useTable(url.listUser, {
    form,
  })

  const columns: TableColumnProps[] = [
    {
      title: '编号',
      dataIndex: 'id',
      render: value => <Typography.Text copyable>{value}</Typography.Text>,
    },
    {
      title: '用户名',
      dataIndex: 'username',
    },
    {
      title: '昵称',
      dataIndex: 'nickname',
    },
    {
      title: '状态',
      dataIndex: 'status',
      render: (value: number) => {
        const status: BadgeProps['status'][] = ['success', 'default']
        return <Badge status={status[value - 1]} text={StatusOptions[value - 1].title}></Badge>
      },
    },
    {
      title: '头像',
      dataIndex: 'avatar',
    },
    {
      title: '是否管理员',
      dataIndex: 'is_admin',
      render: (value: boolean) => {
        return value ? '是' : '否'
      },
    },
  ]

  return (
    <Card
      title="用户列表"
      extra={(
        <Space>
          <Form form={form} layout="inline">
            <Space>
              <Form.Item field="status" noStyle>
                <Select placeholder="用户状态" onChange={search.submit} style={{ width: 100 }}>
                  {StatusOptions.map((option, idx) => (
                    <Select.Option key={idx} value={option.value}>
                      {option.title}
                    </Select.Option>
                  ))}
                </Select>
              </Form.Item>
              <Form.Item field="is_admin" noStyle>
                <Select placeholder="管理员" onChange={search.submit} style={{ width: 100 }}>
                  <Select.Option key={1} value={1}>
                    是
                  </Select.Option>
                  <Select.Option key={2} value={0}>
                    否
                  </Select.Option>
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

export default UserList
