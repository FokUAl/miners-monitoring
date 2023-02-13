import React from 'react'
import './dashboard.scss'

export default function Dashboard(props) {
    console.log(props.devices.devices)
    return (
        <div className="container">
            <div className="dash--title">Dashboard</div>
            <div className="dash--parametrs">
                <div>{props.devices.devices.length}/1000</div>
                <div>{props.devices.devices.length/10}%</div>
            </div>
        </div>
    )
}