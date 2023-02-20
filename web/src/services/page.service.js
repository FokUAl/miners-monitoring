import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const getHome = async () => {
    try {
        return axios.get(API_URL + 'home', {timeout: 2000, headers: authHeader()})
    } catch(error) {
        console.log(error)
    }
}

const getDevice = (props) => {
    return axios.get(API_URL + 'asic' + props, {headers: authHeader()})
}

const PageService = {
    getHome,
    getDevice
}

export default PageService