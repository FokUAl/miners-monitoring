import Container from '@components/Container/Container'
import Button from '@components/Button/Button'
import { useNavigate } from 'react-router-dom'

const Unauthorized = () => {
    const navigate = useNavigate()
    const handleHome = () => {
        navigate('/');
        window.location.reload();
    }
    return (
        <Container>
            <div className="m-bm">
                Unauthorized access
            </div>
            <Button value="To home page" size='m' onClick={handleHome}></Button>
        </Container>
    )
}

export default Unauthorized