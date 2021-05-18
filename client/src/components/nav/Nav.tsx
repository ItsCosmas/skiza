import React from 'react';
import './Nav.scss';

import sprite from 'assets/icons/sprite.svg';

function Nav() {
	return (
		<nav className='nav'>
			<span className='text--bold text--big'>
				{' '}
				<svg className='table-icon delete'>
					<use xlinkHref={`${sprite}#${'icon-satellite'}`} />
				</svg>
				<span style={{ marginLeft: '5px' }}>Skiza</span>
			</span>
			<span className='text--small text--grey'>
				Callback Listener and Aggregator
			</span>
		</nav>
	);
}

export default Nav;
