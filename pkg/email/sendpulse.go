package sendpulse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/AlexanderTurok/beat-store-backend/pkg/cache"
	"github.com/sirupsen/logrus"
)

const (
	baseUrl                = "https://api.sendpulse.com"
	authEndpoint           = "/oauth/access_token"
	addEmailToListEndpoint = "/addressbooks/%s/emails"

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

type addToListRequest struct {
	Emails []emailInfo `json:"emails"`
}

type emailInfo struct {
	Email     string            `json:"email"`
	Variables map[string]string `json:"variables"`
}

func (c *Client) AddEmailToList(input AddEmailInput) error {
	token, err := c.getToken()
	if err != nil {
		return err
	}

	reqData := addToListRequest{
		Emails: []emailInfo{
			{
				Email:     input.Email,
				Variables: input.Variables,
			},
		},
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(addEmailToListEndpoint, input.ListID)

	req, err := http.NewRequest(http.MethodPost, baseUrl+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "aplication/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	logrus.Infof("SendPulse Respons: %s", string(respBody))

	if resp.StatusCode != 200 {
		return errors.New("status code is not OK")
	}

	return nil
}

type authRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
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
		GrantType:    grantType,
		ClientId:     c.id,
		ClientSecret: c.secret,
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(baseUrl+authEndpoint, "aplication/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	var respData authResponse

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(respBody, &respData); err != nil {
		return "", err
	}

	return respData.AccessToken, nil
}
