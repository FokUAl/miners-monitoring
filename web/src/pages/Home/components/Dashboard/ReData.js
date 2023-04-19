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
	const avgTemp = (
		devices
			.filter((device) => device.MinerStatus === 'online')
			.reduce(
				(accum, device) => accum + device.Characteristics.Temperature,
				initialValue
			) / onlineCount
	).toFixed();
	return { onlineCount, totalHashrate, avgTemp };
};

export default ReData;
