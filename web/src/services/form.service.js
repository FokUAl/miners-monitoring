import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const addDevice = (IP, shelf, column, row) => {
    return axios.post(API_URL + '/add-device', {
        headers: authHeader(),
        IP,
        shelf,
        column,
        row
    })
}

const FormService = {
    addDevice
}

export default FormService