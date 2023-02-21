import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const addDevice = (IP, shelf, column, row, owner) => {
    return axios.post(API_URL + '/add', {
        IP,
        shelf,
        column,
        row,
        owner
    }, {
        headers: authHeader()
    })
}

const FormService = {
    addDevice
}

export default FormService