import React from 'react';
import './button.scss';

export default function Button({type, size, float, onClick, value, fluid, icon, className}) {
	const generatorClass = () => {
		const classes = ['btn', `size-${size}`, `float-${float}`]
		if (className) classes.push(className)
		if (fluid) classes.push('fluid')
		return classes.join(' ')
	}
	return (
		<button
			type={type}
			className={generatorClass()}
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
