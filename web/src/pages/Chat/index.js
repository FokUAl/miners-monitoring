import Navbar from '@components/Navbar/Navbar';
import ChatList from './components/ChatList/ChatList';
import ChatArea from './components/ChatArea/ChatArea';

const Chat = ({ isHidden, setIsHidden, username, role, notifications }) => {
	return (
		<div className={isHidden ? 'nav-hidden' : 'nav-full'}>
			<Navbar
				isHidden={isHidden}
				setIsHidden={setIsHidden}
				role={role}
				username={username}
			/>
			<div className="grid-20-80">
				<ChatList notifications={notifications}/>
				<ChatArea />
			</div>
		</div>
	);
};

export default Chat;
