import { useState, useEffect } from 'react'
import ComponentService from '@services/component.service'
import Container from '@components/Container/Container'

const DataLog = ({ IP }) => {
    const [log, setLog] = useState()
    useEffect(() => {
        ComponentService.postLog(IP).then(
            (response) => {
                setLog(response)
            },
            (error) => {
                console.log('DataLog', error)
                setLog('No logs')
            }
        )
    }, [])

    return (
        <Container>
            <pre className="data-log--container">
                {log}
            </pre>
        </Container>
    )
}

export default DataLog