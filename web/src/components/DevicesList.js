import React, { useMemo } from 'react'
import PageService from '../services/page.service'
import TableContainer from './TableContainer'
import ReactTable from 'react-table-6'
// import TableFilter from 'react-table-filter'
import './devicesList.scss'
import "react-table-6/react-table.css"

export default function DevicesList(props) {
    const devices = props.devices.devices
    const data = devices.map(el => {
        return {
            ...el,
            Link: `http://localhost:8008/asic?shelf=${el.Shelf}&row=${el.Row}&column=${el.Column}`
        }
    })

    const handleClick = (props) => {
        PageService.getDevice(props).then((error) => console.log(error))
    }
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
            <ReactTable data={data}  columns={[{
                Header: 'Owner',
                accessor: 'Owner'
            }, {
                Header: 'Miner Model',
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
            },
            {
                Header: 'Link',
                accessor: 'Link',
                Cell: props => <button onClick={() => handleClick(props.value)}>Link</button>
            }]}/>
        </div>
    )
}