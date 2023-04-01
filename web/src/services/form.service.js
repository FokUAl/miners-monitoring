import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const addDevice = (data) => {
    return axios.post(API_URL + 'add', {
        data
    }, {
        headers: authHeader()
    })
}

const editDevice = (minerType, shelf, row, column, owner, IP) => {
    console.log(minerType, shelf, row, column, owner, IP)
    return axios.post(API_URL + 'update-asic-info', {
        minerType, shelf, row, column, owner, IP
    }, {
        headers: authHeader()
    })
}

const deleteDevice = (IP) => {
    return axios.post(API_URL + 'delete-device', {
        IP
    }, {
        headers: authHeader()
    })
}

const addComment = (IP, comment) => {
    console.log(IP, comment)
    return axios.post(API_URL + 'comment-device', {
        IP, comment
    }, {
        headers: authHeader()
    })
}

const FormService = {
    addDevice,
    editDevice,
    deleteDevice,
    addComment
}

export default FormService