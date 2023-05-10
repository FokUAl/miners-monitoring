import { createTheme } from '@mui/material';

const TableTheme = () => {
	const tableTheme = createTheme({
		components: {
			MuiTableRow: {
				styleOverrides: {
					root: {
						'&:hover': {
							backgroundColor: 'black',
						},
					},
				},
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
						backgroundColor: 'black',
					},
                    root: {
						'&.Mui-focused': {
							backgroundColor: 'black',
						},
                    }
				},
			},
		},
	});

	return tableTheme;
};

export default TableTheme();
