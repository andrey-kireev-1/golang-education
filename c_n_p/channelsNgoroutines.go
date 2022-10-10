package main

import (
	"fmt"
	"sync"
	"time"
)

func channelNoGoroutine() { //Канал блокируется за неимением другого потока считывающего данные с канала
	dataChannel := make(chan string)
	dataChannel <- "Some Sample Data"
	fmt.Println(<-dataChannel)
}

func typeNValueOfChannel(dataChannel chan int) {
	fmt.Printf("type of `dataChannel` is %T\n", dataChannel)     //выводит тип данных передаваемый каналом
	fmt.Printf("value of `dataChannel` is %v\n", dataChannel)    //выводит адрес канала в памяти
	fmt.Printf("len of `dataChannel` is %v\n", len(dataChannel)) //длина канала (количество данных)
	fmt.Printf("cap of `dataChannel` is %v\n", cap(dataChannel)) //емкость канала (максимальное кол-во данных для буферизованного канала)
}

func inputChannelIntoFunc(dataChannel chan string) { //Передача канала в функцию
	dataChannel <- "Some Sample Data"
}

func buffChannel() { //Буфферизованный канал не блокируется

	dataChannelBuf := make(chan string, 3)
	dataChannelBuf <- "Some Sample Data"
	dataChannelBuf <- "Some Other Sample Data"
	dataChannelBuf <- "Buffered Channel"
	fmt.Println(<-dataChannelBuf)
	fmt.Println(<-dataChannelBuf)
	fmt.Println(<-dataChannelBuf)
}

func panicInClosedChannel(dataChannel chan int) { //panic "send on closed channel", паника при передаче в закрытый канал
	go func() {
		for i := 0; i < 5; i++ {
			dataChannel <- i
		}
		// close(dataChannel)
		// for i := 5; i < 10; i++ {
		// 	dataChannel <- i
		// }
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-dataChannel)
	}
}

func infLoopNOkValue(dataChannel chan int) { //пример использования бесконечного цикла + ok значения
	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i * i
		}
		close(dataChannel)
	}()
	for {
		v, ok := <-dataChannel
		if !ok {
			break
		} else {
			fmt.Println(v)
		}
	}
	fmt.Println("Inf Loop stopped")
}

func exampleOfNoBuffChannel() { //пример того что в no buff канале может лежать только 1 значение
	dataChannel := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
			dataChannel <- i
		}
	}()
	time.Sleep(time.Second * 3)
}

func channelInChannel() { // Канал в канале
	cc := make(chan chan string) // Канал передающий тип данных Канал
	go func(cc chan chan string) {
		c := make(chan string)
		fmt.Println("Channel <- Channel")
		cc <- c // В канал передаем канал
	}(cc)
	c := <-cc //Принимаем канал в канал
	go func(c chan string) {
		fmt.Println("Channel <- String")
		c <- "String data" // Принимаем строку в принятый канал
	}(c)
	fmt.Println(<-c)
}

func usingSelect() {
	start := time.Now()
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(c chan string) {
		fmt.Println("1 goroutine start", time.Since(start))
		c <- "Hello 1"
	}(chan1)

	go func(c chan string) {
		fmt.Println("2 goroutine start", time.Since(start))
		c <- "Hello 2"
	}(chan2)
	// time.Sleep(time.Nanosecond)
	select {
	case res := <-chan1:
		fmt.Println("First goroutine: ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Second goroutine: ", res, time.Since(start))
	default:
		fmt.Println("No goroutine started")
	}
}

func emptySelect() { //Пустой селект блокирует горутину
	select {}
}

func goroutineWaitGroup() { //использование waitgroup
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, instance int) {
			fmt.Println("Service called on instance", instance)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait() // waitgroup ждет выполнения всех горутин
}

var global int // i == 0

func usingMutex() { //пример захвата горутины мьютексом
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *sync.Mutex) {
			m.Lock() // мьютекс блокирует горутину
			global = global + 1
			m.Unlock() // мьютекс разблокирует горутину
			wg.Done()
		}(&wg, &m)
	}
	wg.Wait()
	fmt.Println("Res of incrementing 1000 times", global)
}
