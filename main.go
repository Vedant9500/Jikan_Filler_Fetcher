package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Episode struct {
	EpisodeID  int    `json:"mal_id"`
	Title      string `json:"title"`
	IsFiller   bool   `json:"filler"`
}

type EpisodesResponse struct {
	Data []Episode `json:"data"`
	Pagination struct {
		HasNextPage bool `json:"has_next_page"`
	} `json:"pagination"`
}

func fetchFillerEpisodes(animeID int) ([]Episode, error) {
	baseURL := fmt.Sprintf("https://api.jikan.moe/v4/anime/%d/episodes", animeID)
	var fillerEpisodes []Episode
	page := 1

	for {
		url := fmt.Sprintf("%s?page=%d", baseURL, page)
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error fetching episodes: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
		}

		var episodesResp EpisodesResponse
		if err := json.NewDecoder(resp.Body).Decode(&episodesResp); err != nil {
			return nil, fmt.Errorf("error decoding response: %v", err)
		}

		for _, episode := range episodesResp.Data {
			if episode.IsFiller {
				fillerEpisodes = append(fillerEpisodes, episode)
			}
		}

		if !episodesResp.Pagination.HasNextPage {
			break
		}
		page++
	}

	return fillerEpisodes, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <anime_id>", os.Args[0])
	}

	animeID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid anime ID: %v", err)
	}

	fmt.Printf("Fetching filler episodes for anime ID %d...\n", animeID)
	fillerEpisodes, err := fetchFillerEpisodes(animeID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if len(fillerEpisodes) == 0 {
		fmt.Println("No filler episodes found.")
		return
	}

	fmt.Printf("Found %d filler episodes:\n", len(fillerEpisodes))
	for _, episode := range fillerEpisodes {
		fmt.Printf("Episode %d: %s\n", episode.EpisodeID, episode.Title)
	}
}
