import { useState, useMemo } from 'react';
import ComponentService from '@services/component.service';
import Container from '@components/Container/Container';
import Button from '@components/Button/Button';
import {
	Paper,
	Table,
	TableBody,
	TableContainer,
	TableCell,
	TableHead,
	TableRow,
} from '@mui/material';
import './allIP.scss';

export default function AllIP({ allIP, setAllIP }) {
	const [loading, setLoading] = useState(true);

	const UpdateIPs = () => {
		console.log(1);
		ComponentService.getAllIP().then(
			(response) => {
				setAllIP(response.data.List);
				console.log('allIP ok ', allIP);
				setLoading(false);
			},
			(error) => {
				console.log('allIP error ', error);
				setLoading(false);
			}
		);
	};

	const HandleLoading = () => {
		setLoading(true);
		UpdateIPs();
	};

	!allIP && UpdateIPs();
	const columns = useMemo(
		() => [
			{
				header: 'IP',
				accessorKey: '1',
				size: 1,
				enableGrouping: false,
			},
			{
				header: 'MAC',
				accessorKey: '0',
				size: 1,
				enableGrouping: false,
			},
		],
		[]
	);

	return (
		<div>
			{loading ? (
				<Container>
					<div className="loader-container">
						<div className="spinner"></div>
					</div>
				</Container>
			) : (
				<Container>
					<div>
						<div className="grid-15-85">
							<Button value="Update IPs" onClick={HandleLoading} size="l" />
							<div className="form--title">All IPs in network</div>
						</div>
						<Container>
							{allIP ? (
								<TableContainer component={Paper} sx={{}}>
									<Table aria-label="caption table">
										<TableHead>
											<TableRow>
												<TableCell align="right">IP</TableCell>
												<TableCell align="right">MAC</TableCell>
											</TableRow>
										</TableHead>
										<TableBody>
											{allIP.map((row) => (
												<TableRow key={row[0]}>
													<TableCell align="right">{row[0]}</TableCell>
													<TableCell align="right">{row[1]}</TableCell>
												</TableRow>
											))}
										</TableBody>
									</Table>
								</TableContainer>
							) : (
								'No Data'
							)}
						</Container>
					</div>
				</Container>
			)}
		</div>
	);
}
