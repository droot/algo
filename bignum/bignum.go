package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type BigNum struct {
	digits []uint8
}

func BigNumFrmString(s string) (*BigNum, error) {
	digits := make([]uint8, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		var v uint8
		if d, err := strconv.Atoi(s[i : i+1]); err != nil {
			return nil, err
		} else {
			v = uint8(d)
		}
		digits = append(digits, v)
	}

	return &BigNum{digits: digits}, nil
}

func BigNumFrmDigit(u uint8) (*BigNum, error) {
	if u > 9 {
		return nil, fmt.Errorf("can't be more than 9 ")
	}
	return &BigNum{
		digits: []uint8{u},
	}, nil
}

func (bn *BigNum) String() string {
	b := new(bytes.Buffer)
	for i := bn.NumDigits() - 1; i >= 0; i-- {
		// figure out a way to do it in a better way
		b.WriteString(fmt.Sprintf("%d", bn.digits[i]))
	}
	s := strings.TrimLeft(b.String(), "0")
	if s == "" {
		return "0"
	}
	return s
}

func (bn *BigNum) NumDigits() int {
	return len(bn.digits)
}

func (bn *BigNum) MultiplyByDigit(m uint8) *BigNum {
	var carry uint8

	digits := []uint8{}
	for _, d := range bn.digits {
		s := m*d + carry
		digits = append(digits, s%10)
		carry = s / 10
	}
	if carry > 0 {
		digits = append(digits, carry)
	}

	return &BigNum{digits: digits}
}

func (bn *BigNum) Multiply(other *BigNum) *BigNum {
	rows := []*BigNum{}
	for i, digit := range other.digits {
		row := bn.MultiplyByDigit(digit).ShiftBy(i)
		fmt.Printf("%v*%d = %s \n", bn.String(), digit, row.String())
		rows = append(rows, row)
	}

	var carry uint8
	digits := []uint8{}
	for i := 0; ; i++ {
		found := false
		var sum uint8
		for _, row := range rows {
			if i < row.NumDigits() {
				sum = sum + row.digits[i]
				found = true
			}
		}
		if !found {
			break
		}
		fmt.Printf("%d ---> sum=%d carry=%d \n", i, sum, carry)
		sum = sum + carry
		digits = append(digits, sum%10)
		carry = sum / 10
	}
	if carry > 0 {
		digits = append(digits, carry)
	}

	return &BigNum{digits: digits}
}

func (bn *BigNum) Sum(other *BigNum) *BigNum {
	return nil
}

func (bn *BigNum) ShiftBy(n int) *BigNum {
	if n == 0 {
		return bn
	}
	digits := []uint8{}
	for i := 0; i < n; i++ {
		digits = append(digits, 0)
	}
	bn.digits = append(digits, bn.digits...)
	return bn
}
