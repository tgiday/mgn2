package mgn2

import (
	"fmt"
	"testing"
)

func TestPrsint(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"+", 0},
		{"-", 0},
		{"-፲፱፻፴፩", -1931},
		{"+፴", 30},
		{"፲፱", 19},
		{"፱", 9},
		{"፻", 100},
		{"፲፻፱፻", 100900},
	}
	//fmt.Println("---Prsint testing---")
	for _, v := range tests {
		got := Prsint(v.in)
		if got != v.want {
			t.Errorf("want %v got %v", v.want, got)
		}
		//fmt.Printf("want %v got %v\n", v.want, got)
	}
}
func TestFmtint(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		//{10000, "፼"},
		{100, "፻"},
		{90, "፺"},
		{10, "፲"},
		{9, "፱"},
		{1, "፩"},
		//{0, "፨"},
		{0, ""},
		{+19, "፲፱"},
		{-1931, "-፲፱፻፴፩"},
		{100900, "፲፻፱፻"},
		//{10000, "ሐ"},
	}
	//fmt.Println("---Fmtint testing---")
	for _, v := range tests {
		got := Fmtint(v.in)
		if got != v.want {
			t.Errorf("want %s got %s", v.want, got)
		}
		//fmt.Printf("want %s got %s\n", v.want, got)
	}
}
func TestFmtfloat(t *testing.T) {
	tests := []struct {
		in   float64
		want string
	}{
		//{10000, "፼"},
		{100.11, "፻.፲፩"},
		{90.2, "፺.፳"},
		{10.345, "፲.፴፬፻፶"},
		{9.41, "፱.፵፩"},
		{1.0, "፩"},
		//{0, "፨"},
		{0.0, ""},
		{+19.0, "፲፱"},
		{-1931.0, "-፲፱፻፴፩"},
		{100900.0, "፲፻፱፻"},
		//{10000, "ሐ"},
	}
	//fmt.Println("---Fmtfloat testing---")
	for _, v := range tests {
		got := Fmtfloat(v.in)
		if got != v.want {
			t.Errorf("want %s got %s", v.want, got)
		}
		//fmt.Printf("want %s got %s\n", v.want, got)
	}

}
func TestPrsfloat(t *testing.T) {
	tests := []struct {
		in   string
		want float64
	}{
		{"", 0},
		{"+", 0},
		{"-", 0},
		{"-፲፱፻፴፩", -1931},
		{"+፴", 30},
		{"፲፱.፲", 19.1},
		{"፱.", 9},
		{"፻", 100},
		{"፻.፲፩", 100.11},
		{"፲፻፱፻", 100900},
	}
	//fmt.Println("---Prsfloat testing---")
	for _, v := range tests {
		got := Prsfloat(v.in)
		if got != v.want {
			t.Errorf("want %v got %v", v.want, got)
		}
		//fmt.Printf("want %v got %v\n", v.want, got)
	}

}

func ExamplePrsint() {
	i := Prsint("፲፱፻፴፩")
	fmt.Println(i)
	//output: 1931
}
func ExampleFmtfloat() {
	i := Fmtfloat(9.41)
	fmt.Println(i)
	//output: ፱.፵፩
}
func ExamplePrsfloat() {
	i := Prsfloat("-፱፻፴.፵፩")
	fmt.Println(i)
	//output: -930.41
}
