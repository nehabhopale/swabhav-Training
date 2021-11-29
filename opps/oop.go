//inheritance by structure embedding
type A struct{
	name string
}
type B struct{
	id string  // here i is small so we can't export it somewhere else directly
	A
}
func (b B)details(){
	fmt.Println(b.id)
	fmt.Println(b.A.name) //we exported Println from fmt package

}
func main(){
	A:=A{"neha"}
	B:=B{"1"}
	B.details()
}
//encapsulation by keeping cap letters if we are exporting from somewhere and if we want to keep it private the small letters