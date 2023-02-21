import React from 'react'
import './button.scss'

export default function Button(props) {
    return (
        <button 
            type={props.type} 
            className={props.className ? props.className : 'btn'}
            onClick={props.onClick}
        >{props.value}</button>
    )
}