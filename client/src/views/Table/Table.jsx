import React from 'react';
import TableList from '../../components/table/TableList';

import sprite from '../../assets/icons/sprite.svg';
import './Table.scss';

function Table() {
	const columns = [
		{
			Header: 'Time',
			accessor: 'timestamp',
		},
		{
			Header: 'Response',
			accessor: 'response',
		},
		{
			Header: 'View',
			accessor: 'view',
			Cell: ({ cell }) => (
				<span role='button'>
					<svg className='table-icon edit'>
						<use xlinkHref={`${sprite}#${'icon-eye'}`} />
					</svg>
				</span>
			),
		},
		{
			Header: 'Delete',
			accessor: 'delete',
			Cell: ({ cell }) => (
				<span role='button'>
					<svg className='table-icon delete'>
						<use xlinkHref={`${sprite}#${'icon-trash-2'}`} />
					</svg>
				</span>
			),
		},
	];

	//  TODO, Realtime Push Here
	const data = [
		{
			timestamp: '10.15',
			response: 'lorem ipsum dolor si',
			id: '89',
		},
		{
			timestamp: '10.15',
			response:
				'lorem ipsum dolor si met paramawata omera lamavamado lorem ipsum dolor si met paramawata omera lamavamado lorem ipsum dolor si met paramawata omera lamavamado',
		},
		{
			timestamp: '10.15',
			response: 'lorem ipsum dolor si met paramawata omera lamavamado',
		},
		{
			timestamp: '10.15',
			response: 'lorem ipsum dolor si met paramawata omera lamavamado',
		},
		{
			timestamp: '10.15',
			response: 'lorem ipsum dolor si met paramawata omera lamavamado',
		},
		{
			timestamp: '10.15',
			response: 'lorem ipsum dolor si met paramawata omera lamavamado',
		},
		{
			timestamp: '10.15',
			response: 'lorem ipsum dolor si met paramawata omera lamavamado',
		},
	];

	return (
		<div className='table'>
			<TableList columns={columns} data={data} />
		</div>
	);
}

export default Table;
