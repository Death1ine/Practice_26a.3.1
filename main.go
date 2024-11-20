package main

import (
	"time"

	"20.2.1/pipeline"
)

const (
	BUFFER_SIZE    = 5               // Размер буфера
	FLUSH_INTERVAL = 5 * time.Second // Интервал опустошения буфера
)

func main() {
	pipeline.Pipeline(BUFFER_SIZE, FLUSH_INTERVAL)
}
