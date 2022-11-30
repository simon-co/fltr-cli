package command

import (
	"fmt"
	"testing"
)

func TestComponentTypesToMap(t *testing.T){
  m := ComponentTypes{}.ToMap()

  for k := range m {
    fmt.Println(k)
  }
}

func TestComponentMapToList(t *testing.T){
  j :=  ComponentTypes{}.ToMap().ToList()

  for _, v := range j {
    fmt.Println(v)
  }
}
