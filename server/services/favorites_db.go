package services

import (
	"fmt"   
	"task/system"
	
)

type Favorite struct {
	ID string `json:"id"` 
    AnimeID      string `json:"anime_id"`
	AnimeTitle   string `json:"anime_title"`
	AnimeImageURL string `json:"anime_image_url"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func dest(fav *Favorite) []interface{} {
    return []interface{}{
        &fav.ID,
        &fav.AnimeID,
		&fav.AnimeTitle,
		&fav.AnimeImageURL,
        &fav.Created,
        &fav.Updated,
    }
}

func selectAllFavorites() ([]Favorite, error) {
    rows, err := system.Db.Query("select * from favorites")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var favs []Favorite = make([]Favorite, 0)
    for rows.Next() {
        fav := Favorite{}
        err = rows.Scan(dest(&fav)...)
        if err != nil {
            return nil, err
        }
        favs = append(favs, fav)
    }
    return favs, nil
}

func InsertFavorite(fav *Favorite) error {
	_, err := system.Db.Exec(
		`INSERT INTO favorites (id, anime_id, anime_title, anime_image_url) 
		 VALUES (?, ?, ?, ?)`,
		fav.ID,
		fav.AnimeID,
		fav.AnimeTitle,
		fav.AnimeImageURL,
		fav.Created,
		fav.Updated,
	)
	return err
}

func DeleteFavorite(id string) error {
	fmt.Println("DeleteFavorite inside", id)
    _, err := system.Db.Exec(`DELETE FROM favorites WHERE id = ?`, id)
    if err != nil {
        fmt.Printf("Error deleting favorite with id %s: %v\n", id, err)
        return err
    }
    return nil
}

func CountFavorites() (int, error) {
	var count int
	err := system.Db.QueryRow("SELECT COUNT(*) FROM favorites").Scan(&count)
	if err != nil {
		fmt.Printf("Error executing QueryRow: %v\n", err)
		return 0, err
	}
	fmt.Println("count", count)
	return count, nil
}