package main

import "fmt"

func main() {
	var angka, now, past, i int
	var naik, turun, both bool

	fmt.Scan(&angka)
	for {
		now = angka % 10
		angka /= 10
		i++
		if i == 1 {
			past = now
		}
		if past > now {
			if turun {
				both = true
			}
			naik = true
		} else if now > past {
			if naik {
				both = true
			}
			turun = true
		}
		past = now
		if angka == 0 {
			break
		}
	}
	if both {
		fmt.Println("tidak terurut")
	} else if turun {
		fmt.Println("descending")
	} else if naik {
		fmt.Println("ascending")
	}

}
