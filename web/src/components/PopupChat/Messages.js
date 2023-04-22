import './messages.scss'

const Messages = ({ messages, username }) => {
	return (
		<div className="messages">
			{messages &&
				messages.map((message, index) =>
					message.MinerType === username ? (
						<div
							className="messages--current-user"
							key={message.MinerType + index}
						>
							{message.MinerType}
						</div>
					) : (
						<div
							className="messages--recipient-user"
							key={message.MinerType + index}
						>
							{message.MinerType}
						</div>
					)
				)}
		</div>
	);
};

export default Messages;
