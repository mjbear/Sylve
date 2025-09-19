<script lang="ts">
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { generateNanoId } from '$lib/utils/string';
	import type { FullAutoFill } from 'svelte/elements';

	interface Props {
		label?: string;
		labelHTML?: boolean;
		value: string | number;
		placeholder: string;
		autocomplete?: FullAutoFill | null | undefined;
		classes: string;
		type?: string;
		textAreaClasses?: string;
		disabled?: boolean;
		onChange?: (value: string | number) => void;
	}

	let {
		value = $bindable(''),
		label = '',
		labelHTML = false,
		placeholder = '',
		autocomplete = 'off',
		classes = 'space-y-1.5',
		type = 'text',
		textAreaClasses = 'min-h-56',
		disabled = false,
		onChange
	}: Props = $props();

	let nanoId = $state(generateNanoId(label));
</script>

<div class={`${classes}`}>
	{#if label}
		<Label class="mt w-full whitespace-nowrap text-sm" for={nanoId}>
			<!-- {labelHTML ? {@html label} : label} -->
			{#if labelHTML}
				{@html label}
			{:else}
				{label}
			{/if}
		</Label>
	{/if}
	{#if type === 'textarea'}
		<Textarea
			class={textAreaClasses}
			id={nanoId}
			{placeholder}
			{autocomplete}
			bind:value
			{disabled}
			oninput={(e) => {
				value = e.target?.value;
				if (onChange) onChange(value);
			}}
		/>
	{:else}
		<Input
			{type}
			id={nanoId}
			{placeholder}
			{autocomplete}
			bind:value
			{disabled}
			oninput={(e) => {
				value = e.target?.value;
				if (onChange) onChange(value);
			}}
		/>
	{/if}
</div>
