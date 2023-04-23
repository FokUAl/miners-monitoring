import Container from '@components/Container/Container';
import Messages from '@components/PopupChat/Messages';

const ChatArea = ({ messages, username }) => {
	return (
		<Container verticalHeight overflowY>
			{messages ? <Messages messages={messages} username={username} /> : 'Open Dialog'}
		</Container>
	);
};

export default ChatArea;
