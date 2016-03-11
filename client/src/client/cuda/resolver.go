package cuda
import "fmt"

var left_direction   int = 1100
var right_direction  int = 1001
var both_directions  int = 2002

type Symbol struct {
    Value           string
    SearchDirection int
    MaxCount        int
    Accepter        func(string)(bool)
    Breaker         func(string)(bool)
}

var URL_SPEC_CHARS = []string {"%","=",":","/","@","?","#"}

type Cyclone struct {
    // line prop
}

func DataHeaderSelector(first_table_string []string)(data [][]int, isTableHeader bool ) {
    return
}

/*func UrlSelector(str []string, delim []int,  data_before []int , data_after []int)(data [][]int, isUrl bool ) {
    fmt.Printf("Delim:%v StrPart:%v", delim,str[delim[0]:delim[1]])
    return
}

func UrlMatcher(str []string, delim []int ) {

    match:=str[delim[0]:delim[1]]
    fmt.Printf("match:%v  str:%v  delim:%v",match,str,delim )


}*/

func StringArrayIsEqual (abc , def []string) (bool) {

    return true


}

func UrlFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    fmt.Printf("\n line: %v \n  delims: %v \n  data: %v \n",lineAsArray,delims,data)
    return ndelims,ndata

}

func ArrayInArrayIndexes (abc []string, phrases ...[]string )(indexes [][]int) {
    // In case when phrases have same prefix result is filling by longest phrase 
    // freaks attack

    if (len(abc) < 1 )||(len(phrases) < 1){return}
    //  first_match := -1
    //  last_match  := -1
    for i := range abc {
        symbol:=abc[i]
        var found [][]string

        for p := range phrases {
            var local_found []string
            phrase:=phrases[p]
            if len(phrase) > 1 {
                for z:= range phrase {
                    zsymbol:=phrase[z]
                    if symbol == zsymbol {
                        //xi:=i
                        for xi := range abc {
                            xi      = i
                            xsymbol := abc[xi]


                        }
                        local_found = append(local_found, symbol)
                    } else {
                        if z == (len(phrase)-1) {

                        }
                    }
                }
            }
        }
    }
    return
}

//func AnalyzeDelims()
