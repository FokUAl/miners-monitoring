import React, { useEffect, useState } from 'react';
import { Routes, Route, useLocation, useNavigate } from 'react-router-dom';
import Home from '@pages/Home';
import AddDevice from '@pages/AddDevice';
import Grid from '@pages/Grid';
import SignIn from '@pages/Auth/SignIn';
import SignUp from '@pages/Auth/SignUp';
import Chat from '@pages/Chat';
import Unauthorized from '@pages/Unauthorized';
import AuthService from '@services/auth.service';
import PrivateRoute from './routes/PrivateRoute';
import OperatorRoute from './routes/OperatorRoute';
import AdminRoute from './routes/AdminRoute';
import Device from '@pages/Device';
import '@scss/app.scss';
import PageService from '@services/page.service';
import jwt_decode from 'jwt-decode';

function App() {
	const [username, setUsername] = useState();
	const [role, setRole] = useState();
	const [isHidden, setIsHidden] = useState(true);
	const location = useLocation();
	const navigation = useNavigate();

	console.log('role', role, typeof role, role === 'admin');

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

	useEffect(() => {
		const token = AuthService.getCurrentUser();
		if (token) {
			const decodedToken = jwt_decode(token);
			const currentTime = new Date();
			if (decodedToken.exp * 1000 < currentTime.getTime()) {
				console.log('Token is expired');
				AuthService.logout();
				navigation('/auth/signIn');
			}
		} else {
			console.log('There is no token');
			AuthService.logout();
			navigation('/auth/signIn');
		}
	}, [location.pathname]);

	return (
		<div className="App">
			<Routes>
				<Route path="/auth/signIn" element={<SignIn />} />
				<Route
					path="/auth/signUp"
					element={
						<PrivateRoute>
							<AdminRoute role={role}>
								<SignUp
									isHidden={isHidden}
									setIsHidden={setIsHidden}
									username={username}
									actionRole={role}
								/>
							</AdminRoute>
						</PrivateRoute>
					}
				/>

				<Route
					path="/"
					element={
						<PrivateRoute role={role}>
							<Home
								isHidden={isHidden}
								setIsHidden={setIsHidden}
								username={username}
								role={role}
							/>
						</PrivateRoute>
					}
				/>
				<Route
					path="/addDevice"
					element={
						<PrivateRoute>
							<OperatorRoute role={role}>
								<AddDevice
									isHidden={isHidden}
									setIsHidden={setIsHidden}
									username={username}
									role={role}
								/>
							</OperatorRoute>
						</PrivateRoute>
					}
				/>
				<Route
					path="/grid"
					element={
						<PrivateRoute>
							<OperatorRoute role={role}>
								<Grid
									isHidden={isHidden}
									setIsHidden={setIsHidden}
									username={username}
									role={role}
								/>
							</OperatorRoute>
						</PrivateRoute>
					}
				/>
				<Route
					path="/device"
					element={
						<PrivateRoute>
							<OperatorRoute role={role}>
								<Device
									isHidden={isHidden}
									setIsHidden={setIsHidden}
									username={username}
									role={role}
								/>
							</OperatorRoute>
						</PrivateRoute>
					}
				/>
				<Route
					path="/chat"
					element={
						<PrivateRoute>
							<OperatorRoute role={role}>
								<Chat
									isHidden={isHidden}
									setIsHidden={setIsHidden}
									username={username}
									role={role}
								/>
							</OperatorRoute>
						</PrivateRoute>
					}
				/>
				<Route path="/unauthorized" element={<Unauthorized />} />
			</Routes>
		</div>
	);
}

export default App;
