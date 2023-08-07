
package main

import "encoding/json"

type Photo struct {
    AlbumID      int    `json:"albumId"`
    ID           int    `json:"id"`
    Title        string `json:"title"`
    URL          string `json:"url"`
    ThumbnailURL string `json:"thumbnailUrl"`
}

func toJson(val []byte) Photo {
    photo := Photo{}
    err := json.Unmarshal(val, &photo)
    if err != nil {
        panic(err)
    }
    return photo
}

