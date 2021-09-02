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
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/user"
	"strings"
	"time"
)

// Post posts the given data, as a byte array, to a secure location, while returning a UUID that allows the user to access the data.
func Post(bytes []byte) (uuid.UUID, error) {
	id := uuid.New()
	err := postWithID(bytes, id, false)
	if err != nil {
		return uuid.Nil, err
	} else {
		return id, nil
	}
}

func postWithID(bytes []byte, id uuid.UUID, alreadyListed bool) error {
	rootDir, err1 := os.UserHomeDir()
	if err1 != nil {
		return err1
	}

	_ = os.Mkdir(rootDir+"/.objs", os.ModePerm)
	rootDir += "/.objs/"

	if alreadyListed {
		err3 := modifyEntry(id.String())
		fmt.Println("Modifying entry!")
		if err3 != nil {
			return err3
		}
	} else {
		err4 := addToList(id.String())
		if err4 != nil {
			return err4
		}
	}

	create, err4 := os.Create(rootDir + id.String() + ".obj")
	if err4 != nil {
		return err4
	}

	defer func(create *os.File) {
		err := create.Close()
		if err != nil {
			panic(err)
		}
	}(create)

	_, err5 := create.Write(bytes)
	if err5 != nil {
		return err5
	}
	return nil
}

func modifyEntry(id string) error {
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
		var toAppend string
		current, err := user.Current()
		if err != nil {
			return err
		}

		if !strings.Contains(line, id) {
			toAppend = line + "\n"
		} else {
			toAppend = id + "\t" + time.Now().String() + "\t" + current.Name + "\t" + "\n"
		}
		newLines = append(newLines, []byte(toAppend)...)
	}
	if err4 := os.Truncate(rootDir+"/.objs/.list.txt", 0); err4 != nil {
		return err4
	}

	if err5 := os.WriteFile(rootDir+"/.objs/.list.txt", newLines, 0666); err5 != nil {
		return err5
	}
	return nil
}
func addToList(id string) error {
	rootDir, err1 := os.UserHomeDir()
	if err1 != nil {
		return err1
	}
	// If the file doesn't exist, create it, or append to the file
	f, err2 := os.OpenFile(rootDir+"/.objs/.list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		return err2
	}
	current, err3 := user.Current()
	if err3 != nil {
		return err3
	}
	if _, err4 := f.Write([]byte(id + "\t" + time.Now().String() + "\t" + current.Name + "\t" + "\n")); err4 != nil {
		return err4
	}
	if err5 := f.Close(); err5 != nil {
		return err5
	}
	return nil
}
