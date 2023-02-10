import React from 'react'
import './dashboard.scss'

export default function Dashboard() {
    return (
        <div className="container">
            <div className="dash--title">Dashboard</div>
            <div className="dash--parametrs">
                <div>0/100</div>
                <div>0%</div>
            </div>
        </div>
    )
}