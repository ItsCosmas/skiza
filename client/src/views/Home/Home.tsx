import React from 'react';
import Nav from 'components/nav/Nav';
import Table from 'views/Table/Table';

import './Home.scss';

function Home() {
	return (
		<div>
			<Nav />
			<main>
				<Table />
			</main>
		</div>
	);
}

export default Home;
