import { useState, useEffect } from 'react';
import ComponentService from '../../services/component.service';
import Container from '../Container/Container';
import Button from '../Button/Button';
import './allIP.scss';

export default function AllIP({ allIP, setAllIP }) {
	const [loading, setLoading] = useState(true);

	const UpdateIPs = () => {
		console.log(1)
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

	(!allIP && UpdateIPs())

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
					<div className="grid-10-90">
						<Button value="Update IPs" onClick={HandleLoading} size="l"/>
						<div className="form--title">All IPs in network</div>
					</div>
						<Container>
							<div className="grid-auto">
								{allIP ? allIP.map(IP => <div key={IP}>{IP[1]}</div>) : 'No Data'}
							</div>
						</Container>
				</Container>
			)}
		</div>
	);
}