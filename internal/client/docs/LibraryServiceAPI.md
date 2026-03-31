# \LibraryServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItems**](LibraryServiceAPI.md#DeleteItems) | **Delete** /Items | Deletes an item from the library and file system
[**DeleteItemsById**](LibraryServiceAPI.md#DeleteItemsById) | **Delete** /Items/{Id} | Deletes an item from the library and file system
[**GetAlbumsByIdSimilar**](LibraryServiceAPI.md#GetAlbumsByIdSimilar) | **Get** /Albums/{Id}/Similar | Finds albums similar to a given album.
[**GetArtistsByIdSimilar**](LibraryServiceAPI.md#GetArtistsByIdSimilar) | **Get** /Artists/{Id}/Similar | Finds albums similar to a given album.
[**GetGamesByIdSimilar**](LibraryServiceAPI.md#GetGamesByIdSimilar) | **Get** /Games/{Id}/Similar | Finds games similar to a given game.
[**GetItemsByIdAncestors**](LibraryServiceAPI.md#GetItemsByIdAncestors) | **Get** /Items/{Id}/Ancestors | Gets all parents of an item
[**GetItemsByIdCriticreviews**](LibraryServiceAPI.md#GetItemsByIdCriticreviews) | **Get** /Items/{Id}/CriticReviews | Gets critic reviews for an item
[**GetItemsByIdDeleteinfo**](LibraryServiceAPI.md#GetItemsByIdDeleteinfo) | **Get** /Items/{Id}/DeleteInfo | Gets delete info for an item
[**GetItemsByIdDownload**](LibraryServiceAPI.md#GetItemsByIdDownload) | **Get** /Items/{Id}/Download | Downloads item media
[**GetItemsByIdFile**](LibraryServiceAPI.md#GetItemsByIdFile) | **Get** /Items/{Id}/File | Gets the original file of an item
[**GetItemsByIdSimilar**](LibraryServiceAPI.md#GetItemsByIdSimilar) | **Get** /Items/{Id}/Similar | Gets similar items
[**GetItemsByIdThememedia**](LibraryServiceAPI.md#GetItemsByIdThememedia) | **Get** /Items/{Id}/ThemeMedia | Gets theme videos and songs for an item
[**GetItemsByIdThemesongs**](LibraryServiceAPI.md#GetItemsByIdThemesongs) | **Get** /Items/{Id}/ThemeSongs | Gets theme songs for an item
[**GetItemsByIdThemevideos**](LibraryServiceAPI.md#GetItemsByIdThemevideos) | **Get** /Items/{Id}/ThemeVideos | Gets theme videos for an item
[**GetItemsCounts**](LibraryServiceAPI.md#GetItemsCounts) | **Get** /Items/Counts | 
[**GetItemsIntros**](LibraryServiceAPI.md#GetItemsIntros) | **Get** /Items/Intros | Gets info to debug intros
[**GetLibrariesAvailableoptions**](LibraryServiceAPI.md#GetLibrariesAvailableoptions) | **Get** /Libraries/AvailableOptions | 
[**GetLibraryMediafolders**](LibraryServiceAPI.md#GetLibraryMediafolders) | **Get** /Library/MediaFolders | Gets all user media folders.
[**GetLibraryPhysicalpaths**](LibraryServiceAPI.md#GetLibraryPhysicalpaths) | **Get** /Library/PhysicalPaths | Gets a list of physical paths from virtual folders
[**GetLibrarySelectablemediafolders**](LibraryServiceAPI.md#GetLibrarySelectablemediafolders) | **Get** /Library/SelectableMediaFolders | Gets all user media folders.
[**GetMoviesByIdSimilar**](LibraryServiceAPI.md#GetMoviesByIdSimilar) | **Get** /Movies/{Id}/Similar | Finds movies and trailers similar to a given movie.
[**GetShowsByIdSimilar**](LibraryServiceAPI.md#GetShowsByIdSimilar) | **Get** /Shows/{Id}/Similar | Finds tv shows similar to a given one.
[**GetTrailersByIdSimilar**](LibraryServiceAPI.md#GetTrailersByIdSimilar) | **Get** /Trailers/{Id}/Similar | Finds movies and trailers similar to a given trailer.
[**PostItemsByIdDelete**](LibraryServiceAPI.md#PostItemsByIdDelete) | **Post** /Items/{Id}/Delete | Deletes an item from the library and file system
[**PostItemsDelete**](LibraryServiceAPI.md#PostItemsDelete) | **Post** /Items/Delete | Deletes an item from the library and file system
[**PostLibraryMediaUpdated**](LibraryServiceAPI.md#PostLibraryMediaUpdated) | **Post** /Library/Media/Updated | Reports that new movies have been added by an external source
[**PostLibraryMoviesAdded**](LibraryServiceAPI.md#PostLibraryMoviesAdded) | **Post** /Library/Movies/Added | Deprecated. Use /Library/Media/Updated
[**PostLibraryMoviesUpdated**](LibraryServiceAPI.md#PostLibraryMoviesUpdated) | **Post** /Library/Movies/Updated | Deprecated. Use /Library/Media/Updated
[**PostLibraryRefresh**](LibraryServiceAPI.md#PostLibraryRefresh) | **Post** /Library/Refresh | Starts a library scan
[**PostLibrarySeriesAdded**](LibraryServiceAPI.md#PostLibrarySeriesAdded) | **Post** /Library/Series/Added | Deprecated. Use /Library/Media/Updated
[**PostLibrarySeriesUpdated**](LibraryServiceAPI.md#PostLibrarySeriesUpdated) | **Post** /Library/Series/Updated | Deprecated. Use /Library/Media/Updated



## DeleteItems

> DeleteItems(ctx).Ids(ids).Execute()

Deletes an item from the library and file system



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ids := "ids_example" // string | Ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.DeleteItems(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.DeleteItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Ids | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemsById

> DeleteItemsById(ctx, id).Execute()

Deletes an item from the library and file system



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.DeleteItemsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.DeleteItemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlbumsByIdSimilar

> QueryResultBaseItemDto GetAlbumsByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Finds albums similar to a given album.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetAlbumsByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetAlbumsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlbumsByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetAlbumsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlbumsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtistsByIdSimilar

> QueryResultBaseItemDto GetArtistsByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Finds albums similar to a given album.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetArtistsByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetArtistsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtistsByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetArtistsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGamesByIdSimilar

> QueryResultBaseItemDto GetGamesByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Finds games similar to a given game.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetGamesByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetGamesByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGamesByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetGamesByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGamesByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdAncestors

> []BaseItemDto GetItemsByIdAncestors(ctx, id).UserId(userId).Execute()

Gets all parents of an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdAncestors(context.Background(), id).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdAncestors`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdCriticreviews

> QueryResultBaseItemDto GetItemsByIdCriticreviews(ctx, id).StartIndex(startIndex).Limit(limit).Execute()

Gets critic reviews for an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdCriticreviews(context.Background(), id).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdCriticreviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdCriticreviews`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdCriticreviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdCriticreviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdDeleteinfo

> LibraryDeleteInfo GetItemsByIdDeleteinfo(ctx, id).Execute()

Gets delete info for an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdDeleteinfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdDeleteinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdDeleteinfo`: LibraryDeleteInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdDeleteinfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdDeleteinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryDeleteInfo**](LibraryDeleteInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdDownload

> GetItemsByIdDownload(ctx, id).Execute()

Downloads item media



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.GetItemsByIdDownload(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdFile

> GetItemsByIdFile(ctx, id).Execute()

Gets the original file of an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.GetItemsByIdFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdSimilar

> QueryResultBaseItemDto GetItemsByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets similar items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdThememedia

> AllThemeMediaResult GetItemsByIdThememedia(ctx, id).InheritFromParent(inheritFromParent).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets theme videos and songs for an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	inheritFromParent := true // bool | Determines whether or not parent items should be searched for theme media. (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdThememedia(context.Background(), id).InheritFromParent(inheritFromParent).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdThememedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThememedia`: AllThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdThememedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThememediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inheritFromParent** | **bool** | Determines whether or not parent items should be searched for theme media. | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**AllThemeMediaResult**](AllThemeMediaResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdThemesongs

> ThemeMediaResult GetItemsByIdThemesongs(ctx, id).InheritFromParent(inheritFromParent).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets theme songs for an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	inheritFromParent := true // bool | Determines whether or not parent items should be searched for theme media. (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdThemesongs(context.Background(), id).InheritFromParent(inheritFromParent).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdThemesongs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThemesongs`: ThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdThemesongs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThemesongsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inheritFromParent** | **bool** | Determines whether or not parent items should be searched for theme media. | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**ThemeMediaResult**](ThemeMediaResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdThemevideos

> ThemeMediaResult GetItemsByIdThemevideos(ctx, id).InheritFromParent(inheritFromParent).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets theme videos for an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	inheritFromParent := true // bool | Determines whether or not parent items should be searched for theme media. (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdThemevideos(context.Background(), id).InheritFromParent(inheritFromParent).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdThemevideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThemevideos`: ThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdThemevideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThemevideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inheritFromParent** | **bool** | Determines whether or not parent items should be searched for theme media. | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**ThemeMediaResult**](ThemeMediaResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsCounts

> ItemCounts GetItemsCounts(ctx).UserId(userId).IsFavorite(isFavorite).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "userId_example" // string | Optional. Get counts from a specific user's library. (optional)
	isFavorite := true // bool | Optional. Get counts of favorite items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsCounts(context.Background()).UserId(userId).IsFavorite(isFavorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsCounts`: ItemCounts
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Get counts from a specific user&#39;s library. | 
 **isFavorite** | **bool** | Optional. Get counts of favorite items | 

### Return type

[**ItemCounts**](ItemCounts.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsIntros

> []PersistenceIntroDebugInfo GetItemsIntros(ctx).Execute()

Gets info to debug intros



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsIntros(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsIntros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsIntros`: []PersistenceIntroDebugInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsIntros`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsIntrosRequest struct via the builder pattern


### Return type

[**[]PersistenceIntroDebugInfo**](PersistenceIntroDebugInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibrariesAvailableoptions

> LibraryOptionsResult GetLibrariesAvailableoptions(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetLibrariesAvailableoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibrariesAvailableoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibrariesAvailableoptions`: LibraryOptionsResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibrariesAvailableoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibrariesAvailableoptionsRequest struct via the builder pattern


### Return type

[**LibraryOptionsResult**](LibraryOptionsResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryMediafolders

> QueryResultBaseItemDto GetLibraryMediafolders(ctx).IsHidden(isHidden).Execute()

Gets all user media folders.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	isHidden := true // bool | Optional. Filter by folders that are marked hidden, or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetLibraryMediafolders(context.Background()).IsHidden(isHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibraryMediafolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryMediafolders`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibraryMediafolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryMediafoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional. Filter by folders that are marked hidden, or not. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryPhysicalpaths

> []string GetLibraryPhysicalpaths(ctx).Execute()

Gets a list of physical paths from virtual folders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetLibraryPhysicalpaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibraryPhysicalpaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryPhysicalpaths`: []string
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibraryPhysicalpaths`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryPhysicalpathsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibrarySelectablemediafolders

> []LibraryMediaFolder GetLibrarySelectablemediafolders(ctx).Execute()

Gets all user media folders.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetLibrarySelectablemediafolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibrarySelectablemediafolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibrarySelectablemediafolders`: []LibraryMediaFolder
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibrarySelectablemediafolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibrarySelectablemediafoldersRequest struct via the builder pattern


### Return type

[**[]LibraryMediaFolder**](LibraryMediaFolder.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoviesByIdSimilar

> QueryResultBaseItemDto GetMoviesByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Finds movies and trailers similar to a given movie.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetMoviesByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetMoviesByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoviesByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetMoviesByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoviesByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShowsByIdSimilar

> QueryResultBaseItemDto GetShowsByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Finds tv shows similar to a given one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetShowsByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetShowsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShowsByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetShowsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShowsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrailersByIdSimilar

> QueryResultBaseItemDto GetTrailersByIdSimilar(ctx, id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Finds movies and trailers similar to a given trailer.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "startItemId_example" // string | Optional. Skip through the list until a given item is found. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetTrailersByIdSimilar(context.Background(), id).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).IsSpecialSeason(isSpecialSeason).AdjacentTo(adjacentTo).StartItemId(startItemId).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetTrailersByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrailersByIdSimilar`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetTrailersByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrailersByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsByIdDelete

> PostItemsByIdDelete(ctx, id).Execute()

Deletes an item from the library and file system



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostItemsByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostItemsByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsDelete

> PostItemsDelete(ctx).Ids(ids).Execute()

Deletes an item from the library and file system



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ids := "ids_example" // string | Ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostItemsDelete(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostItemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Ids | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibraryMediaUpdated

> PostLibraryMediaUpdated(ctx).LibraryPostUpdatedMedia(libraryPostUpdatedMedia).Execute()

Reports that new movies have been added by an external source



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	libraryPostUpdatedMedia := *openapiclient.NewLibraryPostUpdatedMedia() // LibraryPostUpdatedMedia | PostUpdatedMedia

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibraryMediaUpdated(context.Background()).LibraryPostUpdatedMedia(libraryPostUpdatedMedia).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryMediaUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryMediaUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryPostUpdatedMedia** | [**LibraryPostUpdatedMedia**](LibraryPostUpdatedMedia.md) | PostUpdatedMedia | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibraryMoviesAdded

> PostLibraryMoviesAdded(ctx).Execute()

Deprecated. Use /Library/Media/Updated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibraryMoviesAdded(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryMoviesAdded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryMoviesAddedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibraryMoviesUpdated

> PostLibraryMoviesUpdated(ctx).Execute()

Deprecated. Use /Library/Media/Updated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibraryMoviesUpdated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryMoviesUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryMoviesUpdatedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibraryRefresh

> PostLibraryRefresh(ctx).Execute()

Starts a library scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibraryRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryRefreshRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibrarySeriesAdded

> PostLibrarySeriesAdded(ctx).Execute()

Deprecated. Use /Library/Media/Updated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibrarySeriesAdded(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibrarySeriesAdded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibrarySeriesAddedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibrarySeriesUpdated

> PostLibrarySeriesUpdated(ctx).Execute()

Deprecated. Use /Library/Media/Updated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibrarySeriesUpdated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibrarySeriesUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibrarySeriesUpdatedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

