package main;
import "fmt";

func returnSlice(no []int)([]int){
    return no
}

func main(){
    array:=[]int{1,2,3,4,5,6};
    returnData:=returnSlice(array);
    // slicing in go lang
    fmt.Println(returnData[1:3]);
}
