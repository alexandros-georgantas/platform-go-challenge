import { useState } from 'react'

import { useQuery } from 'react-query'
import { useLocation } from 'react-router-dom'
import { App } from 'antd'

import Dashboard from '../ui/Dashboard/Dashboard'
import { useCurrentUser } from '../hooks/useCurrentUser'
import { getUserFavorites } from '../utils/network'

const FavoritesPagePage = () => {
  const { notification } = App.useApp()

  const { currentUser } = useCurrentUser()

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
    queryKey: [
      'favorites',
      currentUser?.id,
      paginationSpecs.page,
      paginationSpecs.pageSize,
    ],
    queryFn: () =>
      getUserFavorites(
        currentUser?.id,
        paginationSpecs.page,
        paginationSpecs.pageSize
      ),
    enabled: !!currentUser,
  })

  if (isError) {
    showNotification()
  }

  return (
    <Dashboard
      items={data?.favorites}
      isLoading={isLoading}
      paginationHandler={onPaginationChange}
      pageSize={paginationSpecs.pageSize}
    />
  )
}

export default FavoritesPagePage
