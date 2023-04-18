import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './navbar.scss';
import { ReactComponent as Logo } from '@assets/images/logo_white.svg';
import { ReactComponent as LogoLess } from '@assets/images/logo_white_less.svg';
import AuthService from '@services/auth.service';
import PageService from '@services/page.service';
import Button from '@components/Button/Button';
import { BsListNested, BsList, BsPlusCircle, BsGrid3X3, BsDoorOpen } from 'react-icons/bs'

export default function Navbar({ isHidden, setIsHidden }) {
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

	const handleHidden = () => {
		setIsHidden(!isHidden);
	};

	return (
		<nav className={isHidden ? 'nav-less' : 'nav'}>
			<div className="nav--logo">
				<Link to="/">
					{isHidden ? <LogoLess /> : <Logo className='nav--logo-img' />}
				</Link>
			</div>
			<ul className="nav--links">
				<li>
					<Link className="nav--link" to="/addDevice">
						<BsPlusCircle color="white" size="25" className='icon'/>{!isHidden && <div className='m-lt'>Add new device</div>}
					</Link>
				</li>
				<li>
					<Link className="nav--link" to="/grid">
						<BsGrid3X3 color="white" size="25" className='icon'/>{!isHidden && <div className='m-lt'>Devices grid</div>}
					</Link>
				</li>
				<li>
					<Link className="nav--link" to="/auth/signIn" onClick={handleLogOut}>
						<BsDoorOpen color="white" size="25" className='icon'/>{!isHidden && <div className='m-lt'>Log out</div>}
					</Link>
				</li>
			</ul>
			{!isHidden && (
				<div className="nav--user">
					<div className="nav--user-nickname">{username}</div>
					<div className="nav--user-role">{role}</div>
				</div>
			)}
			<div className="float-bottom">
				<Button
					value={isHidden ? <BsList /> : <BsListNested />}
					float="center"
					fluid
					size="m"
					onClick={handleHidden}
				/>
			</div>
		</nav>
	);
}
