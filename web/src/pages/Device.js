import React, { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import PageService from '../services/page.service'
import Navbar from '../components/Navbar/Navbar'
import Container from '../components/Container/Container'

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
    }, [data])

    return (
        <div>
            <Navbar />
            <Container>
                <div className="grid-50-50">
                    <div className="grid-hor">
                        <div>Time of work</div>
                        <div>{data.Elapsed} s.</div>
                    </div>
                    <div className="grid-hor">
                        <div>Other</div>
                        <div className="grid-50-50">
                            <div>MAC Address</div>
                            <div>{data.MAC}</div>
                            <div>PowerMode</div>
                            <div>{data.PowerMode}</div>
                        </div>
                    </div>
                    <div className="grid-hor">
                        <div>Average Mega Hashrate</div>
                        <div>{data.MHSav}</div>
                        <div>Temperature</div>
                        <div className="grid-50-50">
                            <div>Average Temperature </div>
                            <div>{data.ChipTempAvg}</div>
                            <div>Max Temperature </div>
                            <div>{data.ChipTempMax}</div>
                            <div>Min Temperature </div>
                            <div>{data.ChipTempMin}</div>
                        </div>
                        <div>Fan speeds</div>
                        <div className="grid-50-50">
                            <div>Fan speed In</div>
                            <div>{data.FanSpeedIn}</div>
                            <div>Fan speed Out</div>
                            <div>{data.FanSpeedOut}</div>
                        </div>
                    </div>
                </div>
            </Container>
        </div>
    )
}