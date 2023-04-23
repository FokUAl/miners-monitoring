import Container from '@components/Container/Container';
import Messages from '@components/PopupChat/Messages';

const ChatArea = ({ chat, username }) => {
	return (
		<Container>
			{chat ? <Messages messages={chat} username={username} /> : 'Open Dialog'}
		</Container>
	);
};

export default ChatArea;
