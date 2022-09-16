import type { SelectProps } from '@arco-design/web-react'
import { Avatar, List, Select, Spin } from '@arco-design/web-react'
import { useRequest } from 'ahooks'
import { request } from '~/api'

type RecordSelectProps = Omit<SelectProps, 'mode'> & {
  service: string
  multiple?: boolean
  fieldNames?: {
    title?: string
    value?: string
    avatar?: string
    description?: string[]
  }
}

function RecordSelect(props: RecordSelectProps) {
  const {
    service,
    multiple,
    fieldNames: fieldNames_,
    ...rest
  } = props
  const fieldNames = { title: 'id', value: 'id', ...fieldNames_ }

  const { data, loading } = useRequest(() => {
    return request({
      url: service,
    })
  })
  const options = data?.data?.data || []

  return (
    <Select
      {...rest}
      mode={multiple ? 'multiple' : undefined}
      allowClear
      notFoundContent={
        loading
          ? (
          <div
            style={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
            }}
          >
            <Spin style={{ margin: 12 }} />
          </div>
            )
          : null
      }
      renderFormat={(option) => {
        return options.find(v => v[fieldNames.value] === option?.value)?.[
          fieldNames.title
        ]
      }}
      options={options.map((item: any) => {
        return {
          label: (
            <List.Item key={item[fieldNames.value]} style={{ lineHeight: 1.2 }}>
              <List.Item.Meta
                avatar={
                  fieldNames.avatar && item[fieldNames.avatar]
                    ? (
                    <Avatar shape="square">
                      <img
                        src={item[fieldNames.avatar]}
                        alt={item[fieldNames.title]}
                      />
                    </Avatar>
                      )
                    : null
                }
                title={item[fieldNames.title]}
                description={
                  fieldNames.description
                    ? fieldNames.description
                      .map(field => item[field])
                      .filter(v => v !== undefined && v !== null)
                      .join(', ')
                    : null
                }
              />
            </List.Item>
          ),
          value: item[fieldNames.value],
        }
      })}
    />
  )
}

export default RecordSelect
