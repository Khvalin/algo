package main

import (
	"fmt"
	"strings"
)

const wall = "#"

type point struct  {
	x,y int
}

func toolAssistedSpeedrun(stage []string) int {
	addFrame := func (stage []string) []string  {
		top := strings.Repeat(wall, len(stage[0]) + 2)
		
		for i := 0; i< len (stage); i++ {
			stage[i] = wall + stage[i] + wall
		}
		stage = append([]string{top} , stage...)
		stage = append( stage, top)

		return stage
	}

	if 0== len(stage) || 0 == len(stage[0]){
		return -1
	}

	stage = addFrame(stage)
//	fmt.Println(stage)
	return -1
}


func main()  {
	fmt.Println("7 == ", toolAssistedSpeedrun([]string{"##########", 
 "##########", 
 "##########", 
 "###F######", 
 "### ######", 
 "S      E #", 
 "# ###### #", 
 "#        #", 
 "##########", 
 "##########"}))
}