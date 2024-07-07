package main;
import "fmt";
import "errors"

func returnType(name string)(string,error){
    if len(name)<5{
        return "" ,errors.New("length should be less than 5");
    }
    return name,nil
}

func main(){
    var name string;
    name="yy";
    returnValue , err:= returnType(name);
    if err!=nil{
        fmt.Println(err);
    }else{
    fmt.Println(returnValue);
    }
    }
