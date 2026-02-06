package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	customMiddleware "github.com/baizhigit/go-ms-examples/httpchi_oapi_codegen/internal/middleware"
	weatherV1 "github.com/baizhigit/go-ms-examples/httpchi_oapi_codegen/pkg/openapi/weather/v1"
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

// WeatherStorage представляет потокобезопасное хранилище данных о погоде
type WeatherStorage struct {
	mu       sync.RWMutex
	weathers map[string]*weatherV1.Weather
}

// NewWeatherStorage создает новое хранилище данных о погоде
func NewWeatherStorage() *WeatherStorage {
	return &WeatherStorage{
		weathers: make(map[string]*weatherV1.Weather),
	}
}

// GetWeather возвращает информацию о погоде по имени города
func (s *WeatherStorage) GetWeather(city string) *weatherV1.Weather {
	s.mu.RLock()
	defer s.mu.RUnlock()
	weather, ok := s.weathers[city]
	if !ok {
		return nil
	}
	return weather
}

// UpdateWeather обновляет данные о погоде для указанного города
func (s *WeatherStorage) UpdateWeather(city string, weather *weatherV1.Weather) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.weathers[city] = weather
}

// WeatherHandler implements weatherV1.ServerInterface для обработки запросов к API погоды
type WeatherHandler struct {
	storage *WeatherStorage
}

// NewWeatherHandler создает новый обработчик запросов к API погоды
func NewWeatherHandler(storage *WeatherStorage) *WeatherHandler {
	return &WeatherHandler{
		storage: storage,
	}
}

// GetWeatherByCity обрабатывает запрос на получение данных о погоде по названию города
func (h *WeatherHandler) GetWeatherByCity(w http.ResponseWriter, r *http.Request, city string) {
	weather := h.storage.GetWeather(city)
	if weather == nil {
		respondJSON(w, http.StatusNotFound, weatherV1.NotFoundError{
			Code:    404,
			Message: "Weather for city '" + city + "' not found",
		})
		return
	}

	respondJSON(w, http.StatusOK, weather)
}

// UpdateWeatherByCity обрабатывает запрос на обновление данных о погоде по названию города
func (h *WeatherHandler) UpdateWeatherByCity(w http.ResponseWriter, r *http.Request, city string) {
	// Manual request parsing
	var req weatherV1.UpdateWeatherByCityJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, weatherV1.BadRequestError{
			Code:    400,
			Message: "Invalid request body",
		})
		return
	}

	weather := &weatherV1.Weather{
		City:        city,
		Temperature: req.Temperature,
		UpdatedAt:   time.Now(),
	}

	h.storage.UpdateWeather(city, weather)

	respondJSON(w, http.StatusOK, weather)
}

// Helper function for JSON responses
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func main() {
	// Создаем хранилище для данных о погоде
	storage := NewWeatherStorage()

	// Создаем обработчик API погоды
	weatherHandler := NewWeatherHandler(storage)

	// Инициализируем роутер Chi
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(customMiddleware.RequestLogger)

	// Монтируем обработчики oapi-codegen
	weatherV1.HandlerFromMux(weatherHandler, r)

	// Запускаем HTTP-сервер
	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout, // Защита от Slowloris атак - тип DDoS-атаки, при которой
		// атакующий умышленно медленно отправляет HTTP-заголовки, удерживая соединения открытыми и истощая
		// пул доступных соединений на сервере. ReadHeaderTimeout принудительно закрывает соединение,
		// если клиент не успел отправить все заголовки за отведенное время.
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
