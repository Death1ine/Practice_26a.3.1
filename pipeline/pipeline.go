package pipeline

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"20.2.1/buffer"
	"20.2.1/filters"
)

// InputDataSource читает данные с консоли и отправляет их в канал
func InputDataSource() <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		scanner := bufio.NewScanner(os.Stdin)
		for {
			// Вывод приглашения для ввода
			fmt.Print("Введите целое число (или 'exit' для выхода): ")
			if !scanner.Scan() {
				break
			}
			input := scanner.Text()

			// Обрабатываем завершение ввода
			if input == "exit" {
				log.Println("Завершение ввода данных.")
				break
			}

			// Пытаемся преобразовать ввод в целое число
			num, err := strconv.Atoi(input)
			if err != nil {
				// Если ввод не корректен, просим ввести целое число
				log.Println("Некорректный ввод. Пожалуйста, введите целое число.")
				continue
			}

			// Отправляем число в канал
			log.Printf("Получено число: %d\n", num)
			output <- num
		}
	}()
	return output
}

// DataConsumer выводит полученные массивы данных в консоль
func DataConsumer(input <-chan []int, done chan<- struct{}) {
	go func() {
		for data := range input {
			// Выводим массив данных
			log.Printf("Получены данные: %v\n", data)
		}
		done <- struct{}{}
	}()
}

// Pipeline реализует конвейер обработки данных
func Pipeline(bufferSize int, flushInterval time.Duration) {
	// Инициализация буфера
	circularBuffer := buffer.NewCircularBuffer(bufferSize, flushInterval)
	log.Printf("Инициализация конвейера с буфером размером %d и интервалом %v\n", bufferSize, flushInterval)
	go circularBuffer.FlushPeriodically()

	// Источник данных
	source := InputDataSource()

	// Применение фильтров
	filteredNegatives := filters.FilterNegativeNumbers(source)
	filteredMultiplesOfThree := filters.FilterNotMultipleOfThree(filteredNegatives)

	// Добавляем отфильтрованные данные в буфер
	done := make(chan struct{})
	DataConsumer(circularBuffer.FlushChan, done) // Передаем канал, а не вызываем его как функцию

	for num := range filteredMultiplesOfThree {
		log.Printf("Добавление числа %d в буфер\n", num)
		circularBuffer.Add(num)
	}

	close(done) // Закрываем канал завершения после обработки
}
