import Tooltip from '../Tooltip/Tooltip';
import { Link } from 'react-router-dom';
import './gridContainer.scss';

export default function GridCell({ children, deviceChar, type }) {
	const classesGenerator = () => {
		const classes = ['grid--cell'];
		if (deviceChar && type === 'onlineMap') {
			classes.push('status-' + deviceChar.MinerStatus);
		} else if (deviceChar && type === 'heatMap') {
			if (deviceChar.Characteristics.ChipTempMax > 69) classes.push('status-normal')
			if (deviceChar.Characteristics.ChipTempMax > 80) classes.push('status-heat')
			if (deviceChar.Characteristics.ChipTempMax < 70) classes.push('status-cold')
			if (deviceChar.Characteristics.ChipTempMax === 0) classes.push('status-notfound')
		} else {
			classes.push('status-undefined');
		}
		return classes.join(' ');
	};

	console.log('GridCell', deviceChar);
	return (
		<div>
			{deviceChar ? (
				<Link
					className="link"
					to={`/device?shelf=${deviceChar.Shelf}&row=${deviceChar.Row}&column=${deviceChar.Column}`}
				>
					<Tooltip content={deviceChar}>
						<div className={classesGenerator()}>{children}</div>
					</Tooltip>
				</Link>
			) : (
				<div className={classesGenerator()}>{children}</div>
			)}
		</div>
	);
}
GridCell.defaultProps = {
	children: '',
	id: 0,
	deviceChar: undefined,
	type: 'onlineMap'
};
