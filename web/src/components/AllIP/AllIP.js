import { useState, useEffect } from 'react';
import ComponentService from '../../services/component.service';
import Container from '../Container/Container';
import Button from '../Button/Button'
import './allIP.scss';

export default function AllIP({ allIP, setAllIP }) {
	const [loading, setLoading] = useState(true);

	const UpdateIPs = () => {	
		useEffect(() => {
			ComponentService.getAllIP().then(
				(response) => {
					console.log(response);
					setAllIP(response.data.List)
					console.log('allIP ok ');
					setLoading(false);
				},
				(error) => {
					console.log('allIP error ', error);
					setLoading(false);
				}
			);
		}, []);
	}

	const HandleLoading = () => {
		setLoading(true)
		UpdateIPs()
	}

	UpdateIPs()
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
					<div className="form--title">All IPs in network</div>
					<div className="grid-10-90">
						<Button value="Update IPs" onClick={HandleLoading} size="l" />	
						<Container><div className="grid-12">{allIP}</div></Container>
					</div>
				</Container>
			)}
		</div>
	);
}
