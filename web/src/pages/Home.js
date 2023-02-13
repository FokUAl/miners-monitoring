import React, { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import Dashboard from '../components/Dashboard'
import DevicesList from '../components/DevicesList'
import PageService from '../services/page.service'

export default function Home() {
    const [ username, setUsername ] = useState()
    const [ role, setRole ] = useState()
    const [ devices, setDevices ] = useState([])
    useEffect(() => {
        PageService.getHome().then(
            (response) => {
                setUsername(response.data.User.username)
                setRole(response.data.User.role)
                setDevices(response.data.Devices)
                console.log(response.data)
            }, (error) => {console.log(error)}
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