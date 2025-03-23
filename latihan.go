package main
import "fmt"

func genap(n int) int {
    if n < 2 {
        return 0
    }
    if n%2 == 0 {
        return n + genap(n-2)
    } else {
        return genap(n - 1)
    }
}

func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(genap(num))
}