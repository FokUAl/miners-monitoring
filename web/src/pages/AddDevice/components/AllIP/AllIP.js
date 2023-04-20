import { useState, useEffect } from 'react';
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
	ThemeProvider,
	createTheme,
} from '@mui/material';
import Input from '@components/Input/Input';
import { ReactComponent as Spinner } from '@assets/images/spinner.svg';
import UbuntuRegular from '@assets/fonts/Ubuntu-Regular.ttf';
import './allIP.scss';

export default function AllIP({ allIP, setAllIP }) {
	const [onSearch, setOnSearch] = useState('');
	const [searched, setSearched] = useState('');
	const [loading, setLoading] = useState(true);

	const handleSearch = (searchedVal) => {
		setSearched(searchedVal);
		const filteredRows = allIP.filter(
			(IP) =>
				IP[0].toLowerCase().includes(searchedVal.toLowerCase()) ||
				IP[1].toLowerCase().includes(searchedVal.toLowerCase())
		);
		setOnSearch(filteredRows);
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

	const HandleLoading = () => {
		setLoading(true);
		UpdateIPs();
	};

	!allIP && UpdateIPs();

	const theme = createTheme({
		typography: {
			fontFamily: 'Ubuntu, Arial',
		},
		components: {
			MuiCssBaseline: {
				styleOverrides: `
					@font-face {
					font-family: 'Ubuntu';
					font-style: normal;
					font-display: swap;
					font-weight: 400;
					src: local('Ubuntu'), local('Ubuntu-Regular'), url(${UbuntuRegular}) format('ttf');
					unicodeRange: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF;
					}
				`,
			},
		},
	});

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
					<Container borderLeft borderRight>
						{allIP ? (
							<div>
								<div className="m-bm">
									<Input
										value={searched}
										setValue={handleSearch}
										size="m"
										placeholder="search"
									/>
								</div>
								<ThemeProvider theme={theme}>
									<TableContainer component={Paper} sx={{}}>
										<Table aria-label="caption table">
											<TableHead>
												<TableRow>
													<TableCell align="right">IP</TableCell>
													<TableCell align="left">MAC</TableCell>
												</TableRow>
											</TableHead>
											<TableBody>
												{onSearch
													? onSearch.map((row) => (
															<TableRow key={row[0]}>
																<TableCell align="right">{row[0]}</TableCell>
																<TableCell align="left">{row[1]}</TableCell>
															</TableRow>
													  ))
													: allIP.map((row) => (
															<TableRow key={row[0]}>
																<TableCell align="right">{row[0]}</TableCell>
																<TableCell align="left">{row[1]}</TableCell>
															</TableRow>
													  ))}
											</TableBody>
										</Table>
									</TableContainer>
								</ThemeProvider>
							</div>
						) : (
							'No Data'
						)}
					</Container>
				</Container>
			)}
		</div>
	);
}
