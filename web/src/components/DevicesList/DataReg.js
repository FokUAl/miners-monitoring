import { useMemo } from 'react';
import { Box } from '@mui/material';
import { Link } from 'react-router-dom';
import Button from '../Button/Button';

function DataReg(devices) {
	const data = devices.map((el) => {
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
				aggregationFn: ['max', 'min'],
				AggregatedCell: ({ cell }) => (
					<div className="table--temperature">
						Max:
						<Box
							sx={{
								display: 'inline',
								color:
									cell.getValue()[0] >= 70 && cell.getValue()[0] <= 80
										? 'green'
										: cell.getValue()[0] < 70
										? 'blue'
										: 'red',
							}}
						>
							{cell.getValue()[0]}
						</Box>
						Min:
						<Box
							sx={{
								display: 'inline',
								color:
									cell.getValue()[1] >= 70 && cell.getValue()[1] <= 80
										? 'green'
										: cell.getValue()[1] < 70
										? 'blue'
										: 'red',
							}}
						>
							{cell.getValue()[1]}
						</Box>
					</div>
				),
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
				header: 'Status',
				accessorKey: 'MinerStatus',
				size: 1,
				enableGrouping: false,
				Cell: ({ cell }) => {
					return (
						<div>
							<Box
								sx={{
									display: 'inline',
									color: cell.getValue() === 'online' ? 'green' : 'red',
								}}
							>
								{cell.getValue()}
							</Box>
						</div>
					);
				},
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
				Cell: ({ cell }) => {
					return (
						<Link className="link" to={cell.getValue()}>
							<Button size="s" value="Link"></Button>
						</Link>
					);
				},
			},
		],
		[]
	);
	return { columns, data };
}

export default DataReg;
