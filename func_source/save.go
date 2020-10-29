package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
	"github.com/nfnt/resize"
	"golang.org/x/oauth2"
)

// UserPayload is the shape of the inbound JSON in the request body
type UserPayload struct {
	Event      string `json:"event"`
	InstanceID string `json:"instance_id"`
	User       struct {
		ID          string `json:"id"`
		Aud         string `json:"aud"`
		Role        string `json:"role"`
		Email       string `json:"email"`
		AppMetadata struct {
			Provider string `json:"provider"`
		} `json:"app_metadata"`
		UserMetadata struct {
			FullName string `json:"full_name"`
		} `json:"user_metadata"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"user"`
}

func main() {
	lambda.Start(handler)
}

func downloadDecode(filepath string, url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return jpeg.Decode(resp.Body)

}

func createResizedFiles(img image.Image, size int) (error, string) {
	m := resize.Resize(uint(size), 0, img, resize.Lanczos3)
	stringSize := strconv.Itoa(size)
	fileName := "profilePicture_" + stringSize + ".jpeg"
	filePath := "/tmp/" + fileName
	out, err := os.Create(filePath)
	if err != nil {
		return err, ""
	}
	defer out.Close()

	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return err, ""
	}
	return nil, fileName
}

func createGHClient() (*github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, nil
}
func uploadFiles(client *github.Client, files []string, id string) error {

	for _, file := range files {
		path := "profile_pictures/" + id + "/" + file

		fileToUpload, err := ioutil.ReadFile("/tmp/" + file)
		if err != nil {
			return err
		}

		opts := &github.RepositoryContentFileOptions{
			Message:   github.String("Uploading profile pictures"),
			Content:   fileToUpload,
			Branch:    github.String("master"),
			Committer: &github.CommitAuthor{Name: github.String(os.Getenv("GITHUB_COMMITTER_NAME")), Email: github.String(os.Getenv("GITHUB_COMMITTER_EMAIL"))},
		}
		if _, _, err := client.Repositories.CreateFile(context.Background(), os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO_NAME"), path, opts); err != nil {
			return err
		}
	}
	return nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var eventBody UserPayload
	var err error
	if err = json.Unmarshal([]byte(request.Body), &eventBody); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode:      500,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            "Unable to unmarshal inbound payload",
			IsBase64Encoded: false,
		}, nil
	}
	log.Print(eventBody)

	email := eventBody.User.Email
	hashedEmail := md5.Sum([]byte(email))
	stringHashedEmail := hex.EncodeToString(hashedEmail[:])
	var files []string

	gravatarURL := "https://www.gravatar.com/avatar/" + stringHashedEmail + "?s=400"
	log.Print(gravatarURL)
	img, err := downloadDecode("/tmp/profilePicture.jpeg", gravatarURL)
	if err != nil {
		log.Print(err, "unable to download file")
	}
	if err, file := createResizedFiles(img, 80); err != nil {
		log.Print(err)
	} else {
		files = append(files, file)
	}
	if err, file := createResizedFiles(img, 200); err != nil {
		log.Print(err)
	} else {
		files = append(files, file)
	}
	if err, file := createResizedFiles(img, 400); err != nil {
		log.Print(err)
	} else {
		files = append(files, file)
	}
	var client *github.Client
	if client, err = createGHClient(); err != nil {
		log.Print(err)
	}

	if err := uploadFiles(client, files, eventBody.User.ID); err != nil {
		log.Print(err)
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            "Files uploaded!",
		IsBase64Encoded: false,
	}, nil
}
