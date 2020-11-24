package model

import (
	database "auth/db"
	"encoding/json"
	"fmt"
	"log"
	"auth/db"
)

type Music struct {
	ID     string
	Name   string
	Path   string
	UserID uint
}

func GetCollectionFromUser(id string) ([]Music, error) {

	log.Println("Trying to recover collection ...")
	var collection []Music
	con := db.Connect()
	query := fmt.Sprintf(`SELECT id, name, path, user_id from music WHERE user_id = %s`, id)
	result, err := con.Query(query)
	if err != nil {
		log.Fatal(err)
		return collection, err
	}
	defer result.Close()

	for result.Next() {
		var music Music
		err := result.Scan(&music.ID, &music.Name, &music.Path, &music.UserID)

		if err != nil {
			log.Fatal(err)
			return collection, err
		} else {
			collection = append(collection, music)
		}
	}
	return collection, nil
}

func (music *Music) SearchingAllSongs() ([]Music, error) {

	var musics []Music
	reply, err := database.Get("music")

	if err != nil {
		log.Println("Searching on mysql")
		var err error
		musics, err := GetCollectionFromUser("1")
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
