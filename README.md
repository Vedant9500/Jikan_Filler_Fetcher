# Jikan Filler Fetcher

This is a simple Go program that uses the [Jikan API](https://jikan.moe) to fetch filler episodes for a given anime. It queries the MyAnimeList (MAL) episode data, identifies filler episodes, and displays them in a readable format.

## Features
- Fetches all episodes of an anime using its MAL ID.
- Identifies and lists filler episodes.
- Handles pagination for anime with many episodes.

## Prerequisites
- Go 1.19 or later.

## Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/Vedant9500/Jikan_Filler_Fetcher
   cd Jikan_Filler_Fetcher
   ```
2. Initialize and download dependencies:
   ```bash
   go mod init jikan_filler_fetcher
   ```

## Usage
1. Build and run the program:
   ```bash
   go run main.go <anime_id>
   ```
   Replace `<anime_id>` with the MAL ID of the anime you want to fetch filler episodes for. For example, Naruto's MAL ID is `20`.

   Example:
   ```bash
   go run main.go 20
   ```
2. The program will output the list of filler episodes for the given anime.

### Sample Output
For Naruto (anime ID `20`):
```
Fetching filler episodes for anime ID 20...
Found 91 filler episodes:
Episode 26: Special Report: Live from the Forest of Death!
Episode 97: Kidnapped! Naruto's Hot Spring Adventure!
...
```

## Error Handling
- If an invalid anime ID is provided, the program will terminate with an error message.
- If the Jikan API is unavailable or responds with an error, the program will notify the user.

## Notes
- Some episodes may not have accurate filler status if the anime data is incomplete on MyAnimeList.

## Contributing
Feel free to open issues or submit pull requests if you'd like to improve this project.

## Acknowledgments
- [Jikan API](https://jikan.moe) for providing the data.
- MyAnimeList for the anime database.
