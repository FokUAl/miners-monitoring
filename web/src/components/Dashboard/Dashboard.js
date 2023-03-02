import React from 'react'
import Container from '../Container/Container'
import './dashboard.scss'

export default function Dashboard(props) {
    return (
        <Container>
            <div className="dash--title">Dashboard</div>
            <div className="dash--parametrs">
                <div>{props.devices.devices.length}/1000</div>
                <div>{props.devices.devices.length/10}%</div>
            </div>
        </Container>
    )
}