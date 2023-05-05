import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import FormService from '@services/form.service';
import Button from '@components/Button/Button';
import Container from '@components/Container/Container';
import Input from '@components/Input/Input';
import './addDeviceForm.scss';

export default function AddDeviceForm({ allIP, allUsers }) {
	const [allOwners, setAllOwners] = useState('');
	const [allTypes, setAllTypes] = useState('');
	const [isLoading, setIsLoading] = useState(false);
	const [addressType, setAddressType] = useState('IP');
	const initialData = [
		{
			minerType: '',
			IP: '',
			MAC: '',
			shelf: '',
			column: '',
			row: '',
			owner: '',
			allOwners: allOwners,
			allTypes: allTypes,
		},
	];
	const [data, setData] = useState(initialData);
	const navigate = useNavigate();

	const handleChange = (index, event) => {
		const { value, name } = event.target;
		const newData = [...data];
		newData[index][name] = value;
		setData(newData);
	};

	useEffect(() => {
		const newData = data.map((data) => ({
			...data,
			allOwners: allOwners,
			allTypes: allTypes,
		}));
		setData(newData);
	}, [allOwners, allTypes]);

	const handleTypes = (event) => {
		setAllTypes(event.target.value);
	};

	const addFormField = () => {
		setData([...data, {
			minerType: allTypes || '',
			IP: '',
			MAC: '',
			shelf: '',
			column: '',
			row: '',
			owner: allOwners || '',
			allOwners: allOwners,
			allTypes: allTypes,
		}]);
	};

	const removeFormField = () => {
		if (data.length > 1) {
			const newData = [...data];
			newData.pop();
			setData(newData);
		}
	};

	const generateAllUsers = allUsers
		? allUsers.map((user) => {
				return <option value={user}>{user}</option>;
		  })
		: undefined;

	const handleAdd = async (e) => {
		e.preventDefault();
		setIsLoading(true);
		for (let i = 0; i < data.length; i++) {
			for (let j = i + 1; j < data.length; j++) {
				if (data[i]['IP'] === data[j]['IP']) {
					setIsLoading(false);
					alert('All IPs must be unique');
					return;
				}
				if (
					data[i]['shelf'] + data[i]['row'] + data[i]['column'] ===
					data[j]['shelf'] + data[j]['row'] + data[j]['column']
				) {
					setIsLoading(false);
					alert('All Locations must be unique');
					return;
				}
			}
		}

		if (addressType === 'MAC') {
			data.forEach((device) => {
				device['MAC'] = device['IP'];
				device['IP'] = '';
			});
		}

		if (allOwners) {
			for (let i = 0; i < data.length; i++) {
				data[i]['owner'] = allOwners;
			}
		}

		if (allTypes) {
			for (let i = 0; i < data.length; i++) {
				data[i]['minerType'] = allTypes;
			}
		}

		console.log(data);
		FormService.addDevice(data).then(
			(response) => {
				navigate('/');
				window.location.reload();
				setIsLoading(false);
			},
			(error) => {
				if (error) {
					console.log('Add device ', error);
					alert(error.response.data.Message);
					setIsLoading(false);
				}
			}
		);
	};

	return (
		<Container>
			<div>
				<form onSubmit={handleAdd}>
					<div className="form--title">Add new Device</div>
					<div className="form--btns">
						<Button
							type="button"
							float="left"
							value="Add"
							onClick={addFormField}
							size="m"
						/>
						<Button
							type="button"
							value="Remove"
							onClick={removeFormField}
							size="m"
							float="left"
						/>
					</div>
					<div className="form--inputs">
						<div></div>
						<input
							type="text"
							value={allTypes}
							name="allTypes"
							onChange={(e) => {
								handleTypes(e);
							}}
						/>
						<div></div>
						<div></div>
						<div></div>
						<div></div>
						{/* <input
							type="text"
							value={allOwners}
							name="allOwners"
							onChange={(e) => {
								handleOwners(e);
							}}
						/> */}
						<select
							name="allOwners"
							className="input--select size-m width-fluid color-primary"
							value={allOwners}
							onChange={(e) => setAllOwners(e.target.value)}
						>
							<option value=""></option>
							{generateAllUsers}
						</select>
					</div>
					<div className="form--labels">
						<label></label>
						<label>Miner Type</label>
						<div onChange={(event) => setAddressType(event.target.value)}>
							<input
								type="radio"
								value="IP"
								name="IP"
								checked={addressType === 'IP'}
							/>
							<label>IP/MAC</label>
							<input
								type="radio"
								value="MAC"
								name="MAC"
								checked={addressType === 'MAC'}
							/>
						</div>
						<label>Shelf</label>
						<label>Row</label>
						<label>Column</label>
						<label>Owner</label>
					</div>
					{data.map((data, index) => (
						<div className="form--inputs" key={index}>
							<label>{index + 1}</label>
							<input
								type="text"
								name="minerType"
								value={data.allTypes ? data.allTypes : data.minerType}
								disabled={data.allTypes}
								onChange={(e) => handleChange(index, e)}
								required
							/>
							<input
								type="text"
								name="IP"
								value={data.IP}
								onChange={(e) => handleChange(index, e)}
								pattern={
									addressType === 'IP'
										? '^([0-9]{1,3}).([0-9]{1,3}).([0-9]{1,3}).([0-9]{1,3})$'
										: addressType === 'MAC'
										? '^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$'
										: ''
								}
								required
							/>
							<input
								type="number"
								name="shelf"
								value={data.shelf}
								onChange={(e) => handleChange(index, e)}
								min="1"
								max="44"
								required
							/>
							<input
								type="number"
								name="row"
								value={data.row}
								onChange={(e) => handleChange(index, e)}
								min="1"
								max="14"
								required
							/>
							<input
								type="number"
								name="column"
								value={data.column}
								onChange={(e) => handleChange(index, e)}
								min="1"
								max="10"
								required
							/>
							<select
								name="owner"
								value={data.allOwners ? data.allOwners : data.owner}
								disabled={data.allOwners}
								className={`input--select size-m width-fluid color-primary ${data.allOwners ? 'disabled' : ''}`}
								onChange={(e) => handleChange(index, e)}
							>
								<option value="" disabled></option>
								{generateAllUsers}
							</select>
						</div>
					))}
					<div className="form--btn-submit">
						<Button
							type="submit"
							value="Add devices"
							loading={isLoading}
							disabled={isLoading}
							size="l"
							float="left"
						/>
					</div>
				</form>
			</div>
		</Container>
	);
}
