import React from 'react'
import './devicesList.scss'

export default function DevicesList() {
    return (
        <div className="container">
            <div className="list--title">ASICs list</div>
            <table className="list">
                <tr>
                    <td>Miner Type</td>
                    <td>Shelf</td>
                    <td>Row</td>
                    <td>Column</td>
                    <td>IP Address</td>
                    <td>MAC Address</td>
                    <td>Status</td>
                    <td>Owner</td>
                    <td>Coin</td>
                </tr>
                <tr>
                    <td>Miner Type</td>
                    <td>Shelf</td>
                    <td>Row</td>
                    <td>Column</td>
                    <td>IP Address</td>
                    <td>MAC Address</td>
                    <td>Status</td>
                    <td>Owner</td>
                    <td>Coin</td>
                </tr>
            </table>
        </div>
    )
}