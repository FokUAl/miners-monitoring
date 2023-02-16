import React, { useMemo } from 'react'
import TableContainer from './TableContainer'
// import TableFilter from 'react-table-filter'
import './devicesList.scss'

export default function DevicesList(props) {
    const devices = props.devices.devices
    const columns = [
        {
            Header: 'Owner',
            accessor: 'Owner'
        },
        {
            Header: 'Miner Type',
            accessor: 'MinerType'
        },
        {
            Header: 'Shelf',
            accessor: 'Shelf'
        },
        {
            Header: 'Row',
            accessor: 'Row'
        },
        {
            Header: 'Column',
            accessor: 'Column'
        },
        {
            Header: 'IP',
            accessor: 'IPAddress'
        },
        {
            Header: 'Status',
            accessor: 'MinerStatus'
        },
        {
            Header: 'Coin',
            accessor: 'Coin'
        }
    ]


    return (
        <div>
            <TableContainer columns={columns} data={devices} />
        </div>
    )
}