package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/AlexanderTurok/beat-store-backend/pkg/cache"
)

const (
	baseUrl      = "https://api.sendpulse.com"
	authEndpoint = "/oauth/access_token"

	grantType = "client_credentials"

	cacheTTL = 3600
)

type Client struct {
	config Config
	cache  cache.Cache
}

func NewClient(config Config, cache cache.Cache) *Client {
	return &Client{
		config: config,
		cache:  cache,
	}
}

type authRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) getToken() (string, error) {
	token, err := c.cache.Get("token")
	if err == nil {
		return token.(string), nil
	}

	token, err = c.authenticate()
	if err != nil {
		return "", err
	}

	if err := c.cache.Set("token", token, cacheTTL); err != nil {
		return "", err
	}

	return token.(string), err
}

func (c *Client) authenticate() (string, error) {
	reqData := authRequest{
		GrantType:    grantType,
		ClientId:     c.config.Id,
		ClientSecret: c.config.Secret,
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	fmt.Printf("req auth data - %s\n", reqBody)

	res, err := http.Post(baseUrl+authEndpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	fmt.Printf("auth status code - %s\n", res.Status)

	if res.StatusCode != 200 {
		return "", errors.New("sendpulse: auth: status code is not OK")
	}

	var resData authResponse

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(resBody, &resData); err != nil {
		return "", err
	}

	fmt.Printf("response auth data - %s\n", resData)

	return resData.AccessToken, nil
}
