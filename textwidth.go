// MIT License
//
// Copyright (c) 2020 Star Brilliant
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package textwidth

import (
	"unicode"

	"golang.org/x/text/width"
)

const tabMask = 7

// GetTextWidth calculates how many columns a string takes on a terminal.
//
// `startColumn` specifies the column that the string starts to print out.
// `wrapColumn` specified the number of columns that the terminal has.
//
// The algorithm takes care of East Asian characters which takes 2 columns to
// display, and inserts a single space at the end of each row when necessary.
//
// To disable text wrapping, specify `wrapColumn = 0`. In this mode, `\n`, `\v`,
// and `\f` will behave as a normal space.
//
// However, this algorithm does not support Middle East languages requiring
// special BiDi shaping algorithms.
func GetTextWidth(s string, startColumn, wrapColumn int) int {
	wrapEnabled := startColumn < wrapColumn
	rows, endColumn := GetTextOffset(s, startColumn, wrapColumn)

	if wrapEnabled {
		return rows*wrapColumn + endColumn - startColumn
	}
	return rows + endColumn - startColumn
}

// GetTextOffset calculates how many rows and columns a string takes on a
// terminal.
//
// `startColumn` specifies the column that the string starts to print out.
// `wrapColumn` specified the number of columns that the terminal has.
//
// The algorithm takes care of East Asian characters which takes 2 columns to
// display, and inserts a single space at the end of each row when necessary.
//
// To disable text wrapping, specify `wrapColumn = 0`. In this mode, ``\v` and
// `\f` behave as if the line is infinity in length.
//
// However, this algorithm does not support Middle East languages requiring
// special BiDi shaping algorithms.
func GetTextOffset(s string, startColumn, wrapColumn int) (rows, endColumn int) {
	wrapEnabled := startColumn < wrapColumn

	row := 0
	column := startColumn

	for _, c := range s {
		switch c {
		case '\b':
			if column > 0 {
				column--
			}
		case '\t':
			newColumn := (column | tabMask) + 1
			if wrapEnabled {
				if newColumn < wrapColumn {
					column = newColumn
				} else if column < wrapColumn {
					column = wrapColumn - 1
				}
			} else {
				column = newColumn
			}
		case '\n':
			if wrapEnabled {
				column = 0
			}
			row++
		case '\v', '\f':
			row++
		case '\r':
			if wrapEnabled {
				column = 0
			}
		default:
			if unicode.IsGraphic(c) {
				kind := width.LookupRune(c).Kind()
				if kind == width.EastAsianWide || kind == width.EastAsianFullwidth {
					if wrapEnabled && column+2 > wrapColumn {
						row++
						column = 0
					}
					column += 2
				} else {
					if wrapEnabled && column+1 > wrapColumn {
						row++
						column = 0
					}
					column++
				}
			}
		}
	}

	return row, column
}
