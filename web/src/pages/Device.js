import React, { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import PageService from '../services/page.service'
import Navbar from '../components/Navbar'

export default function Device() {
    const [data, setData] = useState({})
    const { search } = useLocation()

    useEffect(() => {
        PageService.getDevice(search).then(
            (response) => {
                setData(response.data)
                console.log('device ok ', response.data)
            }, (error) => console.log("device error", error)
        )
    }, [])
    return (
        <div>
            <Navbar />
            <div>{data.MinerType}</div>
        </div>
    )
}