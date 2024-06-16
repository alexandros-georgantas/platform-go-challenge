import React from 'react'
import { App, Button, Space, Layout, Menu, theme } from 'antd'
const { Header, Footer, Content } = Layout
const layoutStyle = {
  overflow: 'hidden',
  width: '100%',
  height: '100%',
}

const MyPage = () => {
  const { message, modal, notification } = App.useApp()

  const showMessage = () => {
    message.success('Success!')
  }

  const showModal = () => {
    modal.warning({
      title: 'This is a warning message',
      content: 'some messages...some messages...',
    })
  }

  const showNotification = () => {
    notification.info({
      message: `Notification topLeft`,
      description: 'Hello, Ant Design!!',
      placement: 'topLeft',
    })
  }

  return (
    <Space wrap>
      <Button type="primary" onClick={showMessage}>
        Open message
      </Button>
      <Button type="primary" onClick={showModal}>
        Open modal
      </Button>
      <Button type="primary" onClick={showNotification}>
        Open notification
      </Button>
    </Space>
  )
}
const items = new Array(3).fill(null).map((_, index) => ({
  key: String(index + 1),
  label: `nav ${index + 1}`,
}))
const MyApp: React.FC = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken()

  return (
    <App style={{ height: '100vh' }}>
      <Layout style={layoutStyle}>
        <Header style={{ display: 'flex', alignItems: 'center' }}>
          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={['1']}
            items={items}
            style={{ flex: 1, minWidth: 0 }}
          />
        </Header>
        <Content style={{ padding: '0 48px', height: 'calc(100vh - 64px)' }}>
          <div
            style={{
              background: colorBgContainer,
              minHeight: 280,
              padding: 24,
              borderRadius: borderRadiusLG,
            }}
          >
            <MyPage />
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          GWI ©{new Date().getFullYear()} Created by Alexandros Georgantas
        </Footer>
      </Layout>
    </App>
  )
}

export default MyApp
