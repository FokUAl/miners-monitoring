import React, { useRef, useState, useEffect } from 'react'
import useAuth from '../../hooks/useAuth'
import { Link, useNavigate, useLocation } from 'react-router-dom'

import axios from 'axios'

export default function SignIn() {
    const { setAuth } = useAuth()

    const navigate = useNavigate()
    const location = useLocation()
    const from = location.state?.from?.pathname || '/'

    const userRef = useRef()
    const errRef = useRef()
    const [nickname, setNickname] = useState("")
    const [password, setPassword] = useState("")
    const [errMsg, setErrMsg] = useState("")
    useEffect(() => {
        userRef.current.focus()
    }, [])
    useEffect(() => {
        setErrMsg("")
    }, [nickname, password])

    const handleLogin = async (e) => {
        e.preventDefault()
        try {
            const response = await axios.post(
                'http://localhost:8008/auth/sign-in',
                JSON.stringify({ nickname, password }),
                {
                    headers: { "Content-Type": "application/json" },
                    withCredentials: true,
                }
            )
            const token = response?.data?.Token
            const role = response?.data?.role
            setAuth({ nickname, password, role, accessToken: token })
            setNickname("")
            setPassword("")
            navigate(from, { replace: true })
        } catch (err) {
            if (!err?.response) {
                setErrMsg("No Server Response");
            } else if (err.response?.status === 400) {
                setErrMsg("Missing Username or Password");
            } else if (err.response?.status === 401) {
                setErrMsg("Unauthorized");
            } else {
                setErrMsg("Login Failed");
            }
            errRef.current.focus();
        }
    };

    return (
        <div className="container">
            <p 
                ref={errRef}
                className={errMsg ? "errmsg" : "offscreen"}
            >{errMsg}</p>
            <form onSubmit={handleLogin}>
                <div className="form--inputs">Sign in to your account</div>
                <div className="form--inputs">
                    <label>Nickname</label>
                    <input name="nickname" type="text" ref={userRef} value={nickname} onChange={(e) => setNickname(e.target.value)}/>
                    <label>Password</label>
                    <input name="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
                    <input type="submit" value="Sign in" />
                </div>
            </form>
        </div>
    )
}