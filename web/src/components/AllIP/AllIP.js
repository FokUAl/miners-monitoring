import { useState, useEffect } from 'react';
import ComponentService from '../../services/component.service';
import Container from '../Container/Container';
import './allIP.scss';

export default function AllIP() {
	const [data, setData] = useState();
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		ComponentService.getAllIP().then(
			(response) => {
				console.log(response);
				setData(response.data.List);
				console.log('allIP ok ');
				setLoading(true);
			},
			(error) => {
				console.log('allIP error', error);
			}
		);
	}, []);

	// const dataArr = data.map(el => {
	//     return (
	//         <div>{el}</div>
	//     )
	// })

	return (
		<div>
			{loading ? (
				<Container>
					<div className="loader-container">
						<div className="spinner"></div>
					</div>
				</Container>
			) : (
				<Container>{data}</Container>
			)}
		</div>
	);
}
