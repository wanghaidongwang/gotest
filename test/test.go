package main

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)
type rect struct {
	width int
	height int
}

type error interface {
	Error() string
}
func (b *rect) area() int {
	return b.width * b.height
}

func main() {
      r :=rect{}
     result , err := r.area()

	//fmt.Println("Please input your full name: ")
	//fmt.Scanln(&firstName,&lastName)
	//// fmt.Scanf(“%s %s”, &firstName, &lastName)
	//fmt.Printf("Hi %s %s!\n", firstName,lastName)
	//fmt.Sscanf(input, format, &f, &i, &s)
	//fmt.Println("From the string we read: ", f, i, s)

	//var str,stt string
	//	for {
	//		fmt.Println("Please input something:")
	//		fmt.Scanln(&str)
	//		fmt.Scanln(&stt) //取址符是必须的
	//		fmt.Printf("str:%s\n" , str)
	//		fmt.Printf("stt:%s\n" , stt)
	//	}


	//r := rect{}
	//fmt.Println("请输入长方形的宽，长")
	//fmt.Scanf("%d,%d", &r.width, &r.height)
	//fmt.Println("area: ", r.area())
//	time.Sleep(10 * time.Second)

	//var   a , b int
	//fmt.Scanf("%d,%d", &a, &b)
	//r := rect{width: a , height: b }
	//fmt.Println("area: ", r.area())
	//time.Sleep(10 * time.Second)
}
