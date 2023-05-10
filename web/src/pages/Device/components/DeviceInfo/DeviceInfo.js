import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Container from '@components/Container/Container';
import DataDisplay from './DataDisplay';
import Button from '@components/Button/Button';
import DataEdit from './DataEdit';
import FormService from '@services/form.service';
import Input from '@components/Input/Input';
import './deviceInfo.scss';

export default function DeviceInfo({ data }) {
	if (data.Miner.Characteristics) {
		data = {
			...data,
			ChipTempAvg: data.Miner.Characteristics.ChipTempAvg,
			ChipTempMax: data.Miner.Characteristics.ChipTempMax,
			ChipTempMin: data.Miner.Characteristics.ChipTempMin,
			Elapsed: data.Miner.Characteristics.Elapsed,
			FanSpeed1: data.Miner.Characteristics.FanSpeed1,
			FanSpeed2: data.Miner.Characteristics.FanSpeed2,
			FanSpeed3: data.Miner.Characteristics.FanSpeed3,
			FanSpeed4: data.Miner.Characteristics.FanSpeed4,
			MAC: data.Miner.MACAddress,
			MHSav: data.Miner.Characteristics.MHSav,
			PowerMode: data.Miner.Characteristics.PowerMode,
			Temperature: data.Miner.Characteristics.Temperature,
			IPAddress: data.Miner.IPAddress,
		};
	}

	const navigate = useNavigate();

	const [isEdit, setIsEdit] = useState(false);

	const [minerType, setMinerType] = useState(data.Miner.MinerType);
	const [shelf, setShelf] = useState(data.Miner.Shelf);
	const [row, setRow] = useState(data.Miner.Row);
	const [column, setColumn] = useState(data.Miner.Column);
	const [owner, setOwner] = useState(data.Miner.Owner);

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
			data.IPAddress,
			data.MAC
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

	return (
		<Container paddingRight>
			<div
				className="grid-50-50"
				style={{ columnGap: '20px', paddingInline: '20px' }}
			>
				<div>
					<div className="device-info--label-1 float-left">
						Main Characteristics
					</div>
					<DataDisplay text={'Time of work'} type="time" data={data.Elapsed} />
					<div className="device-info--label-2 float-left">Temperature</div>
					<DataDisplay
						text={'Current Temperature'}
						data={data.Temperature && data.Temperature + '°'}
					/>
					<DataDisplay
						text={'Chips Temperature Average'}
						data={data.ChipTempAvg.toFixed() && data.ChipTempAvg.toFixed() + '°'}
					/>
					<DataDisplay
						text={'Chips Temperature Max'}
						data={data.ChipTempMax && data.ChipTempMax + '°'}
					/>
					<DataDisplay
						text={'Chips Temperature Min'}
						data={data.ChipTempMin && data.ChipTempMin + '°'}
					/>
					<div className="device-info--label-2 float-left">Fans Speed</div>
					<DataDisplay text={'Fan Speed 1'} data={data.FanSpeed1} />
					<DataDisplay text={'Fan Speed 2'} data={data.FanSpeed2} />
					<DataDisplay text={'Fan Speed 3'} data={data.FanSpeed3} />
					<DataDisplay text={'Fan Speed 4'} data={data.FanSpeed4} />
				</div>
				<div>
					<div className="device-info--label-1 float-left">
						Additional Characteristics
					</div>
					<DataDisplay text={'IP address'} data={data.IPAddress} />
					<DataDisplay text={'MAC address'} data={data.MAC} />
					<DataDisplay text={'Power Mode'} data={data.PowerMode} />
					<DataDisplay
						text={'Mega Hashrate per second Average'}
						data={data.MHSav}
					/>
					<Container borderTop borderBottom>
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
									<div className="device--edit-btns">
										<Button
											value="Cancel"
											size="m"
											onClick={() => setIsEdit(false)}
										/>
										<Button value="Send" size="m" type="submit" />
									</div>
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
									float="center"
									onClick={handleButtonEdit}
								></Button>
							</>
						)}
					</Container>
					<Button
						value="Delete Device"
						size="m"
						float="center"
						onClick={() => handleDelete(data.IPAddress)}
					/>
				</div>
			</div>
		</Container>
	);
}
