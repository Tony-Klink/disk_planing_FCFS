/*
Модель дискового планирования №1
1. Исходные данные:
	- количество цилиндров диска - 1024
	- стартовый цилиндр - 500
	- случайное поступление запросов к дорожкам (генерировать датчиком случайных чисел)
	- среднее время поиска цилиндра - 10 мс
	- скорость вращения - 7200 об/мин
	- дорожка имеет 500 секторов
	- производится чтение файла размером 1 Мбайт
	- дисциплина обслуживания запросов FIFO
2. Результаты работы модели включают:
	- график перемещения головок по цилиндрам
	- количество пересеченных дородек
	- среднее время чтения файла
	- среднее время поиска файла
*/
// head = 500
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"code.google.com/p/draw2d/draw2d" // библиотека, реализующая рисование в стиле HTML5 Canvas
	"image"
	"image/png"
)

const (
	w, h     = 1200, 2000
	findTime = 0.000002778 //преобразованная скорость поиска цилиндра
	speed    = 120.24      //преобразованная скорость вращения
)
// Функция, реализующая алгоритм
func fcfs() {
	var req [400]float64
	var i, n, c, h int
	var head, move float64
	var g [400]float64
	head = 500

	fmt.Println("Текущая позиция головки: ", head)
	h = int(head)

	fmt.Printf("Введите количество запросов: ")
	fmt.Scanf("%d", &n)
	move = 0
	fmt.Println("Общее количество цилиндров: 1024.")

	diskaccess := 0
	for i := 0; i < n; i++ {
		req[i] = float64(rand.Int63n(1023))
	}
	fmt.Printf("Строка запросов: ")
	for i, c = 0, 0; i < n; i, c = i+1, c+1 {
		fmt.Printf("%d ", int(req[i]))
		move += math.Abs(head - req[i])
		head = req[i]
		g[c] = req[i]
		diskaccess = diskaccess + 1
	}
	fmt.Println("\n")
	fmt.Println("Порядок обслуживания:", h)
	for i = 0; i < n; i++ {
		fmt.Printf("%v ", int(g[i]))
	}
	delta_time := move * findTime * speed
	fmt.Println("\n")
	fmt.Println("Количество пересеченных дорожек: ", int(move))
	fmt.Println("Среднее время чтения файла", delta_time)

	//Методы, отвечающие за построение графика движения головок

	d, gc := initGc(w, h)

	gc.MoveTo(0.0, 0.0)
	gc.LineTo(1200.0, 00.0)
	gc.LineTo(1200.0, 1200.0)
	gc.LineTo(0.0, 1200.0)
	gc.LineTo(0.0, 0.0)
	gc.FillStroke()
	gc.MoveTo(10, 10)
	gc.LineTo(10, 490.0)
	gc.LineTo(490, 490)
	gc.Stroke()
	gc.SetLineDash(nil, 0.0)
	gc.MoveTo(float64(h)/2, delta_time*20)
	gc.LineTo(g[0]/2, delta_time*20+delta_time*20)
	for k, v := range g {
		gc.LineTo(v+1/2, float64(k+1)*delta_time*20+delta_time*20)
		if k > n-2 {
			break
		}
	}
	gc.Stroke()
	saveToPngFile("FCFS_graph.png", d)
}
// Функции сохранения изображения графика
func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Файл %s успешно записан.\n", filePath)
}

func initGc(w, h int) (image.Image, draw2d.GraphicContext) {
	d := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(d)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	return d, gc
}

func menu() {
	var f int
req:
	fmt.Println("Реализация модели дискового планирования FCFS")
	fmt.Println("Введите желаемое действие")
	fmt.Println("1) '1' - для выполнения алгоритма")
	fmt.Println("2) '2' - для завершения программы")
	fmt.Printf("Ваше решение: ")
	fmt.Scanf("%d", &f)
	switch f {
	case 1:
		fcfs()
		goto req
	}

}

func main() {
	//fcfs()
	menu()
}
