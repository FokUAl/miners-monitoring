import { useContext } from 'react'
import AuthContext from '../contextAPI/AuthProvider'

const useAuth = () => {
    return useContext(AuthContext)
}

export default useAuth;