import axios from 'axios';
import authHeader from './auth.header';

const API_URL = 'http://localhost:8008/';

const getAllIP = async () => {
	return axios.get(API_URL + 'find-asic-ip', {
		timeout: 1000 * 30,
		headers: authHeader(),
	});
};

const getAllMessages = async () => {
	return axios.get(
		API_URL + 'read-user-messages',
		{ headers: authHeader() }
	);
};

const sendMessage = async (message, receiver) => {
	return axios.post(
		API_URL + 'send-message',
		{ message, receiver },
		{ headers: authHeader() }
	);
};

const postLog = async (IP) => {
	return axios.post(
		API_URL + 'get-kernel-log',
		{ IP },
		{ headers: authHeader()}
	)
    .then((response) => {
        return response.data.KernelLog
    })
}

const ComponentService = {
	getAllIP,
	getAllMessages,
	sendMessage,
	postLog
};

export default ComponentService;
