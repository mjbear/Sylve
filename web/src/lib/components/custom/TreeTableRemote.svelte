<script lang="ts">
	import type { Column, Row } from '$lib/types/components/tree-table';
	import { hasRowsChanged, matchAny } from '$lib/utils/table';
	import { findRow, getAllRows, pruneEmptyChildren } from '$lib/utils/tree-table';
	import { onMount, untrack } from 'svelte';
	import {
		TabulatorFull as Tabulator,
		type ColumnDefinition,
		type RowComponent
	} from 'tabulator-tables';

	let tableComponent: HTMLDivElement | null = null;
	let table: Tabulator | null = $state(null);

	interface Props {
		data: {
			rows: Row[];
			columns: Column[];
		};
		ajaxURL?: string;
		name: string;
		parentActiveRow?: Row[] | null;
		query?: string;
		multipleSelect?: boolean;
	}

	let {
		data,
		name,
		parentActiveRow = $bindable([]),
		query = $bindable(),
		multipleSelect = true,
		ajaxURL
	}: Props = $props();

	let tableInitialized = $state(false);
	let scroll = $state([0, 0]);
	let aboutToClick = $state(false);

	function updateParentActiveRows() {
		if (tableInitialized) {
			parentActiveRow = table?.getSelectedRows().map((r) => r.getData() as Row) || [];
		}
	}

	$effect(() => {
		if (data.rows) {
			untrack(async () => {
				if (query && query !== '') return;

				if (data.rows.length === 0) {
					table?.clearData();
					return;
				}

				const selectedIds = table?.getSelectedRows().map((row) => row.getData().id) || [];
				const treeExpands = getAllRows(table?.getRows() || []).map((row) => ({
					id: row.getData().id,
					expanded: row.isTreeExpanded()
				}));

				if (hasRowsChanged(table, data.rows) && !aboutToClick) {
					if (tableInitialized) {
						await table?.replaceData(pruneEmptyChildren(data.rows));
					}
				}

				selectedIds.forEach((id) => {
					const row = findRow(table?.getRows() || [], id);
					if (row) row.select();
				});

				treeExpands.forEach((treeExpand) => {
					const row = findRow(table?.getRows() || [], treeExpand.id);
					if (row) {
						treeExpand.expanded ? row.treeExpand() : row.treeCollapse();
					}
				});

				updateParentActiveRows();
			});
		}
	});

	onMount(() => {
		if (tableComponent) {
			table = new Tabulator(tableComponent, {
				ajaxURL: ajaxURL ? ajaxURL : undefined,
				ajaxResponse: function (url, params, response) {
					return response.data;
				},
				reactiveData: true,
				columns: data.columns as ColumnDefinition[],
				layout: 'fitColumns',
				selectableRows: multipleSelect ? true : 1,
				dataTreeChildIndent: 16,
				dataTree: true,
				dataTreeChildField: 'children',
				dataTreeStartExpanded: true,
				persistenceID: name,
				paginationMode: 'remote',
				persistence: {
					sort: true,
					page: true,
					filter: true
				},
				placeholder: 'No data available',
				pagination: true,
				paginationSize: 25,
				paginationCounter: 'pages',
				sortMode: 'remote'
			});
		}

		table?.on('rowSelected', updateParentActiveRows);
		table?.on('rowDeselected', updateParentActiveRows);

		table?.on('rowDblClick', (_event: UIEvent, row: RowComponent) => {
			row.toggleSelect();
		});

		table?.on('tableBuilt', () => {
			tableInitialized = true;

			document.querySelector('.tabulator-footer')?.addEventListener('mouseover', () => {
				aboutToClick = true;
			});

			document.querySelector('.tabulator-footer')?.addEventListener('mouseout', () => {
				aboutToClick = false;
			});
		});

		table?.on('scrollVertical', (top) => {
			scroll = [top, scroll[1]];
		});

		table?.on('scrollHorizontal', (left) => {
			scroll = [scroll[0], left];
		});

		table?.on('renderComplete', () => {
			const container = document.querySelector('.tabulator-tableholder') as HTMLDivElement;
			if (container) {
				container.scrollTop = scroll[0];
				container.scrollLeft = scroll[1];
			}
		});
	});

	function tableFilter(query: string) {
		if (table && tableInitialized) {
			if (query === '') {
				table.clearFilter(true);
				return;
			}
			table.setFilter(matchAny, { query });
		}
	}

	$effect(() => {
		tableFilter(query || '');
	});
</script>

<div bind:this={tableComponent} class="flex-1 cursor-pointer" id={name}></div>
