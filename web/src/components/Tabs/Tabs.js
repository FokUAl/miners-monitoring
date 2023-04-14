import { useState } from 'react';
import Container from '@components/Container/Container';
import Tab from './Tab';

function Tabs({ tabs, active, handleActive }) {
	return (
		<Container>
			<div className="tabs">
				{tabs.map((tab) => (
					<Tab
						key={tab.id}
						label={tab.label}
						id={tab.id}
						isActive={active === tab.id}
						handleActive={handleActive}
					/>
				))}
			</div>
		</Container>
	);
}

export default Tabs;
