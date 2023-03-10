import './gridComponent.scss'

export default function GridComponent({ children, gridColumns, gridRaw }) {
    return (
        <div className={`grid grid-columns-${gridColumns} grid-raw-${gridRaw}`}>{children}</div>
    )
}

GridComponent.defaultProps = {
    children: undefined,
    gridColumns: 1,
    gridRaw: 1
}
