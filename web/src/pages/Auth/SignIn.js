import React, { useRef, useState, useEffect, useContext } from 'react'
import AuthContext from '../../contextAPI/AuthProvider'
import axios from 'axios'

export default function SignIn() {
    const userRef = useRef();
    const errRef = useRef();
    const [user, setUser] = useState("");
    const [pwd, setPwd] = useState("");
    const [errMsg, setErrMsg] = useState("");
    const [success, setSuccess] = useState(false);
    useEffect(() => {
        userRef.current.focus();
    }, []);
    useEffect(() => {
        setErrMsg("");
    }, [user, pwd]);

    const { setAuth } = useContext(AuthContext);
    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post(
                'http://localhost:8008/auth/sign-in',
                JSON.stringify({ user, pwd }),
                {
                    headers: { "Content-Type": "application/json" },
                    withCredentials: true,
                }
            );
            const accessToken = response?.data?.accessToken;
            const roles = response?.data?.roles;
            setAuth({ user, pwd, roles, accessToken });
            setUser("");
            setPwd("");
            setSuccess(true);
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
                    <input name="nickname" type="text" ref={userRef} value={user} onChange={(e) => setUser(e.target.value)}/>
                    <label>Password</label>
                    <input name="password" type="password" value={pwd} onChange={(e) => setPwd(e.target.value)}/>
                    <input type="submit" value="Sign in" />
                </div>
            </form>
        </div>
    )
}