package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	dbName := "./bitbucket.sqlite3"
	sevenDays, _ := time.ParseDuration("720h") // 7 days = 168 hours, 30 days = 720 hours

	// delete previous DB (if exists)
	if _, err := os.Stat(dbName); err == nil {
		err := os.Remove(dbName)
		if err != nil {
			log.Fatalf("Error when deleting database: %v", err)
		}
	}

	db, _ :=
		sql.Open("sqlite3", dbName)
	c, _ := db.Prepare("CREATE TABLE IF NOT EXISTS repos (id INTEGER PRIMARY KEY, project TEXT, name TEXT, updated_at INTEGER, updated_at_human TEXT, active TEXT)")
	c.Exec()

	accessToken := getAPIToken()
	currentPage := 0

	for {
		currentPage = currentPage + 1
		currentPageStr := strconv.Itoa(currentPage)
		log.Println("Current page", currentPageStr)
		rr := getRepos(accessToken, currentPageStr)

		for _, repo := range rr.Values {
			now := time.Now()
			slug := repo.Slug
			projectName := repo.Project.Name
			updatedAt := repo.UpdatedOn.Unix()
			updatedAtHuman := repo.UpdatedOn.Local()
			gap := now.Sub(repo.UpdatedOn)
			active := "Yes"

			if gap > sevenDays {
				active = "No"
			}

			log.Printf("Saving repo %v into database\n", slug)
			statement, _ := db.Prepare("INSERT INTO repos (project, name, updated_at, updated_at_human, active) VALUES (?, ?, ?, ?, ?)")
			statement.Exec(projectName, slug, updatedAt, updatedAtHuman, active)
		}

		if rr.Next == "" {
			break
		}
	}
}
