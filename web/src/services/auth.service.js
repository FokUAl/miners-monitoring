import axios from 'axios'

const API_URL = "http://localhost:8008/auth/"

const signUp = async(email, nickname, password, role) => {
  return axios.post(API_URL + 'sign-up', {
    email,
    nickname,
    password,
    role
  })
}

const signIn = async(nickname, password) => {
  return axios.post(API_URL + 'sign-in', {
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

const getCurrentRole = () => {
  return JSON.parse(localStorage.getItem('role'))
}

const AuthService = {
  signUp,
  signIn,
  logout,
  getCurrentUser,
  getCurrentRole
};

export default AuthService
