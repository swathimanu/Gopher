package reverse

// import "fmt"

func Reverse(s string) string{
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i,j = i+1, j-1{
		rev[i], rev[j] = rev[j], rev[i]

	}
	return string(rev)
}
/* func main() {

	fmt.Println(Reverse("Golang"))
	
}*/
