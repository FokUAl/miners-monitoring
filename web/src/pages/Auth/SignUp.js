import React, { useState } from 'react'
import axios from 'axios'

export default function SignUp() {
    const [email, setEmail] = useState()
    const [nickname, setNickname] = useState()
    const [password, setPassword] = useState()

    const handleRegister = async (e) => {
        e.preventDefault()
        await axios.post("http://localhost:8008/auth/sign-up", {email, nickname, password})
            .then((response) => {console.log('success') 
            console.log(response)})
            .catch((exception) => console.log(exception))
    }

    return (
        <div className="container">
            <form onSubmit={handleRegister}>
                <div className="form--title">Create a new account</div>
                <div className="form--inputs">
                    <label>Email</label>
                    <input name="email" type="email" value={email} onChange={e => setEmail(e.target.value)} required />
                    <label>Nickname</label>
                    <input name="nickname" type="text" value={nickname} onChange={e => setNickname(e.target.value)} required />
                    <label>Password</label>
                    <input name="password" type="password" value={password} onChange={e => setPassword(e.target.value)} required />
                    <input type="submit" value="Sign up" />
                </div>
            </form>
        </div>
    )
}