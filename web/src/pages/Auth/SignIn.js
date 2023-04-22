import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import AuthService from '@services/auth.service'
import Button from '@components/Button/Button'
import Input from '@components/Input/Input'
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
        <div className="auth-in">
            <form onSubmit={handleLogin}>
                <div className="form--label auth--label">Sign in to your account</div>
                <div className="form--inputs auth--form">
					<label>Nickname</label>
					<Input
						name="nickname"
						type="text"
						size="m"
						value={nickname}
						setValue={setNickname}
					/>
					<label>Password</label>
					<Input
						name="password"
						type="password"
						size="m"
						value={password}
						setValue={setPassword}
					/>
                    <Button type="submit" value="Sign in" className="btn--less"/>
                </div>
            </form>
        </div>
    )
}