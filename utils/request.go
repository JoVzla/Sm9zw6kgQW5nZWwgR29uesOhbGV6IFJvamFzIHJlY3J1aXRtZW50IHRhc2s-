package utils

import (
	"errors"
	"gogoapps/constants"
	"gogoapps/logger"
	"gogoapps/models/config"
	"gogoapps/models/weathermodel"
	"io/ioutil"
	"net/http"

	"github.com/patrickmn/go-cache"
)

func GetCityData(city string, configuration config.Configuration) (w weathermodel.Weather, e error) {

	//Recover if panics
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				e = errors.New(x)
			case error:
				e = x
			default:
				e = errors.New("Unknown panic")
			}
			logger.Log.Error("Recovered in GetCityData", r)
		}
	}()

	url := constants.BASE_URL + "?q=" + city + "&appid=" + configuration.ApiKey

	data, err := http.Get(url)
	checkError(err)

	body, err := ioutil.ReadAll(data.Body)
	checkError(err)

	result, err := weathermodel.UnmarshalWeather(body)
	checkError(err)

	return result, err
}

func CheckCache(city string, cach *cache.Cache) error {

	_, found := cach.Get(city)

	if found {
		msg := "still in cache"
		logger.Log.Error(msg)
		return errors.New(msg)
	}

	cach.Set(city, city, cache.DefaultExpiration)

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
