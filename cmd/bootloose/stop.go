// SPDX-FileCopyrightText: 2019 Weaveworks Ltd.
// SPDX-FileCopyrightText: 2023 bootloose authors
// SPDX-License-Identifier: Apache-2.0

package bootloose

import (
	"github.com/spf13/cobra"

	"github.com/k0sproject/bootloose/pkg/cluster"
)

func NewStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop cluster machines",
		RunE:  stop,
	}
}

func stop(cmd *cobra.Command, args []string) error {
	cluster, err := cluster.NewFromFile(clusterConfigFile(cmd))
	if err != nil {
		return err
	}
	return cluster.Stop(args)
}
