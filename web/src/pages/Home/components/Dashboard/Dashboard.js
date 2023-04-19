import React from 'react';
import Container from '@components/Container/Container';
import ReData from './ReData';
import { PieChart, Pie, Cell, Tooltip, ResponsiveContainer } from 'recharts';
import './dashboard.scss';

export default function Dashboard({ devices }) {
	const { onlineCount } = ReData(devices);
	const chartData = [
		{ name: 'Online', value: onlineCount },
		{ name: 'Offline', value: devices.length - onlineCount },
	];
    const colors = ['#76a15e', '#c50000']
	return (
		<Container paddingLeft borderBottom>
			<div className="dash--title">Dashboard</div>
			<div className="grid-50-50">
				<div>
					{onlineCount}/{devices.length}
				</div>
				<div>{onlineCount / devices.length}%</div>
			</div>
				<PieChart width={730} height={250}>
					<Pie
						data={chartData}
						dataKey="value"
						nameKey="name"
						cx="50%"
						cy="50%"
						outerRadius={50}
						fill="#8884d8"
                        label
					>
                    {chartData.map((entry, index) => (
                      <Cell key={`cell-${index}`} fill={colors[index % colors.length]} />
                    ))}
                    </Pie>
                    <Tooltip />
				</PieChart>
		</Container>
	);
}
