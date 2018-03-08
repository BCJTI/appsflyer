package appsflyer

func (c *Client) GetMaps(endpoint string) ([]map[string]interface{}, error) {
	body, err := c.DispatchGetRequest(endpoint + AppsFlyerVersion)
	if err != nil {
		return nil, err
	}
	return Map(body)
}

func (c *Client) GetRawData(endpoint string) (entities []Report, err error) {

	var body []byte

	// get response body
	if body, err = c.DispatchGetRequest(endpoint + AppsFlyerVersion); err != nil {
		return
	}

	// parse csv to []Report
	if err = Parse(body, Report{}, func(v interface{}) {
		entities = append(entities, v.(Report))
	}); err != nil {
		return
	}

	return

}
