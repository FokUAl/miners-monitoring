import React, { useMemo } from 'react';
import ReactTable from 'react-table-6';
import { Link } from 'react-router-dom';
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

export default function DevicesList(props) {
	const data = props.devices.devices.map((el) => {
		return {
			...el,
			Link: `/device?shelf=${el.Shelf}&row=${el.Row}&column=${el.Column}`,
			ChipTempAvg: el.Characteristics.ChipTempAvg,
			ChipTempMax: el.Characteristics.ChipTempMax,
			ChipTempMin: el.Characteristics.ChipTempMin,
			Elapsed: el.Characteristics.Elapsed,
			FanSpeedIn: el.Characteristics.FanSpeedIn,
			FanSpeedOut: el.Characteristics.FanSpeedOut,
			MAC: el.Characteristics.MAC,
			MHSav: el.Characteristics.MHSav,
			PowerMode: el.Characteristics.PowerMode,
			Temperature: el.Characteristics.Temperature,
            Location: `${el.Shelf}-${el.Row}-${el.Column}`,
			Characteristics: null,
		};
	});

	const columns = useMemo(
		() => [
			{
				header: 'Owner',
				accessorKey: 'Owner',
                size: 1,
			},
			{
				header: 'Miner Model',
				accessorKey: 'MinerType',
                size: 1,
			},
			{
				header: 'Location',
				accessorKey: 'Location',
                size: 1,
				enableGrouping: false,
			},
            {
                header: 'Temperature',
                accessorKey: 'Temperature',
                size: 1,
				enableGrouping: false,
            },
            {
                header: 'Mega Hashrate Average',
                accessorKey: 'MHSav',
                size: 1,
				enableGrouping: false,
            },
            {
                header: 'Fan Speed In',
                accessorKey: 'FanSpeedIn',
                size: 1,
				enableGrouping: false,
            },
            {
                header: 'Fan Speed Out',
                accessorKey: 'FanSpeedOut',
                size: 1,
				enableGrouping: false,
            },
			{
				header: 'IP',
				accessorKey: 'IPAddress',
                size: 1,
				enableGrouping: false,
			},
            {
                header: 'Link',
                accessorKey: 'Link',
                size: 1,
				enableGrouping: false,
                Cell: ({ cell }) => {return <Link className="link" to={cell.getValue()}><Button className="btn--less" value="Link"></Button></Link>}
            },
		],
		[]
	);

    const theme = createTheme({
      components: {
        MuiTableBodyProps: {
            styleOverrides: {
                root: {
                    backgroundColor: "#ffffff"
                }
            }
        }
      }
    });

	return (
		<Container>
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
                        muiTableContainerProps= {{
                            sx: { padding: '0'}
                        }}
                        muiTableHeadCellProps={{
                            sx: { fontSize: '14px', padding: '0 5px' },
                        }}
                        muiTableBodyCellProps={{
                            sx: { fontSize: '14px', padding: '0 5px', margin: "0"},
                        }} 
                        actions={[
                            rowData => ({
                            icon: () => <Link className="link" to={rowData.Link}>link</Link>,
                            tooltip: 'Link to Device',
                            onClick: (rowData)
                            })
                        ]}
                    />
            </ThemeProvider>
		</Container>
	);
}
