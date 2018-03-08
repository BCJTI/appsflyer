package appsflyer

import (
	"time"
)

// https://support.appsflyer.com/hc/en-us/articles/208387843-Raw-Data-Reports-V5-
type Report struct {
	AttributedTouchType       string `json:"attributed_touch_type"       csv:"Attributed Touch Type"`
	AttributedTouchTime       string `json:"attributed_touch_time"       csv:"Attributed Touch Time"`
	InstallTime               string `json:"install_time"                csv:"Install Time"`
	EventTime                 string `json:"event_time"                  csv:"Event Time"`
	EventName                 string `json:"event_name"                  csv:"Event Name"`
	EventValue                string `json:"event_value"                 csv:"Event Value"`
	EventRevenue              string `json:"event_revenue"               csv:"Event Revenue"`
	EventRevenueCurrency      string `json:"event_revenue_currency"      csv:"Event Revenue Currency"`
	EventRevenueUSD           string `json:"event_revenue_usd"           csv:"Event Revenue USD"`
	EventSource               string `json:"event_source"                csv:"Event Source"`
	IsReceiptValidated        string `json:"is_receipt_validated"        csv:"Is Receipt Validated"`
	Partner                   string `json:"af_prt"                      csv:"Partner"`
	MediaSource               string `json:"media_source"                csv:"Media Source"`
	Channel                   string `json:"af_channel"                  csv:"Channel"`
	Keywords                  string `json:"af_keywords"                 csv:"Keywords"`
	Campaign                  string `json:"campaign"                    csv:"Campaign"`
	CampaignID                string `json:"af_c_id"                     csv:"Campaign ID"`
	Adset                     string `json:"af_adset"                    csv:"Adset"`
	AdsetID                   string `json:"af_adset_id"                 csv:"Adset ID"`
	Ad                        string `json:"af_ad"                       csv:"Ad"`
	AdID                      string `json:"af_ad_id"                    csv:"Ad ID"`
	AdType                    string `json:"af_ad_type"                  csv:"Ad Type"`
	SiteID                    string `json:"af_site_id"                  csv:"Site ID"`
	SubSiteID                 string `json:"af_sub_site_id"              csv:"Sub Site ID"`
	SubParam1                 string `json:"af_sub1"                     csv:"Sub Param 1"`
	SubParam2                 string `json:"af_sub2"                     csv:"Sub Param 2"`
	SubParam3                 string `json:"af_sub3"                     csv:"Sub Param 3"`
	SubParam4                 string `json:"af_sub4"                     csv:"Sub Param 4"`
	SubParam5                 string `json:"af_sub5"                     csv:"Sub Param 5"`
	CostModel                 string `json:"af_cost_model"               csv:"Cost Model"`
	CostValue                 string `json:"af_cost_value"               csv:"Cost Value"`
	CostCurrency              string `json:"af_cost currency"            csv:"Cost Currency"`
	Contributor1Partner       string `json:"contributor1_af_prt"         csv:"Contributor 1 Partner"`
	Contributor1MediaSource   string `json:"contributor1_media_source"   csv:"Contributor 1 Media Source"`
	Contributor1Campaign      string `json:"contributor1_campaign"       csv:"Contributor 1 Campaign"`
	Contributor1TouchType     string `json:"contributor1_touch_type"     csv:"Contributor 1 Touch Type"`
	Contributor1TouchTime     string `json:"contributor1_touch_time"     csv:"Contributor 1 Touch Time"`
	Contributor2Partner       string `json:"contributor1_af_prt"         csv:"Contributor 2 Partner"`
	Contributor2MediaSource   string `json:"contributor2_media_source"   csv:"Contributor 2 Media Source"`
	Contributor2Campaign      string `json:"contributor2_campaign"       csv:"Contributor 2 Campaign"`
	Contributor2TouchType     string `json:"contributor2_touch_type"     csv:"Contributor 2 Touch Type"`
	Contributor2TouchTime     string `json:"contributor2_touch_time"     csv:"Contributor 2 Touch Time"`
	Contributor3Partner       string `json:"contributor3_af_prt"         csv:"Contributor 3 Partner"`
	Contributor3MediaSource   string `json:"contributor3_media_source"   csv:"Contributor 3 Media Source"`
	Contributor3Campaign      string `json:"contributor3_campaign"       csv:"Contributor 3 Campaign"`
	Contributor3TouchType     string `json:"contributor3_touch_type"     csv:"Contributor 3 Touch Type"`
	Contributor3TouchTime     string `json:"contributor3_touch_time"     csv:"Contributor 3 Touch Time"`
	Region                    string `json:"region"                      csv:"Region"`
	CountryCode               string `json:"country_code"                csv:"Country Code"`
	State                     string `json:"state"                       csv:"State"`
	City                      string `json:"city"                        csv:"City"`
	PostalCode                string `json:"postal_code"                 csv:"Postal Code"`
	DMA                       string `json:"dma"                         csv:"DMA"`
	IP                        string `json:"ip"                          csv:"IP"`
	WIFI                      string `json:"wifi"                        csv:"WIFI"`
	Operator                  string `json:"operator"                    csv:"Operator"`
	Carrier                   string `json:"carrier"                     csv:"Carrier"`
	Language                  string `json:"language"                    csv:"Language"`
	AppsFlyerID               string `json:"appsflyer_id"                csv:"AppsFlyer ID"`
	AdvertisingID             string `json:"advertising_id"              csv:"Advertising ID"`
	IDFA                      string `json:"idfa"                        csv:"IDFA"`
	AndroidID                 string `json:"android_id"                  csv:"Android ID"`
	CustomerUserID            string `json:"customer_user_id"            csv:"Customer User ID"`
	IMEI                      string `json:"imei"                        csv:"IMEI"`
	IDFV                      string `json:"idfv"                        csv:"IDFV"`
	Platform                  string `json:"platform"                    csv:"Platform"`
	DeviceType                string `json:"device_type"                 csv:"Device Type"`
	OSVersion                 string `json:"os_version"                  csv:"OS Version"`
	AppVersion                string `json:"app_version"                 csv:"App Version"`
	SDKVersion                string `json:"sdk_version"                 csv:"SDK Version"`
	AppID                     string `json:"app_id"                      csv:"App ID"`
	AppName                   string `json:"app_name"                    csv:"App Name"`
	BundleID                  string `json:"bundle_id"                   csv:"Bundle ID"`
	IsRetargeting             string `json:"is_retargeting"              csv:"Is Retargeting"`
	RetargetingConversionType string `json:"retargeting_conversion_type" csv:"Retargeting Conversion Type"`
	AttributionLookback       string `json:"af_attribution_lookback"     csv:"Attribution Lookback"`
	ReengagementWindow        string `json:"af_reengagement_window"      csv:"Reengagement Window"`
	IsPrimaryAttribution      string `json:"is_primary_attribution"      csv:"Is Primary Attribution"`
	UserAgent                 string `json:"user_agent"                  csv:"User Agent"`
	HTTPReferrer              string `json:"http_referrer"               csv:"HTTP Referrer"`
	OriginalURL               string `json:"original_url"                csv:"Original URL"`
	InstallAppStore           string `json:"install_app_store"           csv:"Install App Store"`
}

func (r *Report) GetAttributedTouchTime() (time.Time, error) {
	return ParseDateTimeFormat(r.AttributedTouchTime)
}

func (r *Report) GetInstallTime() (time.Time, error) {
	return ParseDateTimeFormat(r.InstallTime)
}

func (r *Report) GetEventTime() (time.Time, error) {
	return ParseDateTimeFormat(r.EventTime)
}
