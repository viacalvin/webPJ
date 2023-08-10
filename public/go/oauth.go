package handlers

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "net/http"
    "fmt"
    "io/ioutil"
    "context"
    "log"
    "encoding/base64"
    "crypto/rand"
    "os"
    "time"
)

// 구글 2단계 인증 로직
var googleOauthConfig = &oauth2.Config{
    RedirectURL:  "http://localhost:8000/auth/google/callback",
    ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
    ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
    oauthState := generateStateOauthCookie(w)
    u := googleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
    oauthState, _ := r.Cookie("oauthstate")

    if r.FormValue("state") != oauthState.Value {
        log.Println("invalid oauth google state")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
	// resv

    data, err := getUserDataFromGoogle(r.FormValue("code"))
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
    var expiration = time.Now().Add(365 * 24 * time.Hour)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)
    cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
    http.SetCookie(w, &cookie)

    return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {

    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        return nil, fmt.Errorf("교환 실패: %s", err.Error())
    }
    response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
    if err != nil {
        return nil, fmt.Errorf("유저정보 인증 실패: %s", err.Error())
    }
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("응답 오류: %s", err.Error())
    }
    return contents, nil
}
