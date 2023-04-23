import { useState, useEffect } from 'react'
import Navbar from '@components/Navbar/Navbar';
import ChatList from './components/ChatList/ChatList';
import ChatArea from './components/ChatArea/ChatArea';
import PageService from '@services/page.service'

const Chat = ({ isHidden, setIsHidden, username, role }) => {
	const [dialog, setDialog] = useState()
    const [notifications, setNotifications] = useState()
	const [messages, setMessages] = useState()

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
				setMessages(response.data.Messages)
				console.log('get dialog', messages)
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
				<ChatArea messages={messages} username={username}/>
			</div>
		</div>
	);
};

export default Chat;
