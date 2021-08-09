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
	"errors"
	"fmt"
	"github.com/avr1/objs/obj"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(postCommand)
}

var postCommand = &cobra.Command{
	Use:   "post",
	Short: "Posts the given file to the object store",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args[1:]) != 2 {
			return errors.New("not enough arguments in post call")
		} else {
			bytes, err := os.ReadFile(os.Args[2])
			if err != nil {
				return err
			}
			n, err := fmt.Println(obj.Post(bytes))
			if err != nil {
				return err
			}
			fmt.Println(n)
			return nil
		}
	},
}
