import './deviceInfo.scss';

export default function DataDisplay({ text, data, type, date }) {
	function ToTimeString(totalSeconds) {
		const totalMs = totalSeconds * 1000;
		const result = new Date(totalMs).toISOString().slice(11, 19);

		return result;
	}
	let newDate
	if (date) {
		newDate = new Date(date).toUTCString()
	}

	return (
		<>
			<div className="grid-50-50">
				<div className="float-left">{text}</div>
				{type === 'time' ? (
					<div className="float-right">{ToTimeString(data)}</div>
				) : (
					<div className="float-right">{data}</div>
				)}
			</div>
			{date ? (
				<div className="data-display-wrapper-plus">
					<div className="float-left">{newDate}</div>
				</div>
			) : (
				<div className="data-display-wrapper" />
			)}
		</>
	);
}

DataDisplay.defaultProps = {
	text: '',
	data: '',
	type: undefined,
	date: '',
};
