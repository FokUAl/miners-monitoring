import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import FormService from '../../../../services/form.service';
import Button from '../../../../components/Button/Button';
import './addDeviceForm.scss';
import Container from '../../../../components/Container/Container';

export default function AddDeviceForm({ allIP }) {
	const [allOwners, setAllOwners] = useState('');

	const initialData = [
		{
			minerType: '',
			IP: '',
			shelf: '',
			column: '',
			row: '',
			owner: '',
			allOwners: allOwners,
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
		const newData = data.map((data) => ({ ...data, allOwners: allOwners }));
		setData(newData);
	}, [allOwners]);

	const handleOwners = (event) => {
		setAllOwners(event.target.value);
		console.log(allOwners, data);
	};

	const addFormField = () => {
		setData([
			...data,
			{
				minerType: '',
				IP: '',
				shelf: '',
				column: '',
				row: '',
				owner: '',
				allOwners: allOwners,
			},
		]);
	};

	const removeFormField = () => {
		if (data.length > 1) {
			const newData = [...data];
			newData.pop();
			setData(newData);
		}
	};

	const handleAdd = async (e) => {
		e.preventDefault();
		for (let i = 0; i < data.length; i++) {
			for (let j = i + 1; j < data.length; j++) {
				if (data[i]['IP'] === data[j]['IP']) {
					alert('All IPs must be unique');
					return;
				}
				if (
					data[i]['shelf'] + data[i]['row'] + data[i]['column'] ===
					data[j]['shelf'] + data[j]['row'] + data[j]['column']
				) {
					alert('All Locations must be unique');
					return;
				}
			}
			// if (!allIP.includes(data[i]['IP'])) {
			// 	alert('There is no such IP');
			// 	return;
			// }
		}

		if (allOwners) {
			for (let i = 0; i < data.length; i++) {
				data[i]['owner'] = allOwners;
			}
		}

		console.log(data);
		FormService.addDevice(data).then(
			(response) => {
				navigate('/');
				window.location.reload();
			},
			(error) => {
				if (error) {
					console.log('Add device ', error);
					alert(error.response.data.Message);
				}
			}
		);
	};

	return (
		<Container>
			<div className="centrilized-cont">
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
						<input
							type="text"
							value={allOwners}
							name="allOwners"
							onChange={(e) => {
								handleOwners(e);
							}}
							className="form--input-allOwners"
						/>
					</div>
					<div className="form--labels">
						<label></label>
						<label>Miner Type</label>
						<label>IP/MAC</label>
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
								value={data.minerType}
								onChange={(e) => handleChange(index, e)}
								required
							/>
							<input
								type="text"
								name="IP"
								value={data.IP}
								onChange={(e) => handleChange(index, e)}
								pattern="^([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})|^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
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
							<input
								type="text"
								name="owner"
								value={data.allOwners ? data.allOwners : data.owner}
								disabled={data.allOwners}
								onChange={(e) => handleChange(index, e)}
								required
							/>
						</div>
					))}
					<div className="form--btn-submit">
						<Button type="submit" value="Add Devices" size="l" float="left" />
					</div>
				</form>
			</div>
		</Container>
	);
}
