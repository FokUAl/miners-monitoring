const ReData = ({ devices }) => {
	const onlineCount = devices
		? devices.filter((device) => device.MinerStatus === 'online').length
		: 0;
	console.log(onlineCount);
	return { onlineCount };
};

export default ReData;
