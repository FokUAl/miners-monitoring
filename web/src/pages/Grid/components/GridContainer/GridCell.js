import Tooltip from '../Tooltip/Tooltip';
import { Link } from 'react-router-dom';
import './gridContainer.scss';

export default function GridCell({ children, deviceChar, type }) {
	const classesGenerator = () => {
		const classes = ['grid--cell'];
		if (deviceChar && type === 'onlineMap') {
			classes.push('status-' + deviceChar.MinerStatus);
		} else {
			classes.push('status-undefined');
		}
		return classes.join(' ');
	};

	const number = deviceChar ? deviceChar.Characteristics.ChipTempMax / 100 : 0;

	const getColor = (value) => {
		var hue = ((1 - value) * 270).toString(10);
		return ['hsl(', hue, ',100%,50%)'].join('');
	};

	const style = {
		background: getColor(number),
		minWidth: '10px',
		minHeight: '20px',
	};

	console.log('GridCell', type);
	return (
		<div>
			{deviceChar ? (
				<Link
					className="link"
					to={`/device?shelf=${deviceChar.Shelf}&row=${deviceChar.Row}&column=${deviceChar.Column}`}
				>
					<Tooltip content={deviceChar}>
						{type === 'onlineMap' ? (
							<div className={classesGenerator()}>{children}</div>
						) : (
							<div style={style}>{children}</div>
						)}
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
	type: 'onlineMap',
};
