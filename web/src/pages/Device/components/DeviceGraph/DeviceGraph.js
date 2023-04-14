import { useState } from 'react';
import Container from '@components/Container/Container';
import { LineChart, Line, XAxis, Tooltip, ResponsiveContainer } from 'recharts';
import Tabs from '@components/Tabs/Tabs';
import './deviceGraph.scss';

export default function DeviceGraph({ charHistory }) {
	const [active, setActive] = useState(0);
	const dataMinutes = [];
	const isLonger = charHistory.length > 80 ? true : false;
	function padTo2Digits(num) {
		return String(num).padStart(2, '0');
	}
	for (
		let i = isLonger ? charHistory.length - 80 : 0;
		i < charHistory.length;
		i++
	) {
		const newDate = new Date(charHistory[i].Time);
		const hoursAndMinutes =
			padTo2Digits(newDate.getHours() - 6) +
			':' +
			padTo2Digits(newDate.getMinutes());
		const stats = {
			time: hoursAndMinutes,
			temperature: charHistory[i].ChipTempMax,
			hash: charHistory[i].MHSav,
			fanIn: charHistory[i].FanSpeedIn,
			fanOut: charHistory[i].FanSpeedOut,
		};
		dataMinutes.push(stats);
	}

	const RATIO_THRESHOLD = 70;
	var min_ratio = Math.min(...charHistory.map((item) => item.temperature));
	var max_ratio = Math.max(...charHistory.map((item) => item.temperature));
	min_ratio = 0;
	var threshold = min_ratio + ((max_ratio - min_ratio) * RATIO_THRESHOLD) / 100;

	const renderLineChartTemp = (
		<LineChart width={1400} height={400} data={dataMinutes} margin={{}}>
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
	);

	const renderLineChartFans = (
		<LineChart width={1400} height={400} data={dataMinutes} margin={{}}>
			<Line type="monotone" dataKey="fanIn" stroke="#eb4034" />
			<Line type="monotone" dataKey="fanOut" stroke="#5EF24C" />
			<XAxis dataKey="time" />
			<Tooltip />
		</LineChart>
	);

	const handleActive = (id) => {
		setActive(id);
		console.log(id);
	};

	return (
		<div style={{ marginTop: '20px' }}>
			<Tabs
				tabs={[
					{
						id: 0,
						label: 'Temperature',
					},
					{
						id: 1,
						label: 'Hashrate',
					},
					{
						id: 2,
						label: 'Fans speed',
					},
				]}
				active={active}
				handleActive={handleActive}
			/>
			{active === 0 && (
				<Container>
					<div className="graph--container">
						<ResponsiveContainer>{renderLineChartTemp}</ResponsiveContainer>
					</div>
				</Container>
			)}
			{active === 1 && (
				<Container>
					<div className="graph--container">
						<ResponsiveContainer>{renderLineChartHash}</ResponsiveContainer>
					</div>
				</Container>
			)}
			{active === 2 && (
				<Container>
					<div className="graph--container">
						<ResponsiveContainer>{renderLineChartFans}</ResponsiveContainer>
					</div>
				</Container>
			)}
		</div>
	);
}
