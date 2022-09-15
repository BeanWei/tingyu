import type { FormInstance, TableColumnProps } from '@arco-design/web-react'
import { Button, Card, Form, Input, InputNumber, Space, Table, Typography } from '@arco-design/web-react'
import { IconPlus } from '@arco-design/web-react/icon'
import { useRef, useState } from 'react'
import { useRequest } from 'ahooks'
import { request, url } from '~/api'
import { useTable } from '~/hooks'

const { useForm } = Form

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
        <Form ref={formRef} initialValues={props.initialValues || { rank: 9999 }}>
          <Space direction="vertical">
            <Form.Item field="title" label="话题标题" rules={[{ required: true }]}>
              <Input />
            </Form.Item>
            <Form.Item field="rank" label="推荐值">
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
    },
    {
      title: '话题标题',
      dataIndex: 'title',
    },
    {
      title: '图标',
      dataIndex: 'icon',
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
      title="话题"
      extra={(
        <Space>
          <Form form={form}>
            <Form.Item field="search" noStyle>
              <Input.Search placeholder="搜索" onSearch={search.submit}/>
            </Form.Item>
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
