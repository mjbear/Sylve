import type { Column, Row } from '$lib/types/components/tree-table';
import type { SwitchList } from '$lib/types/network/switch';

export function generateTableData(switches: SwitchList | undefined): {
	rows: Row[];
	columns: Column[];
} {
	const columns: Column[] = [
		{
			field: 'id',
			visible: false,
			title: 'ID'
		},
		{
			field: 'name',
			title: 'Name'
		},
		{
			field: 'bridge',
			title: 'Bridge'
		}
	];

	const rows: Row[] = [];
	if (switches && switches['manual']) {
		console.log(switches['manual']);
		for (const sw of switches['manual']) {
			const row: Row = {
				id: sw.id,
				name: sw.name,
				bridge: sw.bridge || '-'
			};

			rows.push(row);
		}
	}

	return { rows, columns };
}
