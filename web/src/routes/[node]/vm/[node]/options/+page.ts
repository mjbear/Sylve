import { getVMDomain, getVMs } from '$lib/api/vm/vm';
import { SEVEN_DAYS } from '$lib/utils';
import { cachedFetch } from '$lib/utils/http';

export async function load({ params }) {
	const vmId = Number(params.node);
	const cacheDuration = SEVEN_DAYS;
	const [vms, domain] = await Promise.all([
		cachedFetch('vm-list', async () => await getVMs(), cacheDuration),
		cachedFetch(`vm-domain-${vmId}`, async () => await getVMDomain(vmId), cacheDuration)
	]);

	const vm = vms.find((vm) => vm.vmId === vmId);

	return {
		vm,
		domain,
		vms
	};
}
