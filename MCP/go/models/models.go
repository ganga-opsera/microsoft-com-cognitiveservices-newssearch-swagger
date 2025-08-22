package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// VideoObject represents the VideoObject schema from the OpenAPI specification
type VideoObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Video VideoObject `json:"video,omitempty"` // Defines a video object that is relevant to the query.
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g the source URL for the image).
	Height int `json:"height,omitempty"` // The height of the source media object, in pixels.
	Width int `json:"width,omitempty"` // The width of the source media object, in pixels.
}

// News represents the News schema from the OpenAPI specification
type News struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
}

// ImageObject represents the ImageObject schema from the OpenAPI specification
type ImageObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Video VideoObject `json:"video,omitempty"` // Defines a video object that is relevant to the query.
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g the source URL for the image).
	Height int `json:"height,omitempty"` // The height of the source media object, in pixels.
	Width int `json:"width,omitempty"` // The width of the source media object, in pixels.
}

// NewsTopic represents the NewsTopic schema from the OpenAPI specification
type NewsTopic struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// MediaObject represents the MediaObject schema from the OpenAPI specification
type MediaObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Video VideoObject `json:"video,omitempty"` // Defines a video object that is relevant to the query.
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Displaytext string `json:"displayText,omitempty"` // The display version of the query term. This version of the query term may contain special characters that highlight the search term found in the query string. The string contains the highlighting characters only if the query enabled hit highlighting
	Searchlink string `json:"searchLink,omitempty"` // The URL that you use to get the results of the related search. Before using the URL, you must append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header. Use this URL if you're displaying the results in your own user interface. Otherwise, use the webSearchUrl URL.
	Text string `json:"text"` // The query string. Use this string as the query term in a new search request.
	Thumbnail ImageObject `json:"thumbnail,omitempty"` // Defines an image
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL that takes the user to the Bing search results page for the query.Only related search results include this field.
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// NewsArticle represents the NewsArticle schema from the OpenAPI specification
type NewsArticle struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Video VideoObject `json:"video,omitempty"` // Defines a video object that is relevant to the query.
	Wordcount int `json:"wordCount,omitempty"` // The number of words in the text of the Article.
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// TrendingTopics represents the TrendingTopics schema from the OpenAPI specification
type TrendingTopics struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// Article represents the Article schema from the OpenAPI specification
type Article struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Video VideoObject `json:"video,omitempty"` // Defines a video object that is relevant to the query.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
}
