import React, { useMemo } from 'react';
import Button from '@components/Button/Button';
import Container from '@components/Container/Container';
import './devicesList.scss';
import 'react-table-6/react-table.css';
import MaterialReactTable from 'material-react-table';
import { ThemeProvider } from '@mui/material';
import PageService from '@services/page.service';
import DataReg from './DataReg';
import TableTheme from './TableTheme';

export default function DevicesList({ devices, setDevices }) {
	const { columns, data } = DataReg(devices);

	const tableTheme = TableTheme

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
		<Container paddingRight paddingLeft>
			<div className="m-bm">
				<Button size="m" value="Update Table" onClick={UpdateRequest} />
			</div>
			<ThemeProvider theme={tableTheme}>
				<MaterialReactTable
					columns={columns}
					data={data}
					enableGrouping
					enableStickyHeader
					enableStickyFooter
					enableDensityToggle={false}
					muiTableContainerProps={{ sx: { maxHeight: '600px' } }}
					manualPagination
					enableBottomToolbar={false}
					enableColumnActions={false}
					initialState={{
						density: 'compact',
						expanded: true,
						grouping: ['Owner', 'MinerType'],
						pagination: { pageIndex: 0, PageSize: 100 },
						sorting: [{ id: 'Owner', desc: false }],
					}}
					muiTableHeadCellProps={{
						sx: {
							backgroundColor: '#2D2D2D',
							color: 'white',
							fontSize: '12px',
							padding: '0 2px',
							borderTop: 'none',
							borderInline: 'none',
							borderColor: 'black',
						},
					}}
					muiTableBodyCellProps={{
						sx: {
							backgroundColor: '#5a5a5a',
							color: 'white',
							fontSize: '12px',
							padding: '0 2px',
							borderTop: 'none',
							borderInline: 'none',
							borderColor: 'black',
						},
					}}
					muiTableProps={{
						sx: {
							padding: '0px 0px 0px 0px',
							border: 'none',
						},
					}}
					muiTablePaperProps={{
						sx: {
							// color: 'white',
							fontSize: '12px',
							backgroundColor: '#353535',
						},
					}}
					muiTopToolbarProps={{
						sx: {
							backgroundColor: '#353535',
						}
					}}
					muiToolbarAlertBannerChipProps={{
						sx: {
							backgroundColor: '#5a5a5a',
							// color: 'white',
							fontSize: '12px',
						},
					}}
					muiToolbarAlertBannerProps={{
						sx: {
							backgroundColor: '#333333',
							color: 'white',
							fontSize: '12px',
						},
					}}
				/>
			</ThemeProvider>
		</Container>
	);
}
