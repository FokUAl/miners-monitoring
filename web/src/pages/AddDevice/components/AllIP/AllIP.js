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
import Input from '@components/Input/Input'
import { ReactComponent as Spinner } from '@assets/images/spinner.svg';
import './allIP.scss';

export default function AllIP({ allIP, setAllIP }) {
	const [data, setData] = useState(allIP);
	const [searched, setSearched] = useState('');
	const [loading, setLoading] = useState(true);

	const requestSearch = (searchedVal) => {
		const filteredRows = data.filter((row) => {
			return row.name.toLowerCase().includes(searchedVal.toLowerCase());
		});
		setData(filteredRows);
	};

	const UpdateIPs = () => {
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

	const cancelSearch = () => {
		setSearched('');
		requestSearch(searched);
	};

	const HandleLoading = () => {
		setLoading(true);
		UpdateIPs();
	};

	!allIP && UpdateIPs();

	return (
		<div>
			{loading ? (
				<Container>
					<Spinner className="page-spinner" />
				</Container>
			) : (
				<Container>
					<div className="grid-15-85 m-lt">
						<Button value="Update IPs" onClick={HandleLoading} size="l" />
						<div className="form--title">All IPs in network</div>
					</div>
					<Container>
						{allIP ? (
							<Paper>
								<Input
									value={searched}
									onChange={(searchVal) => requestSearch(searchVal)}
									onCancelSearch={() => cancelSearch()}
								/>
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
							</Paper>
						) : (
							'No Data'
						)}
					</Container>
				</Container>
			)}
		</div>
	);
}
