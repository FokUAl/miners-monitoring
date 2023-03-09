import Container from '../Container/Container'
import GridComponent from '../GridComponent/GridComponent'
import GridCell from './GridCell'
import './gridContainer.scss'

export default function GridContainer({ devices }) {
    const GridRowBuilder = () => {
        const gridRows = []
        for (let i = 0; i < 60; i++) {
            gridRows.push(<GridCell key={i}></GridCell>)
        }
        return gridRows
    }
    const GridRows = GridRowBuilder()
    
    const GridShelfBuilder = ({ children, shelf }) => {
        const gridShelfs = []
        for (let i = 0; i < 11; i++) {
            gridShelfs.push(<GridComponent gridColumns="10" id={i+1} key={`${i}`}>{children}</GridComponent>)
        }
        return gridShelfs
    }
    return (
        <Container>
            <GridComponent gridColumns="11">
                <GridShelfBuilder >
                    {GridRows}
                </GridShelfBuilder>
            </ GridComponent>
        </Container>
    )
}