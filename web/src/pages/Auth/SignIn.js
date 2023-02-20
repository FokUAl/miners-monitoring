import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import AuthService from '../../services/auth.service'
import Button from '../../components/Button/Button'
import './signIn.scss'

export default function SignIn() {
    const [nickname, setNickname] = useState("")
    const [password, setPassword] = useState("")
    let navigate = useNavigate()

    const handleLogin = (e) => {
        e.preventDefault()
        AuthService.signIn(nickname, password).then(
            () => {
                navigate('/')
                window.location.reload()
            },
            (error) => {
                console.log(error)
                AuthService.logout()
            }
        )
    }

    return (
        <div className="container sign-in">
            <form onSubmit={handleLogin}>
                <div className="form--label sign-in--label">Sign in to your account</div>
                <div className="form--inputs sign-in--form">
                    <label>Nickname</label>
                    <input name="nickname" type="text" value={nickname} onChange={(e) => setNickname(e.target.value)}/>
                    <label>Password</label>
                    <input name="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
                    <Button type="submit" value="Sign in" className="btn--less"/>
                </div>
            </form>
        </div>
    )
}