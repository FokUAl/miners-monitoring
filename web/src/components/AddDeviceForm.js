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
        FormService.addDevice(IP, shelf, column, row, owner).catch(
            (error) => { if (error) console.log('Add device ', error)}
        )
    }

    return (
        <div className="container">
            <form onSubmit={handleAdd}>
                <div className="form--title">Add new Device</div>
                <div className="form--labels">
                    <label>IP</label>
                    <label>Shelf</label>
                    <label>Column</label>
                    <label>Row</label>
                    <label>Owner</label>
                </div>
                <div className="form--inputs">
                    <input type="text" value={IP} onChange={e => setIP(e.target.value)} required/>
                    <input type="text" value={shelf} onChange={e => setShelf(e.target.value)} required/>
                    <input type="text" value={column} onChange={e => setColumn(e.target.value)} required/>
                    <input type="text" value={row} onChange={e => setRow(e.target.value)} required/>
                    <input type="text" value={owner} onChange={e => setOwner(e.target.value)} required/>
                </div>
                <Button type="submit" value="Add" className="btn--less form--btn"/>
            </form>
        </div>
    )
}