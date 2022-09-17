import type { BadgeProps, DescriptionsProps, SelectOptionProps, TableColumnProps } from '@arco-design/web-react'
import { Badge, Button, Card, Descriptions, Form, Input, Select, Space, Table, Typography } from '@arco-design/web-react'
import { useRequest } from 'ahooks'
import { useRef, useState } from 'react'
import Ellipsis from 'react-ellipsis-component'
import { request, url } from '~/api'
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

function PostForm(props: {
  initialValues: AnyObject
  onBack: (reload?: boolean) => void
}) {
  const isChanged = useRef<boolean>(false)
  const values = useRef<AnyObject>(props.initialValues)
  const { loading, run } = useRequest((data: AnyObject) => {
    return request({
      url: url.updatePost,
      data,
    })
  }, {
    manual: true,
    onSuccess: () => {
      isChanged.current = true
    },
  })

  const data: DescriptionsProps['data'] = [
    {
      label: '编号',
      value: values.current.id,
    },
    {
      label: '审核状态',
      value: <Badge status={['success', 'default', 'warning'][values.current.status - 1] as BadgeProps['status']} text={StatusOptions[values.current.status - 1].title}></Badge>,
    },
    {
      label: 'IP归属地',
      value: values.current.ip_loc,
    },
    {
      label: '是否置顶',
      value: values.current.is_top ? '是' : '否',
    },
    {
      label: '是否精华',
      value: values.current.is_excellent ? '是' : '否',
    },
    {
      label: '是否锁定',
      value: values.current.is_lock ? '是' : '否',
    },
    {
      label: '评论数量',
      value: values.current.comment_count,
    },
    {
      label: '最后回复时间',
      value: values.current.latest_replied_at,
    },
    {
      label: '帖子内容',
      value: extractText(JSON.parse(values.current.content).root),
      span: 3,
    },
  ]

  function handleUpdate(field: string) {
    if (field === 'status')
      values.current[field] = values.current[field] === 1 ? 2 : 1
    else
      values.current[field] = !values.current[field]
    run(values.current)
  }

  return (
    <Card
      title={'审核帖子'}
      extra={
        <Space>
          <Button onClick={() => props.onBack(isChanged.current)}>返回</Button>
          <Button type="outline" loading={loading} onClick={() => handleUpdate('is_top')}>
            {values.current.is_top ? '取消置顶' : '置顶'}
          </Button>
          <Button type="outline" loading={loading} onClick={() => handleUpdate('is_excellent')}>
            {values.current.is_excellent ? '取消加精' : '加精'}
          </Button>
          <Button type="outline" loading={loading} onClick={() => handleUpdate('is_lock')}>
            {values.current.is_lock ? '取消锁定' : '锁定'}
          </Button>
          <Button type="outline" loading={loading} onClick={() => handleUpdate('status')}>
            {values.current.status === 1 ? '不通过' : '通过'}
          </Button>
        </Space>
      }
      headerStyle={{
        paddingTop: 20,
        paddingBottom: 6,
      }}
    >
      <div style={{ maxWidth: 650, margin: '0 auto' }}>
        <Descriptions
          column={1}
          data={data}
        />
      </div>
    </Card>
  )
}

function PostList() {
  const [form] = useForm()
  const { tableProps, search } = useTable(url.listPost, {
    form,
  })
  const [curRecord, setCurRecord] = useState<AnyObject>()

  if (curRecord) {
    return (
      <PostForm
        initialValues={curRecord}
        onBack={(reload: boolean) => {
          setCurRecord(undefined)
          if (reload)
            search.reset()
        }}
      />
    )
  }

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
      title: '帖子内容',
      dataIndex: 'content',
      render: value => (
        <Ellipsis text={extractText(JSON.parse(value).root)} maxLine={1} ellipsis />
      ),
    },
    {
      title: '操作',
      dataIndex: 'operations',
      headerCellStyle: { paddingLeft: '15px' },
      width: 100,
      render: (_, record) => (
        <Button
          type="text"
          size="small"
          onClick={() => {
            setCurRecord(record)
          }}
        >
          查看
        </Button>
      ),
    },
  ]

  return (
    <Card
      title="帖子列表"
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

export default PostList
