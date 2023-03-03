import React from 'react'
import Navbar from '../components/Navbar/Navbar'
import AddDeviceForm from '../components/AddDeviceForm/AddDeviceForm'
import AllIP from '../components/AllIP/AllIP'

export default function AddDevice() {
    return (
        <>
            <Navbar />
            <AddDeviceForm />
            <AllIP />
        </>
    )
}