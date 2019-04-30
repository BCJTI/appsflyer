package appsflyer

const (
	AppsFlyerVersion  = "v5"
	defaultAPIBaseURL = "https://hq.appsflyer.com"
)

// https://support.appsflyer.com/hc/en-us/articles/207034346-Pull-APIs-Pulling-AppsFlyer-Reports-by-APIs
// https://support.appsflyer.com/hc/en-us/articles/207034216-Partners-Media-Source-Report
// https://support.appsflyer.com/hc/en-us/articles/207034256-Partners-by-Date-Report
// https://support.appsflyer.com/hc/en-us/articles/207034226-Geo-Performance-Report

const (
	// Performance Reports
	PartnersReport      = "partners_report/"
	PartnersDailyReport = "partners_by_date_report/"
	DailyReport         = "daily_report/"
	GeoReport           = "geo_report/"
	GeoDailyReport      = "geo_by_date_report/"

	// Raw Data Reports
	InstallsReport        = "installs_report/"
	InAppEventsReport     = "in_app_events_report/"
	UninstallsReport      = "uninstall_events_report/"
	OrganicInstallsReport = "organic_installs_report/"
	OrganicInAppReport    = "organic_in_app_events_report/"

	// Protect360 Reports
	BlockedInstallsReport    = "blocked_installs_report/"
	BlockedInAppEventsReport = "blocked_in_app_events_report/"
	BlockedClicksReport      = "blocked_clicks_report/"

	// Validation Rules Reports
	InvalidInstallsReport    = "invalid_installs_report/"
	InvalidInAppEventsReport = "invalid_in_app_events_report/"
)
