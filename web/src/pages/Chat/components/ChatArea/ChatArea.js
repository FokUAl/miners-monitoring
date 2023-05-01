import { useState } from 'react';
import Container from '@components/Container/Container';
import Messages from '@components/PopupChat/Messages';
import ComponentService from '@services/component.service';
import { BsSend } from 'react-icons/bs';
import './chatArea.scss'

const ChatArea = ({ messages, username, dialog }) => {
	const [message, setMessage] = useState();
	const receiver = dialog

	const handleSend = async (e) => {
		e.preventDefault();
		ComponentService.sendMessage(message, receiver).then(
			(response) => {
				console.log('chat send ok', message, receiver);
			},
			(error) => {
				console.log('chat send', error);
			}
		);
		setMessage('');
	};
	return (
		<Container verticalHeight>
			<div className="chat-area">
				<div className='chat-area--container'>
					{messages ? (
						<Messages messages={messages} username={username} />
					) : (
						'Open Dialog'
					)}
				</div>
                <div></div>
				<form onSubmit={handleSend} className="chat--form">
					<input
						className="chat--input"
						type="text"
						value={message}
						required
						onChange={(e) => setMessage(e.target.value)}
					/>
					<button className="chat--submit" type="submit">
						<BsSend color="black" size="15" />
					</button>
				</form>
			</div>
		</Container>
	);
};

export default ChatArea;
