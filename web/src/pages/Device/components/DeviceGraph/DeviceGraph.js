import Container from '@components/Container/Container';
import { LineChart, Line, XAxis, Tooltip, ResponsiveContainer } from 'recharts';
import './deviceGraph.scss';

export default function DeviceGraph({ charHistory }) {
	const dataMinutes = [];
	const graphLength = charHistory.length > 80 ? 80 : charHistory.length;
	function padTo2Digits(num) {
		return String(num).padStart(2, '0');
	}
	for (let i = 0; i < graphLength; i++) {
		const newDate = new Date(charHistory[i].Time);
		const hoursAndMinutes =
			padTo2Digits(newDate.getHours()-6) +
			':' +
			padTo2Digits(newDate.getMinutes());
			// newDate.getHours() + ":" + newDate.getMinutes()
		const temp = {
			time: hoursAndMinutes,
			temperature: charHistory[i].ChipTempMax,
		};
		dataMinutes.push(temp);
	}

	const RATIO_THRESHOLD = 70;
	var min_ratio = Math.min(...charHistory.map((item) => item.temperature));
	var max_ratio = Math.max(...charHistory.map((item) => item.temperature));
	min_ratio = 0;
	var threshold = min_ratio + ((max_ratio - min_ratio) * RATIO_THRESHOLD) / 100;
	console.log('Ratio', min_ratio, max_ratio, threshold);

	const renderLineChart = (
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

	return (
		<div style={{ marginTop: '20px' }}>
			<Container>
				<div className="graph--container">
					<ResponsiveContainer>{renderLineChart}</ResponsiveContainer>
				</div>
			</Container>
		</div>
	);
}
