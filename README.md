google-photo-uploader
===

## Usage
### 0. Configure environment variables
|name||
|---|---|
|GOOGLE_CLIENT_ID|required|
|GOOGLE_CLIENT_SECRET|required|
|OAUTH_ACCESS_TOKEN|required|
|OAUTH_REFRESH_TOKEN|required|

#### Get OAuth Client ID and Secret
See this article.

[Get started with REST  |  Google Photos APIs  |  Google Developers](https://developers.google.com/photos/library/guides/get-started)

#### Get OAuth access token and refresh token
Example code to get OAuth tokens.

```go
package main

import (
	"context"
	photos "github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func AuthorizationUrl() {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	oc := photos.NewOAuthConfig(photos.APIAppCredentials{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	})
	oc.RedirectURL = "http://localhost:8080/auth/google/callback"

	//Get authorization url
	url := oc.AuthCodeURL("", oauth2.AccessTypeOffline)
	logger.Printf("url=%+v", url)
}

func Token() {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	ctx := context.Background()
	oc := photos.NewOAuthConfig(photos.APIAppCredentials{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	})

	//Get access token and refresh token
	token, err := oc.Exchange(ctx, "", oauth2.AccessTypeOffline)
	if err != nil {
		logger.Fatalf("failed to exchange: %+v", err)
		return
	}
	logger.Printf("token: %+v", token)
}
```

### 1. Create album
```shell script
$ google-photo-uploader create-album  --name 'Album name'
2020/04/16 23:14:36 Album ID: AC...
```

### 2. Add photo to album
```shell script
$ google-photo-uploader upload --file path/to/photo.jpg  --album AC...
2020/04/16 23:18:53 https://photos.google.com/lr/album/AC.../photo/AC...
```

- You can upload photos to only albums created by the same client ID.
