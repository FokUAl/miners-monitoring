import axios from 'axios'
import authHeader from './auth.header'

const API_URL = 'http://localhost:8008/'

const getAllIP = async () => {
    try {
        // return axios.get(API_URL + 'find-asic-ip', {timeout: 1000 * 30, headers: authHeader()})
        return axios({
            method: 'get',
            url: API_URL + 'find-asic-ip',
            timeout: 30000,
            headers: authHeader()
          })
    } catch(error) {
        console.log('allIP: ' + error)
    }
}

const ComponentService = {
    getAllIP
}

export default ComponentService