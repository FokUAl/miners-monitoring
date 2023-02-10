import React from 'react'
import { Link } from 'react-router-dom'
import './navbar.scss'
import {ReactComponent as Logo} from '../assets/images/logo_black.svg'

export default function Navbar() {
    return (
        <nav>
            <Link to="/"><Logo /></Link>
            <div className="nav--user">
                <div className="nav--user-nickname">Nickname</div>
                <div className="nav--user-role">Operator</div>
            </div>
            <div className="nav--links">
                <Link to="/addDevice">Add new device</Link>
                <Link to="/grid">Devices grid</Link>
                <Link to="/">Log out</Link>
            </div>
            
        </nav>
    )
}