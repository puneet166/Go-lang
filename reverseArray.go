package main;

import "fmt";

func reverseArray(array []int)([]int){
    var arrayInFunction=make([]int , len(array));
   
    for i:=len(array)-1;i>=0;i--{
       arrayInFunction[len(array)-(i+1)]=array[i];
    }
    return arrayInFunction;
}
func main(){
    array:=[]int{-1,2,3,4,-5,4,-2};
    fmt.Println(reverseArray(array));
}
