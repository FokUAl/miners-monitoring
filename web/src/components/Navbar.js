import React from 'react'
import { Link } from 'react-router-dom'
import './navbar.scss'
import {ReactComponent as Logo} from '../assets/images/logo_black.svg'
import AuthService from '../services/auth.service'

export default function Navbar() {
    const handleLogOut = () => {
        AuthService.logout()
    }

    return (
        <nav>
            <Link to="/"><Logo /></Link>
            <div className="nav--user">
                <div className="nav--user-nickname">{localStorage.getItem('nickname')}</div>
                <div className="nav--user-role">Operator</div>
            </div>
            <div className="nav--links">
                <Link to="/addDevice">Add new device</Link>
                <Link to="/grid">Devices grid</Link>
                <Link to="/auth/signIn" onClick={handleLogOut}>Log out</Link>
            </div>
            
        </nav>
    )
}