import Input from '@components/Input/Input'

const DataEdit = ({text, value, width, setValue}) => {
    return (
        <>
            <div className="grid-50-50">
                <div className="float-left">{text}</div>
                <div className="float-right">
                    <Input size='s' width={width} color='secondary' textAlign='right' value={value} setValue={setValue}/>
                </div>
            </div>
            <div className="data-display-wrapper" />
        </>
    )
}

DataEdit.defaultProps = {
    text: '',
}

export default DataEdit