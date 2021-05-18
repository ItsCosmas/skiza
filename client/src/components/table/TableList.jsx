import React, { useState } from 'react';
import { useTable } from 'react-table';
import ReactModal from 'react-modal';

import sprite from 'assets/icons/sprite.svg';

import './TableList.scss';

function TableList({ columns, data }) {
	const [showModal, setShowModal] = useState(false);

	const triggerModal = () => {
		setShowModal(!showModal);
	};

	const {
		getTableProps,
		getTableBodyProps,
		headerGroups,

		rows,
		prepareRow,
	} = useTable({
		columns,
		data,
	});

	const customStyles = {
		content: {
			top: '50%',
			left: '50%',
			right: 'auto',
			bottom: 'auto',
			marginRight: '-50%',
			transform: 'translate(-50%, -50%)',
			// background: 'rgba(255,255,255,1)',
		},
		overlay: {
			background: 'rgba(0,0,0,.5)',
		},
	};

	// http://reactcommunity.org/react-modal/accessibility/
	ReactModal.setAppElement('#root');

	return (
		<table {...getTableProps()}>
			<thead>
				{headerGroups.map((headerGroup) => (
					<tr {...headerGroup.getHeaderGroupProps()}>
						{headerGroup.headers.map((column) => (
							<th {...column.getHeaderProps()}>{column.render('Header')}</th>
						))}
					</tr>
				))}
			</thead>
			<tbody {...getTableBodyProps()}>
				{rows.map((row, i) => {
					prepareRow(row);
					return (
						<tr {...row.getRowProps()}>
							{row.cells.map((cell) => {
								return (
									<>
										<td {...cell.getCellProps()} onClick={() => triggerModal()}>
											{cell.render('Cell')}
										</td>
										<ReactModal
											// overlayClassName={'modal-overlay'}
											isOpen={showModal}
											style={customStyles}
											contentLabel=' Modal'>
											<div className='dialog'>
												<h4 className='heading--primary u-margin-top-small u-center-text'>
													<span className='text--bold text--big'>
														<svg className='table-icon'>
															<use xlinkHref={`${sprite}#${'icon-clock'}`} />
														</svg>
														<span style={{ marginLeft: '5px' }}>Time</span>
													</span>
												</h4>
												<div className='dialog u-center-text u-margin-top-small'>
													Lorem ipsum dolor sit amet consectetur adipisicing
													elit. Nemo voluptatum autem ex, fugiat unde delectus
													laborum cumque asperiores dolorum eos inventore sit id
													rem? Nulla deserunt tenetur fugiat earum illo!
												</div>
												<div className='dialog--actions u-margin-top-small u-center-text'>
													<button
														className='btn btn--icon btn-spread btn--red'
														block='true'
														type='submit'
														onClick={() => triggerModal()}>
														<svg className='btn--icon-shape'>
															<use xlinkHref={`${sprite}#${'icon-trash-2'}`} />
														</svg>
														<span>Delete</span>
													</button>
													<button
														className='btn btn--icon btn-spread btn--green'
														block='true'
														type='submit'
														onClick={() => triggerModal()}>
														<svg className='btn--icon-shape'>
															<use xlinkHref={`${sprite}#${'icon-cancel'}`} />
														</svg>
														<span>Close</span>
													</button>
												</div>
											</div>
										</ReactModal>
									</>
								);
							})}
						</tr>
					);
				})}
			</tbody>
		</table>
	);
}

export default TableList;
