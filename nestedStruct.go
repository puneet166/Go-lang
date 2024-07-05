package main 

import "fmt"
// const(
//     black =iota,
//     white )
type colorType struct{
     colorTypeinBool bool;
}
type structname struct{
    name string;
    age uint64;
    height float64;
    color colorType;
}
func initStruct(name string, age uint64,height float64, colorTypebool bool) structname{
    structNameDupplicate:=structname{name:"puneet singh",age:65,height:5.3,color:colorType{colorTypeinBool:colorTypebool}};
    // fmt.Println("-----------structNameDupplicate-----",structNameDupplicate);
    return structNameDupplicate;
}
func main(){
    fmt.Println("-----name ---------");
    fmt.Println(initStruct("puneet",65,5.9,true));
}
