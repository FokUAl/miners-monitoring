import { useState, useEffect } from 'react';
import Container from '../Container/Container';
import DataDisplay from './DataDisplay';
import Button from '../Button/Button';
import DataEdit from './DataEdit'
import FormService from '../../services/form.service';
import './deviceInfo.scss';

export default function DeviceInfo({ data }) {
	if (data.Characteristics) {
		data = {
			...data,
			ChipTempAvg: data.Characteristics.ChipTempAvg,
			ChipTempMax: data.Characteristics.ChipTempMax,
			ChipTempMin: data.Characteristics.ChipTempMin,
			Elapsed: data.Characteristics.Elapsed,
			FanSpeedIn: data.Characteristics.FanSpeedIn,
			FanSpeedOut: data.Characteristics.FanSpeedOut,
			MAC: data.Characteristics.MAC,
			MHSav: data.Characteristics.MHSav,
			PowerMode: data.Characteristics.PowerMode,
			Temperature: data.Characteristics.Temperature,
		};
	}

	const [isEdit, setIsEdit] = useState(false);

	const [minerType, setMinerType] = useState(data.MinerType);
	const [shelf, setShelf] = useState(data.Shelf);
	const [row, setRow] = useState(data.Row);
	const [column, setColumn] = useState(data.Column);
	const [owner, setOwner] = useState(data.Owner)

	const handleEdit = async(e) => {
		e.preventDefault();
		FormService.editDevice(minerType, shelf, row, column, owner).then(
		    response => {
		        window.location.reload()
		    },
		    error => { if (error) {
		        console.log('Add device ', error)
		        // alert(error.response.data.Message)
		    }}
		)
	}

	return (
		<Container>
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
						data={data.ChipTempAvg && data.ChipTempAvg + '°'}
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
					<DataDisplay text={'Fan Speed In'} data={data.FanSpeedIn} />
					<DataDisplay text={'Fan Speed Out'} data={data.FanSpeedOut} />
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
					<Container borderTop borderRight borderBottom borderLeft>
						{isEdit ? (
                            <>
                                <div className="device-info--label-1 float-left">
                                    Location and Miner Type
                                </div>
								<form onSubmit={handleEdit}>
									<DataEdit text="Miner Type" value={minerType} width='l' setValue={setMinerType}/>
									<DataEdit text="Owner" value={owner} width='l' setValue={setOwner}/>
									<DataEdit text="Shelf" value={shelf} width='l' setValue={setShelf}/>
									<DataEdit text="Row" value={row} width='l' setValue={setRow}/>
									<DataEdit text="Column" value={column} width='l' setValue={setColumn}/>
									<Button 
										value="Cancel"
										size="m"
										onClick={() => setIsEdit(false)}
									/>
									<Button 
										value="OK"
										size="m"
										type="submit"
										onClick={() => setIsEdit(false)}
									/>
								</form>
                            </>
						) : (
							<>
								<div className="device-info--label-1 float-left">
									Location and Miner Type
								</div>
								<DataDisplay text={'Miner Type'} data={data.MinerType} />
								<DataDisplay text={'Owner'} data={data.Owner} />
								<DataDisplay text={'Shelf'} data={data.Shelf} />
								<DataDisplay text={'Row'} data={data.Row} />
								<DataDisplay text={'Column'} data={data.Column} />
								<Button
									value="Edit"
									size="m"
									onClick={() => setIsEdit(true)}
								></Button>
							</>
						)}
					</Container>
				</div>
			</div>
		</Container>
	);
}
