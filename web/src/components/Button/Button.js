import React from 'react';
import './button.scss';

export default function Button({type, size, float, onClick, value, className}) {
	return (
		<button
			type={type}
			className={`btn size-${size} float-${float} ${className}`}
			onClick={onClick}
		>
			{value}
		</button>
	);
}

Button.defaultProps = {
	type: 'button',
	size: 's',
	float: 'center'
};
