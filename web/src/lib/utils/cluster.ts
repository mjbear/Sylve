/**
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2025 The FreeBSD Foundation.
 *
 * This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
 * of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
 * under sponsorship from the FreeBSD Foundation.
 */

import type { ClusterDetails, ClusterNode } from '$lib/types/cluster/cluster';

export function getQuorumStatus(
	details: ClusterDetails,
	nodes: ClusterNode[]
): 'ok' | 'warning' | 'error' {
	const voters = details.nodes.filter((n) => (n.suffrage ?? 'Voter').toLowerCase() !== 'nonvoter');
	const totalVoters = voters.length;

	if (totalVoters === 0) return 'error';

	const onlineVoters = voters.filter((rn) =>
		nodes.some(
			(n) => n.nodeUUID === rn.id && n.status.toLowerCase() === 'online' // adjust if your API differs
		)
	).length;

	const quorum = Math.floor(totalVoters / 2) + 1;
	const hasLeader = Boolean(details.leaderId) || details.nodes.some((n) => n.isLeader === true);

	if (!hasLeader) return 'error';

	if (onlineVoters < quorum) {
		return 'error';
	}

	if (onlineVoters < totalVoters) {
		return 'warning';
	}

	return 'ok';
}
