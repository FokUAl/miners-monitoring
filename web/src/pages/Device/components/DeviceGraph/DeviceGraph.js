import Container from '@components/Container/Container';
import { LineChart, Line, XAxis, Tooltip, ResponsiveContainer } from 'recharts';
import './deviceGraph.scss';

export default function DeviceGraph({ charHistory }) {
	const dataMinutes = [];
	const isLonger = charHistory.length > 8 ? true : false
	function padTo2Digits(num) {
		return String(num).padStart(2, '0');
	}
	for (let i = (isLonger ? (charHistory.length - 80) : 0); i < (isLonger ? charHistory.length : 80); i++) {
		const newDate = new Date(charHistory[i].Time);
		const hoursAndMinutes =
			padTo2Digits(newDate.getHours()-6) +
			':' +
			padTo2Digits(newDate.getMinutes());
		const stats = {
			time: hoursAndMinutes,
			temperature: charHistory[i].ChipTempMax,
			hash: charHistory[i].MHSav
		};
		dataMinutes.push(stats);
	}

	const RATIO_THRESHOLD = 70;
	var min_ratio = Math.min(...charHistory.map((item) => item.temperature));
	var max_ratio = Math.max(...charHistory.map((item) => item.temperature));
	min_ratio = 0;
	var threshold = min_ratio + ((max_ratio - min_ratio) * RATIO_THRESHOLD) / 100;
	console.log('Ratio', min_ratio, max_ratio, threshold);

	const renderLineChartTemp = (
		<LineChart width={1400} height={400} data={dataMinutes} margin={{}}>
			<defs>
				<linearGradient id="color50pct" x1="0%" x2="0%" y2="0%" y1="100%">
					<stop offset="0%" stopColor="red" />
					<stop offset={`${threshold - 0.01}%`} stopColor="red" />
					<stop offset={`${threshold + 0.01}%`} stopColor="green" />

					{/* <stop offset="49.9%" stopColor="red" />
                    <stop offset="50%" stopColor="green" /> */}

					<stop offset="100%" stopColor="green" />
				</linearGradient>
			</defs>
			<Line type="monotone" dataKey="temperature" stroke="#eb4034" />
			<XAxis dataKey="time" />
			<Tooltip />
		</LineChart>
	);

	const renderLineChartHash = (
		<LineChart width={1400} height={400} data={dataMinutes} margin={{}}>
			<Line type="monotone" dataKey="hash" stroke="#eb4034" />
			<XAxis dataKey="time" />
			<Tooltip />
		</LineChart>
	)

	return (
		<div style={{ marginTop: '20px' }}>
			<Container>
				<div className="graph--container">
					<ResponsiveContainer>{renderLineChartTemp}</ResponsiveContainer>
				</div>
			</Container>
			<Container>
				<div className="graph--container">
					<ResponsiveContainer>{renderLineChartHash}</ResponsiveContainer>
				</div>
			</Container>
		</div>
	);
}
