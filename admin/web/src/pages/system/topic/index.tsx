import type { BadgeProps, FormInstance, SelectOptionProps, TableColumnProps } from '@arco-design/web-react'
import { Badge, Button, Card, Form, Input, InputNumber, Select, Space, Table, Typography } from '@arco-design/web-react'
import { IconPlus } from '@arco-design/web-react/icon'
import { useRef, useState } from 'react'
import { useRequest } from 'ahooks'
import { request, url } from '~/api'
import { useTable } from '~/hooks'
import RecordSelect from '~/components/RecordSelect'

const { useForm } = Form

const StatusOptions: SelectOptionProps[] = [
  {
    title: '上线',
    value: 1,
  },
  {
    title: '下线',
    value: 2,
  },
  {
    title: '待审核',
    value: 3,
  },
]

function TopicForm(props: {
  initialValues: AnyObject
  onBack: (reload?: boolean) => void
}) {
  const formRef = useRef<FormInstance>()
  const { loading, run } = useRequest((data: AnyObject) => {
    return request({
      url: url.createTopic,
      data,
    })
  }, {
    manual: true,
    onSuccess: () => {
      props.onBack(true)
    },
  })

  function handleSubmit() {
    formRef.current.validate().then((values) => {
      run(values)
    })
  }

  function handleRest() {
    formRef.current.resetFields()
  }

  return (
    <Card
      title={props.initialValues.id ? '更新话题' : '新建话题'}
      extra={
        <Space>
          <Button onClick={() => props.onBack()}>返回</Button>
          <Button onClick={handleRest}>重置</Button>
          <Button type="primary" onClick={handleSubmit} loading={loading}>提交</Button>
        </Space>
      }
      headerStyle={{
        paddingTop: 20,
      }}
    >
      <div style={{ maxWidth: 650, margin: '0 auto' }}>
        <Form ref={formRef} initialValues={props.initialValues || { rec_rank: 9999, status: 1 }}>
          <Space direction="vertical">
            <Form.Item field="topic_category_id" label="分类" rules={[{ required: true }]}>
              <RecordSelect
                service={url.listCategory}
                fieldNames={{
                  title: 'name',
                }}
              />
            </Form.Item>
            <Form.Item field="title" label="话题标题" rules={[{ required: true }]}>
              <Input />
            </Form.Item>
            <Form.Item field="icon" label="图标">
              <Input />
            </Form.Item>
            <Form.Item field="description" label="描述">
              <Input.TextArea />
            </Form.Item>
            <Form.Item field="status" label="状态">
              <Select>
                {StatusOptions.map((option, idx) => (
                  <Select.Option key={idx} value={option.value}>
                    {option.title}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
            <Form.Item field="rec_rank" label="推荐值">
              <InputNumber />
            </Form.Item>
          </Space>
        </Form>
      </div>
    </Card>
  )
}

function TopicList() {
  const [form] = useForm()
  const { tableProps, search } = useTable(url.listTopic, {
    form,
  })
  const [curRecord, setCurRecord] = useState<AnyObject>()

  if (curRecord) {
    return (
      <TopicForm
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
      render: value => <Typography.Text copyable>{value}</Typography.Text>,
    },
    {
      title: '分类',
      dataIndex: 'topic_category_id',
      render: (_col, record) => <Typography.Text>{record.edges?.topic_category?.name || '-'}</Typography.Text>,
    },
    {
      title: '话题标题',
      dataIndex: 'title',
    },
    {
      title: '状态',
      dataIndex: 'status',
      render: (value: number) => {
        const status: BadgeProps['status'][] = ['success', 'default', 'warning']
        return <Badge status={status[value - 1]} text={StatusOptions[value - 1].title}></Badge>
      },
    },
    {
      title: '帖子数量',
      dataIndex: 'post_count',
    },
    {
      title: '关注数量',
      dataIndex: 'follower_count',
    },
    {
      title: '参与者数量',
      dataIndex: 'attender_count',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      headerCellStyle: { paddingLeft: '15px' },
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
      title="话题列表"
      extra={(
        <Space>
          <Form form={form} layout="inline">
            <Space>
              <Form.Item field="status" noStyle>
                <Select placeholder="话题状态" onChange={search.submit} style={{ width: 100 }}>
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
          <Button
            type="primary"
            icon={<IconPlus />}
            onClick={() => {
              setCurRecord({})
            }}
          >
            新建
          </Button>
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

export default TopicList
