package day04
import "testing"
import "fmt"



type Person struct {
    Name string
    Age int
    Sex string
}

func (p Person) SayHi() {
    fmt.Printf("Hi, my name is %s, I'm %d years old\n", p.Name, p.Age)
}
func TestStruct6(t *testing.T) {
	var p Person
	p.Name = "Tom"
	p.Age = 18
	p.Sex = "男"
	p.SayHi()  
}

// 结构体作为函数参数  构造函数
func newPerson(name string, age int, sex string) Person {
    return Person{
        Name: name,
        Age: age,
        Sex: sex,
    }
}

func TestStruct7(t *testing.T) {
    p := newPerson("Tom", 18, "男")
    p.SayHi()
}

//recover default value
func setPersonInfo(p *Person) {
    p.Name = "Tom"
    p.Age = 20
    p.Sex = "男"
}

func (p *Person) NewYear() {
    p.Age++
}

func TestStruct8(t *testing.T) {
    var p Person
    setPersonInfo(&p)
    p.SayHi()
}

func TestStruct9(t *testing.T) {
    var p Person
	p = newPerson("Tom", 18, "男")
    p.NewYear()
    p.SayHi()
}