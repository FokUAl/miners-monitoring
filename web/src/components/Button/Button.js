import React from 'react';
import './button.scss';

export default function Button(props) {
	return (
		<button
			type={props.type}
			className={`btn size-${props.size} float-${props.float}`}
			onClick={props.onClick}
		>
			{props.value}
		</button>
	);
}

Button.defaultProps = {
	type: 'button',
	size: 's',
	float: 'center'
};
