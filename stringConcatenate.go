package main;

import "fmt";

func concatenateString(str1 []string , str2 []string)([]string){
        // var traceIndex int =0;
        concatString := make([]string, len(str1)+len(str2));
        for i:=0;i<len(str1);i++{
            concatString[i]=str1[i];
        }
        for i:=0;i<len(str2);i++{
            concatString[len(str1)+i]=str2[i];
        }
        return concatString;
        
    
}
func main(){
    first:=[]string{"puneet","singh","h","v","ss"};
    second:=[]string{"array","array12","e","t"};
    fmt.Println("-----line no 19-----",concatenateString(first,second))
}
