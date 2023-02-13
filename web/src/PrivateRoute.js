import React from 'react'
import { Navigate } from 'react-router-dom'
import AuthService from './services/auth.service'

const PrivateRoute = ({ children }) => {
    if (AuthService.getCurrentUser()) {
        return children
    }

    return <Navigate to='/auth/signIn' />
}

export default PrivateRoute