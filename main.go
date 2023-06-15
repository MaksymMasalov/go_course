package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем текущее время с использованием NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// Обработка ошибки: выводим ее в STDERR
		log.Println("Ошибка при получении времени NTP:", err)
		// Возвращаем ненулевой код выхода
		os.Exit(1)
	}

	// Преобразуем полученное время в локальное время
	localTime := ntpTime.In(time.Local)

	// Печатаем текущее время
	fmt.Println("Текущее время:", localTime)
}
