// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package network

import (
	"fmt"
	"net"
)

func TryBindToPort(ip string, port int, proto string) error {
	addr := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen(proto, addr)
	if err != nil {
		return err
	}

	defer listener.Close()
	return nil
}
