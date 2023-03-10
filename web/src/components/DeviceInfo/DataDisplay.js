import './deviceInfo.scss';

export default function DataDisplay({ text, data, type }) {
	function ToTimeString(totalSeconds) {
		const totalMs = totalSeconds * 1000;
		const result = new Date(totalMs).toISOString().slice(11, 19);
	  
		return result;
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
			<div className="wrapper" />
		</>
	);
}

DataDisplay.defaultProps = {
	text: '',
	data: '',
	type: undefined,
};
