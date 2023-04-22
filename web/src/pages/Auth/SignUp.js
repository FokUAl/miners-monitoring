import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Button from '@components/Button/Button';
import Input from '@components/Input/Input';
import Navbar from '@components/Navbar/Navbar'
import AuthService from '../../services/auth.service';

export default function SignUp({ isHidden, setIsHidden, username, actionRole}) {
	const [email, setEmail] = useState();
	const [nickname, setNickname] = useState();
	const [password, setPassword] = useState();
	const [role, setRole] = useState();

	const navigate = useNavigate();
	const handleRegister = async (e) => {
		e.preventDefault();
		AuthService.signUp(email, nickname, password, role).then(
			(response) => {
				navigate('/');
				window.location.reload();
				console.log(response.data.message);
			},
			(error) => {
				console.log(error);
			}
		);
	};

	return (
		<div className={isHidden ? 'nav-hidden' : 'nav-full'}>
			<Navbar
				isHidden={isHidden}
				setIsHidden={setIsHidden}
				role={actionRole}
				username={username}
			/>
			<div className="auth-up">
				<form onSubmit={handleRegister}>
					<div className="form--label auth--label">Sign up new user</div>
					<div className="form--inputs auth--form">
						<label>Email</label>
						<Input
							name="email"
							type="text"
							size="m"
							value={email}
							setValue={setEmail}
							required
							pattern={`[a-z0-9]+@[a-z]+\.[a-z]{2,3}`}
						/>
						<label>Nickname</label>
						<Input
							name="nickname"
							type="text"
							size="m"
							value={nickname}
							setValue={setNickname}
							required
						/>
						<label>Password</label>
						<Input
							name="password"
							type="password"
							size="m"
							value={password}
							setValue={setPassword}
							required
						/>
						<label>Role</label>
						<Input
							name="role"
							type="select"
							options={['admin', 'operator', 'user']}
							setValue={setRole}
						/>
						<Button type="submit" value="Sign up" className="btn--less" />
					</div>
				</form>
			</div>
		</div>
	);
}
