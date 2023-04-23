import { useState, useEffect } from 'react'
import Navbar from '@components/Navbar/Navbar';
import ChatList from './components/ChatList/ChatList';
import ChatArea from './components/ChatArea/ChatArea';
import PageService from '@services/page.service'

const Chat = ({ isHidden, setIsHidden, username, role }) => {
    const [notifications, setNotifications] = useState()

	useEffect(() => {
		PageService.getNotifications().then(
			(response) => {
				setNotifications(response.data.SenderStat)
				console.log('notifications ok', notifications)
			},
			(error) => {
				console.log('notifications', error)
			}
		)
	}, [])

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
