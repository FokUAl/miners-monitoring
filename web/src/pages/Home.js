import React, { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import Dashboard from '../components/Dashboard'
import DevicesList from '../components/DevicesList'
import PageService from '../services/page.service'

export default function Home() {
    const [content, setContent] = useState()
    useEffect(() => {
        PageService.getHome().then(
            (response) => {
                setContent(response.data)
            },
            (error) => {
                console.loe(error)
            }
        )
    }, [])
    console.log(content)
    return (
        <>
            <Navbar />
            <Dashboard />
            <DevicesList />
        </>
    )
}