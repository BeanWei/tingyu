import type { FormInstance } from '@arco-design/web-react'
import { Button, Card, Form, Input, InputNumber, Space, Table } from '@arco-design/web-react'
import { IconPlus } from '@arco-design/web-react/icon'
import { useRef, useState } from 'react'
import { columns } from './constants'
import { url } from '~/api'
import { useTable } from '~/hooks'

const { useForm } = Form

function CategoryForm(props: {
  initialValues: AnyObject
  onBack: () => void
}) {
  const formRef = useRef<FormInstance>()

  return (
    <Card
      title={props.initialValues.id ? '更新分类' : '新建分类'}
      extra={
        <Space>
          <Button onClick={props.onBack}>返回</Button>
          <Button>重置</Button>
          <Button type="primary">提交</Button>
        </Space>
      }
      headerStyle={{
        paddingTop: 20,
      }}
    >
      <div style={{ maxWidth: 650, margin: '0 auto' }}>
        <Form ref={formRef}>
          <Space direction="vertical">
            <Form.Item field="name" label="分类名称">
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
        onBack={() => setCurRecord(undefined)}
      />
    )
  }

  return (
    <Card
      title="话题分类"
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
