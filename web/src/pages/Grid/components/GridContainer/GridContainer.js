import { useState } from 'react';
import Container from '@components/Container/Container';
import GridComponent from '@components/GridComponent/GridComponent';
import GridCell from './GridCell';
import Tabs from '@components/Tabs/Tabs';
import './gridContainer.scss';

export default function GridContainer({ devices, shelf }) {
	const [tabsActive, setTabsActive] = useState(0);

	const handleTabsActive = (id) => {
		setTabsActive(id);
	};

	const GridRowBuilder = (devices, type) => {
		if (devices === null) return;
		const gridRows = [];
		for (let i = 14; i >= 0; i--) {
			const currentRow = devices.filter((device) => device.Row === i);
			for (let j = 0; j < 10; j++) {
				const currentDevice = currentRow.find(
					(device) => device.Column === j + 1
				);
				gridRows.push(
					<GridCell
						deviceChar={currentDevice ? currentDevice : null}
						key={'' + i + j + shelf}
						type={type}
					/>
				);
			}
		}
		return gridRows;
	};

	const GridShelfBuilder = ({ devices, start, end, type }) => {
		if (devices === null) return null;
		const gridShelfs = [];
		for (let i = start; i <= end; i++) {
			const currentShelfDevices = devices.filter(
				(device) => device.Shelf === i
			);
			gridShelfs.push(
				<GridComponent gridColumns="10" key={`${i}`}>
					{GridRowBuilder(
						currentShelfDevices.length ? currentShelfDevices : null,
						type,
						i
					)}
				</GridComponent>
			);
		}
		return gridShelfs;
	};
	return (
		<Container paddingLeft>
			<Tabs
				tabs={[
					{
						id: 0,
						label: 'Online map',
					},
					{
						id: 1,
						label: 'Heat map',
					},
				]}
				active={tabsActive}
				handleActive={handleTabsActive}
			/>
			{tabsActive === 0 && (
				<GridComponent gridColumns="11">
					<GridShelfBuilder
						devices={devices}
						start={1}
						end={11}
						type={'onlineMap'}
					/>
					<GridShelfBuilder
						devices={devices}
						start={11}
						end={22}
						type={'onlineMap'}
					/>
					<GridShelfBuilder
						devices={devices}
						start={22}
						end={33}
						type={'onlineMap'}
					/>
					<GridShelfBuilder
						devices={devices}
						start={33}
						end={44}
						type={'onlineMap'}
					/>
				</GridComponent>
			)}
			{tabsActive === 1 && (
				<GridComponent gridColumns="11">
					<GridShelfBuilder
						devices={devices}
						start={1}
						end={11}
						type={'heatMap'}
					/>
					<GridShelfBuilder
						devices={devices}
						start={11}
						end={22}
						type={'heatMap'}
					/>
					<GridShelfBuilder
						devices={devices}
						start={22}
						end={33}
						type={'heatMap'}
					/>
					<GridShelfBuilder
						devices={devices}
						start={33}
						end={44}
						type={'heatMap'}
					/>
				</GridComponent>
			)}
		</Container>
	);
}
