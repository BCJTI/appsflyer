package appsflyer

func (c *Client) GetMaps(endpoint string) ([]map[string]interface{}, error) {
	body, err := c.DispatchGetRequest(endpoint + AppsFlyerVersion)
	if err != nil {
		return nil, err
	}
	return Map(body)
}

func (c *Client) GetRawData(endpoint string) (entities []RawData, err error) {

	var body []byte

	// get response body
	if body, err = c.DispatchGetRequest(endpoint + AppsFlyerVersion); err != nil {
		return
	}

	// parse csv to []Report
	if err = Parse(body, RawData{}, func(v interface{}) {
		entities = append(entities, v.(RawData))
	}); err != nil {
		return
	}

	return

}
