package handler

var (
	weatherAPIBaseURL string
	apiKey            string
)

func Init() {
	weatherAPIBaseURL = "https://api.weatherapi.com/v1/current.json"
	apiKey = "F1015a1c9ab04d269cb12534250511"
	// apiKey = os.Getenv("WEATHER_API_KEY")
}
