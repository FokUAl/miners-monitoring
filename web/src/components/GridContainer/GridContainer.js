import Container from '../Container/Container';
import GridComponent from '../GridComponent/GridComponent';
import GridCell from './GridCell';
import './gridContainer.scss';

export default function GridContainer({devices}) {
	const GridRowBuilder = () => {
		const gridRows = [];
		for (let i = 0; i < 140; i++) {
			gridRows.push(<GridCell id={i+1} key={i}></GridCell>);
		}
		return gridRows;
	};
	const GridShelfBuilder = (devices) => {
        // console.log(devices.devices)
		const gridShelfs = [];
		for (let i = 0; i < 11; i++) {
            // const devicesCurrentShelf = devices.devices.filter((el) => el === i.toString())
            devices.devices.forEach(el => console.log(el))
			gridShelfs.push(
				<GridComponent gridColumns="10" id={`${i + 1}`} key={`${i}`}>
					{GridRowBuilder()}
				</GridComponent>
			);
		}
		return gridShelfs;
	};
	return (
		<Container>
			<GridComponent gridColumns="11">
				<GridShelfBuilder devices={devices}/>
			</GridComponent>
		</Container>
	);
}
