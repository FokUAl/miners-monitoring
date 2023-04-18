import React from 'react';
import Button from '@components/Button/Button';
import Container from '@components/Container/Container';
import './devicesList.scss';
import 'react-table-6/react-table.css';
import MaterialReactTable from 'material-react-table';
import PageService from '@services/page.service';
import DataReg from './DataReg';

export default function DevicesList({ devices, setDevices }) {
	const { columns, data } = DataReg(devices);

	const UpdateRequest = () => {
		PageService.getHome().then(
			(response) => {
				setDevices(response.data.Devices);
				console.log('update ok ');
			},
			(error) => {
				console.log('update error: ', error);
			}
		);
	};

	return (
		<Container>
			<div className="m-bm">
				<Button size="m" value="Update Table" onClick={UpdateRequest} />
			</div>
			<MaterialReactTable
				columns={columns}
				data={data}
				enableGrouping
				enableStickyHeader
				enableStickyFooter
				enableRowVirtualization
				manualPagination
				enableBottomToolbar={false}
				initialState={{
					density: 'compact',
					expanded: true,
					grouping: ['Owner'],
					pagination: { pageIndex: 0, PageSize: 100 },
					sorting: [{ id: 'Owner', desc: false }],
				}}
				// muiTableHeadCellProps={{
				// 	sx: {
				// 		backgroundColor: '#888787',
				// 		color: 'white',
				// 		fontSize: '12px',
				// 	},
				// }}
				// muiTableBodyCellProps={{
				// 	sx: {
				// 		backgroundColor: '#888787',
				// 		color: 'white',
				// 		fontSize: '12px',
				// 	}
				// }}
				// muiTablePaperProps={{
				// 	sx: {
				// 		backgroundColor: '#888787',
				// 		color: 'white',
				// 		fontSize: '12px',
				// 	}
				// }}
				// muiTopToolbarProps={{
				// 	sx: {
				// 		backgroundColor: '#1e1e1e',
				// 		color: 'white',
				// 		fontSize: '10px'
				// 	}
				// }}
			/>
		</Container>
	);
}
