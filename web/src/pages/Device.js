import React, { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import PageService from '../services/page.service'
import Navbar from '../components/Navbar/Navbar'
import DeviceInfo from '../components/DeviceInfo/DeviceInfo'

export default function Device() {
    const [data, setData] = useState({})
    const { search } = useLocation()

    useEffect(() => {
        PageService.getDevice(search).then(
            (response) => {
                setData(response.data)
                console.log(data)
            }, (error) => console.log("device error", error)
        )
    }, [data])
    console.log(data)
    return (
        <div>
            <Navbar />
            <DeviceInfo data={data} />
        </div>
    )
}