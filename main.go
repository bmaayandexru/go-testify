package main

import (
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	// читаем строку в countStr
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		// count строки нет -> ошибка
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	// преобразуем строку в число
	count, err := strconv.Atoi(countStr)
	if err != nil {
		// ошибка преобразования
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	// читаем город в строку
	city := req.URL.Query().Get("city")
	// извлекаем из карты значение
	cafe, ok := cafeList[city]
	if !ok {
		// нет значения в карте
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	// тут в cafe слайс строк кафе
	if count > len(cafe) {
		/*
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("запрошено больше чем есть кафе"))
			return
		*/
		// так было в старом обработчике
		count = len(cafe)
	}

	// формируем ответ слияние слайса строк через ","
	answer := strings.Join(cafe[:count], ",")

	// возвращаем ствтус и ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}
