import React from 'react'
import './Container.scss'

export default function Container({ children, borderTop, borderRight, borderBottom, borderLeft }) {
    const generateClasses = () => {
        const classes = ['container']
        if (borderTop) classes.push('border-top')
        if (borderRight) classes.push('border-right')
        if (borderBottom) classes.push('border-bottom')
        if (borderLeft) classes.push('border-left')
        return classes.join(' ')
    }
    return (
        <div className={generateClasses()}>{children}</div>
    )
}

Container.defaultProps = {
    children: <></>,
    borderTop: false,
    borderRight: false,
    borderBottom: false,
    borderLeft: false
}