package main;

import "fmt";

func leftRight(array []int)([]int){
    var arrayInFunction=make([]int , len(array));
    var left =0;
    var right= len(array)-1;
    for i:=0;i<len(array);i++{
        if(array[i]>=0){
            arrayInFunction[left]=array[i];
            left++;
        }else{
            arrayInFunction[right]=array[i];
            right--;
        }
    }
    return arrayInFunction;
}
func main(){
    array:=[]int{-1,2,3,4,-5,4,-2};
    fmt.Println(leftRight(array));
}
