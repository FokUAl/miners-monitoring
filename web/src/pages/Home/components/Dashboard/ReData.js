const ReData = ({ devices }) => {
	const onlineCount = devices
		? devices.filter((device) => device.MinerStatus === 'online').length
		: 0;
    const totalHashrate = devices ? devices.reduce((accum, device) => accum + device.Characteristics.MHSav) : '0'
	const avgTemp = devices ? devices.reduce((accum, device) => accum + device.Characteristics.Temperature) / devices.length : '0'
	return { onlineCount , totalHashrate, avgTemp};
};

export default ReData;
