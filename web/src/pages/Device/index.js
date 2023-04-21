import React, { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import PageService from '@services/page.service'
import Navbar from '@components/Navbar/Navbar'
import DeviceInfo from './components/DeviceInfo/DeviceInfo'
import Comments from './components/Comments/Comments'
import DeviceGraph from './components/DeviceGraph/DeviceGraph'

export default function Device({isHidden, setIsHidden, role, username}) {
    const [data, setData] = useState({})
    const [seconds, setSeconds] = useState(0);
    const { search } = useLocation()

    useEffect(() => {
        PageService.getDevice(search).then(
            (response) => {
                setData(response.data)
                console.log(data)
            }, (error) => console.log("device error", error)
        )
        const interval = setInterval(() => {
          setSeconds(seconds => seconds + 1);
        }, 5000);
        return () => clearInterval(interval);
    }, [seconds])
    console.log(data)
    return (
		<div className={isHidden? "nav-hidden" : "nav-full"}>
            <Navbar isHidden={isHidden} setIsHidden={setIsHidden} role={role} username={username}/>
            <div>
                {data.Miner && <DeviceInfo data={data} />}
                {data.CharacteristicsHistory && <DeviceGraph charHistory={data.CharacteristicsHistory}/>}
                <Comments data={data} />
            </div>
        </div>
    )
}