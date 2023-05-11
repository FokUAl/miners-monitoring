import { useState } from 'react'
import './tooltip.scss'

const Tooltip = ({children, content}) => {
    const [active, setActive] = useState(false)

    const showTip = () => {
        setActive(true)
    }

    const hideTip = () => {
        setActive(false)
    }

    const classesGenerator = () => {
        const classes = ['Tooltip--tip']
        if (content.Shelf === 1|| content.Shelf === 12 || content.Shelf === 23 || content.Shelf === 34 ) classes.push('right')
        if ( content.Shelf === 11 || content.Shelf === 22 || content.Shelf === 33 || content.Shelf === 44) classes.push('left')
        console.log('cl', classes)
        return classes.join(' ')
    }

    return (
        <div className="Tooltip--wrapper" onMouseEnter={showTip} onMouseLeave={hideTip}>
            {children}
            {active && (
                <div className={classesGenerator()}>
                    {`Location:${content.Shelf}-${content.Row}-${content.Column} 
                    IP:${content.IPAddress} 
                    Status:${content.MinerStatus}
                    Temperature:${content.Characteristics.Temperature}
                    TH/s:${content.Characteristics.MHSav}`}
                </div>
            )}
        </div>
    )
}

export default Tooltip