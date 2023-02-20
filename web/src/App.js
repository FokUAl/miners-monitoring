import React, { useEffect } from 'react'
import { Routes, Route, useLocation, useNavigate } from 'react-router-dom' 
import Home from './pages/Home'
import AddDevice from './pages/AddDevice'
import Grid from './pages/Grid'
import SignIn from './pages/Auth/SignIn'
import SignUp from './pages/Auth/SignUp'
import AuthService from './services/auth.service'
import PrivateRoute from './PrivateRoute'
import Device from './pages/Device'
import "./scss/app.scss"
import PageService from './services/page.service'

function App() {
  const location = useLocation()
  const navigation = useNavigate()
  useEffect(() => {
    PageService.getHome().catch((error) => {
      if (error) {
        console.log('Auth error ', error)
        AuthService.logout()
        navigation('/auth/signIn')
      }
    })
  }, [location.pathname])
  return (
    <div className="App">
      <Routes>
        <Route 
          path="/auth/signIn" 
          element={<SignIn />} 
        />
        <Route 
          path="/auth/signUp" 
          element={<SignUp />} 
        />

        <Route 
          path='/' 
          element={
            <PrivateRoute><Home /></PrivateRoute>
          } 
        />
        <Route 
          path='/addDevice' 
          element={
            <PrivateRoute><AddDevice /></PrivateRoute>
          } 
        />
        <Route 
          path='/grid' 
          element={
            <PrivateRoute><Grid /></PrivateRoute>
          } 
        />
        <Route 
          path='/device' 
          element={
            <PrivateRoute><Device /></PrivateRoute>
          } 
        />
      </Routes>
    </div>
  );
}

export default App;
