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

package textwidth_test

import (
	"testing"

	textwidth "github.com/m13253/go-textwidth"
)

func TestASCII(t *testing.T) {
	w := textwidth.GetTextWidth("1234", 0, 0)
	if w != 4 {
		t.Errorf("width(1234) != 4")
	}
}

func TestAsian(t *testing.T) {
	w := textwidth.GetTextWidth("你好", 0, 0)
	if w != 4 {
		t.Errorf("width(你好) != 4")
	}
}

func TestASCIIWrap(t *testing.T) {
	w := textwidth.GetTextWidth("1234567890123456789012345", 0, 10)
	if w != 25 {
		t.Errorf("width(1234567890123456789012345) != 25")
	}

	w = textwidth.GetTextWidth("1234567890123456789012345", 5, 10)
	if w != 25 {
		t.Errorf("width(1234567890123456789012345) != 25")
	}
}

func TestAsianWrap(t *testing.T) {
	w := textwidth.GetTextWidth("一二三四/一二三四五/一二", 0, 10)
	if w != 25 {
		t.Errorf("width(一二三四/一二三四五/一二) != 25")
	}

	w = textwidth.GetTextWidth("一二三四/一二三四五/一二", 5, 10)
	if w != 27 {
		t.Errorf("width(一二三四/一二三四五/一二) != 27")
	}
}

func TestFullWidth(t *testing.T) {
	w := textwidth.GetTextWidth("1２四6７九", 0, 0)
	if w != 10 {
		t.Errorf("width(1２四6７九) != 10")
	}
}

func TestBackspace(t *testing.T) {
	w := textwidth.GetTextWidth("123456789X\b0", 0, 10)
	if w != 10 {
		t.Errorf("width(123456789X\\b0) != 10")
	}
}

func TestTab(t *testing.T) {
	w := textwidth.GetTextWidth("a\tb\tc\td", 0, 10)
	if w != 11 {
		t.Errorf("width(a\\tb\\tc\\td) != 11")
	}
	w = textwidth.GetTextWidth("a\tb\tc\td", 0, 0)
	if w != 25 {
		t.Errorf("width(a\\tb\\tc\\td) != 25")
	}
}

func TestLF(t *testing.T) {
	w := textwidth.GetTextWidth("12345\n1234567890\n12345", 0, 10)
	if w != 25 {
		t.Errorf("width(12345\\n1234567890\\n12345) != 25")
	}
	w = textwidth.GetTextWidth("12345\n1234567890\n12345", 0, 0)
	if w != 22 {
		t.Errorf("width(12345\\n1234567890\\n12345) != 22")
	}
}

func TestVT(t *testing.T) {
	w := textwidth.GetTextWidth("12345\v67890\n12345\v67890\f12345", 0, 10)
	if w != 55 {
		t.Errorf("width(12345\\v67890\\n12345\\v67890\\f12345) != 55")
	}
}

func TestCR(t *testing.T) {
	w := textwidth.GetTextWidth("1234567890\r12345", 0, 10)
	if w != 5 {
		t.Errorf("width(1234567890\\r12345) != 5")
	}
}
