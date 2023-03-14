import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import FormService from '../../services/form.service';
import Button from '../Button/Button';
import './addDeviceForm.scss';

export default function AddDeviceForm({ allIP }) {
	const initialData = [
		{
			IP: '',
			shelf: '',
			column: '',
			row: '',
			owner: '',
		},
	];
	const [data, setData] = useState(initialData);
	const navigate = useNavigate();
	// const [allOwners, setAllOwners] = useState('')

	const handleChange = (index, event) => {
		const { value, name } = event.target;
		const newData = [...data];
		newData[index][name] = value;
		setData(newData);
	};

	const addFormField = () => {
		setData([
			...data,
			{
				IP: '',
				shelf: '',
				column: '',
				row: '',
				owner: '',
				minerType: '',
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
			}
			// if (!allIP.includes(data[i]['IP'])) {
			// 	alert('There is no such IP');
			// 	return;
			// }
		}
		FormService.addDevice(data).then(
		    response => {
		        navigate('/')
		        window.location.reload()
		    },
		    error => { if (error) {
		        console.log('Add device ', error)
		        alert(error.response.data.Message)
		    }}
		)
	};

	return (
		<div className="container">
			<form onSubmit={handleAdd}>
				<div className="form--title">Add new Device</div>
				<div className="form--btns">
					<Button type="button" value="Add" onClick={addFormField} size="m" />
					<Button
						type="button"
						value="Remove"
						onClick={removeFormField}
						size="m"
					/>
					{/* <input type="text" value={allOwners} onChange={e => {setAllOwners(e.target.value)}} className="form--input-allOwners"/> */}
				</div>
				<div className="form--labels">
					<label>Miner Type</label>
					<label>IP</label>
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
							pattern="^([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})$"
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
							value={data.owner}
							onChange={(e) => handleChange(index, e)}
							required
						/>
					</div>
				))}
				<div className="form--btn-submit">
					<Button type="submit" value="Add Devices" size="l" />
				</div>
			</form>
		</div>
	);
}
