package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Forest struct {
	Trees []*Tree
}

func TrainForest(data [][]interface{}, labels []string, samples, features, trees int) *Forest {
	rand.Seed(time.Now().UnixNano())
	forest := &Forest{}
	forest.Trees = make([]*Tree, trees)
	done_flag := make(chan bool)
	mutex := &sync.Mutex{}
	for i := 0; i < trees; i++ {
		go func(x int) {
			fmt.Printf("Entrenando árbol %v \n", x+1)
			forest.Trees[x] = TrainTree(data, labels, samples, features)
			fmt.Printf("Arbol %v está listo\n", x+1)
			mutex.Lock()
			mutex.Unlock()
			done_flag <- true
		}(i)
	}

	for i := 1; i <= trees; i++ {
		<-done_flag
	}

	return forest
}
