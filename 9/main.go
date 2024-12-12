package main

import (
	"fmt"
	"strconv"
)

type file struct {
	length int
	id int
	moved bool
}

func main() {
	input := sampleInput
	input = realInput

	files := make([]int, 1000000)
	inFile := true
	fileIndex := 0
	pos := 0
	spaces := make([]int, 0, 1000000)

	files2 := make([]file, 100000)
	pos2 := 0

	for _,c := range input {
		num, _ := strconv.Atoi(string(c))
		val := fileIndex
		if !inFile { 
			val = -1 
		}
		for range num {
			files[pos] = val
			if !inFile {
				spaces = append(spaces, pos)
			}
			pos++
		}
		files2[pos2] = file{num, val, false}
		pos2++
		inFile = !inFile
		if inFile { fileIndex++}
	}

	files2 = files2[:pos2]
/*
	for i := range pos { 
		if files[i] < 0 { fmt.Print(".")} else { fmt.Print(files[i]) }
	}
	fmt.Println()
*/
	for _,f := range files2 {
		for j:=0;j<f.length;j++ {
			if f.id < 0 { fmt.Print(".")} else { fmt.Print(f.id)}
		}
	}
	fmt.Println()

	/*
	origPos := pos
	spacePos := 0
	for spacePos < len(spaces) && spaces[spacePos] < pos {
		pos--
		for(files[pos] == -1) { pos--}
		s := spaces[spacePos]
		if s >= pos { break }
		spacePos++
		files[s] = files[pos]
		files[pos] = -1
	}*/

	for pi := pos2-1;pi>=0;pi-- {
		f := files2[pi]
		if f.moved || f.id < 0 { continue}
		for qi :=0; qi < pi; qi++ {
			f2 := files2[qi]
			if f2.id == -1 && f2.length >= f.length {
				diff := f2.length - f.length
				f.moved = true
				files2[qi] = f
				files2[pi] = file{f.length, -1, true}
				if diff != 0 {
					files3 := append([]file{}, files2[:qi+1]...)
					files3 = append(files3, file{diff, -1, false})
					files2 = append(files3, files2[qi+1:]...)
				}

				/*
				for _,f := range files2 {
					for j:=0;j<f.length;j++ {
						if f.id < 0 { fmt.Print(".")} else { fmt.Print(f.id)}
					}
				}*/
				fmt.Print(".")
				//fmt.Println()
				break
			}
		}
	}

	checksum := 0
	pos = 0
	for _,f := range files2 {
		for j:=0;j<f.length;j++ {
			if f.id < 0 { 
				fmt.Print(".")
			} else { 
				fmt.Print(f.id)
				checksum += pos*f.id
			}
			pos++
		}
	}
	fmt.Println()
	fmt.Println(checksum)


	_ = input
}

