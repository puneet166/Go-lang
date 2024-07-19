package main;
import "fmt";


func countrepeatedNo(array [10]int)(map[int]int){
    mapping:=make(map[int]int);
    var count =0;
    for i:=0;i<len(array);i++{
        // mapping[array[i]]=0;
        for j:=i;j<len(array);j++{
            if(array[i]==array[j]){
                count=count+1;
            }
        }
        // fmt.Println("------",mapping[array[i]])
        if(mapping[array[i]]==0){
            mapping[array[i]]=count;
        }
                // fmt.Println("--nnn----",mapping[array[i]])

        count=0;
    }
    return mapping;
}

func main(){
    array:=[10]int{3,2,3,4,2,9,9,9,9,0};
   mappingOfCount:=countrepeatedNo(array);
    fmt.Println("------line no 12----",mappingOfCount);
}
