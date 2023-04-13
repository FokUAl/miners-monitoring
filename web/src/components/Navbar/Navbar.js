import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './navbar.scss';
import { ReactComponent as Logo } from '@assets/images/logo_white.svg';
import AuthService from '@services/auth.service';
import PageService from '@services/page.service';

export default function Navbar(props) {
	const [username, setUsername] = useState();
	const [role, setRole] = useState();
	useEffect(() => {
		PageService.userInfo().then(
			(response) => {
				setUsername(response.data.username);
				setRole(response.data.role);
				console.log('navbar ok ');
			},
			(error) => {
				console.log('navbar error', error);
			}
		);
	}, []);

	const handleLogOut = () => {
		AuthService.logout();
	};

	return (
		<nav>
			<Link to="/">
				<Logo />
			</Link>
			<div className="nav--user">
				<div className="nav--user-nickname">{username}</div>
				<div className="nav--user-role">{role}</div>
			</div>
			<ul className="nav--links">
				<li>
					<Link className="nav--link" to="/addDevice">
						Add new Device
					</Link>
				</li>
				<li>
					<Link className="nav--link" to="/grid">
						Devices grid
					</Link>
				</li>
				<li>
					<Link className="nav--link" to="/auth/signIn" onClick={handleLogOut}>
						Log out
					</Link>
				</li>
			</ul>
		</nav>
	);
}
