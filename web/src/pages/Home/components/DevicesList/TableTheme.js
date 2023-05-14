import { createTheme } from '@mui/material';

const TableTheme = () => {
	const tableTheme = createTheme({
		components: {
			MuiTableRow: {
				styleOverrides: {
					root: {
						'&.MuiTableRow-root:hover': {
							td: {
								backgroundColor: '#3d3d3d'
							}
						}
					}
				}
			},
			MuiButtonBase: {
				styleOverrides: {
					root: {
						'& *': {
							color: 'white',
						},
					},
				},
			},
			MuiInputBase: {
				styleOverrides: {
					input: {
						backgroundColor: '#2f2f2f',
						color: 'white',
						'&:focus': {
							backgroundColor: '#454545',
							color: 'white',
						},
					},
                    root: {
						'&.Mui-focused': {
							backgroundColor: '#454545',
						},
                    }
				},
			},
		},
	});

	return tableTheme;
};

export default TableTheme();
