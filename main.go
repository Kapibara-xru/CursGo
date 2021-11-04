package main

import (
	http "CursGo/http"
	tasks "CursGo/tasks"
	"fmt"
	"os"
	"time"
)

const (
	task1    = "Циклическая ротация"
	task2    = "Поиск отсутствующего элемента"
	task3    = "Проверка последовательности"
	task4    = "Чудные вхождения в массив"
	taskAll  = "Решить все задачи"
	nickName = "Kapibara_xru"
)

func main() {
	var task int
	for {
		fmt.Println("Выберите задачу для решения:")
		fmt.Printf("1) %s \n", task1)
		fmt.Printf("2) %s \n", task2)
		fmt.Printf("3) %s \n", task3)
		fmt.Printf("4) %s \n", task4)
		fmt.Printf("5) %s \n", taskAll)
		fmt.Println("Для выхода введите 0")
		fmt.Fscan(os.Stdin, &task)

		switch task {
		case 1:
			fmt.Println(solution(task1))
		case 2:
			fmt.Println(solution(task2))
		case 3:
			fmt.Println(solution(task3))
		case 4:
			fmt.Println(solution(task4))
		case 5:
			solutionAll()
			time.Sleep(time.Millisecond)
		case 0:
			return
		}
		time.Sleep(time.Second)
	}
}

func solutionFirstTask(req []interface{}) []interface{} {
	var k int
	var numbers []int
	var resArray []interface{}
	var numArray []interface{}
	var result []interface{}

	for _, res := range req {
		resArray = res.([]interface{})
		numArray = resArray[0].([]interface{})
		for _, num := range numArray {
			numbers = append(numbers, int(num.(float64)))
		}
		k = int(resArray[1].(float64))

		result = append(result, tasks.Solution1(numbers, k))
		numbers = nil
	}
	return result
}

func solutionTasks(task string, req []interface{}) []interface{} {
	var numbers []int
	var resArray []interface{}
	var numArray []interface{}
	var result []interface{}

	for _, res := range req {
		resArray = res.([]interface{})
		numArray = resArray[0].([]interface{})
		for _, num := range numArray {
			numbers = append(numbers, int(num.(float64)))
		}
		switch task {
		case task2:
			result = append(result, tasks.Solution2(numbers))
		case task3:
			result = append(result, tasks.Solution3(numbers))
		case task4:
			result = append(result, tasks.Solution4(numbers))
		}
		numbers = nil
	}
	return result
}

func solutionAll() {
	go func(msg string) {
		fmt.Println(msg)
	}(solution(task1))

	go func(msg string) {
		fmt.Println(msg)
	}(solution(task2))

	go func(msg string) {
		fmt.Println(msg)
	}(solution(task3))

	go func(msg string) {
		fmt.Println(msg)
	}(solution(task4))
}

func solution(task string) (res string) {
	switch task {
	case task1:
		url := "http://116.203.203.76:3000/tasks/" + task
		payload := http.HttpGetRequest(url)
		result := solutionFirstTask(payload)
		res = http.HttpPostRequst(nickName, task, payload, result)
	default:
		url := "http://116.203.203.76:3000/tasks/" + task
		payload := http.HttpGetRequest(url)
		result := solutionTasks(task, payload)
		res = http.HttpPostRequst(nickName, task, payload, result)
	}
	return
}
