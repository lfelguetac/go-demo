package utils

import (
	. "go-demo/model"
	"os"
	"strconv"
)

func GetStringEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetBoolEnv(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if ok {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return boolValue
	} else {
		return defaultValue
	}
}

func DeleteFirstClient(userSessionsData []SessionData, client string) []SessionData {
	for i := range userSessionsData {
		if filterClient(userSessionsData[i], client) {
			return append(userSessionsData[:i], userSessionsData[i+1:]...)
		}
	}
	return userSessionsData
}

func filterClient(sessionData SessionData, client string) bool {
	return client != "" && sessionData.Client == client
}
