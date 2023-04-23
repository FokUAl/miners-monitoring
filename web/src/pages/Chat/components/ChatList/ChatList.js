import { useState } from 'react';
import Container from '@components/Container/Container';
import './chatList.scss';

const ChatList = ({ notifications, setDialog }) => {
	console.log('chatlist', notifications);
	const [isActiveDialog, setIsActiveDialog] = useState();

	const handleActive = ({ index }) => {
		setIsActiveDialog(index);
	};
	const generatorClasses = ({ index }) => {
		const classes = ['chat-list--dialog'];
        console.log('generator', index, isActiveDialog)
		if (isActiveDialog === index) classes.push('active');
		return classes.join(' ');
	};
	return (
		<Container overflowY verticalHeight borderRight>
			{notifications &&
				notifications.map((notification, index) => (
					<div
						className={() => generatorClasses(index)}
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
