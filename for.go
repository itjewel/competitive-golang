// normal for
for i:0; i<10; i++ {
	fmt.Println(i)
}

// only condition
i := 0
for i<5 {
	fmt.Println(i)
	i++
}

// infinite loop

for {
	Println("infinite loop")
}

i := 0
// no initial values
for ; i<10; {
	fmt.Println(i)
	i++
}

// Like while

i: 0;
for i<10 {
	fmt.Println(i)
	i++
}

// If with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}