package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"html/template"
	
	"github.com/gorilla/mux"
)



// from https://mholt.github.io/json-to-go/
// converts JSON into a Go type definition

type WeatherData struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type ForecastData struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC          float64 `json:"maxtemp_c"`
				MaxtempF          float64 `json:"maxtemp_f"`
				MintempC          float64 `json:"mintemp_c"`
				MintempF          float64 `json:"mintemp_f"`
				AvgtempC          float64 `json:"avgtemp_c"`
				AvgtempF          float64 `json:"avgtemp_f"`
				MaxwindMph        float64 `json:"maxwind_mph"`
				MaxwindKph        float64 `json:"maxwind_kph"`
				TotalprecipMm     float64 `json:"totalprecip_mm"`
				TotalprecipIn     float64 `json:"totalprecip_in"`
				TotalsnowCm       float64 `json:"totalsnow_cm"`
				AvgvisKm          float64 `json:"avgvis_km"`
				AvgvisMiles       float64 `json:"avgvis_miles"`
				Avghumidity       float64 `json:"avghumidity"`
				DailyWillItRain   int     `json:"daily_will_it_rain"`
				DailyChanceOfRain int     `json:"daily_chance_of_rain"`
				DailyWillItSnow   int     `json:"daily_will_it_snow"`
				DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
				Condition         struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				Uv float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise          string `json:"sunrise"`
				Sunset           string `json:"sunset"`
				Moonrise         string `json:"moonrise"`
				Moonset          string `json:"moonset"`
				MoonPhase        string `json:"moon_phase"`
				MoonIllumination string `json:"moon_illumination"`
			} `json:"astro"`
			Hour []struct {
				TimeEpoch int     `json:"time_epoch"`
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				IsDay     int     `json:"is_day"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				WindMph      float64 `json:"wind_mph"`
				WindKph      float64 `json:"wind_kph"`
				WindDegree   int     `json:"wind_degree"`
				WindDir      string  `json:"wind_dir"`
				PressureMb   float64 `json:"pressure_mb"`
				PressureIn   float64 `json:"pressure_in"`
				PrecipMm     float64 `json:"precip_mm"`
				PrecipIn     float64 `json:"precip_in"`
				Humidity     int     `json:"humidity"`
				Cloud        int     `json:"cloud"`
				FeelslikeC   float64 `json:"feelslike_c"`
				FeelslikeF   float64 `json:"feelslike_f"`
				WindchillC   float64 `json:"windchill_c"`
				WindchillF   float64 `json:"windchill_f"`
				HeatindexC   float64 `json:"heatindex_c"`
				HeatindexF   float64 `json:"heatindex_f"`
				DewpointC    float64 `json:"dewpoint_c"`
				DewpointF    float64 `json:"dewpoint_f"`
				WillItRain   int     `json:"will_it_rain"`
				ChanceOfRain int     `json:"chance_of_rain"`
				WillItSnow   int     `json:"will_it_snow"`
				ChanceOfSnow int     `json:"chance_of_snow"`
				VisKm        float64 `json:"vis_km"`
				VisMiles     float64 `json:"vis_miles"`
				GustMph      float64 `json:"gust_mph"`
				GustKph      float64 `json:"gust_kph"`
				Uv           float64 `json:"uv"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}


type ResponseCurrentData struct {
	Name		string		`json:"name"`
	TempC		float64 	`json:"temp_c"`
	Text		string		`json:"text"`
	Icon		string		`json:"icon"`
}


type ResponseForecastData struct {
	Name		string	`json:"name"`
	TempAvgs	[]float64	`json:"temp_avgs"`
}


// get the weather(temperature) for the city today
func weatherapicurrent(city string) ResponseCurrentData {
	// http://api.weatherapi.com/v1/current.json?key=6810fd11c3c545fe97504231220211&q=Beijing
	key := "6810fd11c3c545fe97504231220211"
	url := "https://api.weatherapi.com/v1/"
	page := "current.json"
	url = url + page
	
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("q", city)
	q.Add("key", key)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherdata WeatherData
	err = json.Unmarshal(responseBody, &weatherdata)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: we can use "-" in WeatherData to ignore fields that we don't want export to json.
	var responsecurrentdata ResponseCurrentData
	responsecurrentdata.Name = city
	responsecurrentdata.TempC = weatherdata.Current.TempC
	responsecurrentdata.Icon = weatherdata.Current.Condition.Icon
	responsecurrentdata.Text = weatherdata.Current.Condition.Text
	return responsecurrentdata
}

// get the weather(average temperature) for the city for the next few days
func weatherapiforecast(city string, forecast string) ResponseForecastData {
	// http://api.weatherapi.com/v1/current.json?key=6810fd11c3c545fe97504231220211&q=Beijing&days=3
	key := "6810fd11c3c545fe97504231220211"
	url := "https://api.weatherapi.com/v1/"
	page := "forecast.json"
	url = url + page
	
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("q", city)
	q.Add("key", key)
	q.Add("days", forecast)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var forecastdata ForecastData
	err = json.Unmarshal(responseBody, &forecastdata)
	if err != nil {
		log.Fatal(err)
	}

	var responseforecastdata ResponseForecastData
	responseforecastdata.Name = city;
	for i := 0; i < len(forecastdata.Forecast.Forecastday); i++ {
		responseforecastdata.TempAvgs = append(responseforecastdata.TempAvgs, forecastdata.Forecast.Forecastday[i].Day.AvgtempC)
	}
	
	return  responseforecastdata;
}

// main page
func welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}


func queryCurrentUI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	responsedata := weatherapicurrent(city)

	t, err := template.ParseFiles("weather-today.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, responsedata)
	if err != nil {
		log.Fatal(err)
	}
}


func queryCurrent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	responsedata := weatherapicurrent(city)

	json.NewEncoder(w).Encode(responsedata)
}


func queryForecastUI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	forecast := vars["forecast"]
	responsedata := weatherapiforecast(city, forecast)

	t, err := template.ParseFiles("weather-forecast.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, responsedata)
	if err != nil {
		log.Fatal(err)
	}

}

func queryForecast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	forecast := vars["forecast"]
	responsedata := weatherapiforecast(city, forecast)

	json.NewEncoder(w).Encode(responsedata)
}


func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", welcome)
	myRouter.HandleFunc("/api/{city}", queryCurrent)
	myRouter.HandleFunc("/api/{city}/{forecast:[1-9]}", queryForecast)
	myRouter.HandleFunc("/ui/{city}", queryCurrentUI)
	myRouter.HandleFunc("/ui/{city}/{forecast:[1-9]}", queryForecastUI)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

//              |------->welcome
// main->handleRequests->queryCurret
//              |_____-->queryForecast
func main() {
	handleRequests()
}
