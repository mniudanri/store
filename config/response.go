package config

import (
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
)


func SetResponseByInterface(w http.ResponseWriter, T interface{}, statusCode int, message string) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)

  // Default response structure
  json.NewEncoder(w).Encode(bson.M{"data": T,"message": message})
}

func WriteResponseMessage(w http.ResponseWriter, statusCode int, message string) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)

  // Default response structure
  json.NewEncoder(w).Encode(bson.M{"message": message})
}
