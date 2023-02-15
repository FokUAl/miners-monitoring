import axios from 'axios'

const API_URL = "http://localhost:8008/auth/"

const signUp = async(nickname, email, password) => {
  return axios.post(API_URL + 'sign-up', {
    nickname,
    email,
    password
  })
}

const signIn = async(nickname, password) => {
  return axios
    .post(API_URL + 'sign-in', {
      nickname,
      password,
    })
    .then((response) => {
        localStorage.setItem('token', JSON.stringify(response.data.Token))
        return response.data
    })
}

const logout = () => {
  localStorage.removeItem('token')
}

const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem('token'))
}

const AuthService = {
  signUp,
  signIn,
  logout,
  getCurrentUser,
};

export default AuthService
