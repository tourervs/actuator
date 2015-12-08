package cuda
import "strings"
//import "fmt"

/*
var splitted_by_space  int = 0
var splitted_by_colon  int = 1
*/

var OPEN_SECTION_SQUARE      int = 0   //   [
var CLOSE_SECTION_SQUARE     int = 1   //   ]
var OPEN_SECTION_TRIANGLE    int = 2   //   <
var OPEN_SL_SECTION_TRIANGLE int = 3   //   </
var CLOSE_SECTION_TRIANGLE   int = 4   //   >
var OPEN_SECTION_ROUND       int = 5   //   (
var CLOSE_SECTION_ROUND      int = 6   //   )
var OPEN_SECTION_CURLY       int = 7   //   {
var CLOSE_SECTION_CURLY      int = 8   //   {

var SINGLE_QUOTE             int = 8   //   '
var DOUBLE_QUOTE             int = 9   //   "
var GRAVE_QUOTE              int = 10   //   `

var HYPHEN                   int = 11  // -
var MINUS                    int = 11  // -
var PLUS                     int = 12  // +
var UNDERSCORE               int = 13  // _
var EQUAL                    int = 14  // = 
var COLON                    int = 15  // : 
var SEMICOLON                int = 16  // ;
var COMMA                    int = 17  // ,
var DOT                      int = 18  // .

var SLASH                    int = 19  // /
var BACKSLASH                int = 20  // \
var PIPE                     int = 21  // |
var ASTERISK                 int = 22  // *
var NUMBER                   int = 23  // #
var PENIS                    int = 24  // o|o
var DOLLAR                   int = 25  // $
var AMPERSAND                int = 26  // &
var SUFFIX                   int = 27  // ^
var PERCENT                  int = 28  // %
var MAIL                     int = 29  // @
var EXCLAM                   int = 30  // !
var TILDE                    int = 31  // ~

//var SIGNS = make(map[int]string)
//SIGNS[OPEN_SECTION_SQUARE]="["
//SIGNS := map[int]string{ EXCLAM : "!" }




func SignMap()(signs map[int]string) {
    signs=make(map[int]string)
    signs[OPEN_SECTION_SQUARE]      ="["
    signs[CLOSE_SECTION_SQUARE]     ="]"
    signs[OPEN_SECTION_TRIANGLE]    ="<"
    signs[OPEN_SL_SECTION_TRIANGLE] ="</"
    signs[CLOSE_SECTION_TRIANGLE]   =">"
    signs[OPEN_SECTION_ROUND]       ="("
    signs[CLOSE_SECTION_ROUND]      =")"
    signs[OPEN_SECTION_CURLY]       ="{"
    signs[CLOSE_SECTION_CURLY]      ="}"

    signs[SINGLE_QUOTE]             ="'"
    signs[DOUBLE_QUOTE]             =`"`
    signs[GRAVE_QUOTE]              ="`"

    signs[HYPHEN]                   ="-"
    signs[PLUS]                     ="+"
    signs[UNDERSCORE]               ="_"
    signs[EQUAL]                    ="="
    signs[COLON]                    =":"
    signs[SEMICOLON]                =";"
    signs[COMMA]                    =","
    signs[DOT]                      ="."

    signs[SLASH]                    ="/"
    signs[BACKSLASH]                =`\`
    signs[PIPE]                     ="|"
    signs[ASTERISK]                 ="*"
    signs[NUMBER]                   ="#"
    signs[DOLLAR]                   ="$"
    signs[AMPERSAND]                ="&"
    signs[SUFFIX]                   ="^"
    signs[PERCENT]                  ="%"
    signs[MAIL]                     ="@"
    signs[EXCLAM]                   ="!"
    signs[TILDE]                    ="~"

    return


}

func GetKeyByValue(signs map[int]string, string_value string) (key int) {


    for key, value :=range signs {
        if value == string_value {
            return key
        }

    }
    return -1
}

func ValueExists(signs map[int]string,value string)(found bool ) {

    values:=GetMapValues(signs)
    for i := range values {
        if values[i]==value {
            found=true
        }
    }
    return found

}

func GetMapValues(signs map[int]string)(values []string ){

    for _, value := range signs {
        values=append(values, value)
    }
    return values
}

//var SQ_CU
func GetSignIndex(entry string)(map[int][]int) {

   sign_map:=SignMap()
   sign_indexes:=make(map[int][]int)
   lineAsArray:=strings.Split(entry,"")
   for i := range lineAsArray {
       char:=lineAsArray[i]
       charSignKey:=GetKeyByValue(sign_map, char)
       if charSignKey > 0 {
           if _, ok := sign_indexes[charSignKey]; ok==false {
               sign_indexes[charSignKey]= []int {}
           }
           sign_indexes[charSignKey]=append(sign_indexes[charSignKey], i)

       }

    }
    return sign_indexes

}



func AcidPriorityAnalyzer (entry string) () {
    //lineSplitBySpace:=
    //lineSplitByQuote:=
    //lineSplitByColon:=

}

func AcidSequencer(entry string) ()  {


}

//func 

/*

var SECTION_SQUARE_OPEN = int  0 







*/

func SortByNested ( entry string ) () {




}

func GetSignScope( lineAsArray []string, sign int, sign_pos int) (scope [][2]int) {

    switch {
        case sign==EQUAL:
            var first_part [2]int
            var last_part  [2]int
            first_part[0] = 0
            first_part[1] = sign_pos-1
            last_part[0] = sign_pos+1
            last_part[1] = len(lineAsArray)-1
            scope=append(scope,first_part)
            scope=append(scope,last_part)
            return scope
    }
    return scope

}