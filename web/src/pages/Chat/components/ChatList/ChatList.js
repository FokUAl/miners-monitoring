import { useState } from 'react';
import Container from '@components/Container/Container';
import './chatList.scss';

const ChatList = ({ notifications, setDialog }) => {
	console.log('chatlist', notifications);
	const [isActiveDialog, setIsActiveDialog] = useState();

	const handleActive = ({ index }) => {
		setIsActiveDialog(index);
	};
	return (
		<Container overflowY verticalHeight borderRight>
			{notifications &&
				notifications.map((notification, index) => (
					<div
                        className={`chat-list--dialog${index === isActiveDialog ? 'active': ''}`}
						key={index}
						onClick={() => {
							setDialog(notification.Name);
							handleActive(index);
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
