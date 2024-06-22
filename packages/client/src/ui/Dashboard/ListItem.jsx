import React from 'react'
import { Avatar, List, Space } from 'antd'
import { StarOutlined, AreaChartOutlined } from '@ant-design/icons'

const IconText = ({ icon, text }) => (
  <Space>
    {React.createElement(icon)}
    {text}
  </Space>
)
const ListItem = ({ item }) => {
  return (
    <List.Item
      key={item.ID}
      actions={[
        <IconText
          icon={StarOutlined}
          text="Favorite"
          key="list-vertical-star-o"
        />,
      ]}
    >
      <List.Item.Meta
        avatar={<Avatar size="large" icon={<AreaChartOutlined />} />}
        title={item.Description}
        description={item.Description}
      />
      {item.RelatedType}
    </List.Item>
  )
}

export default ListItem
