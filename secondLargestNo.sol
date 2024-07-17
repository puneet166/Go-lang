package main;
import "fmt";

func secondLargest(array [10]int)(int){
    var firstNo=0;
    var secondNo=0;
    for i:=0;i<len(array);i++{
        if(array[i]>firstNo){
            secondNo=firstNo;
            firstNo=array[i];
        }else{
            if(array[i]>secondNo && array[i]<firstNo){
         secondNo=array[i];   
        }
        }
        
    }
    return secondNo;
}

func main(){
    array:=[10]int{0,1,2,3,4,5,6,7,8,9};
    fmt.Println("-------",secondLargest(array));
    
}
