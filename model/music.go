package model

import (
	database "auth/db"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
)

type Music struct {
	gorm.Model

	ID     uint
	Name   string
	Path   string
	UserID uint
}

func (music *Music) SearchingAllSongs() ([]Music, error) {

	conn := database.Connect()
	defer conn.Close()

	var musics []Music
	reply, err := database.Get("music")

	if err != nil {
		log.Println("Searching on mysql")
		var err error
		var musics []Music
		err = conn.Debug().Model(&Music{}).Limit(100).Find(&musics).Error
		if err != nil {
			return []Music{}, err
		}
		musicsBytes, _ := json.Marshal(musics)
		database.Set("music", musicsBytes)
		return musics, nil
	} else {
		log.Println("Searching on redis")
		json.Unmarshal(reply, &musics)
		return musics, nil
	}
}
