import { useEffect, useState, useRef, useMemo } from 'react';
import Button from '@components/Button/Button';
import Input from '@components/Input/Input';
import Messages from './Messages';
import ComponentService from '@services/component.service';
import { BsSend } from 'react-icons/bs';
import './popupChat.scss';
import PageService from '../../services/page.service';

const PopupChat = ({ username }) => {
	const [isChatHidden, setIsChatHidden] = useState(true);
	const [messages, setMessages] = useState();
	const [message, setMessage] = useState();
	const chatRef = useRef();
	const messagesResponseRef = useRef()

	const handleHidden = () => {
		setIsChatHidden(!isChatHidden);
	};

	useEffect(() => {
		ComponentService.getAllMessages().then(
			(response) => {
				// setMessages(response.data.Messages);
				messagesResponseRef.current = response.data.Messsages;
				console.log('popup chat ok', messages);
			},
			(error) => {
				console.log('popup chat', error);
			}
		);
	}, []);

	const msgs = useMemo(() => {

	}, [messagesResponseRef])


	const handleSend = async (e) => {
		e.preventDefault();
		ComponentService.sendMessage(message).then(
			(response) => {
				console.log('popup chat send ok');
			},
			(error) => {
				console.log('popup chat send', error);
			}
		);
		setMessage('');
	};
	// chatRef.current.scrollTop = chatRef.current.scrollHeight;

	return (
		<div className="popup-chat">
			{isChatHidden ? (
				<></>
			) : (
				<div className="popup-chat--window">
					<label>Chat</label>
					<div ref={chatRef} className="popup-chat--messages">
						<Messages messages={messages} username={username} />
					</div>
					<form onSubmit={handleSend} className="popup-chat--form">
						<input
							className="popup-chat--input"
							type="text"
							value={message}
							required
							onChange={(e) => setMessage(e.target.value)}
						/>
						<button className="popup-chat--submit" type="submit">
							<BsSend color="black" size="15" />
						</button>
					</form>
				</div>
			)}
			<Button
				value="chat"
				color="secondary"
				size="m"
				float="right"
				className="popup-chat--btn"
				onClick={handleHidden}
			/>
		</div>
	);
};

export default PopupChat;
