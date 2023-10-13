package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

type Graph struct {
	readersCount atomic.Int32      // Лічильник читачів
	writersCount atomic.Int32      // Лічильник письменників
	nodesCount   int               // Кількість вузлів в графі
	graph        map[int]list.List // Мапа для зберігання списку сусідніх вузлів кожного вузла
}

func NewGraph() (*Graph, error) {
	graph := &Graph{}
	graph.graph = make(map[int]list.List)
	return graph, nil
}

func (graph *Graph) updatePrice() {
	// Метод для оновлення цін на ребрах графу.
	fmt.Println("updatePrice() started")
	defer fmt.Println("updatePrice() finished")

	for graph.readersCount.Load() > 0 && graph.writersCount.Load() > 0 {
		// Очікування на виконання інших читачів та письменників.
	}

	graph.writersCount.Add(1)
	for _, edges := range graph.graph {
		front := edges.Front()
		for i := 0; i < edges.Len(); i++ {
			var shouldUpdate bool
			shouldUpdate = rand.Intn(2)%2 == 1

			if shouldUpdate {
				front.Value.([]int)[1] = rand.Int()
				fmt.Printf("Updated price of an edge.")
			}
			fmt.Printf("No need for Updated price of an edge.")
		}
	}
	graph.writersCount.Add(-1)
}

func (graph *Graph) addOrRemoveEdges() {
	// Метод для додавання або видалення ребер графу.
	fmt.Println("addOrRemoveEdges() started")
	defer fmt.Println("addOrRemoveEdges() finished")
	for graph.readersCount.Load() > 0 && graph.writersCount.Load() > 0 {
		// Очікування на виконання інших читачів та письменників.
	}

	graph.writersCount.Add(1)
	for _, edges := range graph.graph {
		if rand.Intn(2)%2 == 1 {
			index := rand.Intn(edges.Len())
			position := edges.Front()
			for i := 0; i <= index; i++ {
				position = position.Next()
			}
			edges.Remove(position)
			fmt.Printf("Removed an edge.")
		}
		if rand.Intn(2)%2 == 1 {
			index := rand.Intn(graph.nodesCount)
			var edge [2]int
			edge[0] = index
			edge[1] = rand.Int()
			edges.PushBack(edge)
			fmt.Printf("Added an edge.")
		}
	}
	graph.writersCount.Add(-1)
}

func (graph *Graph) addOrRemoveNode() {
	// Метод для додавання або видалення вузла графу.
	fmt.Println("addOrRemoveNode() started")
	defer fmt.Println("addOrRemoveNode() finished")
	for graph.readersCount.Load() > 0 && graph.writersCount.Load() > 0 {
		// Очікування на виконання інших читачів та письменників.
	}

	graph.writersCount.Add(1)
	if rand.Intn(2)%2 == 1 {
		graph.nodesCount = graph.nodesCount + 1
		fmt.Printf("Added a node.")
	}
	graph.writersCount.Add(-1)
}

func (graph *Graph) isThereAPath(from int, to int) bool {
	// Метод для перевірки існування шляху між двома вузлами в графі.
	fmt.Println("isThereAPath() started")
	defer fmt.Println("isThereAPath() finished")
	for graph.readersCount.Load() > 0 {
		// Очікування на виконання інших читачів.
	}

	graph.readersCount.Add(1)
	defer graph.readersCount.Add(-1)

	if from < 0 || from >= graph.nodesCount || to < 0 || to >= graph.nodesCount {
		// Перевірка, чи індекси входять в діапазон доступних вузлів.
		return false
	}

	var visited = make([]bool, graph.nodesCount)
	dfs(from, &visited, graph)
	return visited[to]
}

func dfs(v int, visited *[]bool, graph *Graph) {
	// Рекурсивний метод для обходу графу в глибину (DFS).
	(*visited)[v] = true

	l := graph.graph[v]
	front := l.Front()

	for front != nil {
		to := front.Value.([2]int)[0]
		if !(*visited)[to] {
			dfs(to, visited, graph)
		}
	}
}

func main() {
	graph, _ := NewGraph()
	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		graph.updatePrice()
		wg.Done()
	}()
	go func() {
		graph.isThereAPath(0, 10)
		wg.Done()
	}()
	go func() {
		graph.addOrRemoveEdges()
		wg.Done()
	}()
	go func() {
		graph.addOrRemoveNode()
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Всі горутини завершили виконання. Програма завершена.")
}
