import './deviceInfo.scss'

export default function CommentDisplay ({username, date, content}) {
	const newDate = new Date(date).toUTCString()

	return (
        <div>
            <div className="grid-50-50">
                <div className="float-left">{username}</div>
                <div className="float-right">{newDate}</div>
            </div>
            <div className="data-display-wrapper" />
            <div className="float-left">{content}</div>
        </div>
	);
}