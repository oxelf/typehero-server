package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)


const recaptchaURL = "https://www.google.com/recaptcha/api/siteverify"

// Struct for incoming JSON
type Submission struct {
	Username    string `json:"username"`
	Score       int    `json:"score"`
	CaptchaToken string `json:"captchaToken"`
}

// Struct for reCAPTCHA response
type RecaptchaResponse struct {
	Success bool    `json:"success"`
	Score   float64 `json:"score"`
}

func verifyCaptcha(token string) (float64, error) {
	secret := os.Getenv("RECAPTCHA_TOKEN")

	// Prepare the form data for the POST request
	form := url.Values{}
	form.Add("secret", secret)
	form.Add("response", token)

	// Send POST request to reCAPTCHA server
	resp, err := http.PostForm(recaptchaURL, form)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var recaptchaResponse RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&recaptchaResponse); err != nil {
		return 0, err
	}

	if !recaptchaResponse.Success {
		return 0, nil
	}

	return recaptchaResponse.Score, nil
}
