import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const addDevice = (data) => {
    console.log(JSON.stringify(data))
    return axios.post(API_URL + 'add', {
        data
    }, {
        headers: authHeader()
    })
}

const FormService = {
    addDevice
}

export default FormService