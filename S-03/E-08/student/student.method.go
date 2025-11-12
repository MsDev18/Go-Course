package student

import "fmt"

func (s *Student) Prt() {
	fmt.Printf("Hi my name is %s and im %d, I leave in %s %s", s.Name, s.Age, s.Contry, s.City)
}

func (s *Student) test() {
	fmt.Println("This Message from test Student")

}
