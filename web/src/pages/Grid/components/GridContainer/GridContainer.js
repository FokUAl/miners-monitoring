import Container from '@components/Container/Container';
import GridComponent from '@components/GridComponent/GridComponent';
import GridCell from './GridCell';
import './gridContainer.scss';

export default function GridContainer({devices, shelf}) {
	const GridRowBuilder = (devices) => {
		if (devices === null) return
		const gridRows = [];
		console.log('row', devices)
		for (let i = 14; i >= 0; i--) {
			const currentRow = devices.filter(device => device.Row === i)
			console.log('row1', currentRow)
			for (let j = 0; j < 10; j++) {
				const currentDevice = currentRow.find(device => device.Column === j + 1)
				gridRows.push(<GridCell deviceChar={currentDevice ? currentDevice : null} key={''+i+j+shelf}></GridCell>);
			}
		}
		return gridRows;
	};

	const GridShelfBuilder = ({ devices, start, end }) => {
		if (devices === null) return null
		console.log('ex', devices)
		const gridShelfs = [];
		for (let i = start; i <= end; i++) {
            const currentShelfDevices = devices.filter(device => device.Shelf === i)
			console.log('shelf', i, currentShelfDevices)
			gridShelfs.push(
				<GridComponent gridColumns="10" key={`${i}`}>
					{GridRowBuilder(currentShelfDevices.length ? currentShelfDevices : null, i)}
				</GridComponent>
			);
		}
		return gridShelfs;
	};
	return (
		<Container paddingLeft>
			<GridComponent gridColumns="11">
				<GridShelfBuilder devices={devices} start={1} end={11}/>
				<GridShelfBuilder devices={devices} start={11} end={22}/>
				<GridShelfBuilder devices={devices} start={22} end={33}/>
				<GridShelfBuilder devices={devices} start={33} end={44}/>
			</GridComponent>
		</Container>
	);
}
