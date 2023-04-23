import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const getHome = async () => {
    return axios.get(API_URL + 'home', {timeout: 3000, headers: authHeader()})
}

const getDevice = (props) => {
    return axios.get(API_URL + 'asic' + props, {headers: authHeader()})
}

const userInfo = () => {
    return axios.get(API_URL + 'user-info', {timeout: 2000, headers: authHeader()})
}

const getAllUsers = () => {
    return axios.get(API_URL + 'get-all-users', {headers: authHeader()})
}

const getNotifications = () => {
    return axios.get(API_URL + 'get-senders', {headers: authHeader()})
}

const PageService = {
    getHome,
    getDevice,
    userInfo,
    getAllUsers,
    getNotifications
}

export default PageService