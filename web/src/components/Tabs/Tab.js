import './tabs.scss'

const Tab = ({ id, label, isActive, handleActive }) => {
	return (
		<button
			onClick={() => handleActive(id)}
			className={`tab${isActive ? ' active' : ''}`}
		>
			{label}
		</button>
	);
};

export default Tab;
