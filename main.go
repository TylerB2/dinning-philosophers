package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The Dinning philosopher problem is well known in computer science circles.
Five philosophers, numbered from 0 to 4 live in a house where the table
is laid for them, Each philosopher has thier own place at the table.
Their only difficulty - besides those of philosophy - is that the dish
served is a very difficult kind of spaghetti which has to be eaten with
two forks, There are two forks next to each plate, so that presents no
difficulty.
As a consequence, however this means that no two neighbours may be eating simultaneaosly
*/

//List of philosophers
var philosophers = []string{"plato", "Socretes", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderFinished []string
var orderMutex sync.Mutex

//The number of meals
const hunger = 3

func main() {
	//Print intro
	fmt.Println("The Dinning Philosophers Problem")
	fmt.Println("=================================")

	wg.Add(len(philosophers))

	leftFork := &sync.Mutex{}

	//one go routine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		//create mutex for right fork
		rightFork := &sync.Mutex{}

		go dinningProblem(philosophers[i], leftFork, rightFork)

		leftFork = rightFork
	}
	wg.Wait()

	fmt.Println("Every One has left the table....")
	fmt.Println("")
	//the order in which philosophers finished
	fmt.Println("Philosopher in order of leaving")
	fmt.Println("===============================")

	for i, phil := range orderFinished {
		fmt.Printf("%d) %s\n", i, phil)
	}

}

func dinningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	//Print a message
	fmt.Printf("%s is seated\n", philosopher)
	time.Sleep(sleepTime)

	//philospher meals
	for i := hunger; i > 0; i-- {
		fmt.Printf("%s is hungry\n", philosopher)
		time.Sleep(sleepTime)

		//Lock both forks for philosphere to eat
		leftFork.Lock()
		fmt.Printf("\t%s picked up the left fork\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the right fork\n", philosopher)

		//philosopher is now eating
		fmt.Printf("%s has both forks and is eating\n", philosopher)
		time.Sleep(eatTime)

		//give the philosopher some time to think
		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)

		//Unlock the forks
		rightFork.Unlock()
		fmt.Printf("\t%s has put down the right fork\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s has put down the left fork\n", philosopher)
	}

	//print out done message
	fmt.Printf("%s is satisfied\n", philosopher)
	time.Sleep(sleepTime)
	fmt.Printf("%s has left the table\n", philosopher)
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}
