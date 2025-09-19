// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package clusterModels

type Cluster struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Enabled       bool   `json:"enabled"`
	Key           string `json:"key"`
	RaftBootstrap *bool  `json:"raftBootstrap"`
	RaftIP        string `json:"raftIP"`
	RaftPort      int    `json:"raftPort"`
}
