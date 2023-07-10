package unsplash

type UnsplashService interface {
	GetPhotos(page int, perPage int) ([]Photo, error)
	SearchPhotos(query string, page int, perPage int) (SearchResponse, error)
}