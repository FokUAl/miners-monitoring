import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const getHome = () => {
    return axios.get(API_URL + 'home', {header: authHeader()})
}

const PageService = {
    getHome
}

export default PageService