package main

import "fmt"

//欧几里德公式，求最大公约数
func main() {
	var p, q = 65, 64
	fmt.Printf("p:%v,q:%v,gcd:%v", p, q, gcd(p, q))
}

func gcd(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(p, r)
}
