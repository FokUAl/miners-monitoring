import { useState, useEffect } from 'react'
import ComponentService from '../../services/component.service'

export default function AllIP () {
    const [ data, setData ] = useState()

    useEffect(() => {
        ComponentService.getAllIP().then(
            (response) => {
                console.log(response)
                setData(response.data.List)
                console.log('allIP ok ');
            },
            (error) => {
                console.log('allIP error', error);
            }
        );
    }, []);

    // const dataArr = data.map(el => {
    //     return (
    //         <div>{el}</div>
    //     )
    // })    

    return (
        <div>
            {/* {dataArr} */}
            {data}
        </div>
    )
}