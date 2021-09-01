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
	"path/filepath"
	"strings"
)

//Get returns the byte array of the file, as stored in the object store.
//Only returns the data if the person accessing it has permission to get the data.
func Get(id uuid.UUID) ([]byte, error) {
	pathToRead, err1 := os.UserHomeDir()
	if err1 != nil {
		return nil, err1
	}

	if b := getValidFromList(id.String()); b != nil {
		return nil, b
	}

	pathToRead = filepath.Join(pathToRead, ".objs", id.String()+".obj")
	file, err2 := os.ReadFile(pathToRead)
	if err2 != nil {
		return nil, errors.New("no object with that UUID could be found")
	}

	return file, nil
}

func getValidFromList(id string) error {
	rootDir, err1 := os.UserHomeDir()
	if err1 != nil {
		return err1
	}

	f, err2 := os.ReadFile(rootDir + "/.objs/.list.txt")

	if err2 != nil {
		return err2
	}

	lines := strings.Split(string(f), "\n")

	for _, line := range lines {
		if strings.Contains(line, id) {
			data := strings.Split(line, "\t")
			current, err3 := user.Current()
			if err3 != nil {
				return err3
			}
			if current.Name != "System Administrator" && data[2] != current.Name {
				return errors.New("cannot get a file you did not create")
			}
		}
	}

	return nil
}
