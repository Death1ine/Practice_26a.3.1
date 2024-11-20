package buffer

import (
	"sync"
	"time"
)

type CircularBuffer struct {
	buffer       []int
	size         int
	mutex        sync.Mutex
	flushChan    chan []int
	flushTimeout time.Duration
}

// NewCircularBuffer создает новый кольцевой буфер
func NewCircularBuffer(size int, flushTimeout time.Duration) *CircularBuffer {
	return &CircularBuffer{
		buffer:       make([]int, 0, size),
		size:         size,
		flushChan:    make(chan []int),
		flushTimeout: flushTimeout,
	}
}

// Add добавляет элемент в буфер и, при необходимости, выполняет его опустошение
func (cb *CircularBuffer) Add(item int) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.buffer = append(cb.buffer, item)
	if len(cb.buffer) >= cb.size {
		cb.flush()
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

// flush очищает буфер и отправляет данные в канал для обработки
func (cb *CircularBuffer) flush() {
	if len(cb.buffer) > 0 {
		cb.flushChan <- cb.buffer
		cb.buffer = make([]int, 0, cb.size)
	}
}

// FlushChan возвращает канал для отправки данных потребителю
func (cb *CircularBuffer) FlushChan() <-chan []int {
	return cb.flushChan
}
