import React, { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import Dashboard from '../components/Dashboard'
import DevicesList from '../components/DevicesList'
import PageService from '../services/page.service'
import axios from 'axios'

export default function Home() {
    const [content, setContent] = useState()
    const token = localStorage.getItem('token')
    useEffect(() => {
        axios.get('http://localhost:8008/home', {headers: {'Authorization' : token}})
        .then(response => {console.log(response)})
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