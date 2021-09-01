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
	"github.com/avr1/objs/obj"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(putCommand)
}

var putCommand = &cobra.Command{
	Use:   "put UUID File",
	Short: "Puts the contents of the given file in the object with the given UUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args[1:]) != 3 {
			return errors.New("not enough arguments in get call")
		} else {
			id, err1 := uuid.Parse(os.Args[2])
			if err1 != nil {
				return err1
			}

			bytes, err := os.ReadFile(os.Args[3])
			if err != nil {
				return err
			}

			err2 := obj.Put(id, bytes)
			if err2 != nil {
				return err2
			}

			return nil
		}
	},
}
