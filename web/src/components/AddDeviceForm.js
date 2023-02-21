import React, { useState } from 'react'
import FormService from '../services/form.service'
import Button from './Button/Button'
import './addDeviceForm.scss'

export default function AddDeviceForm() {
    const [ IP, setIP ] = useState()
    const [ shelf, setShelf ] = useState()
    const [ column, setColumn ] = useState()
    const [ row, setRow ] = useState()
    const [ owner, setOwner ] = useState()

    const handleAdd = async(e) => {
        e.preventDefault()
        FormService.addDevice(IP, shelf, column, row, owner).then(
            (error) => { console.log(error)}
        )
    }

    return (
        <div className="container">
            <form onSubmit={handleAdd}>
                <div className="form--title">Add new Device</div>
                <div className="form--inputs">
                    <label>IP</label>
                    <input type="text" value={IP} onChange={e => setIP(e.target.value)} required/>
                    <label>Shelf</label>
                    <input type="text" value={shelf} onChange={e => setShelf(e.target.value)} required/>
                    <label>Column</label>
                    <input type="text" value={column} onChange={e => setColumn(e.target.value)} required/>
                    <label>Row</label>
                    <input type="text" value={row} onChange={e => setRow(e.target.value)} required/>
                    <label>Owner</label>
                    <input type="text" value={owner} onChange={e => setOwner(e.target.value)} required/>
                    <Button type="submit" value="Add" className="btn--less"/>
                </div>
            </form>
        </div>
    )
}