import type { FormInstance, TableColumnProps } from '@arco-design/web-react'
import { Button, Card, Form, Input, InputNumber, Space, Table, Typography } from '@arco-design/web-react'
import { IconPlus } from '@arco-design/web-react/icon'
import { useRef, useState } from 'react'
import { useRequest } from 'ahooks'
import { request, url } from '~/api'
import { useTable } from '~/hooks'

const { useForm } = Form

function CategoryForm(props: {
  initialValues: AnyObject
  onBack: (reload?: boolean) => void
}) {
  const formRef = useRef<FormInstance>()
  const { loading, run } = useRequest((data: AnyObject) => {
    return request({
      url: url.createCategory,
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
      title={props.initialValues.id ? '更新分类' : '新建分类'}
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
            <Form.Item field="name" label="分类名称" rules={[{ required: true }]}>
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

function CategoryList() {
  const [form] = useForm()
  const { tableProps, search } = useTable(url.listCategory, {
    form,
  })
  const [curRecord, setCurRecord] = useState<AnyObject>()

  if (curRecord) {
    return (
      <CategoryForm
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
      title="分类列表"
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

export default CategoryList
