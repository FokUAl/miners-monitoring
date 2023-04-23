import './messages.scss';

const Messages = ({ messages, username }) => {
	console.log('Messages', messages);
	return (
		<div className="messages">
			{messages && messages.map((message, index) =>
				message.Sender === username ? (
					<div className="messages--container">
						<div className="messages--current-user" key={index}>
							{message.Content}
						</div>
					</div>
				) : (
					<div className="messages--container">
						<div className="messages--recipient-user" key={index}>
							{message.Content}
						</div>
					</div>
				)
			)}
		</div>
	);
};

export default Messages;
