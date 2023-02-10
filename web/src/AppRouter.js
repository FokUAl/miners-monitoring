import React from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

import Home from './pages/Home'
import AddDevice from './pages/AddDevice'
import Grid from './pages/Grid'
import SignIn from './pages/Auth/SignIn'
import SignUp from './pages/Auth/SignUp'

export default function AppRouter() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" exact element={<Home />} />
                <Route path="/addDevice" element={<AddDevice />} />
                <Route path="/grid" element={<Grid />} />
                <Route path="/auth/signIn" element={<SignIn />} />
                <Route path="/auth/signUp" element={<SignUp />} />
            </Routes>
        </BrowserRouter>
    )
}