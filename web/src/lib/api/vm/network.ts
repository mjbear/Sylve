import { APIResponseSchema, type APIResponse } from '$lib/types/common';
import { apiRequest } from '$lib/utils/http';

export async function detachNetwork(vmId: number, switchId: number): Promise<APIResponse> {
	return await apiRequest(`/vm/network/detach`, APIResponseSchema, 'POST', {
		vmId,
		networkId: switchId
	});
}

export async function attachNetwork(
	vmId: number,
	switchName: string,
	emulation: string,
	macId: number
): Promise<APIResponse> {
	return await apiRequest(`/vm/network/attach`, APIResponseSchema, 'POST', {
		vmId,
		switchName,
		emulation,
		macId
	});
}
