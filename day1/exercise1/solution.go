package main

import "fmt"

type Matrix struct {
	nrows int
	ncolumns int
	elements [][]float64
}

func (m Matrix) getRows() int{
	return m.nrows
}

func (m Matrix) getColumns() int{
	return m.ncolumns
}

func (m *Matrix) setElements(i int,j int,val float64){
	m.elements[i][j]=val
}

func (m3 *Matrix) addMatrix(m1 Matrix,m2 Matrix) {

	for i := range(m1.elements) {
		for j := range (m1.elements[0]) {
			m3.elements[i][j] = m1.elements[i][j] + m2.elements[i][j]

		}
	}
}

func (m Matrix) printMatrixJSON(){
	fmt.Printf("Matrix:{")
	for i := range(m.elements) {
		fmt.Printf("{ ")
		for j := range (m.elements[0]) {
			fmt.Printf("%v ", m.elements[i][j])
		}
		fmt.Printf("}")
	}
	fmt.Printf("}")
}

func main(){
	m1:=Matrix{nrows: 2,ncolumns: 3, elements: [][]float64{{1,2,3},{2,3,4},}}
	m2:= Matrix{nrows: 2,ncolumns: 3, elements: [][]float64{{5,6,7},{8,9,10},}}
	fmt.Println(m1.getRows())
	fmt.Println(m1.getColumns())
	m2.setElements(1,2,6)
	m3:=Matrix{nrows: m1.nrows, ncolumns:m1.ncolumns,elements: [][]float64{{0,0,0},{0,0,0},}}
	m3.addMatrix(m1,m2)
	fmt.Println(m3.elements)
	m1.printMatrixJSON()
}
