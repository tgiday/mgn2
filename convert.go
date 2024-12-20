package mgn2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Prsfloat return the float64 value parsing geez string g.
func Prsfloat(g string) float64 {
	if !strings.Contains(g, ".") {
		//i, _ := strconv.Atoi(str)
		i := Prsint(g)
		f := float64(i)
		return f
	}
	strs := strings.Split(g, ".")
	s0 := Prsint(strs[0])
	s1 := Prsint(strs[1])
	res := fmt.Sprint(s0) + "." + fmt.Sprint(s1)
	f, _ := strconv.ParseFloat(res, 64)
	return f
}

// Prsint return the intiger value parsing geez string g.
func Prsint(g string) int {
	if strings.Contains(g, ".") {
		return 0
	}
	var signed bool
	if len(g) > 1 && g[0] == '-' {
		g = g[1:]
		signed = true
	}
	lst := strings.SplitAfter(g, "፻")
	var zz []int
	for idx, x := range lst {
		var z1, z2 int
		if strings.Contains(x, "፻") {
			if x == "፻" {
				y := prsrune('፩')
				z1 = y * int(math.Pow10(2*((len(lst)-1)-idx)))
			} else {
				l := strings.Split(x, "፻")
				for _, v := range l[0] {
					y := prsrune(v)
					y = y * int(math.Pow10(2*((len(lst)-1)-idx)))
					z1 += y
				}
			}
			zz = append(zz, z1)
		} else {
			for _, v := range x {
				y := prsrune(v)
				y = y * int(math.Pow10(2*((len(lst)-1)-idx)))
				z2 += y
			}
			zz = append(zz, z2)
		}
	}
	var i int
	for _, v := range zz {
		i += v
	}
	if signed {
		i = -1 * i
	}
	return i
}

func prsrune(g rune) int {
	dif := '፩' - '1'
	if g == '፻' {
		return 100
	}
	if '፱' < g && g < '፻' {
		r := g - dif - 9
		i, _ := strconv.Atoi(string(r) + "0")
		return i
	}
	if '፨' <= g && g <= '፱' {
		r := g - dif
		i, _ := strconv.Atoi(string(r))
		return i
	}
	return 0
}

// Fmtint return string representation of integer i
func Fmtint(i int) string {
	var signed bool
	if i < 0 {
		i = -1 * i
		signed = true
	}
	str := strconv.Itoa(i)
	var zz []rune
	for idx, v := range str {
		b := (len(str) - 1) - idx
		q := b % 2
		if q != 0 {
			i, _ := strconv.Atoi(string(v))
			x := fmtrune(i * 10)
			zz = append(zz, x)
		} else {
			i, _ := strconv.Atoi(string(v))
			x := fmtrune(i)
			if idx != len(str)-1 {
				zz = append(zz, x, '፻')
			} else {
				zz = append(zz, x)
			}
		}
	}
	var ans []rune
	for _, v := range zz {
		if v != '፨' {
			ans = append(ans, v)
		}
	}
	if len(ans) >= 2 && ans[0] == '፩' && ans[1] == '፻' {
		ans = ans[1:]
	}
	var answer string
	switch signed {
	case true:
		answer = "-" + string(ans)
	case false:
		answer = string(ans)
	}
	return answer
}

func fmtrune(i int) rune {
	dif := '፩' - '1'
	if i == 100 {
		return '፻'
	}
	if 10 <= i && i <= 90 {
		r := dif + '1' + rune(i)/10 + 8
		return r
	}
	if 0 < i && i <= 9 {
		r := dif + '1' + rune(i) - 1
		return r
	}
	return '፨'
}

// Fmtfloat return string representation of float64 n
func Fmtfloat(n float64) string {
	str := strconv.FormatFloat(n, 'f', -1, 64)
	if !strings.Contains(str, ".") {
		i, _ := strconv.Atoi(str)
		s := Fmtint(i)
		return s
	}
	strs := strings.Split(str, ".")
	i, _ := strconv.Atoi(strs[0])
	s0 := Fmtint(i)
	q := len(strs[1]) % 2
	j, _ := strconv.Atoi(strs[1])
	if q != 0 {
		j, _ = strconv.Atoi(strs[1] + "0")
	}
	s1 := Fmtint(j)
	return s0 + "." + s1
}
