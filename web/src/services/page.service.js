import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const getHome = () => {
    return axios.get(API_URL + 'home', {headers: authHeader()})
}

const getDevice = (props) => {
    return axios.get(props, {headers: authHeader()})
}

const PageService = {
    getHome,
    getDevice
}

export default PageService