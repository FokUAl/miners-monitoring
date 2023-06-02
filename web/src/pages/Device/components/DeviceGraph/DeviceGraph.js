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
			padTo2Digits(newDate.getUTCHours()) +
			':' +
			padTo2Digits(newDate.getUTCMinutes());
		const stats = {
			time: hoursAndMinutes,
			temperature: charHistory[i].ChipTempMax,
			hash: charHistory[i].THSav,
			fan1: charHistory[i].FanSpeed1,
			fan2: charHistory[i].FanSpeed2,
			fan3: charHistory[i].FanSpeed3,
			fan4: charHistory[i].FanSpeed4,
		};
		dataMinutes.push(stats);
	}
	
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
			<Line type="monotone" dataKey="fan1" stroke="#eb4034" />
			<Line type="monotone" dataKey="fan2" stroke="#5EF24C" />
			<Line type="monotone" dataKey="fan3" stroke="#001aff" />
			<Line type="monotone" dataKey="fan4" stroke="#ff00c3" />
			<XAxis dataKey="time" />
			<Tooltip />
		</LineChart>
	);

	const handleActive = (id) => {
		setActive(id);
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
				<Container paddingLeft paddingRight>
					<div className="graph--container">
						<ResponsiveContainer>{renderLineChartTemp}</ResponsiveContainer>
					</div>
				</Container>
			)}
			{active === 1 && (
				<Container paddingLeft paddingRight>
					<div className="graph--container">
						<ResponsiveContainer>{renderLineChartHash}</ResponsiveContainer>
					</div>
				</Container>
			)}
			{active === 2 && (
				<Container paddingLeft paddingRight>
					<div className="graph--container">
						<ResponsiveContainer>{renderLineChartFans}</ResponsiveContainer>
					</div>
				</Container>
			)}
		</div>
	);
}
