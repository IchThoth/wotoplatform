/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
 * Copyright (c) 2021 ALiwoto.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package wotoStrong

import wv "wp-server/wotoPacks/utils/wotoValues"

// _setValue will set the bytes value of the StrongString.
func (_s *StrongString) _setValue(str string) {
	if _s._value == nil {
		_s._value = make([]rune, wv.BaseIndex)
	}

	for _, current := range str {
		_s._value = append(_s._value, current)
	}
}
