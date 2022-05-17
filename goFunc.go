package main

// type Employee struct{
// 	fname string
// 	Iname string
// }
// func (emp Employee) fullname(){
// 	fmt.Println(emp.fname+" "+emp.Iname)
// }

// func main(){
// 	e1:=Employee{"John","Ponting"}
// 	e1.fullname()
// }
// func fun() int {
// 	return 123456789
// }
// func main(){
// 	x:=fun()
// 	fmt.Println(x)
// }
// func main(){
// 	fmt.Println(addAll(10,15,20,25,30))
// }

// func addAll(args ... int)(int,int)  {
// 	finalAddValue:=0
// 	finalSubValue:=0
// 	for _,value  := range args{
// 	   finalAddValue += value
// 	   finalSubValue -= value
// 	}
// 	return finalAddValue,finalSubValue
//  }
// func main(){
// 	number :=10
// 	squareNum := func()(int){
// 		number*=number
// 		return number
// 	}
// 	fmt.Println(squareNum())
// 	fmt.Println(squareNum())
// }
// func main() {
// 	var x [5]int
// 	var i,j int
// 	for i = 0;i<5;i++{
// 		x[i] = i+10
// }
// for j=0;j<5;j++{
// 	fmt.Printf("Element[%d]=%d\n",j,x[j])
// }
// }
// func main(){
// 	var a = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
// 	var i,j,p int
// 	for i=0;i<3;i++ {
// 		for p=0;p<i;p++ {
// 			fmt.Print(" * ")
// 			}
// 		for j=0;j<3;j++{

// 			fmt.Print(" ",a[i][j])

// 		}
// 		fmt.Println()

// 	}
// }
// type person struct
// {
// 	firstName string
// 	lastName string
// 	age int
// }
// func main(){
// 	x:=person{age:30,firstName: "vijay",lastName: "Anders",}
// 	fmt.Println(x)
// 	fmt.Println(x.firstName)
// }

// type person struct {
// 	fname string
// 	iname string
// }
// type employee struct {
// 	person
// 	empid int
// }
// func (p person) details(){
// 	fmt.Println(p,""+"I am a Person")
// }
// func (e employee) details(){
// 	fmt.Println(e,""+"I am A Employ")
// }
// func main(){
// 	p1:=person{"Raj","Kumar"}
// 	p1.details()
// 	e1:=employee{person: person{"John","Ponting"},empid: 11}
// 	e1.details()
// }
// type vehicle interface{
// 	accelerate()
// }
// func foo(v vehicle){
// 	fmt.Println(v)
// }
// type car struct{
// 	model string
// 	color string
// }
// func (c car ) accelerate() {
// 	fmt.Println("Accelrating?")
// }
// type toyota struct {
// 	model string
// 	color string
// 	speed int
// }
// func (t toyota) accelerate(){
// 	fmt.Println("I am toyota,I accelerate fast?")
// }
// func main(){
// 	c1:=car{"suzuli","blue"}
// 	t1:=toyota{"Toyota","red",100}
// 	c1.accelerate()
// 	t1.accelerate()
// 	foo(c1)
// 	foo(t1)
// }
