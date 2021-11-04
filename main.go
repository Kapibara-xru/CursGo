package main

import (
	http "CursGo/http"
	tasks "CursGo/tasks"
	"fmt"
	"os"
	"strings"
)

const task1 = "Циклическая ротация"
const task2 = "Поиск отсутствующего элемента"
const task3 = "Проверка последовательности"
const task4 = "Чудные вхождения в массив"
const taskAll = "Решить все задачи"
const nickName = "Kapibara_xru"

func main() {
	var task int
	fmt.Println("Выберите задачу для решения:")
	fmt.Printf("1) %s \n", task1)
	fmt.Printf("2) %s \n", task2)
	fmt.Printf("3) %s \n", task3)
	fmt.Printf("4) %s \n", task4)
	fmt.Printf("5) %s \n", taskAll)
	fmt.Fscan(os.Stdin, &task)

	switch task {
	case 1:
		url := "http://116.203.203.76:3000/tasks/" + task1
		payload := http.HttpGetRequest(url)
		result := strings.ReplaceAll(fmt.Sprintf("%v", solutionFirstTask(payload)), " ", ",")
		fmt.Println(http.HttpPostRequst(nickName, task1, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result))
	case 2:
		url := "http://116.203.203.76:3000/tasks/" + task2
		payload := http.HttpGetRequest(url)
		result := strings.ReplaceAll(fmt.Sprintf("%v", solutionTasks(task2, payload)), " ", ",")
		fmt.Println(http.HttpPostRequst(nickName, task2, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result))
	case 3:
		url := "http://116.203.203.76:3000/tasks/" + task3
		payload := http.HttpGetRequest(url)
		fmt.Println("Получил GET")
		result := strings.ReplaceAll(fmt.Sprintf("%v", solutionTasks(task3, payload)), " ", ",")
		fmt.Println("Получил результат")
		fmt.Println(http.HttpPostRequst(nickName, task3, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result))
	case 4:
		url := "http://116.203.203.76:3000/tasks/" + task4
		payload := http.HttpGetRequest(url)
		result := strings.ReplaceAll(fmt.Sprintf("%v", solutionTasks(task4, payload)), " ", ",")
		fmt.Println(http.HttpPostRequst(nickName, task4, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result))
	case 5:
		msg := make(chan string, 4)

		go func() {
			url := "http://116.203.203.76:3000/tasks/" + task1
			payload := http.HttpGetRequest(url)
			result := strings.ReplaceAll(fmt.Sprintf("%v", solutionFirstTask(payload)), " ", ",")
			msg <- http.HttpPostRequst(nickName, task1, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result)
		}()

		go func() {
			url := "http://116.203.203.76:3000/tasks/" + task2
			payload := http.HttpGetRequest(url)
			result := strings.ReplaceAll(fmt.Sprintf("%v", solutionTasks(task2, payload)), " ", ",")
			msg <- http.HttpPostRequst(nickName, task2, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result)
		}()

		go func() {
			url := "http://116.203.203.76:3000/tasks/" + task3
			payload := http.HttpGetRequest(url)
			result := strings.ReplaceAll(fmt.Sprintf("%v", solutionTasks(task3, payload)), " ", ",")
			msg <- http.HttpPostRequst(nickName, task3, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result)
		}()

		go func() {
			url := "http://116.203.203.76:3000/tasks/" + task4
			payload := http.HttpGetRequest(url)
			result := strings.ReplaceAll(fmt.Sprintf("%v", solutionTasks(task4, payload)), " ", ",")
			msg <- http.HttpPostRequst(nickName, task4, strings.ReplaceAll(fmt.Sprintf("%v", payload), " ", ","), result)
		}()

		fmt.Println(<-msg)
		fmt.Println(<-msg)
		fmt.Println(<-msg)
		fmt.Println(<-msg)
	}
}

func solutionFirstTask(req []interface{}) [][]int {
	var k int
	var numbers []int
	var resArray []interface{}
	var numArray []interface{}
	var result [][]int

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

func solutionTasks(task string, req []interface{}) []int {
	var numbers []int
	var resArray []interface{}
	var numArray []interface{}
	var result []int

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
