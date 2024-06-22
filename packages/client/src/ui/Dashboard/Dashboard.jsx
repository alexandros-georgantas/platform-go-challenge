import styled from 'styled-components'
import List from './List'

const StyledList = styled(List)`
  height: 100%;
`
const Dashboard = ({ items, isLoading, pageSize, paginationHandler }) => (
  <StyledList
    items={items}
    isLoading={isLoading}
    paginationHandler={paginationHandler}
    pageSize={pageSize}
  />
)

export default Dashboard
