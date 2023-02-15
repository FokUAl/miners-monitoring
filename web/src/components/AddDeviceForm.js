import React, { useState } from 'react'
import FormService from '../services/form.service'

export default function AddDeviceForm() {
    const [ IP, setIP ] = useState()
    const [ shelf, setShelf ] = useState()
    const [ column, setColumn ] = useState()
    const [ raw, setRaw ] = useState()

    const handleAdd = async(e) => {
        e.preventDefault()
        FormService.addDevice(IP, shelf, column, raw).then(
            (error) => { console.log(error)}
        )
    }

    return (
        <div className="container">
            <form onSubmit={handleAdd}>
                <div class="form--title">Add new Device</div>
                <div class="form--inputs">
                    <label>IP</label>
                    <input type="text" value={IP} onChange={e => setIP(e.target.value)} required/>
                    <label>Shelf</label>
                    <input type="text" value={shelf} onChange={e => setShelf(e.target.value)} required/>
                    <label>Column</label>
                    <input type="text" value={column} onChange={e => setColumn(e.target.value)} required/>
                    <label>Raw</label>
                    <input type="text" value={raw} onChange={e => setRaw(e.target.value)} required/>
                    <input type="submit" />
                </div>
            </form>
        </div>
    )
}