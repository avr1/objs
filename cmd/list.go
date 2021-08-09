/*
 *
 * Copyright © 2021 Arjun Ramachandrula <avrdev77@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Returns the list of all used UUIDs currently stored.",
	Run: func(cmd *cobra.Command, args []string) {
		rootDir, err1 := os.UserHomeDir()
		if err1 != nil {
			panic("Not supposed to happen.")
		}
		file, err2 := os.ReadFile(rootDir + "/.objs/.list.txt")
		if err2 != nil {
			panic("Not supposed to happen.")
		}
		fmt.Println(string(file))
	},
}
