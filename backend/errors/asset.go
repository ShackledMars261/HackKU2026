package errors

import "net/http"

var ErrAssetNotFound = New("asset_not_found", "asset not found", http.StatusNotFound)
