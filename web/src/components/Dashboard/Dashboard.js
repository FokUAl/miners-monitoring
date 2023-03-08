import React from 'react'
import Container from '../Container/Container'
import './dashboard.scss'

export default function Dashboard({ devices }) {
    return (
        <Container>
            <div className="dash--title">Dashboard</div>
            <div className="dash--parametrs">
                <div>{devices.length}/1000</div>
                <div>{devices.length/10}%</div>
            </div>
        </Container>
    )
}