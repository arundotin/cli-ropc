package adfs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/term"
	"io/ioutil"
	"marunk20/cli-ropc/utils"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"syscall"
)

type AdfsResponseType struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int64  `json:"expires_in"`
	Resource              string `json:"resource"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	IdToken               string `json:"id_token"`
}

func LoginAndGetUserFullName() string {
	cdsId, cdsPassword := getCredentailsFromUser()
	adfsResponse := getTokenFromADFS(cdsId, cdsPassword)
	accessToken := parseAdfsResponseAndGetAccessToken(adfsResponse)
	return utils.ExtractFullNameFromUserToken(accessToken)
}


func parseAdfsResponseAndGetAccessToken(adfsResponse []byte) string {
	var adfsResponseJson AdfsResponseType

	error := json.Unmarshal(adfsResponse, &adfsResponseJson)

	if error != nil {
		fmt.Println("ADFS Authentication failed. Please try again")
		os.Exit(1)
	}



	if adfsResponseJson.AccessToken == "" {
		fmt.Println("Unable to login .. Something seems to be wrong with your credentials.")
		os.Exit(1)
	}

	return adfsResponseJson.AccessToken
}

func getTokenFromADFS(cdsId, cdsPassword string) []byte {
	adfsEndpoint := "https://maskeddomain.com/adfs/oauth2/token"
	data := url.Values{}

	data.Set("grant_type", "password")
	data.Set("response_type", "token")
	data.Set("client_id", "masked-client-id")
	data.Set("resource", "masked-password")
	data.Set("username", cdsId)
	data.Set("password", cdsPassword)

	httpClient := &http.Client{}
	request, error := http.NewRequest(
		"POST",
		adfsEndpoint,
		strings.NewReader(data.Encode()))

	if error != nil {
		return nil
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, error := httpClient.Do(request)

	if error != nil {
		return nil
	}

	defer response.Body.Close()
	responseBody, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return nil
	}

	return responseBody
}

func getCredentailsFromUser() (cdsId, cdsPassword string) {
	fmt.Println("Enter your CDS-Id : ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		cdsId = scanner.Text()
	}

	fmt.Println("Enter your CDS-Password : ")

	password, _ := term.ReadPassword(syscall.Stdin)

	return cdsId, string(password)

}
