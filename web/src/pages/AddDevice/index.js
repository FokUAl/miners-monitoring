import React, { useEffect, useState } from 'react'
import Navbar from '../../components/Navbar/Navbar'
import AddDeviceForm from './components/AddDeviceForm/AddDeviceForm'
import PageService from '../../services/page.service'
import AllIP from './components/AllIP/AllIP'

export default function AddDevice({isHidden, setIsHidden, role, username}) {
    const [allIP, setAllIP] = useState()
    const [allUsers, setAllUsers] = useState()
    useEffect(() => {
        PageService.getAllUsers().then((response)=> {
            setAllUsers(response.data.AllUsers)
            console.log('get all users', allUsers)
        }, (error) => {
            console.log('get all users', error)
        })
    }, [])
    return (
		<div className={isHidden? "nav-hidden" : "nav-full"}>
            <Navbar isHidden={isHidden} setIsHidden={setIsHidden} role={role} username={username}/>
            <div className="grid-70-30 column-gap-10">
                <AddDeviceForm allIP={allIP} allUsers={allUsers} />
                <AllIP setAllIP={setAllIP} allIP={allIP} />
            </div>
        </div>
    )
}