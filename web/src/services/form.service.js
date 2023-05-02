import axios from 'axios';
import authHeader from './auth.header';

const API_URL = 'http://localhost:8008/';

const addDevice = (data) => {
	return axios.post(
		API_URL + 'add',
		{
			data,
		},
		{
			headers: authHeader(),
		}
	);
};

const editDevice = (minerType, shelf, row, column, owner, IP, MAC) => {
	console.log(minerType, shelf, row, column, owner, IP, MAC);
	return axios.post(
		API_URL + 'update-asic-info',
		{
			minerType,
			shelf,
			row,
			column,
			owner,
			IP,
			MAC
		},
		{
			headers: authHeader(),
		}
	);
};

const deleteDevice = (IP) => {
	return axios.post(
		API_URL + 'delete-device',
		{
			IP,
		},
		{
			headers: authHeader(),
		}
	);
};

const addComment = (IP, comment) => {
	console.log(IP, comment);
	return axios.post(
		API_URL + 'comment-device',
		{
			IP,
			comment,
		},
		{
			headers: authHeader(),
		}
	);
};

const deleteComment = (username, content, creationDate) => {
	return axios.post(
		API_URL + 'delete-comment',
		{
			username,
			content,
			creationDate,
		},
		{
			headers: authHeader(),
		}
	);
};

const editComment = (username, content, creationDate) => {
	return axios.post(
		API_URL + 'edit-comment',
		{
			username,
			content,
			creationDate,
		},
		{
			headers: authHeader(),
		}
	);
};

const FormService = {
	addDevice,
	editDevice,
	deleteDevice,
	addComment,
	deleteComment,
	editComment,
};

export default FormService;
