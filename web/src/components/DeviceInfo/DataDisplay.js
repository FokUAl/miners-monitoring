import './deviceInfo.scss';

export default function DataDisplay({ text, data }) {
	return (
		<>
			<div className="grid-50-50">
				<div className="float-left">{text}</div>
				<div className="float-right">{data}</div>
			</div>
			<div className="wrapper" />
		</>
	);
}
