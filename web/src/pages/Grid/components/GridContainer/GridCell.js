import Tooltip from '../Tooltip/Tooltip';
import { Link } from 'react-router-dom';
import './gridContainer.scss';

export default function GridCell({ children, deviceChar, type }) {
	const classesGenerator = () => {
		const classes = ['grid--cell'];
		if (deviceChar && type === 'onlineMap') {
			classes.push('status-' + deviceChar.MinerStatus);
		} else if (deviceChar && type === 'heatMap') {
			if (deviceChar.Characteristics.ChipTempMax > 59) classes.push('status-normal')
			if (deviceChar.Characteristics.ChipTempMax > 90) classes.push('status-heat')
			if (deviceChar.Characteristics.ChipTempMax < 60) classes.push('status-cold')
			if (deviceChar.Characteristics.ChipTempMax === 0) classes.push('status-notfound')
		} else {
			classes.push('status-undefined');
		}
		return classes.join(' ');
	};
	const multiplyColor = () => {
		let hexColor = '#000000'
		// Удаление символа # из шестнадцатеричного значения
		hexColor = hexColor.replace('#', '');
		
		// Разделение шестнадцатеричного значения на составляющие компоненты цвета
		var red = parseInt(hexColor.substring(0, 2), 16);
		var green = parseInt(hexColor.substring(2, 4), 16);
		var blue = parseInt(hexColor.substring(4, 6), 16);
		
		// Умножение компонент цвета на коэффициент
		red = Math.round(red * deviceChar.Characteristics.ChipTempMax);
		green = Math.round(green * deviceChar.Characteristics.ChipTempMax);
		blue = Math.round(blue * deviceChar.Characteristics.ChipTempMax);
		
		// Ограничение значений компонент цвета в пределах 0-255
		red = Math.min(Math.max(red, 0), 255);
		green = Math.min(Math.max(green, 0), 255);
		blue = Math.min(Math.max(blue, 0), 255);
		
		// Преобразование компонент цвета обратно в шестнадцатеричное значение
		var multipliedHexColor = '#' + red.toString(16) + green.toString(16) + blue.toString(16);
		
		return multipliedHexColor;
	  }

	console.log('GridCell', deviceChar);
	return (
		<div>
			{deviceChar ? (
				<Link
					className="link"
					to={`/device?shelf=${deviceChar.Shelf}&row=${deviceChar.Row}&column=${deviceChar.Column}`}
				>
					<Tooltip content={deviceChar}>
						<div className={classesGenerator()} style={{backgroundColor:{multiplyColor}}}>{children}</div>
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
