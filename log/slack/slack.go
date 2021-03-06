package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ResultadosDigitais/x9/config"
	"github.com/ResultadosDigitais/x9/log"
)

func Send(repository, vulnerability, filename, values, id, issueURL string) {
	slackWebHook := config.Opts.SlackWebhook
	interactiveActions := getInteractiveActionsFields(id, issueURL)
	if slackWebHook != "" {
		values := map[string]interface{}{
			"text": ":warning: *Vulnerability found*",
			"attachments": []map[string]interface{}{
				map[string]interface{}{
					"color": "#AB3117",
					"fields": []map[string]interface{}{
						map[string]interface{}{
							"title": "Repository",
							"value": repository,
							"short": false,
						},
						map[string]interface{}{
							"title": "Vulnerability",
							"value": vulnerability,
							"short": false,
						},
						map[string]interface{}{
							"title": "File",
							"value": filename,
							"short": false,
						},
						map[string]interface{}{
							"title": "Values",
							"value": values,
							"short": false,
						},
					},
				},
				interactiveActions,
			},
		}
		jsonValue, _ := json.Marshal(values)
		resp, err := http.Post(slackWebHook, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			log.Error(err.Error(), nil)
		} else if resp.StatusCode != http.StatusOK {
			log.Error(fmt.Sprintf("cannot send message to slack [status %d]", resp.StatusCode), nil)
		}
	}
}

func getInteractiveActionsFields(id, issueURL string) map[string]interface{} {
	falsePositiveButton := map[string]string{
		"name":  "False Positive",
		"text":  "False Positive",
		"type":  "button",
		"value": "false_positive",
	}

	if issueURL != "" {
		return map[string]interface{}{
			"color": "#3AA3E3",
			"fields": []map[string]interface{}{
				map[string]interface{}{
					"title": "Issue",
					"value": issueURL,
					"short": false,
				},
			},
		}

	}
	return map[string]interface{}{
		"callback_id":     id,
		"title":           "Actions",
		"color":           "#3AA3E3",
		"attachment_type": "default",
		"actions": []map[string]string{
			map[string]string{
				"name":  "Open Issue",
				"text":  "Open Issue",
				"type":  "button",
				"value": "open_issue",
			},
			falsePositiveButton,
		},
	}
}
