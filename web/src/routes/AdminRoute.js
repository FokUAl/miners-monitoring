import React from 'react'
import { Navigate } from 'react-router-dom'

const PrivateRoute = ({ children, role }) => {
    if (role === 'Admin') {
        return children
    }

    return <Navigate to='/unauthorized'/>
}

export default PrivateRoute