import React, { useState } from 'react'
import Navbar from '../components/Navbar/Navbar'
import AddDeviceForm from '../components/AddDeviceForm/AddDeviceForm'
import AllIP from '../components/AllIP/AllIP'

export default function AddDevice() {
    const [ allIP, setAllIP] = useState()
    return (
        <>
            <Navbar />
            <div className="grid-hor">
                <AddDeviceForm allIP={allIP} />
                <AllIP setAllIP={setAllIP} allIP={allIP} />
            </div>
        </>
    )
}