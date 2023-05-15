import React from 'react'
import { Navigate } from 'react-router-dom'
import AuthService from '@services/auth.service'

const PrivateRoute = ({ children, role }) => {
    if (AuthService.getCurrentRole() === 'Admin') {
        return children
    } 

    return <Navigate to='/unauthorized'/>
}

export default PrivateRoute