import React, { useState, useEffect } from 'react';
import Navbar from '@components/Navbar/Navbar';
import Dashboard from './components/Dashboard/Dashboard';
import DevicesList from './components/DevicesList/DevicesList';
import Empty from './components/Empty/Empty';
import PageService from '@services/page.service';
import PopupChat from '@components/PopupChat/PopupChat';

export default function Home({ isHidden, setIsHidden, role, username }) {
	const [devices, setDevices] = useState();
	const [time, setTime] = useState(Date.now())
	const [delay, setDelay] = useState()
	const [allUsers, setAllUsers] = useState()
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
		console.log('1')
		if (delay && delay > 0) {
			const interval = setInterval(() => {
				setTime(Date.now())
			}, delay);
			return () => clearInterval(interval);
		}
	}, [time]);

    useEffect(() => {
        PageService.getAllUsers().then((response)=> {
            setAllUsers(response.data.AllUsers)
            console.log('get all users', allUsers)
        }, (error) => {
            console.log('get all users', error)
        })
    }, [])

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
						<DevicesList devices={devices} setDevices={setDevices} delay={delay} setDelay={setDelay} allUsers={allUsers}/>
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
