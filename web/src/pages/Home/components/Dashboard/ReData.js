const ReData = (devices) => {
	const initialValue = 0;
	const onlineCount = devices.filter(
		(device) => device.MinerStatus === 'online'
	).length;
	const totalHashrate =
		devices
			.reduce(
				(accum, device) => accum + device.Characteristics.MHSav,
				initialValue
			)
			.toFixed() + ' MHS';
	const avgTemperature = (
		devices
			.filter((device) => device.MinerStatus === 'online')
			.reduce(
				(accum, device) => accum + device.Characteristics.Temperature,
				initialValue
			) / onlineCount
	).toFixed();
	const avgTemp = isNaN(avgTemperature)? 0 : avgTemperature
	let maxTemp = 0
	devices.forEach(device => device.Characteristics.Temperature > maxTemp ? maxTemp = device.Characteristics.Temperature : '')
	return { onlineCount, totalHashrate, avgTemp, maxTemp };
};

export default ReData;
