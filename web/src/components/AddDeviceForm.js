import React from 'react'

export default function AddDeviceForm() {
    return (
        <div className="container">
            <form>
                <div class="form--title">Add new Device</div>
                <div class="form--inputs">
                    <select name="model">
                        <option value=""></option>
                    </select>
                    <div className="connections">
                        <input name="connection" type="radio" />
                        <input name="connection" type="radio" />
                    </div>
                    <input type="text" />
                    <input type="text" />
                    <input type="text" />
                    <input type="text" />
                    <input type="text" />
                </div>
            </form>
        </div>
    )
}