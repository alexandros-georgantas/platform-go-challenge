import { useState } from 'react'

import { useQuery } from 'react-query'
import { useLocation } from 'react-router-dom'
import { App } from 'antd'

import Dashboard from '../ui/Dashboard/Dashboard'

import { getAssets } from '../utils/network'

const AudiencePage = () => {
  const { notification } = App.useApp()
  const location = useLocation()
  const { pathname } = location
  const assetType = pathname.replace(/^\/+/g, '')

  const [paginationSpecs, setPaginationSpecs] = useState({
    page: 1,
    pageSize: 10,
  })

  const showNotification = (eMsg) => {
    notification.error({
      message: 'Oops, something went wrong',
      description: eMsg,
      placement: 'topRight',
    })
  }

  const onPaginationChange = (page, pageSize) => {
    setPaginationSpecs({ page, pageSize })
  }

  const { isLoading, isError, data } = useQuery({
    queryKey: [assetType, paginationSpecs.page, paginationSpecs.pageSize],
    queryFn: () =>
      getAssets(
        paginationSpecs.page,
        paginationSpecs.pageSize,
        assetType !== 'assets' ? assetType : null
      ),
  })

  if (isError) {
    showNotification()
  }

  return (
    <Dashboard
      items={data?.[assetType]}
      isLoading={isLoading}
      paginationHandler={onPaginationChange}
      pageSize={paginationSpecs.pageSize}
    />
  )
}

export default AudiencePage
