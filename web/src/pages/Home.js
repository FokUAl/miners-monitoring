import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Navbar from '../components/Navbar'
import Dashboard from '../components/Dashboard'
import DevicesList from '../components/DevicesList'
import PageService from '../services/page.service'

export default function Home() {
    const [ devices, setDevices ] = useState([])
    useEffect(() => {
        PageService.getHome().then(
            (response) => {
                setDevices(response.data.Devices)
                console.log('home ok ')
            }, (error) => {
                console.log(error)
            }
        )
    }, [])
    return (
        <>
            <Navbar />
            <Dashboard devices={{devices}} />
            <DevicesList devices={{devices}} />
        </>
    )
}