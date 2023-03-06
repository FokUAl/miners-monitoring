import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Navbar from '../components/Navbar/Navbar'
import Dashboard from '../components/Dashboard/Dashboard'
import DevicesList from '../components/DevicesList/DevicesList'
import PageService from '../services/page.service'
import Container from '../components/Container/Container'

export default function Home() {
    const [ devices, setDevices ] = useState()
    useEffect(() => {
        PageService.getHome().then(
            (response) => {
                setDevices(response.data.Devices)
                console.log('home ok ', devices)
            }, (error) => {
                console.log('home error: ', error)
            }
        )
    }, [])
    return (
        <>
            <Navbar />
                {devices ? 
                    <div className="grid-hor">
                        <Dashboard devices={{devices}} />
                        <DevicesList devices={{devices}} />
                    </div>
                : <div className="grid-hor">
                    <div/>
                    <div/>
                </div>
                }
        </>
    )
}