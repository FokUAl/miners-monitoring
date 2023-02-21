import React, { useState } from 'react'
import FormService from '../services/form.service'
import Button from './Button/Button'
import './addDeviceForm.scss'

export default function AddDeviceForm() {
    const initialData = [
        {
            IP: '',
            shelf: '',
            column: '',
            row: '',
            owner: ''
        }
    ]
    const [data, setData] = useState(initialData)

    const handleChange = (index, event) => {
        const {value, name} = event.target
        const newData = [...data]
        newData[index][name] = value
        console.log('changed')
        setData(newData)
    }

    const addFormField = () => {
        setData([...data, {
            IP: '',
            shelf: '',
            column: '',
            row: '',
            owner: ''
        }])
    }

    const removeFormField = () => {
        if (data.length > 1) {
            const newData = [...data]
            newData.pop()
            setData(newData)
        }
    }

    const handleAdd = async(e) => {
        e.preventDefault()
        FormService.addDevice(data).catch(
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
                {data.map((data, index) => (
                    <div className="form--inputs" key={index}>
                        <label>{index+1}</label>
                        <input type="text" name='IP' value={data.IP || ''} onChange={e => handleChange(index, e)} required/>
                        <input type="text" name='shelf' value={data.shelf || ''} onChange={e => handleChange(index, e)} required/>
                        <input type="text" name='column' value={data.column || ''} onChange={e => handleChange(index, e)} required/>
                        <input type="text" name='row' value={data.row || ''} onChange={e => handleChange(index, e)} required/>
                        <input type="text" name='owner' value={data.owner || ''} onChange={e => handleChange(index, e)} required/>
                    </div>
                ))}
                <div className="form--btns">
                    <Button type="submit" value="Add Devices" className="btn--less form--btn"/>
                    <Button type="button" value="Add" onClick={addFormField} className="btn--less form--btn"/>
                    <Button type="button" value="Remove" onClick={removeFormField} className="btn--less form--btn"/>
                </div>
            </form>
        </div>
    )
}