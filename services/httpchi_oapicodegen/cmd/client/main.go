package main

import (
	"context"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"

	weatherV1 "github.com/baizhigit/go-ms-examples/httpchi_oapi_codegen/pkg/openapi/weather/v1"
)

const (
	serverURL       = "http://localhost:8080"
	defaultCityName = "Almaty"
	defaultMinTemp  = -10
	defaultMaxTemp  = 40
)

func main() {
	ctx := context.Background()

	// Инициализация oapi-codegen клиента
	client, err := weatherV1.NewClientWithResponses(serverURL)
	if err != nil {
		log.Fatalf("❌ Ошибка при создании клиента: %v", err)
	}

	log.Println("=== Тестирование API для работы с данными о погоде ===")
	log.Println()

	// 1. Пытаемся получить данные о погоде (которых пока нет)
	log.Printf("🌦️ Получение данных о погоде для города %s\n", defaultCityName)
	log.Println("===================================================")

	weatherResp, err := client.GetWeatherByCityWithResponse(ctx, defaultCityName)
	if err != nil {
		log.Printf("❌ Ошибка при получении погоды: %v\n", err)
		return
	}

	// Проверяем статус ответа
	if weatherResp.StatusCode() == http.StatusNotFound {
		if weatherResp.JSON404 != nil {
			log.Printf("ℹ️ Данные о погоде для города %s не найдены: %s\n",
				defaultCityName, weatherResp.JSON404.Message)
		} else {
			log.Printf("ℹ️ Данные о погоде для города %s не найдены\n", defaultCityName)
		}
	} else if weatherResp.JSON200 != nil {
		log.Printf("Данные о погоде для города %s: %+v\n", defaultCityName, weatherResp.JSON200)
	}

	// 2. Обновляем данные о погоде
	log.Printf("🔄 Обновление данных о погоде для города %s\n", defaultCityName)
	log.Println("=====================================================")

	// Создаем запрос на обновление погоды
	temperature := gofakeit.Float32Range(defaultMinTemp, defaultMaxTemp)
	updateRequest := weatherV1.UpdateWeatherByCityJSONRequestBody{
		Temperature: temperature,
	}

	updatedResp, err := client.UpdateWeatherByCityWithResponse(ctx, defaultCityName, updateRequest)
	if err != nil {
		log.Printf("❌ Ошибка при обновлении погоды: %v\n", err)
		return
	}

	if updatedResp.StatusCode() != http.StatusOK {
		log.Printf("❌ Неожиданный статус: %d\n", updatedResp.StatusCode())
		return
	}

	if updatedResp.JSON200 != nil {
		log.Printf("✅ Данные о погоде обновлены: %+v\n", updatedResp.JSON200)
	}

	// 3. Получаем обновленные данные о погоде
	log.Printf("🌦️ Получение обновленных данных о погоде для города %s\n", defaultCityName)
	log.Println("===========================================================")

	weatherResp, err = client.GetWeatherByCityWithResponse(ctx, defaultCityName)
	if err != nil {
		log.Printf("❌ Ошибка при получении погоды: %v\n", err)
		return
	}

	if weatherResp.StatusCode() == http.StatusOK && weatherResp.JSON200 != nil {
		log.Printf("✅ Получены данные о погоде: %+v\n", weatherResp.JSON200)
		log.Println("Тестирование завершено успешно!")
	} else {
		log.Printf("❌ Неожиданный ответ: статус %d\n", weatherResp.StatusCode())
	}
}
