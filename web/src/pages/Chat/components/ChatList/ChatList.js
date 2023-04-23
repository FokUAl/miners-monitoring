import { useState } from 'react';
import Container from '@components/Container/Container';
import './chatList.scss';

const ChatList = ({ notifications, setDialog, sender }) => {
	console.log('chatlist', notifications);

	return (
		<Container overflowY verticalHeight borderRight>
			{notifications &&
				notifications.map((notification, index) => (
					<div
                        className={`chat-list--dialog${notification.Name === sender ? ' active': ''}`}
						key={index}
						onClick={() => {
							setDialog(notification.Name);
						}}
					>
						<div className="chat-list--dialog-name">{notification.Name}</div>
						{notification.isRead && (
							<div className="chat-list--dialog-notification">!</div>
						)}
					</div>
				))}
		</Container>
	);
};

export default ChatList;
