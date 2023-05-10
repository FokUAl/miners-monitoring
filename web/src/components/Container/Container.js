import React from 'react'
import './Container.scss'

export default function Container({ children, borderTop, borderRight, borderBottom, borderLeft, paddingTop, paddingRight, paddingBottom, paddingLeft, verticalHeight, overflowY, className }) {
    const generateClasses = () => {
        const classes = ['container']
        if (borderTop) classes.push('border-top')
        if (borderRight) classes.push('border-right')
        if (borderBottom) classes.push('border-bottom')
        if (borderLeft) classes.push('border-left')
        if (paddingTop) classes.push('padding-top')
        if (paddingRight) classes.push('padding-right')
        if (paddingBottom) classes.push('padding-bottom')
        if (paddingLeft) classes.push('padding-left')
        if (verticalHeight) classes.push('vertical-height')
        if (overflowY) classes.push('overflow-y')
        if (className) classes.push({className})
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