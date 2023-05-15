import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const getHome = async () => {
    return axios.get(API_URL + 'home', {timeout: 3000, headers: authHeader()})
}

const getDevice = (props) => {
    return axios.get(API_URL + 'asic' + props, {headers: authHeader()})
}

const userInfo = async () => {
    return axios.get(API_URL + 'user-info', {headers: authHeader()})
}

const getAllUsers = () => {
    return axios.get(API_URL + 'get-all-users', {headers: authHeader()})
}

const getNotifications = () => {
    return axios.get(API_URL + 'get-senders', {headers: authHeader()})
}

const postDialog = (dialog, username) => {
    return axios.post(API_URL + 'read-messages-from', {dialog, username},{headers: authHeader()})
}

const PageService = {
    getHome,
    getDevice,
    userInfo,
    getAllUsers,
    getNotifications,
    postDialog
}

export default PageService