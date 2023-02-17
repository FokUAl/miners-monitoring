import React, { StrictMode, useMemo } from 'react'
import PageService from '../services/page.service'
import TableContainer from './TableContainer'
import ReactTable from 'react-table-6'
import TableFilter from 'react-table-filter'
import './devicesList.scss'
import "react-table-6/react-table.css"

export default function DevicesList(props) {
    const devices = props.devices.devices
    const models = []
    const data = devices.map(el => {
        models.push(el.MinerType)
        return {
            ...el,
            Link: `http://localhost:8008/asic?shelf=${el.Shelf}&row=${el.Row}&column=${el.Column}`
        }
    })

    const modelsOptions = models.map((el, index) => {
        return (
            <option value={el} key={index}>{el}</option>
        )
    })

    console.log(modelsOptions)

    const handleClick = (props) => {
        PageService.getDevice(props).then((error) => console.log(error))
    }

    return (
        <div>
            <ReactTable data={data} filterable={true} columns={[
                {
                    Header: 'Owner',
                    accessor: 'Owner'
                }, 
                {
                    Header: 'Miner Model',
                    accessor: 'MinerType',
                    filterMethod: (filter, row) => {
                        if (filter.value === "all") {
                            return true;
                        } else {
                            return row[filter.id] === filter.value;
                        }
                    },
                    Filter: ({ filter, onChange }) => (
                        <select
                            onChange={(event) => onChange(event.target.value)}
                            style={{ width: "100%" }}
                            value={filter ? filter.value : "all"}
                        >
                            <option value="all">Show All</option>
                            {modelsOptions}
                        </select>
                    )
                },
                {
                    Header: 'Shelf',
                    accessor: 'Shelf',
                    filterable: false
                },
                {
                    Header: 'Row',
                    accessor: 'Row',
                    filterable: false
                },
                {
                    Header: 'Column',
                    accessor: 'Column',
                    filterable: false
                },
                {
                    Header: 'IP',
                    accessor: 'IPAddress',
                    filterable: false
                },
                {
                    Header: 'Status',
                    accessor: 'MinerStatus',
                    filterable: false
                },
                {
                    Header: 'Coin',
                    accessor: 'Coin',
                    filterable: false
                },
                {
                    Header: 'Link',
                    accessor: 'Link',
                    filterable: false,
                    Cell: props => <button onClick={() => handleClick(props.value)}>Link</button>
                }
            ]}/>
        </div>
    )
}