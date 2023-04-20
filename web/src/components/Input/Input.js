import './input.scss';

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
	color,
	options,
	pattern
}) => {
	const optionsTag = options
		? options.map((option) => {
				return <option value={option}>{option}</option>;
		  })
		: undefined;
	return (
		<div className="input--wrapper">
			{label && <label>{label}</label>}
			{type === 'select' ? (
				<select
					className={`input--select size-${size} width-${width} text-align-${textAlign} color-${color}`}
					onChange={(e) => setValue(e.target.value)}
				>
					{optionsTag}
				</select>
			) : (
				<input
					className={`input size-${size} width-${width} text-align-${textAlign} color-${color}`}
					type={type}
					name={name}
					value={value}
					onChange={(e) => {
						setValue(e.target.value);
					}}
					placeholder={placeholder}
					required={required}
					pattern={pattern}
				/>
			)}
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
	color: 'primary',
	pattern: false
};

export default Input;
