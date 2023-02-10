import React from 'react'
import Navbar from '../components/Navbar'
import Dashboard from '../components/Dashboard'
import DevicesList from '../components/DevicesList'

export default function Home() {
    return (
        <>
            <Navbar />
            <Dashboard />
            <DevicesList />
        </>
    )
}