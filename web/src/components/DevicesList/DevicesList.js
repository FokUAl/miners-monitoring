import React from 'react';
import Button from '../Button/Button';
import Container from '../Container/Container';
import './devicesList.scss';
import 'react-table-6/react-table.css';
import {
	Box,
	Stack,
	useTheme,
	createTheme,
	ThemeProvider,
} from '@mui/material';
import MaterialReactTable from 'material-react-table';
import PageService from '../../services/page.service';
import DataReg from './DataReg'

export default function DevicesList({ devices, setDevices }) {
	const {columns, data} = DataReg(devices)

	const theme = createTheme({
		components: {
			MuiTableBodyProps: {
				styleOverrides: {
					root: {
						backgroundColor: '#ffffff',
					},
				},
			},
		},
	});
	
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
			<Button size="m" value="Update Table" onClick={UpdateRequest} />
			<ThemeProvider theme={theme}>
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
					muiTableContainerProps={{
						sx: { padding: '0' },
					}}
					muiTableHeadCellProps={{
						sx: { fontSize: '14px', padding: '0 5px' },
					}}
					muiTableBodyCellProps={{
						sx: { fontSize: '14px', padding: '0 5px', margin: '0' },
					}}
					muiTableBodyProps={{
						sx: {
							backgroundColor: 'seconday',
						},
					}}
				/>
			</ThemeProvider>
		</Container>
	);
}
