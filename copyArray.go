package main;
import "fmt";

    type  copyArrayInterface interface{
         copyArray(array []int)([]int);
    }
    func copyArray(array []int)([]int){
        // use this make function when you dont know the size of array.
        arrayCopy := make([]int, len(array))
        for i:=0;i<len(array);i++{
            arrayCopy[i]=array[i];
        }
        return arrayCopy
    }
    func main(){
    array:=[]int{1,2,3,4,5};
    fmt.Println("----",copyArray(array));
}
