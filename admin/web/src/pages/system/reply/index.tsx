import type { BadgeProps, SelectOptionProps, TableColumnProps } from '@arco-design/web-react'
import { Badge, Card, Form, Input, Select, Space, Table, Typography } from '@arco-design/web-react'
import Ellipsis from 'react-ellipsis-component'
import { url } from '~/api'
import { useTable } from '~/hooks'
import { extractText } from '~/utils/lexical'

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

function ReplyList() {
  const [form] = useForm()
  const { tableProps, search } = useTable(url.listReply, {
    form,
  })

  const columns: TableColumnProps[] = [
    {
      title: '编号',
      dataIndex: 'id',
      width: 180,
      render: value => <Typography.Text copyable>{value}</Typography.Text>,
    },
    {
      title: 'IP归属地',
      dataIndex: 'ip_loc',
      width: 180,
    },
    {
      title: '审核状态',
      dataIndex: 'status',
      width: 180,
      render: (value: number) => {
        const status: BadgeProps['status'][] = ['success', 'default', 'warning']
        return <Badge status={status[value - 1]} text={StatusOptions[value - 1].title}></Badge>
      },
    },
    {
      title: '回复内容',
      dataIndex: 'content',
      render: value => (
        <Ellipsis text={extractText(JSON.parse(value).root)} maxLine={1} ellipsis />
      ),
    },
  ]

  return (
    <Card
      title="回复列表"
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

export default ReplyList
