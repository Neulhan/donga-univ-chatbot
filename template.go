package main

type j map[string]interface{}

func textTemplate(input []string) interface{} {
	textSlice := []j{}

	for _, text := range input {
		cardObj := j{
			"simpleText": j{
				"text": text,
			},
		}

		textSlice = append(textSlice, cardObj)

	}

	return j{
		"version": "2.0",
		"template": j{
			"outputs": textSlice,
		},
	}
}

func cardTemplate(title string, desc string, imgURL string, buttons []j) interface{} {
	cardList := []j{}

	cardObj := j{
		"basicCard": j{
			"title":       title,
			"description": desc,
			"thumbnail": j{
				"imageUrl": imgURL,
			},
			"buttons": buttons,
		},
	}

	cardList = append(cardList, cardObj)

	return j{
		"version": "2.0",
		"template": j{
			"outputs": cardList,
		},
	}
}

func buttonTemplate(action string, label string, webLinkURL string) (re j) {
	re = j{
		"action":     action,
		"label":      label,
		"webLinkUrl": webLinkURL,
	}
	return
}
