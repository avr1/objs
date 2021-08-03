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
	"github.com/google/uuid"
	"log"
	"os"
	"os/user"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Post posts the given data, as a byte array, to a secure location, while returning a UUID that allows the user to access the data.
func Post(bytes []byte) uuid.UUID {
	id := uuid.New()
	rootDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir(rootDir+"/.objs", os.ModePerm)
	rootDir += "/.objs/"

	addToList(id.String())
	create, err1 := os.Create(rootDir + id.String() + ".obj")
	check(err1)

	defer func(create *os.File) {
		err := create.Close()
		check(err)
	}(create)
	_, err2 := create.Write(bytes)
	check(err2)
	return id
}

func addToList(id string) {
	rootDir, err1 := os.UserHomeDir()
	if err1 != nil {
		log.Fatal(err1)
	}
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(rootDir+"/.objs/.list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	current, err := user.Current()
	if err != nil {
		panic("What's messed up about your user?")
	}
	if _, err := f.Write([]byte(id + "\t" + time.Now().String() + "\t" + current.Name + "\t" + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
