import Container from '../Container/Container';
import DataDisplay from './DataDisplay'
import './deviceInfo.scss'

export default function DeviceInfo({ data }) {
	return (
		<Container>
            <div className="grid-50-50" style={{columnGap: "20px", paddingInline: '20px'}}>
                <div>
                    <div className="device-info--label-1 float-left">Main Characteristics</div>
                    <DataDisplay text={"Time of work"} data={data.Elapsed} />
                    <div className="device-info--label-2 float-left">Temperature</div>
                    <DataDisplay text={"Current Temperature"} data={data.Temperature + '°'} />
                    <DataDisplay text={"Temperature Average"} data={data.ChipTempAvg + '°'} />
                    <DataDisplay text={"Temperature Max"} data={data.ChipTempMax + '°'} />
                    <DataDisplay text={"Temperature Min"} data={data.ChipTempMin + '°'} />
                    <div className="device-info--label-2 float-left">Fans Speed</div>
                    <DataDisplay text={"Fan Speed In"} data={data.FanSpeedIn} />
                    <DataDisplay text={"Fan Speed Out"} data={data.FanSpeedOut} />
                </div>
                <div>
                <div className="device-info--label-1 float-left">Additional Characteristics</div>
                <DataDisplay text={"IP address"} data={data.IP} />
                <DataDisplay text={"MAC address"} data={data.MAC} />
                <DataDisplay text={"Power Mode"} data={data.PowerMode} />
                <DataDisplay text={"Mega Hashrate per second Average"} data={data.MHSav} />
                </div>
            </div>
		</Container>
	);
}
