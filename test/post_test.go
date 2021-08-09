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

package test

import (
	"github.com/avr1/objs/obj"
	"os"
	"testing"
)

func TestPostAndGetCall(t *testing.T) {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	id, err1 := obj.Post(b)
	if err1 != nil {
		t.Errorf("Not posted successfully.")
	}
	rootDir, err2 := os.UserHomeDir()
	if err2 != nil {
		t.Errorf("Root directory not determined successfully.")
	}

	file, err3 := os.ReadFile(rootDir + "/.objs/" + id.String() + ".obj")

	if err3 != nil {
		t.Errorf("Couldn't read the file successfully.")
	}

	for i, r := range file {
		if r != b[i] {
			t.Errorf("The files don't match exactly")
		}
	}

	t.Cleanup(func() {
		if err := os.Remove(rootDir + "/.objs/" + id.String() + ".obj"); err != nil {
			return
		}
	})
	/*newBytes := obj.Get(id)
	for i, byt := range b {
		if byt !=  newBytes[i] {
			t.Errorf("The bytes were not the same as the ones posted.", total, 10)
		}
	}*/
}
