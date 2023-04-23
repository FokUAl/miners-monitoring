import { useState, useEffect } from 'react'
import Navbar from '@components/Navbar/Navbar';
import ChatList from './components/ChatList/ChatList';
import ChatArea from './components/ChatArea/ChatArea';
import PageService from '@services/page.service'

const Chat = ({ isHidden, setIsHidden, username, role }) => {
	const [dialog, setDialog] = useState()
    const [notifications, setNotifications] = useState()
	const [chat, setChat] = useState()

	useEffect(() => {
		PageService.getNotifications().then(
			(response) => {
				setNotifications(response.data.List)
				console.log('notifications ok', notifications)
			},
			(error) => {
				console.log('notifications', error)
			}
		)
	}, [])

	useEffect(() => {
		PageService.postDialog(dialog, username).then(
			(response) => {
				setChat(response.data)
				console.log('get dialog', chat)
			},
			(error) => {
				console.log('get dialog', error)
			}
		)
	}, [dialog])

	return (
		<div className={isHidden ? 'nav-hidden' : 'nav-full'}>
			<Navbar
				isHidden={isHidden}
				setIsHidden={setIsHidden}
				role={role}
				username={username}
			/>
			<div className="grid-20-80">
				<ChatList notifications={notifications} setDialog={setDialog}/>
				<ChatArea chat={chat} username={username}/>
			</div>
		</div>
	);
};

export default Chat;
