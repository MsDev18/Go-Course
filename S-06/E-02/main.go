package main

type User struct {
	ID   uint
	Name string
}

func main() {
	//fmt.Println(square(1))
	//fmt.Println(square(2))
	//fmt.Println(square(3))
	//fmt.Println(square(4))
}

//func square(i int) int {
//	if i > 9988 && i < 9999 {
//		return i
//	}
//	return i * i
//}

func dayOfWeek(i int) string {
	switch i {
	case 1:
		return "شنبه"
	case 2:
		return "یک شنبه"
	case 3:
		return "دو شنبه"
	case 4:
		return "سه شنبه"
	case 5:
		return "چهار شنبه"
	case 6:
		return "پنج شنبه"
	case 7:
		return "جمعه"
	default:
		return ""
	}
}
