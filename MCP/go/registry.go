package main

import (
	"github.com/news-search-client/mcp-server/config"
	"github.com/news-search-client/mcp-server/models"
	tools_newscategory "github.com/news-search-client/mcp-server/tools/newscategory"
	tools_newssearch "github.com/news-search-client/mcp-server/tools/newssearch"
	tools_newstrendingtopics "github.com/news-search-client/mcp-server/tools/newstrendingtopics"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_newscategory.CreateNews_categoryTool(cfg),
		tools_newssearch.CreateNews_searchTool(cfg),
		tools_newstrendingtopics.CreateNews_trendingTool(cfg),
	}
}
