import React, { useEffect } from 'react';
import { Routes, Route, useLocation, useNavigate } from 'react-router-dom';
import Home from '@pages/Home';
import AddDevice from '@pages/AddDevice';
import Grid from '@pages/Grid';
import SignIn from '@pages/Auth/SignIn';
import SignUp from '@pages/Auth/SignUp';
import AuthService from '@services/auth.service';
import PrivateRoute from './PrivateRoute';
import Device from '@pages/Device';
import '@scss/app.scss';
import PageService from '@services/page.service';
import jwt_decode from 'jwt-decode';

function App() {
	const location = useLocation();
	const navigation = useNavigate();

	useEffect(() => {
		const token = AuthService.getCurrentUser();
		if (token) {
			const decodedToken = jwt_decode(token);
      const currentTime = new Date()
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
				<Route path="/auth/signUp" element={<SignUp />} />

				<Route
					path="/"
					element={
						<PrivateRoute>
							<Home />
						</PrivateRoute>
					}
				/>
				<Route
					path="/addDevice"
					element={
						<PrivateRoute>
							<AddDevice />
						</PrivateRoute>
					}
				/>
				<Route
					path="/grid"
					element={
						<PrivateRoute>
							<Grid />
						</PrivateRoute>
					}
				/>
				<Route
					path="/device"
					element={
						<PrivateRoute>
							<Device />
						</PrivateRoute>
					}
				/>
			</Routes>
		</div>
	);
}

export default App;
