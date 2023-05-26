package wework

type TemplateCardCommon struct {
	CardType string `json:"card_type"`
	Source   struct {
		IconUrl   string `json:"icon_url"`
		Desc      string `json:"desc"`
		DescColor int    `json:"desc_color"`
	} `json:"source"`
	MainTitle struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"main_title"`
	QuoteArea struct {
		Type      int    `json:"type"`
		Url       string `json:"url"`
		Appid     string `json:"appid"`
		Pagepath  string `json:"pagepath"`
		Title     string `json:"title"`
		QuoteText string `json:"quote_text"`
	} `json:"quote_area"`
	HorizontalContentList []struct {
		Keyname string `json:"keyname"`
		Value   string `json:"value"`
		Type    int    `json:"type,omitempty"`
		Url     string `json:"url,omitempty"`
		MediaId string `json:"media_id,omitempty"`
	} `json:"horizontal_content_list"`
	JumpList []struct {
		Type     int    `json:"type"`
		Url      string `json:"url,omitempty"`
		Title    string `json:"title"`
		Appid    string `json:"appid,omitempty"`
		Pagepath string `json:"pagepath,omitempty"`
	} `json:"jump_list"`
	CardAction struct {
		Type     int    `json:"type"`
		Url      string `json:"url"`
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
	} `json:"card_action"`
}
