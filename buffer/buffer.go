package buffer

import (
	"log"
	"sync"
	"time"
)

type CircularBuffer struct {
	buffer       []int
	size         int
	mutex        sync.Mutex
	FlushChan    chan []int // Измените на с заглавной буквы
	flushTimeout time.Duration
}

// NewCircularBuffer создает новый кольцевой буфер
func NewCircularBuffer(size int, flushTimeout time.Duration) *CircularBuffer {
	log.Println("Создан новый кольцевой буфер с размером:", size)
	return &CircularBuffer{
		buffer:       make([]int, 0, size),
		size:         size,
		FlushChan:    make(chan []int), // И здесь тоже
		flushTimeout: flushTimeout,
	}
}

// Add добавляет элемент в буфер и, при необходимости, выполняет его опустошение
func (cb *CircularBuffer) Add(item int) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	log.Printf("Добавление элемента %d в буфер\n", item)
	cb.buffer = append(cb.buffer, item)
	if len(cb.buffer) >= cb.size {
		cb.flush()
	}
}

// flush очищает буфер и отправляет данные в канал для обработки
func (cb *CircularBuffer) flush() {
	if len(cb.buffer) > 0 {
		log.Printf("Опустошение буфера. Отправка данных в канал: %v\n", cb.buffer)
		cb.FlushChan <- cb.buffer // Здесь тоже
		cb.buffer = make([]int, 0, cb.size)
	}
}

// FlushPeriodically запускает процесс регулярной очистки буфера с заданным интервалом времени
func (cb *CircularBuffer) FlushPeriodically() {
	ticker := time.NewTicker(cb.flushTimeout)
	for {
		<-ticker.C
		cb.mutex.Lock()
		cb.flush()
		cb.mutex.Unlock()
	}
}
