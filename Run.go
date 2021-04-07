package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"text/tabwriter"
)

func main() {
	fmt.Println("平方根: ", Sqrt2(2), Sqrt2Other(2), math.Sqrt(2))
	SetAccountInformation("demo", "demo")
	fmt.Println("全局变量: ", GetAccount(), GetPassWord())
	fmt.Println(os.Getwd())
	i := 10
	j := 11
	fmt.Println(rand.Int(), math.Sqrt2, i, j)

	w := tabwriter.NewWriter(os.Stdout, 15, 4, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmith\t")
	w.Flush()
}
