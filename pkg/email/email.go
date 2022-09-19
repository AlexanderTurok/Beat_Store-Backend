package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	addEmailToListEndpoint = "/addressbooks/%d/emails"
)

type AddEmailToList struct {
	ListId int         `json:"id"`
	Emails []EmailData `json:"emails"`
}

type EmailData struct {
	Email     string            `json:"email"`
	Variables map[string]string `json:"variables"`
}

type Result struct {
	Result bool `json:"result"`
}

func (c *Client) AddEmailToList(input AddEmailToList) (Result, error) {
	token, err := c.getToken()
	if err != nil {
		return Result{}, err
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return Result{}, err
	}

	path := fmt.Sprintf(baseUrl+addEmailToListEndpoint, input.ListId)
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(reqBody))
	if err != nil {
		return Result{}, err
	}

	req.Header.Set("Content-Type", "aplication/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Result{}, err
	}

	defer res.Body.Close()

	var resData Result

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return Result{}, err
	}

	if res.StatusCode != 200 {
		return Result{}, errors.New("sendpulse: status code is not OK")
	}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return Result{}, err
	}

	return resData, nil
}
