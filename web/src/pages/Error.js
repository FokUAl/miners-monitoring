import { useRouteError } from 'react-reouter-dom'

export default function Error() {
    const error = useRouteError()
    console.error(error)

    return (
        <div id="error-page">
            <h1>Something gone wrong</h1>
            <p>
                <i>
                {error.statusText || error.message }    
                </i>
            </p>
        </div>
    )
}