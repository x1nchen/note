package main

func main() {

	println("start")
	i := 0
LOOP:
	for {
		println("outer loop start", i)
		for {
			println("inner loop start", i)
			if i++; i == 3 {
				break LOOP
			}
			println("inner loop over", i)
		}
		println("outer loop over", i)
	}
}
