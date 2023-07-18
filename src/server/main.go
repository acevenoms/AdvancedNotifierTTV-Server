package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/api/helix"
	"github.com/gorilla/mux"
)

var Model = make(map[string][]string)

var TwitchAPI = twitch.API(os.Getenv("TTV_CLIENT_ID"))

func main() {
	initTwitch()

	r := mux.NewRouter()
	registerChannelRoutes(r)
	registerDebugRoutes(r)

	http.ListenAndServe(":80", r)
}

func initTwitch() {
	token, err := getTwitchToken(TwitchAPI.ID, os.Getenv("TTV_CLIENT_SECRET"))
	if err != nil {
		fmt.Printf("Failed to connect to Twitch...%s\n", err)
	} else {
		TwitchAPI = TwitchAPI.NewBearer(token)
		testUserQuery := helix.UserOpts{Logins: []string{"vinesauce"}}
		users, err := TwitchAPI.Helix().GetUsers(testUserQuery)
		if err == nil {
			fmt.Printf("Successfully connected to twitch and got users %+v\n", users)
		} else {
			fmt.Printf("Failed to get users from Twitch...%s\n", err)
		}
	}
}

func getTwitchToken(client_id string, client_secret string) (string, error) {
	tokenEndpoint, _ := url.ParseRequestURI("https://id.twitch.tv/oauth2/token")

	tokenParams := url.Values{}
	tokenParams.Add("client_id", client_id)
	tokenParams.Add("client_secret", client_secret)
	tokenParams.Add("grant_type", "client_credentials")
	tokenEndpoint.RawQuery = tokenParams.Encode()

	body := bytes.NewReader([]byte(tokenEndpoint.RawQuery))
	res, err := http.Post(tokenEndpoint.String(), "application/x-www-form-urlencoded", body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var tokenResponse map[string]string
	rawResponse, _ := io.ReadAll(res.Body)
	json.Unmarshal(rawResponse, &tokenResponse)
	return tokenResponse["access_token"], nil
}

func registerChannelRoutes(r *mux.Router) {
	r.HandleFunc("/channel/{channel_login}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channel_login := vars["channel_login"]

		result := channelAdd(channel_login)
		if result {
			fmt.Fprintln(w, "Added!")
		} else {
			fmt.Fprintln(w, "Created!")
		}
	})
}

func channelAdd(channel_login string) bool {
	users, ok := Model[channel_login]
	if ok {
		users = append(users, "NewUser")
		return true
	} else {
		users = append(users, "FirstUser")
		return false
	}
}

func registerDebugRoutes(r *mux.Router) {
	r.HandleFunc("/debug/dumpModel", func(w http.ResponseWriter, r *http.Request) {
		debugDumpModel(w)
	})
}

func debugDumpModel(w io.Writer) {
	jsonStr, err := json.Marshal(Model)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	} else {
		fmt.Fprintln(w, string(jsonStr))
	}
}
