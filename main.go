package main

import (
	"fmt"
	"sort"
	"log"
)

type Data struct {
	ID, ParentID int
	Payload      string
}

type InMemory struct {
	ID, ParentID int
	Payload      string
	Children     []*InMemory
}

var input = []Data{
	{1, 0, "str1"},
	{2, 1, "str2"},
	{3, 1, "str3"},
	{4, 3, "str4"},
	{5, 3, "str5"}}

func main() {
	var IDsToSort []int
	if len(input) == 0 {
		log.Fatalln("Invalid input slice")
	}
	for _, v := range input {
		IDsToSort = append(IDsToSort, v.ParentID)
	}
	sort.Ints(IDsToSort)
	convertToTree(input, IDsToSort[0]).RecursivePrint()
}

func convertToTree(input []Data, firstID int) *InMemory {
	inMemory := new(InMemory)
	for i := range input {
		inMemory.RecursiveAnalyze(input[i], firstID)
	}
	return inMemory
}

func (imForAnalyze *InMemory) RecursiveAnalyze(dataForInsert Data, firstID int) {
	if dataForInsert.ParentID == firstID {
		imForAnalyze.ID = dataForInsert.ID
		imForAnalyze.ParentID = dataForInsert.ParentID
		imForAnalyze.Payload = dataForInsert.Payload
	} else {
		if dataForInsert.ParentID == imForAnalyze.ID {
			imForAnalyze.Children = append(imForAnalyze.Children, &InMemory{ID: dataForInsert.ID, ParentID: dataForInsert.ParentID, Payload: dataForInsert.Payload})
		} else {
			if imForAnalyze.Children != nil {
				for _, v := range imForAnalyze.Children {
					v.RecursiveAnalyze(dataForInsert, firstID)
				}
			}
		}
	}
}

func (imForAnalyze *InMemory) RecursivePrint() {
	fmt.Printf("RecursivePrint node = %#v\n", imForAnalyze)
	if imForAnalyze.Children != nil {
		for _, v := range imForAnalyze.Children {
			v.RecursivePrint()
		}
	}
}
