import type { FormInstance, PaginationProps, TableProps } from '@arco-design/web-react'
import { useMemoizedFn, useRequest, useUpdateEffect } from 'ahooks'
import { useState } from 'react'
import { request } from '~/api'

export interface Params {
  page?: number
  limit?: number
  sorter?: any
  filter?: any
  [key: string]: any
}

export interface TableOptions {
  form?: FormInstance
  pagination?: PaginationProps
  [key: string]: any
}

const useTable = (service: string, options: TableOptions = {}): {
  tableProps: TableProps
  search: {
    submit: () => void
    reset: () => void
  }
} => {
  const [formParams, setFormParams] = useState<AnyObject>()
  const [pagination, setPatination] = useState<PaginationProps>({
    pageSize: options?.pagination?.pageSize || 20,
    current: 1,
  })
  const { data, loading, run } = useRequest((params: Params) => {
    return request({
      url: service,
      params,
    })
  }, {
    onSuccess: (result) => {
      setPatination({
        ...pagination,
        total: result?.data?.total,
      })
    },
  })

  useUpdateEffect(() => {
    run({
      page: pagination.current,
      limit: pagination.pageSize,
      filter: formParams,
    })
  }, [pagination.current, pagination.pageSize, formParams])

  return {
    tableProps: {
      rowKey: 'id',
      loading,
      data: data?.data?.data,
      pagination: {
        sizeCanChange: true,
        showTotal: true,
        pageSize: 10,
        current: 1,
        pageSizeChangeResetCurrent: true,
        ...options.pagination,
        ...pagination,
      },
      // TODO: 支持排序和自定义筛选
      onChange: ({ current, pageSize }) => {
        setPatination({
          current,
          pageSize,
        })
      },
    },
    search: {
      submit: useMemoizedFn(() => {
        if (!options.form)
          return
        const values = options.form.getFieldsValue()
        setFormParams(values)
        setPatination({
          ...pagination,
          current: 1,
        })
      }),
      reset: useMemoizedFn(() => {
        if (!options.form)
          return
        options.form.resetFields()
        setFormParams(undefined)
        setPatination({
          ...pagination,
          current: 1,
        })
      }),
    },
  }
}

export default useTable
