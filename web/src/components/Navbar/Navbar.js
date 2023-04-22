import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './navbar.scss';
import { ReactComponent as Logo } from '@assets/images/logo_white.svg';
import { ReactComponent as LogoLess } from '@assets/images/logo_white_less.svg';
import AuthService from '@services/auth.service';
import PageService from '@services/page.service';
import Button from '@components/Button/Button';
import {
	BsListNested,
	BsList,
	BsPlusCircle,
	BsGrid3X3,
	BsBoxArrowRight,
	BsBoxArrowInRight,
	BsChatLeftText
} from 'react-icons/bs';

export default function Navbar({ isHidden, setIsHidden, username, role }) {
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
					{isHidden ? <LogoLess /> : <Logo className="nav--logo-img" />}
				</Link>
			</div>
			<ul className="nav--links">
				{(role === 'Operator' || role === 'Admin') && (
					<li>
						<Link className="nav--link" to="/addDevice">
							<BsPlusCircle color="white" size="25" className="icon" />
							{!isHidden && <div className="m-lt">Add new device</div>}
						</Link>
					</li>
				)}
				{(role === 'Operator' || role === 'Admin') && (
					<li>
						<Link className="nav--link" to="/grid">
							<BsGrid3X3 color="white" size="25" className="icon" />
							{!isHidden && <div className="m-lt">Devices grid</div>}
						</Link>
					</li>
				)}
				{(role === 'Admin') && (
					<li>
						<Link className="nav--link" to="/auth/signUp">
							<BsBoxArrowInRight color="white" size="25" className="icon" />
							{!isHidden && <div className="m-lt">Register new user</div>}
						</Link>
					</li>
				)}
				{(role === 'Operator' || role === 'Admin') && (
					<li>
						<Link className="nav--link" to="/chat">
							<BsChatLeftText color="white" size="25" className="icon" />
							{!isHidden && <div className="m-lt">Chat</div>}
						</Link>
					</li>
				)}
				<li>
					<Link className="nav--link" to="/auth/signIn" onClick={handleLogOut}>
						<BsBoxArrowRight color="white" size="25" className="icon" />
						{!isHidden && <div className="m-lt">Log out</div>}
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
