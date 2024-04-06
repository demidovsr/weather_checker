package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"` // Температура
	} `json:"main"`
	Name string `json:"name"` // Название города
}

func main() {

	err := godotenv.Load("secret.env")
	if err != nil {
		fmt.Println("Ошибка загрузки файла secret.env:", err)
		return
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("API ключ не найден.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input city: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in connection to the Server: ", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Impossible to get wether for this city!")
		return
	}

	var weatherData WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		fmt.Println("Error in decoding data: ", err)
		return
	}

	fmt.Printf("Current weather in %s: %.2f°C\n", weatherData.Name, weatherData.Main.Temp)

}
