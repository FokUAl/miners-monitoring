import React, { useState, useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import './navbar.scss'
import { ReactComponent as Logo } from '../assets/images/logo_black.svg'
import AuthService from '../services/auth.service'
import PageService from '../services/page.service'
import Button from './Button/Button'


export default function Navbar(props) {
    const [username, setUsername] = useState()
    const [role, setRole] = useState()
    useEffect(() => {
        PageService.getHome().then(
            (response) => {
                setUsername(response.data.User.username)
                setRole(response.data.User.role)
                console.log('navbar ok ')
            }, (error) => {
                console.log("navbar error", error)
            }
        )
    }, [])

    const handleLogOut = () => {
        AuthService.logout()
    }
    
    return (
        <nav>
            <Link to="/"><Logo /></Link>
            <div className="nav--user">
                <div className="nav--user-nickname">{username}</div>
                <div className="nav--user-role">{role}</div>
            </div>
            <div className="nav--links">
                <Link className="link" to="/addDevice"><Button value='Add new device'></Button></Link>
                <Link className="link" to="/grid"><Button value='Devices grid'></Button></Link>
                <Link className="link" to="/auth/signIn" onClick={handleLogOut}><Button value='Log out'></Button></Link>
            </div>
        </nav>
    )
}