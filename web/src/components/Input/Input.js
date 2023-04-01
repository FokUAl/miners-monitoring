import './input.scss'

const Input = ({
	type,
	name,
	value,
	setValue,
	placeholder,
	required,
	size,
	width,
	textAlign,
	label,
	color
}) => {
	return (
		<div className="input--wrapper">
			{label && <label>{label}</label>}
			<input
				className={`input size-${size} width-${width} text-align-${textAlign} color-${color}`}
				type={type}
				name={name}
				value={value}
				onChange={e => {setValue(e.target.value)}}
				placeholder={placeholder}
				required={required}
				textAlign={textAlign}
			/>
		</div>
	);
};

Input.defaultProps = {
	type: 'text',
	name: '',
	value: '',
	setValue: undefined,
	placeholder: '',
	required: false,
	size: 's',
	width: 'fluid',
	textAlign: 'left',
	label: undefined,
	color: 'primary'	
};

export default Input;
