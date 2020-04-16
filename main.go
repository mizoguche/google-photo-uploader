package main

import (
	"context"
	photos "github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
	"log"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "upload",
				Action: upload,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "file path to upload",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "album",
						Usage:    "album ID to add(it's required to be created by the same client id)",
						Required: true,
					},
				},
			},
			{
				Name:   "create-album",
				Action: createAlbum,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "album name to create",
						Required: true,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createClient(ctx context.Context) (*photos.Client, error) {
	oc := photos.NewOAuthConfig(photos.APIAppCredentials{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	})

	httpc := oc.Client(ctx, &oauth2.Token{
		AccessToken:  os.Getenv("OAUTH_ACCESS_TOKEN"),
		RefreshToken: os.Getenv("OAUTH_REFRESH_TOKEN"),
		Expiry:       time.Now().Add(-1000),
	})

	return photos.NewClient(httpc)
}

func upload(c *cli.Context) error {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		return err
	}

	filename := c.String("file")
	albumID :=  c.String("album")
	item, err := client.AddMediaItem(ctx, filename, albumID)
	if err != nil {
		return err
	}

	logger.Printf("%+v", item.ProductUrl)
	return nil
}

func createAlbum(c *cli.Context) error {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		return err
	}

	name := c.String("name")
	item, err := client.GetOrCreateAlbumByName(name)
	if err != nil {
		return err
	}

	logger.Printf("Album ID: %+v", item.Id)
	return nil
}
