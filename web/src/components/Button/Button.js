import React from 'react';
import './button.scss';

export default function Button(props) {
	return (
		<button
			type={props.type}
			className={`btn size-${props.size}`}
			onClick={props.onClick}
		>
			{props.value}
		</button>
	);
}

Button.defaultProps = {
	type: 'button',
	size: 's',
};
