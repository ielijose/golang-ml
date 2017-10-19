package main

import (
	"fmt"
	"os"
	"strconv"
	"./packages"
)

func main() {
	var slice []int
	var t *tree.Tree

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			p, err := strconv.Atoi(os.Args[i])
		
			if err != nil {
				fmt.Print("Input error!")
			} else {
				slice = append(slice, p)				
			}			
		}

		for i := 0; i < len(slice); i++ {
			t = tree.Add(t, slice[i])
		}

		tree.InOrder(t)
	}	
}