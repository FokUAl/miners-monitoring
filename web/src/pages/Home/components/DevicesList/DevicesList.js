import React, { useMemo, useState } from 'react';
import { Link } from 'react-router-dom';
import Button from '@components/Button/Button';
import Container from '@components/Container/Container';
import './devicesList.scss';
import 'react-table-6/react-table.css';
import MaterialReactTable from 'material-react-table';
import { ThemeProvider } from '@mui/material';
import { Box } from '@mui/material';
import PageService from '@services/page.service';
import FormService from '@services/form.service';
import DataReg from './DataReg';
import TableTheme from './TableTheme';
import { IoEnter } from 'react-icons/io5';
import { BsTrashFill } from 'react-icons/bs';

export default function DevicesList({
	devices,
	setDevices,
	delay,
	setDelay,
	allUsers,
}) {
	const { columns, data } = DataReg(devices, allUsers);

	const tableTheme = TableTheme;

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

	const handleDelete = async (IP) => {
		const confirmed = window.confirm("Delete device?");
		if (confirmed) {
			FormService.deleteDevice(IP).then(
				(response) => {
					UpdateRequest();
				},
				(error) => {
					console.log('Add device ', error);
				}
			);
		}
	};

	return (
		<Container paddingRight paddingLeft>
			<div className="grid-50-50 m-bm">
				<Button size="m" value="Update Table" onClick={UpdateRequest} />
				<div className="grid-50-50 align-items-center">
					<label className="float-right m-rt">Auto Update</label>
					<select
						className="input--select size-s width-m color-primary float-left"
						value={delay}
						onChange={(e) => setDelay(e.target.value)}
					>
						<option value=""></option>
						<option value="60000">1 min</option>
						<option value="300000">5 min</option>
						<option value="600000">10 min</option>
					</select>
				</div>
			</div>
			<ThemeProvider theme={tableTheme}>
				<MaterialReactTable
					columns={columns}
					data={data}
					enableGrouping
					enableStickyHeader
					manualPagination
					enableDensityToggle={false}
					enableBottomToolbar={false}
					enableColumnActions={false}
					initialState={{
						density: 'compact',
						expanded: false,
						grouping: ['Owner', 'MinerType'],
						pagination: { pageIndex: 0, PageSize: 100 },
						sorting: [{ id: 'Owner', desc: false }],
					}}
					enableRowActions
					positionActionsColumn="last"
					renderRowActions={({ row, table }) => (
						<Box sx={{ display: 'flex', flexWrap: 'nowrap', gap: '8px' }}>
							<Link
								className="link"
								to={`/device?shelf=${row.original.Shelf}&row=${row.original.Level}&column=${row.original.Miner}`}
							>
								<Button size="s" value={<IoEnter />} />
							</Link>
							<Button
								size="s"
								value={<BsTrashFill />}
								onClick={() => handleDelete(row.original.IPAddress)}
							/>
						</Box>
					)}
					muiTableContainerProps={{ sx: { maxHeight: '600px' } }}
					muiTableHeadCellProps={{
						sx: {
							backgroundColor: '#2D2D2D',
							color: 'white',
							fontSize: '12px',
							padding: '0 0 0 5px',
							border: 'none',
						},
					}}
					muiTableBodyCellProps={{
						sx: {
							backgroundColor: '#5a5a5a',
							color: 'white',
							fontSize: '12px',
							fontWeight: '500',
							padding: '0 0 0 5px',
							borderTop: 'none',
							borderRight: 'none',
							borderColor: '#1e1e1e',
						},
					}}
					muiExpandButtonProps={{
						sx: {
							color: 'white',
							'&.Mui-disabled': {
								color: 'transparent',
								backgroundColor: 'transparent',
							},
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
							fontSize: '12px',
							backgroundColor: '#353535',
						},
					}}
					muiTopToolbarProps={{
						sx: {
							backgroundColor: '#353535',
						},
					}}
					muiToolbarAlertBannerChipProps={{
						sx: {
							backgroundColor: '#5a5a5a',
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
