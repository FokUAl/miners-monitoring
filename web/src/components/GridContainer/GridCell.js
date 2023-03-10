import './gridContainer.scss'

export default function GridCell({ children, id }) {
    return (
        <div className="grid--cell" id={id}>{children}</div>
    )
}

GridCell.defaultProps = {
    children: '',
    id: 0,
}