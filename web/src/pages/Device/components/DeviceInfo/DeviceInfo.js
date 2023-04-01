import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Container from '../../../../components/Container/Container';
import DataDisplay from './DataDisplay';
import Button from '../../../../components/Button/Button';
import DataEdit from './DataEdit';
import FormService from '../../../../services/form.service';
import Input from '../../../../components/Input/Input';
import './deviceInfo.scss';

export default function DeviceInfo({ data }) {
	if (data.Miner.Characteristics) {
		data = {
			...data,
			ChipTempAvg: data.Miner.Characteristics.ChipTempAvg,
			ChipTempMax: data.Miner.Characteristics.ChipTempMax,
			ChipTempMin: data.Miner.Characteristics.ChipTempMin,
			Elapsed: data.Miner.Characteristics.Elapsed,
			FanSpeedIn: data.Miner.Characteristics.FanSpeedIn,
			FanSpeedOut: data.Miner.Characteristics.FanSpeedOut,
			MAC: data.Miner.Characteristics.MAC,
			MHSav: data.Miner.Characteristics.MHSav,
			PowerMode: data.Miner.Characteristics.PowerMode,
			Temperature: data.Miner.Characteristics.Temperature,
		};
	}

	const navigate = useNavigate();

	const [isEdit, setIsEdit] = useState(false);

	const [minerType, setMinerType] = useState(data.Miner.MinerType);
	const [shelf, setShelf] = useState(data.Miner.Shelf);
	const [row, setRow] = useState(data.Miner.Row);
	const [column, setColumn] = useState(data.Miner.Column);
	const [owner, setOwner] = useState(data.Miner.Owner);

	const [comment, setComment] = useState('');

	const handleButtonEdit = () => {
		setMinerType(data.Miner.MinerType);
		setShelf(data.Miner.Shelf.toString());
		setRow(data.Miner.Row.toString());
		setColumn(data.Miner.Column.toString());
		setOwner(data.Miner.Owner);
		setIsEdit(true);
	};

	const handleEdit = async (e) => {
		e.preventDefault();
		FormService.editDevice(
			minerType,
			shelf,
			row,
			column,
			owner,
			data.Miner.IPAddress
		).then(
			(response) => {
				navigate(`/device?shelf=${shelf}&row=${row}&column=${column}`);
				window.location.reload();
			},
			(error) => {
				console.log('Add device ', error);
				alert(error.response.data.Miner.Message);
			}
		);
	};

	const handleDelete = async (IP) => {
		FormService.deleteDevice(IP).then(
			(response) => {
				navigate('/');
				window.location.reload();
			},
			(error) => {
				console.log('Add device ', error);
			}
		);
	};

	const handleAddComm = async (e) => {
		e.preventDefault()
		FormService.addComment(data.Miner.IPAddress, comment).then(
			(response) => {
				console.log('Add comment ok');
				// navigate(`/device?shelf=${data.Miner.shelf}&row=${data.Miner.row}&column=${data.Miner.column}`)
				// window.location.reload();
				setComment('')
			},
			(error) => {
				console.log('Add comment', error);
			}
		);
	};

	return (
		<div className="grid-hor">
			<Container>
				<div
					className="grid-50-50"
					style={{ columnGap: '20px', paddingInline: '20px' }}
				>
					<div>
						<div className="device-info--label-1 float-left">
							Main Characteristics
						</div>
						<DataDisplay
							text={'Time of work'}
							type="time"
							data={data.Miner.Elapsed}
						/>
						<div className="device-info--label-2 float-left">Temperature</div>
						<DataDisplay
							text={'Current Temperature'}
							data={data.Miner.Temperature && data.Miner.Temperature + '°'}
						/>
						<DataDisplay
							text={'Chips Temperature Average'}
							data={data.Miner.ChipTempAvg && data.Miner.ChipTempAvg + '°'}
						/>
						<DataDisplay
							text={'Chips Temperature Max'}
							data={data.Miner.ChipTempMax && data.Miner.ChipTempMax + '°'}
						/>
						<DataDisplay
							text={'Chips Temperature Min'}
							data={data.Miner.ChipTempMin && data.Miner.ChipTempMin + '°'}
						/>
						<div className="device-info--label-2 float-left">Fans Speed</div>
						<DataDisplay text={'Fan Speed In'} data={data.Miner.FanSpeedIn} />
						<DataDisplay text={'Fan Speed Out'} data={data.Miner.FanSpeedOut} />
					</div>
					<div>
						<div className="device-info--label-1 float-left">
							Additional Characteristics
						</div>
						<DataDisplay text={'IP address'} data={data.Miner.IPAddress} />
						<DataDisplay text={'MAC address'} data={data.Miner.MAC} />
						<DataDisplay text={'Power Mode'} data={data.Miner.PowerMode} />
						<DataDisplay
							text={'Mega Hashrate per second Average'}
							data={data.Miner.MHSav}
						/>
						<Container borderTop borderRight borderBottom borderLeft>
							{isEdit ? (
								<>
									<div className="device-info--label-1 float-left">
										Location and Miner Type
									</div>
									<form onSubmit={handleEdit}>
										<DataEdit
											text="Miner Type"
											value={minerType}
											width="l"
											setValue={setMinerType}
										/>
										<DataEdit
											text="Owner"
											value={owner}
											width="l"
											setValue={setOwner}
										/>
										<DataEdit
											text="Shelf"
											value={shelf}
											width="l"
											setValue={setShelf}
										/>
										<DataEdit
											text="Row"
											value={row}
											width="l"
											setValue={setRow}
										/>
										<DataEdit
											text="Column"
											value={column}
											width="l"
											setValue={setColumn}
										/>
										<Button
											value="Cancel"
											size="m"
											onClick={() => setIsEdit(false)}
										/>
										<Button value="OK" size="m" type="submit" />
									</form>
								</>
							) : (
								<>
									<div className="device-info--label-1 float-left">
										Location and Miner Type
									</div>
									<DataDisplay text={'Miner Type'} data={data.Miner.MinerType} />
									<DataDisplay text={'Owner'} data={data.Miner.Owner} />
									<DataDisplay text={'Shelf'} data={data.Miner.Shelf} />
									<DataDisplay text={'Row'} data={data.Miner.Row} />
									<DataDisplay text={'Column'} data={data.Miner.Column} />
									<Button
										value="Edit"
										size="m"
										onClick={handleButtonEdit}
									></Button>
								</>
							)}
						</Container>
						<Button
							value="Delete Device"
							size="m"
							onClick={() => handleDelete(data.Miner.IPAddress)}
						/>
					</div>
				</div>
			</Container>
			<Container>
				<div className="grid-50-50">
					<div className="grid-hor">
						Comments section
						{data.Comments}
					</div>
					<form onSubmit={handleAddComm} className="grid-85-15">
						<Input
							type="text"
							value={comment}
							setValue={setComment}
							width="fluid"
							placeholder="Type your comment here"
						/>
						<Button value="add comment" type="submit" />
					</form>
				</div>
			</Container>
		</div>
	);
}
