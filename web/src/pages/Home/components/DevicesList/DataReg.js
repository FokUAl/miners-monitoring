import { useMemo } from 'react';
import { Box } from '@mui/material';
import { Link } from 'react-router-dom';
import Button from '@components/Button/Button';

function DataReg(devices, allUsers) {
	const data = devices.map((el) => {
		return {
			...el,
			Link: `/device?shelf=${el.Shelf}&row=${el.Row}&column=${el.Column}`,
			ChipTempAvg: el.Characteristics.ChipTempAvg,
			ChipTempMax: el.Characteristics.ChipTempMax,
			ChipTempMin: el.Characteristics.ChipTempMin,
			Elapsed: el.Characteristics.Elapsed,
			FanSpeeds: `${el.Characteristics.FanSpeed1}-${el.Characteristics.FanSpeed2}-${el.Characteristics.FanSpeed3}-${el.Characteristics.FanSpeed4}`,
			MAC: el.Characteristics.MAC,
			THSav: el.Characteristics.THSav,
			PowerMode: el.Characteristics.PowerMode,
			Temperature: el.Characteristics.Temperature,
			Shelf: el.Shelf,
			Level: el.Row,
			Miner: el.Column,
			Characteristics: null,
		};
	});

	const columns = useMemo(
		() => [
			{
				header: 'Owner',
				accessorKey: 'Owner',
				size: 1,
				filterVariant: 'select',
				filterSelectOptions: allUsers && allUsers,
			},
			{
				header: 'Miner Model',
				accessorKey: 'MinerType',
				size: 1,
			},
			{
				header: 'Shelf',
				accessorKey: 'Shelf',
				size: 1,
				enableGrouping: false,
				filterVarian: 'range',
				filterFn: 'betweenInclusive',
			},
			{
				header: 'Level',
				accessorKey: 'Level',
				size: 1,
				enableGrouping: false,
				filterVarian: 'range',
				filterFn: 'betweenInclusive',
			},
			{
				header: 'Miner',
				accessorKey: 'Miner',
				size: 1,
				enableGrouping: false,
				filterVarian: 'range',
				filterFn: 'betweenInclusive',
			},
			{
				header: 'Temperature',
				accessorKey: 'Temperature',
				size: 1,
				enableGrouping: false,
				aggregationFn: ['max', 'min'],
				AggregatedCell: ({ cell }) => (
					<div className="table--temperature">
						Max:{' '}
						<Box
							sx={{
								display: 'inline',
								color:
									cell.getValue()[0] >= 70 && cell.getValue()[0] <= 80
										? '#76a15e' //green
										: cell.getValue()[0] < 70
										? '#3f8ae0' //blue
										: '#f83b3b', //red
							}}
						>
							{cell.getValue()[0]}
						</Box>{' '}
						Min:{' '}
						<Box
							sx={{
								display: 'inline',
								color:
									cell.getValue()[1] >= 70 && cell.getValue()[1] <= 80
										? '#76a15e' //green
										: cell.getValue()[1] < 70
										? '#3f8ae0' //blue
										: '#f83b3b', //red
							}}
						>
							{cell.getValue()[1]}
						</Box>
					</div>
				),
			},
			{
				header: 'TH/s',
				accessorKey: 'THSav',
				size: 1,
				enableGrouping: false,
				aggregationFn: 'sum',
				AggregatedCell: ({ cell }) => <div>Total THS: {cell.getValue()}</div>,
			},
			{
				header: 'Fan Speeds',
				accessorKey: 'FanSpeeds',
				size: 1,
				enableGrouping: false,
				enableColumnFilter: false,
			},
			{
				header: 'Status',
				accessorKey: 'MinerStatus',
				size: 1,
				enableGrouping: false,
				enableColumnFilter: false,
				aggregationFn: ['unique', 'count'],
				AggregatedCell: ({ cell, column, row }) => {
					return (
					<div>
						<Box sx={{display:'inline'}}>{row.getValue(column.id).filter(el => el[0] === 'online').length}</Box>
						{'/'}
						<Box sx={{display:'inline'}}>{cell.getValue()[1]}</Box>
						{' '}
					</div>
				)},
				Cell: ({ cell }) => {
					return (
						<div>
							<Box
								sx={{
									display: 'inline',
									color: cell.getValue() === 'online' ? '#76a15e' : '#f83b3b', //green & red
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
				Cell: ({ cell }) => {
					return (
						<a
							className="table--link"
							href={`http://${cell.getValue()}`}
							target="_blank"
							rel="noreferrer"
						>
							{cell.getValue()}
						</a>
					);
				},
				enableColumnFilter: false,
			},
			{
				header: 'Link',
				accessorKey: 'Link',
				size: 1,
				enableColumnFilter: false,
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
