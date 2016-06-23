package common

import "fmt"
import "crypto/rand"
//import "io/ioutil"
//import "strings"

func GenId()(uuid string,err error) {
    b := make([]byte, 16)
    _,err= rand.Read(b)
    if err!= nil {
        return "",err
    }
    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4],b[4:6],b[6:8],b[8:10],b[10:])
    return uuid, nil
}
/*
func ReadFileLines (filename string) (lines []string,err error){


    content, err := ioutil.ReadFile(filename)

    if err != nil {

        return lines, err

    }

    lines = strings.Split(string(content), "\n")

    return lines,err

}
*/
