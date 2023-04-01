import Tooltip from '../Tooltip/Tooltip';
import { Link } from 'react-router-dom'
import './gridContainer.scss';

export default function GridCell({ children, deviceChar }) {
	const classesGenerator = () => {
		const classes = ['grid--cell'];
		if (deviceChar) {
			classes.push('status-' + deviceChar.MinerStatus);
		} else {
			classes.push('status-undefined');
		}
		return classes.join(' ');
	};

	console.log('GridCell', deviceChar);
	return (
		<div>
			{deviceChar ? (
                <Link className='link' to={`/device?shelf=${deviceChar.Shelf}&row=${deviceChar.Row}&column=${deviceChar.Column}`}>
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
};
