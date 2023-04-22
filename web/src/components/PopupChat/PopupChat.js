import { useEffect, useState } from 'react';
import Button from '@components/Button/Button';
import Input from '@components/Input/Input'
import Messages from './Messages'
import ComponentService from '@services/component.service'
import './popupChat.scss';
import PageService from '../../services/page.service';

const PopupChat = ({ username }) => {
	const [isChatHidden, setIsChatHidden] = useState(true);
    const [messages, setMessages] = useState()
    const [message, setMessage] = useState()

    const handleHidden = () => {
        setIsChatHidden(!isChatHidden)
    }

    useEffect(() => {
        ComponentService.getAllMessages().then(
            (response) => {
                setMessages(response.data.Messages)
                console.log('popup chat ok', messages)
            }, 
            (error) => {
                console.log('popup chat', error)
            }
        )
    }, [])

    const handleSend = ({ message }) => {
        ComponentService.sendMessage(message).then(
            (response) => {
                console.log('popup chat send ok')
            }, 
            (error) => {
                console.log('popup chat send', error)
            }
        )
        setMessage('')
    }

	return (
		<div className="popup-chat">
			{isChatHidden ? <></> : 
            <div className="popup-chat--window">
                <label>Chat</label>
                <div className="popup-chat--messages">
                    <Messages messages={messages} username={username} />
                </div>
                <div className="grid-85-15">
                    <Input size="s" value={message} setValue={setMessage}/>
                    <Button value="send" size='s' onClick={() => handleSend(message)}/>
                </div>
            </div>
            }
			<Button value="chat" color="secondary" size="m" float="right" fluid onClick={handleHidden}/>
		</div>
	);
};

export default PopupChat;
