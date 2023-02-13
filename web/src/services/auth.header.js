export default function authHeader() {
    const user = JSON.parse(localStorage.getItem('token'))

    if (user && user.Token) {
        return {Authorization: user.Token}
    } else {
        return {}
    }
}