package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job — задание на обработку одного URL.
type Job struct {
	ID  int
	URL string
}

// Result — результат обработки задания.
type Result struct {
	Job      Job
	Status   string
	Duration time.Duration
}

// Количество одновременно работающих воркеров (Worker Pool).
const workerCount = 5

// worker читает задания из канала jobs, имитирует HTTP-запрос
// и отправляет результат в канал results.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		start := time.Now()

		// Заглушка HTTP-запроса: случайная задержка от 100 до 500 мс.
		delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
		time.Sleep(delay)

		results <- Result{
			Job:      job,
			Status:   "обработан",
			Duration: time.Since(start),
		}
	}
}

func main() {
	// Список URL-адресов для обработки.
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://netology.ru",
		"https://wikipedia.org",
		"https://youtube.com",
		"https://reddit.com",
		"https://stackoverflow.com",
		"https://amazon.com",
	}

	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup

	// Fan-out: запускаем фиксированное число воркеров.
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Отправляем все задания и закрываем канал jobs,
	// чтобы воркеры завершили цикл for range.
	for i, url := range urls {
		jobs <- Job{ID: i + 1, URL: url}
	}
	close(jobs)

	// Fan-in: когда все воркеры закончат, закрываем канал результатов.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Собираем результаты в слайс.
	var collected []Result
	for res := range results {
		collected = append(collected, res)
	}

	// Итоговый отчёт.
	fmt.Println("=== Отчёт по обработке URL ===")
	var total time.Duration
	for _, res := range collected {
		fmt.Printf("[%2d] %-30s %-10s %v\n",
			res.Job.ID, res.Job.URL, res.Status, res.Duration.Round(time.Millisecond))
		total += res.Duration
	}

	fmt.Println("------------------------------")
	fmt.Printf("Всего обработано: %d\n", len(collected))
	if len(collected) > 0 {
		avg := total / time.Duration(len(collected))
		fmt.Printf("Среднее время выполнения: %v\n", avg.Round(time.Millisecond))
	}
}
