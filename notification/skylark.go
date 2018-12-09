package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	apiConfig      = viper.GetStringMap("skylark_api")
	apiUrl         = apiConfig["push_url"].(string)
	clientId       = apiConfig["client_id"].(string)
	slpNamespaceId = apiConfig["namespace_id"].(int)
	clientSecret   = apiConfig["client_secret"].(string)
	templateId     = apiConfig["template_id"].(string)
)

type NotifyBody struct {
	Openids        []string               `json:"openids"`
	TemplateEntity map[string]interface{} `json:"template_entity"`
}

func NotifyBySkylark(openids []string, options map[string]string) bool {
	notifyBody := NotifyBody{Openids: openids, TemplateEntity: buildTemplateEntity(options)}

	notifyBodyJson, _ := json.Marshal(notifyBody)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(notifyBodyJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", buildAuthorizationHeader())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to post skylark api: %v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func buildAuthorizationHeader() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"namespace_id": slpNamespaceId,
		"exp":          time.Now().Unix() + 60*5,
	})

	tokenString, err := token.SignedString([]byte(clientSecret))

	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%s:%s", clientId, tokenString)
}

func buildTemplateEntity(options map[string]string) (result map[string]interface{}) {
	templateEntityTpl := `
		{
			"template_id": "%s",
			"url": "",
			"data": {
				"first": {
					"value": "%s",
					"color": "#173177"
				},
				"keyword1": {
					"value": "%s",
					"color": "#173177"
				},
				"keyword2": {
					"value": "%s",
					"color": "#173177"
				},
				"remark": {
					"value": "%s",
					"color": "#173177"
				}
			}
		}

	`

	templateEntityStr := fmt.Sprintf(templateEntityTpl, templateId, options["first"], options["keyword1"], options["keyword2"], options["remark"])
	err := json.Unmarshal([]byte(templateEntityStr), &result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}
