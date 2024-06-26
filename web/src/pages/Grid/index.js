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
		<div className={isHidden? "nav-hidden" : "nav-full"}>
            <Navbar isHidden={isHidden} setIsHidden={setIsHidden} role={role} username={username}/>
            <GridContainer devices={devices}/>
        </div>
    )
}