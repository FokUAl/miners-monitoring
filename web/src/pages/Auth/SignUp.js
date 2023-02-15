import React, { useState } from 'react'
import AuthService from '../../services/auth.service'

export default function SignUp() {
    const [email, setEmail] = useState()
    const [nickname, setNickname] = useState()
    const [password, setPassword] = useState()

    const handleRegister = async (e) => {
        e.preventDefault()
        AuthService.signUp(email, nickname, password).then(
            (response) => {
                console.log(response.data.message)
            }, (error) => {
                console.log(error)
                AuthService.logout()
            }
        )
    }

    return (
        <div className="container">
            <form onSubmit={handleRegister}>
                <div className="form--title">Create a new account</div>
                <div className="form--inputs">
                    <label>Email</label>
                    <input type="email" value={email} onChange={e => setEmail(e.target.value)} required />
                    <label>Nickname</label>
                    <input type="text" value={nickname} onChange={e => setNickname(e.target.value)} required />
                    <label>Password</label>
                    <input type="password" value={password} onChange={e => setPassword(e.target.value)} required />
                    <input type="submit" value="Sign up" />
                </div>
            </form>
        </div>
    )
}