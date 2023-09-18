ch := make(chan int, 3) // Создание буферизированного канала с емкостью 3

go func() {
	ch <- 1
	ch <- 2
	ch <- 3
}()

// Горутина может продолжать выполнение, так как буфер еще не заполнен

value1 := <-ch // Получаем данные из канала
value2 := <-ch
value3 := <-ch

fmt.Println(value1, value2, value3)