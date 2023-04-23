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

const sendMessage = async (message) => {
	return axios.post(
		API_URL + 'send-message',
		{ message },
		{ headers: authHeader() }
	);
};

const ComponentService = {
	getAllIP,
	getAllMessages,
	sendMessage,
};

export default ComponentService;
