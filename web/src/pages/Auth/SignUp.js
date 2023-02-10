import React, { useState } from 'react'
import axios from 'axios'

export default function SignUp() {
    const [nickname, setNickname] = useState()
    const [password, setPassword] = useState()

    // const handleRegister = async (e) => {
    //     e.preventDefault()

    //     const response = await fetch('http://localhost:8008/auth/sign-up', {
    //         method: 'POST',
    //         headers: {'Content-Type': 'application/json'},
    //         body: JSON.stringify({
    //             nickname,
    //             password
    //         })
    //     })

    //     const content = await response.json()
    //     console.log(content)
    // }

    const handleRegister = (e) => {
        e.preventDefault()
        axios.post(`http://localhost:8008/auth/sign-up` , {nickname: nickname, password: password})
    }

    return (
        <div className="container">
            <form onSubmit={handleRegister}>
                <div className="form--title">Create a new account</div>
                <div className="form--inputs">
                    <label>Nickname</label>
                    <input name="nickname" type="text" onChange={e => setNickname(e.target.value)} />
                    <label>Password</label>
                    <input name="password" type="password" onChange={e => setPassword(e.target.value)} />
                    <input type="submit" value="Sign up" />
                </div>
            </form>
        </div>
    )
}