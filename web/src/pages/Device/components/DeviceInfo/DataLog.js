import { useState, useEffect } from 'react'
import ComponentService from '@services/component.service'
import Container from '@components/Container/Container'

const DataLog = ({ IP }) => {
    const [log, setLog] = useState()
    useEffect(() => {
		// PageService.getDevice(search).then(
		// 	(response) => {
		// 		setData(response.data);
		// 		console.log('device ok', data);
		// 	},
		// 	(error) => console.log('device error', error)
		// );
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
            {log}
        </Container>
    )
}

export default DataLog