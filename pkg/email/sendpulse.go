package sendpulse

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	id     string
	secret string
	cache  cache.Cache
}

func NewClient(id, secret string, cache cache.Cache) *Client {
	return &Client{
		id:     id,
		secret: secret,
		cache:  cache,
	}
}

type authRequest struct {
	grantType    string `json:"grant_type"`
	clientId     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
}

type authResponse struct {
	accessToken string `json:"access_token"`
	tokenType   string `json:"token_type"`
	expiresIn   string `json:"expires_in"`
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

	return token.(string), nil
}

func (c *Client) authenticate() (string, error) {
	reqData := authRequest{
		grantType:    grantType,
		clientId:     c.id,
		clientSecret: c.secret,
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(baseUrl+authEndpoint, "aplication/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", err
	}

	var respData authResponse

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(respBody, &respData); err != nil {
		return "", err
	}

	return respData.accessToken, nil
}
