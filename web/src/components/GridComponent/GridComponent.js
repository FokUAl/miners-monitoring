import './gridComponent.scss'

export default function GridComponent({ children, gridColumns, id }) {
    return (
        <div className={`grid grid-columns-${gridColumns}`} id={id}>{children}</div>
    )
}

GridComponent.defaultProps = {
    children: undefined,
    gridColumns: 1,
    id: undefined
}
