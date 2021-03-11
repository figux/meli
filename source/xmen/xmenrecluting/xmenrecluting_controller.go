package xmenrecluting

import (
	"fmt"
	"../person"
)


//parameter: customField string
func IsMutant( p person.Person) bool {

	var mutant bool = false
	var arraySize int = len(p.Dna)
	var stringDna int = 0
	var right int = 0
	var down int = 0
	var diag int = 0

	var counter = 0
	
	for raw := 0; raw < arraySize; raw++ {
		for col := 0; col < arraySize; col++ {
			//moves
			right = col + 3
			down = raw + 3
			diag = col - 3
			
			//right
			if right < arraySize {
				if p.Dna[raw][col] == p.Dna[raw][col+1] && p.Dna[raw][col] == p.Dna[raw][col+2] && p.Dna[raw][col] == p.Dna[raw][col+3]{
					stringDna++
				}
			}

			//down
			if down < arraySize {
				if p.Dna[raw][col] == p.Dna[raw+1][col] && p.Dna[raw][col] == p.Dna[raw+2][col] && p.Dna[raw][col] == p.Dna[raw+3][col]{
					stringDna++
				}
			}

			//diag 1
			if right < arraySize && down < arraySize{
				if p.Dna[raw][col] == p.Dna[raw+1][col+1] && p.Dna[raw][col] == p.Dna[raw+2][col+2] && p.Dna[raw][col] == p.Dna[raw+3][col+3]{
					stringDna++
				}	
			}

			//diag 2
			if down < arraySize && diag >= 0{
				if p.Dna[raw][col] == p.Dna[raw+1][col-1] && p.Dna[raw][col] == p.Dna[raw+2][col-2] && p.Dna[raw][col] == p.Dna[raw+3][col-3]{
					stringDna++
				}	
			}


			counter++


			if stringDna > 1{
				raw = arraySize + 1
				col = arraySize + 1
				mutant = true
			}

		}
	}


	//fmt.Printf("%q ", p.Dna[0][0])
	fmt.Println("Numero de veces en el ciclo:", counter)


	return mutant
}