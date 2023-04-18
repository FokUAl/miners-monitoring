import React from 'react';
import { ReactComponent as Spinner } from '@assets/images/spinner.svg';
import './button.scss';

export default function Button({
	type,
	size,
	float,
	onClick,
	value,
	fluid,
	icon,
	loading,
	disabled,
	className,
}) {
	const generatorClass = () => {
		const classes = ['btn', `size-${size}`, `float-${float}`];
		if (className) classes.push(className);
		if (fluid) classes.push('fluid');
		return classes.join(' ');
	};
	return (
		<button
			type={type}
			className={generatorClass()}
			onClick={onClick}
			disabled={disabled}
		>
			{loading ? <Spinner className="btn-spinner" /> : value}
		</button>
	);
}

Button.defaultProps = {
	type: 'button',
	size: 's',
	float: 'center',
};
