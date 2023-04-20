import React, { useState, useEffect } from 'react'
import GridContainer from './components/GridContainer/GridContainer'
import Navbar from '../../components/Navbar/Navbar'
import PageService from '../../services/page.service';

export default function Grid({isHidden, setIsHidden, role, username}) {
	const [devices, setDevices] = useState([]);
	useEffect(() => {
		PageService.getHome().then(
			(response) => {
				setDevices(response.data.Devices);
				console.log('grid ok ');
			},
			(error) => {
				console.log('grid error: ', error);
			}
		);
	}, []);
    return (
		<div className={isHidden? "nav-hidden" : "nav-full"} role={role} username={username}>
            <Navbar isHidden={isHidden} setIsHidden={setIsHidden}/>
            <GridContainer devices={devices}/>
        </div>
    )
}