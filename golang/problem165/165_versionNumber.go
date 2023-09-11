package main

import (
	"fmt"
	"strings"
  "strconv"
)

func main() {
	fmt.Println(compareVersion("1.0", "1.0.0"))
}


func compareVersion(version1 string, version2 string) int {
  version1ToList, version2ToList := strings.Split(version1, "."), strings.Split(version2, ".");
  

  for i:=0; i < len(version1ToList) || i < len(version2ToList); i++{
    version1i, version2i  := 0, 0

    if(len(version1ToList) > i){
      version1i, _= strconv.Atoi(version1ToList[i])
    }

    if(len(version2ToList) > i){
      version2i, _= strconv.Atoi(version2ToList[i])
    }

    if(version1i > version2i){
      return 1
    }

    if(version1i < version2i){
      return -1
    }
  }
  
  return 0;
}


/**
  Note for this problem:

*/
