package models

import "mime/multipart"

type File struct {
	Path string
	File *multipart.FileHeader
}

var (
	FilePathSettings      = "settings"
	FilePathRestaurants   = "restaurants"
	FilePathPages         = "pages"
	FilePathNews          = "news"
	FilePathGalleries     = "galleries"
	FilePathSliders       = "sliders"
	FilePathNotifications = "notifications"
	FilePathGroups        = "groups"
	FilePathProducts      = "products"
)
