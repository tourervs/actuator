package main
import "client/cuda"
import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    //line:="  2  2 2 2   2 22    32 32 3 23 2 32                               333"
    line2:=`Defaults        secure_path="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"`
    //line3:="                                     "
    //line3:="}}  }}"
    lineAsArray2:=strings.Split(line2, "")
    //lineAsArray3:=strings.Split(line3, "")
    cuda.DebugPrintCharCounter(line2)
    //parser:=cuda.MakeParser("[")

    delims,data:=cuda.GetIndexes(lineAsArray2)
    fmt.Printf("\ndelims: %v\n data: %v \n" , delims , data)

    /*for i := range delims_indexes{

        fmt.Printf("--\n%v\n--",cuda.GetFixedArrayChars(lineAsArray2, delims_indexes[i]))



    }*/
    //fmt.Printf("\nEscapeSpaces|%v| len|%d| \n",cuda.Escape_Spaces(lineAsArray3), len(cuda.Escape_Spaces(lineAsArray3)))

    /*data:=cuda.PrepareData(lineAsArray2, delims_indexes)

    fmt.Printf("\n data :\n %v \n",data)
    */
    //fmt.Printf("--%v--",cuda.Escape_Spaces(lineAsArray2))
}