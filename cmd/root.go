/*
Copyright © 2023 Galactica Network

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "galactica-guardian",
	Short: "galactica-guardian is a CLI tool for managing and verifying user data on the Galactica blockchain",
	Long: `galactica-guardian is a powerful command-line interface (CLI) tool designed for
guardians in the Galactica blockchain ecosystem. Guardians play a pivotal role
in verifying and managing user data on the blockchain. 

In the initial stage, galactica-guardian primarily focuses on Know Your
Customer (KYC) verification. Users can visit the Galactica website or app to
input their personal information. Guardians then perform document verification
using national IDs, passports, or driver's licenses. Once verified, guardians
issue zkKYC attestations on-chain, which are stored in a Merkle tree registry.

As Galactica evolves, guardians will have the capability to issue various
other attestations, such as ZkCerts for education diplomas, gym memberships,
conference tickets, and more. This CLI empowers guardians to efficiently
manage these verification processes and maintain a secure and trusted
ecosystem.

For detailed usage instructions, available commands, and options, please refer
to the respective sections in the documentation.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
