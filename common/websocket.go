package common

import "github.com/olahol/melody"

var Mel *melody.Melody

func CreateWebSocket() *melody.Melody {
	Mel = melody.New()
	return Mel
}

func GetWebSocket() *melody.Melody {
	return Mel
}
