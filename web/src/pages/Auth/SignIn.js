import React from 'react'

export default function SignIn() {
    return (
        <div className="container">
            <form>
                <div className="form--inputs">Sign in to your account</div>
                <div className="form--inputs">
                    <label>Nickname</label>
                    <input name="nickname" type="text"/>
                    <label>Password</label>
                    <input name="password" type="password"/>
                    <input type="submit" value="Sign in" />
                </div>
            </form>
        </div>
    )
}