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

package obj

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"os/user"
	"strings"
)

// Remove the object with the given UUID from the object store, provided that the
// id exists, and either the user who posted the object or the superuser is
// trying to remove it.
func Remove(id uuid.UUID) error {
	err1 := removeFromList(id.String())
	if err1 != nil {
		return err1
	}

	rootDir, err2 := os.UserHomeDir()
	if err2 != nil {
		return err2
	}

	err3 := os.Remove(rootDir + "/.objs/" + id.String() + ".obj")
	if err3 != nil {
		return err3
	}

	return nil
}

//removeFromList removes the id passed into it from the list, provided no errors are generated.
func removeFromList(id string) error {
	rootDir, err1 := os.UserHomeDir()
	if err1 != nil {
		return err1
	}

	f, err2 := os.ReadFile(rootDir + "/.objs/.list.txt")
	if err2 != nil {
		return err2
	}

	lines := strings.Split(string(f), "\n")
	var newLines []byte
	for _, line := range lines {
		if !strings.Contains(line, id) {
			newLines = append(newLines, []byte(line)...)
		} else {
			data := strings.Split(line, "\t")
			current, err3 := user.Current()
			if err3 != nil {
				return err3
			}
			if current.Name != "System Administrator" && data[2] != current.Name {
				return errors.New("you are not the user who created this file, and so you cannot remove it")
			}
		}
	}

	if err4 := os.Truncate(rootDir+"/.objs/.list.txt", 0); err4 != nil {
		return err4
	}

	if err5 := os.WriteFile(rootDir+"/.objs/.list.txt", newLines, 0666); err5 != nil {
		return err5
	}
	return nil
}
