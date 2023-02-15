import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Navbar from '../components/Navbar'
import Dashboard from '../components/Dashboard'
import DevicesList from '../components/DevicesList'
import PageService from '../services/page.service'
import AuthService from '../services/auth.service'

export default function Home() {
    const [ username, setUsername ] = useState()
    const [ role, setRole ] = useState()
    const [ devices, setDevices ] = useState([])
    const navigate = useNavigate()
    useEffect(() => {
        PageService.getHome().then(
            (response) => {
                setUsername(response.data.User.username)
                setRole(response.data.User.role)
                setDevices(response.data.Devices)
            }, (error) => {
                console.log(error)
                AuthService.logout()
                navigate('/auth/signIn')
            }
        )
    }, [])
    return (
        <>
            <Navbar username={username} role={role} />
            <Dashboard devices={{devices}} />
            <DevicesList devices={{devices}} />
        </>
    )
}