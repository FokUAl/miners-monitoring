import React, { useState } from 'react'
import Navbar from '../../components/Navbar/Navbar'
import AddDeviceForm from './components/AddDeviceForm/AddDeviceForm'
import AllIP from './components/AllIP/AllIP'

export default function AddDevice({isHidden, setIsHidden}) {
    const [allIP, setAllIP] = useState()
    return (
		<div className={isHidden? "grid-5-95" : "grid-15-85"}>
            <Navbar isHidden={isHidden} setIsHidden={setIsHidden}/>
            <div className="grid-hor">
                <AddDeviceForm allIP={allIP} />
                <AllIP setAllIP={setAllIP} allIP={allIP} />
            </div>
        </div>
    )
}