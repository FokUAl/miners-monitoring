import React, { useState } from 'react'
import Navbar from '../../components/Navbar/Navbar'
import AddDeviceForm from './components/AddDeviceForm/AddDeviceForm'
import AllIP from './components/AllIP/AllIP'

export default function AddDevice({isHidden, setIsHidden, role, username}) {
    const [allIP, setAllIP] = useState()
    return (
		<div className={isHidden? "nav-hidden" : "nav-full"}>
            <Navbar isHidden={isHidden} setIsHidden={setIsHidden} role={role} username={username}/>
            <div className="grid-70-30 column-gap-10">
                <AddDeviceForm allIP={allIP} />
                <AllIP setAllIP={setAllIP} allIP={allIP} />
            </div>
        </div>
    )
}