import { useState } from 'react'
import Container from '@components/Container/Container'
import './chatList.scss'

const ChatList = ({notifications}) => {
    console.log(notifications)
    const [isActiveDialog, setIsActiveDialog] = useState()
    return (
        <Container overflowY verticalHeight borderRight>
            {notifications && notifications.map((notification) =>
            <div className="chat-list--dialog">
                <div className="chat-list--dialog-name">{notification.Name}</div>
                {notification.isRead && <div className="chat-list--dialog-notification">!</div>}
            </div>)}
        </Container>
    )
}

export default ChatList