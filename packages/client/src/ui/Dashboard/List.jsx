import { List as AntList } from 'antd'
import styled from 'styled-components'

import ListItem from './ListItem'

const StyledList = styled(AntList)`
  height: 100%;
  > div {
    height: calc(100% - 120px);
    overflow-y: auto;
  }
`

const List = ({ items, isLoading, paginationHandler, pageSize }) => {
  return (
    <StyledList
      itemLayout="vertical"
      size="large"
      loading={isLoading}
      pagination={{
        onChange: paginationHandler,
        defaultCurrent: 1,
        total: items?.totalCount,
        pageSize: pageSize,
        align: 'center',
      }}
      dataSource={items?.items}
      renderItem={(item) => <ListItem item={item} />}
    />
  )
}

export default List
