package server

import (
  "github.com/BkSearch/Crawl-Server/common"
  "github.com/BkSearch/Crawl-Server/db"
  "github.com/gocolly/colly/v2"
  _ "github.com/lib/pq"
  "github.com/joho/godotenv"
)

type Server struct {
  Database *db.TopicDB

}
