import React, { useState, useEffect } from 'react';
import Navbar from '@components/Navbar/Navbar';
import Dashboard from './components/Dashboard/Dashboard';
import DevicesList from './components/DevicesList/DevicesList';
import Empty from './components/Empty/Empty';
import PageService from '@services/page.service';
import PopupChat from '@components/PopupChat/PopupChat';

export default function Home({ isHidden, setIsHidden, role, username }) {
	const [devices, setDevices] = useState();
	useEffect(() => {
		PageService.getHome().then(
			(response) => {
				setDevices(response.data.Devices);
				console.log('home ok ', response.data.Devices);
			},
			(error) => {
				console.log('home error: ', error);
			}
		);
	}, []);

	return (
		<div>
			<div className={isHidden ? 'nav-hidden' : 'nav-full'}>
				<Navbar
					isHidden={isHidden}
					setIsHidden={setIsHidden}
					role={role}
					username={username}
				/>
				{devices ? (
					<div className="grid-hor">
						<Dashboard devices={devices} />
						{(role === 'Admin' || role === 'Operator') && (
							<DevicesList devices={devices} setDevices={setDevices} />
						)}
					</div>
				) : (
					<div className="grid-hor">
						<Empty />
					</div>
				)}
			</div>
			{role === 'User' && <PopupChat username={username} />}
		</div>
	);
}
