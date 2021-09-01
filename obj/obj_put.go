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

import "github.com/google/uuid"

//Put puts the byte
func Put(id uuid.UUID, bytes []byte) error {
	_, err1 := Get(id)
	if err1 != nil {
		return err1
	} else {
		err2 := postWithID(bytes, id, true)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
