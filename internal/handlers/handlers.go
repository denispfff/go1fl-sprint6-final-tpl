package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(res http.ResponseWriter, req *http.Request, logger *log.Logger) {
	// http.ServeFile(res, req, "index.html") - вроде и так можно
	if req.Method != http.MethodGet {
		// Ошибка на стороне клиента, зачем логировать на сервере
		http.Error(res, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("index.html")
	if err != nil {
		errText := "Ошибка загрузки страницы"
		logger.Printf("%s: %v", errText, err)

		http.Error(res, errText, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf8")
	res.WriteHeader(http.StatusOK)

	res.Write(data)
}

func HandleUpload(res http.ResponseWriter, req *http.Request, logger *log.Logger) {
	if req.Method != http.MethodPost {
		// Ошибка на стороне клиента, зачем логировать на сервере
		http.Error(res, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	req.ParseMultipartForm(10 << 20)

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		errText := "Ошибка при получении файла"
		logger.Printf("%s: %v", errText, err)

		http.Error(res, errText, http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		errText := "Ошибка при чтении файла"
		logger.Printf("%s: %v", errText, err)

		http.Error(res, errText, http.StatusInternalServerError)
		return
	}
	morseReverse := service.ReverseMorse(string(data))
	// Ошибка в требованиях - предлагают использовать time.Now().UTC().String()
	// Расширение берём из расширения исходного файла, хоть в тестовых файлах его нет
	fileName := time.Now().UTC().Format("2006-01-02T15-04-05") + filepath.Ext(handler.Filename)
	resFile, err := os.Create(fileName)
	if err != nil {
		errText := "Ошибка при создании файла"
		logger.Printf("%s: %v", errText, err)
		http.Error(res, errText, http.StatusInternalServerError)
		return
	}

	defer resFile.Close()

	_, err = fmt.Fprint(resFile, morseReverse)
	if err != nil {
		errText := "Ошибка при записи в файл"
		logger.Printf("%s: %v", errText, err)

		http.Error(res, errText, http.StatusInternalServerError)
		return
	}

	//Вернуть результат конвертации строки.
	res.Header().Set("Content-Type", "text/plain; charset=utf8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(morseReverse))
}
