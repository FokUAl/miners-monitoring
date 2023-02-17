import React, { useState, useEffect } from 'react'
import { Routes, Route } from 'react-router-dom' 
import Home from './pages/Home'
import AddDevice from './pages/AddDevice'
import Grid from './pages/Grid'
import SignIn from './pages/Auth/SignIn'
import SignUp from './pages/Auth/SignUp'
import AuthService from './services/auth.service'
import PrivateRoute from './PrivateRoute'
import "./scss/app.scss"

function App() {
  const [currentUser, setCurrentUser] = useState(undefined)

  useEffect(() => {
    const user = AuthService.getCurrentUser()

    if(user) {
      setCurrentUser(user)
    }

  }, [])

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
      </Routes>
    </div>
  );
}

export default App;
