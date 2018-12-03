// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package pagespeedonline provides access to the PageSpeed Insights API.
//
// See https://developers.google.com/speed/docs/insights/v5/get-started
//
// Usage example:
//
//   import "google.golang.org/api/pagespeedonline/v5"
//   ...
//   pagespeedonlineService, err := pagespeedonline.New(oauthHttpClient)
package pagespeedonline // import "google.golang.org/api/pagespeedonline/v5"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "pagespeedonline:v5"
const apiName = "pagespeedonline"
const apiVersion = "v5"
const basePath = "https://www.googleapis.com/pagespeedonline/v5/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Pagespeedapi = NewPagespeedapiService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Pagespeedapi *PagespeedapiService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewPagespeedapiService(s *Service) *PagespeedapiService {
	rs := &PagespeedapiService{s: s}
	return rs
}

type PagespeedapiService struct {
	s *Service
}

type GoogleprotobufValue interface{}

type LighthouseAuditResultV5 struct {
	// Description: The description of the audit.
	Description string `json:"description,omitempty"`

	// Details: Freeform details section of the audit.
	Details googleapi.RawMessage `json:"details,omitempty"`

	// DisplayValue: The value that should be displayed on the UI for this
	// audit.
	DisplayValue string `json:"displayValue,omitempty"`

	// ErrorMessage: An error message from a thrown error inside the audit.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Explanation: An explanation of the errors in the audit.
	Explanation string `json:"explanation,omitempty"`

	// Id: The audit's id.
	Id string `json:"id,omitempty"`

	Score interface{} `json:"score,omitempty"`

	// ScoreDisplayMode: The enumerated score display mode.
	ScoreDisplayMode string `json:"scoreDisplayMode,omitempty"`

	// Title: The human readable title.
	Title string `json:"title,omitempty"`

	Warnings interface{} `json:"warnings,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseAuditResultV5) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseAuditResultV5
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type LighthouseCategoryV5 struct {
	// AuditRefs: An array of references to all the audit members of this
	// category.
	AuditRefs []*LighthouseCategoryV5AuditRefs `json:"auditRefs,omitempty"`

	// Description: A more detailed description of the category and its
	// importance.
	Description string `json:"description,omitempty"`

	// Id: The string identifier of the category.
	Id string `json:"id,omitempty"`

	// ManualDescription: A description for the manual audits in the
	// category.
	ManualDescription string `json:"manualDescription,omitempty"`

	Score interface{} `json:"score,omitempty"`

	// Title: The human-friendly name of the category.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuditRefs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditRefs") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseCategoryV5) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseCategoryV5
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type LighthouseCategoryV5AuditRefs struct {
	// Group: The category group that the audit belongs to (optional).
	Group string `json:"group,omitempty"`

	// Id: The audit ref id.
	Id string `json:"id,omitempty"`

	// Weight: The weight this audit's score has on the overall category
	// score.
	Weight float64 `json:"weight,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Group") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Group") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseCategoryV5AuditRefs) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseCategoryV5AuditRefs
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LighthouseCategoryV5AuditRefs) UnmarshalJSON(data []byte) error {
	type NoMethod LighthouseCategoryV5AuditRefs
	var s1 struct {
		Weight gensupport.JSONFloat64 `json:"weight"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Weight = float64(s1.Weight)
	return nil
}

type LighthouseResultV5 struct {
	// Audits: Map of audits in the LHR.
	Audits map[string]LighthouseAuditResultV5 `json:"audits,omitempty"`

	// Categories: Map of categories in the LHR.
	Categories map[string]LighthouseCategoryV5 `json:"categories,omitempty"`

	// CategoryGroups: Map of category groups in the LHR.
	CategoryGroups map[string]LighthouseResultV5CategoryGroups `json:"categoryGroups,omitempty"`

	// ConfigSettings: The configuration settings for this LHR.
	ConfigSettings *LighthouseResultV5ConfigSettings `json:"configSettings,omitempty"`

	// Environment: Environment settings that were used when making this
	// LHR.
	Environment *LighthouseResultV5Environment `json:"environment,omitempty"`

	// FetchTime: The time that this run was fetched.
	FetchTime string `json:"fetchTime,omitempty"`

	// FinalUrl: The final resolved url that was audited.
	FinalUrl string `json:"finalUrl,omitempty"`

	// I18n: The internationalization strings that are required to render
	// the LHR.
	I18n *LighthouseResultV5I18n `json:"i18n,omitempty"`

	// LighthouseVersion: The lighthouse version that was used to generate
	// this LHR.
	LighthouseVersion string `json:"lighthouseVersion,omitempty"`

	// RequestedUrl: The original requested url.
	RequestedUrl string `json:"requestedUrl,omitempty"`

	// RunWarnings: List of all run warnings in the LHR. Will always output
	// to at least `[]`.
	RunWarnings []interface{} `json:"runWarnings,omitempty"`

	// RuntimeError: A top-level error message that, if present, indicates a
	// serious enough problem that this Lighthouse result may need to be
	// discarded.
	RuntimeError *LighthouseResultV5RuntimeError `json:"runtimeError,omitempty"`

	// Timing: Timing information for this LHR.
	Timing *LighthouseResultV5Timing `json:"timing,omitempty"`

	// UserAgent: The user agent that was used to run this LHR.
	UserAgent string `json:"userAgent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audits") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audits") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LighthouseResultV5CategoryGroups: A grouping contained in a category
// that groups similar audits together.
type LighthouseResultV5CategoryGroups struct {
	// Description: An optional human readable description of the category
	// group.
	Description string `json:"description,omitempty"`

	// Title: The title of the category group.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5CategoryGroups) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5CategoryGroups
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LighthouseResultV5ConfigSettings: The configuration settings for this
// LHR.
type LighthouseResultV5ConfigSettings struct {
	// EmulatedFormFactor: The form factor the emulation should use.
	EmulatedFormFactor string `json:"emulatedFormFactor,omitempty"`

	// Locale: The locale setting.
	Locale string `json:"locale,omitempty"`

	OnlyCategories interface{} `json:"onlyCategories,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EmulatedFormFactor")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EmulatedFormFactor") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5ConfigSettings) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5ConfigSettings
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LighthouseResultV5Environment: Environment settings that were used
// when making this LHR.
type LighthouseResultV5Environment struct {
	// BenchmarkIndex: The benchmark index number that indicates rough
	// device class.
	BenchmarkIndex float64 `json:"benchmarkIndex,omitempty"`

	// HostUserAgent: The user agent string of the version of Chrome used.
	HostUserAgent string `json:"hostUserAgent,omitempty"`

	// NetworkUserAgent: The user agent string that was sent over the
	// network.
	NetworkUserAgent string `json:"networkUserAgent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BenchmarkIndex") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BenchmarkIndex") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5Environment) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5Environment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LighthouseResultV5Environment) UnmarshalJSON(data []byte) error {
	type NoMethod LighthouseResultV5Environment
	var s1 struct {
		BenchmarkIndex gensupport.JSONFloat64 `json:"benchmarkIndex"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.BenchmarkIndex = float64(s1.BenchmarkIndex)
	return nil
}

// LighthouseResultV5I18n: The internationalization strings that are
// required to render the LHR.
type LighthouseResultV5I18n struct {
	// RendererFormattedStrings: Internationalized strings that are
	// formatted to the locale in configSettings.
	RendererFormattedStrings *LighthouseResultV5I18nRendererFormattedStrings `json:"rendererFormattedStrings,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "RendererFormattedStrings") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RendererFormattedStrings")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5I18n) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5I18n
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LighthouseResultV5I18nRendererFormattedStrings: Internationalized
// strings that are formatted to the locale in configSettings.
type LighthouseResultV5I18nRendererFormattedStrings struct {
	// AuditGroupExpandTooltip: The tooltip text on an expandable chevron
	// icon.
	AuditGroupExpandTooltip string `json:"auditGroupExpandTooltip,omitempty"`

	// CrcInitialNavigation: The label for the initial request in a critical
	// request chain.
	CrcInitialNavigation string `json:"crcInitialNavigation,omitempty"`

	// CrcLongestDurationLabel: The label for values shown in the summary of
	// critical request chains.
	CrcLongestDurationLabel string `json:"crcLongestDurationLabel,omitempty"`

	// ErrorLabel: The label shown next to an audit or metric that has had
	// an error.
	ErrorLabel string `json:"errorLabel,omitempty"`

	// ErrorMissingAuditInfo: The error string shown next to an erroring
	// audit.
	ErrorMissingAuditInfo string `json:"errorMissingAuditInfo,omitempty"`

	// LabDataTitle: The title of the lab data performance category.
	LabDataTitle string `json:"labDataTitle,omitempty"`

	// LsPerformanceCategoryDescription: The disclaimer shown under
	// performance explaning that the network can vary.
	LsPerformanceCategoryDescription string `json:"lsPerformanceCategoryDescription,omitempty"`

	// ManualAuditsGroupTitle: The heading shown above a list of audits that
	// were not computerd in the run.
	ManualAuditsGroupTitle string `json:"manualAuditsGroupTitle,omitempty"`

	// NotApplicableAuditsGroupTitle: The heading shown above a list of
	// audits that do not apply to a page.
	NotApplicableAuditsGroupTitle string `json:"notApplicableAuditsGroupTitle,omitempty"`

	// OpportunityResourceColumnLabel: The heading for the estimated page
	// load savings opportunity of an audit.
	OpportunityResourceColumnLabel string `json:"opportunityResourceColumnLabel,omitempty"`

	// OpportunitySavingsColumnLabel: The heading for the estimated page
	// load savings of opportunity audits.
	OpportunitySavingsColumnLabel string `json:"opportunitySavingsColumnLabel,omitempty"`

	// PassedAuditsGroupTitle: The heading that is shown above a list of
	// audits that are passing.
	PassedAuditsGroupTitle string `json:"passedAuditsGroupTitle,omitempty"`

	// ScorescaleLabel: The label that explains the score gauges scale
	// (0-49, 50-89, 90-100).
	ScorescaleLabel string `json:"scorescaleLabel,omitempty"`

	// ToplevelWarningsMessage: The label shown preceding important warnings
	// that may have invalidated an entire report.
	ToplevelWarningsMessage string `json:"toplevelWarningsMessage,omitempty"`

	// VarianceDisclaimer: The disclaimer shown below a performance metric
	// value.
	VarianceDisclaimer string `json:"varianceDisclaimer,omitempty"`

	// WarningHeader: The label shown above a bulleted list of warnings.
	WarningHeader string `json:"warningHeader,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AuditGroupExpandTooltip") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditGroupExpandTooltip")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5I18nRendererFormattedStrings) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5I18nRendererFormattedStrings
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LighthouseResultV5RuntimeError: A top-level error message that, if
// present, indicates a serious enough problem that this Lighthouse
// result may need to be discarded.
type LighthouseResultV5RuntimeError struct {
	// Code: The enumerated Lighthouse Error code.
	Code string `json:"code,omitempty"`

	// Message: A human readable message explaining the error code.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5RuntimeError) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5RuntimeError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LighthouseResultV5Timing: Timing information for this LHR.
type LighthouseResultV5Timing struct {
	// Total: The total duration of Lighthouse's run.
	Total float64 `json:"total,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Total") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Total") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LighthouseResultV5Timing) MarshalJSON() ([]byte, error) {
	type NoMethod LighthouseResultV5Timing
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LighthouseResultV5Timing) UnmarshalJSON(data []byte) error {
	type NoMethod LighthouseResultV5Timing
	var s1 struct {
		Total gensupport.JSONFloat64 `json:"total"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Total = float64(s1.Total)
	return nil
}

type PagespeedApiLoadingExperienceV5 struct {
	// Id: The url, pattern or origin which the metrics are on.
	Id string `json:"id,omitempty"`

	InitialUrl string `json:"initial_url,omitempty"`

	Metrics map[string]PagespeedApiLoadingExperienceV5Metrics `json:"metrics,omitempty"`

	OverallCategory string `json:"overall_category,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiLoadingExperienceV5) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiLoadingExperienceV5
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiLoadingExperienceV5Metrics: The type of the metric.
type PagespeedApiLoadingExperienceV5Metrics struct {
	Category string `json:"category,omitempty"`

	Distributions []*PagespeedApiLoadingExperienceV5MetricsDistributions `json:"distributions,omitempty"`

	Percentile int64 `json:"percentile,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Category") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Category") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiLoadingExperienceV5Metrics) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiLoadingExperienceV5Metrics
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiLoadingExperienceV5MetricsDistributions struct {
	Max int64 `json:"max,omitempty"`

	Min int64 `json:"min,omitempty"`

	Proportion float64 `json:"proportion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Max") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiLoadingExperienceV5MetricsDistributions) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiLoadingExperienceV5MetricsDistributions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *PagespeedApiLoadingExperienceV5MetricsDistributions) UnmarshalJSON(data []byte) error {
	type NoMethod PagespeedApiLoadingExperienceV5MetricsDistributions
	var s1 struct {
		Proportion gensupport.JSONFloat64 `json:"proportion"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Proportion = float64(s1.Proportion)
	return nil
}

type PagespeedApiPagespeedResponseV5 struct {
	// AnalysisUTCTimestamp: The UTC timestamp of this analysis.
	AnalysisUTCTimestamp string `json:"analysisUTCTimestamp,omitempty"`

	// CaptchaResult: The captcha verify result
	CaptchaResult string `json:"captchaResult,omitempty"`

	// Id: Canonicalized and final URL for the document, after following
	// page redirects (if any).
	Id string `json:"id,omitempty"`

	// Kind: Kind of result.
	Kind string `json:"kind,omitempty"`

	// LighthouseResult: Lighthouse response for the audit url as an object.
	LighthouseResult *LighthouseResultV5 `json:"lighthouseResult,omitempty"`

	// LoadingExperience: Metrics of end users' page loading experience.
	LoadingExperience *PagespeedApiLoadingExperienceV5 `json:"loadingExperience,omitempty"`

	// OriginLoadingExperience: Metrics of the aggregated page loading
	// experience of the origin
	OriginLoadingExperience *PagespeedApiLoadingExperienceV5 `json:"originLoadingExperience,omitempty"`

	// Version: The version of PageSpeed used to generate these results.
	Version *PagespeedApiPagespeedResponseV5Version `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AnalysisUTCTimestamp") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnalysisUTCTimestamp") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV5) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV5
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV5Version: The version of PageSpeed used
// to generate these results.
type PagespeedApiPagespeedResponseV5Version struct {
	// Major: The major version number of PageSpeed used to generate these
	// results.
	Major int64 `json:"major,omitempty"`

	// Minor: The minor version number of PageSpeed used to generate these
	// results.
	Minor int64 `json:"minor,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Major") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Major") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV5Version) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV5Version
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "pagespeedonline.pagespeedapi.runpagespeed":

type PagespeedapiRunpagespeedCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Runpagespeed: Runs PageSpeed analysis on the page at the specified
// URL, and returns PageSpeed scores, a list of suggestions to make that
// page faster, and other information.
func (r *PagespeedapiService) Runpagespeed(url string) *PagespeedapiRunpagespeedCall {
	c := &PagespeedapiRunpagespeedCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("url", url)
	return c
}

// Category sets the optional parameter "category": A Lighthouse
// category to run; if none are given, only Performance category will be
// run
//
// Possible values:
//   "accessibility"
//   "best-practices"
//   "performance"
//   "pwa"
//   "seo"
func (c *PagespeedapiRunpagespeedCall) Category(category ...string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.SetMulti("category", append([]string{}, category...))
	return c
}

// Locale sets the optional parameter "locale": The locale used to
// localize formatted results
func (c *PagespeedapiRunpagespeedCall) Locale(locale string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("locale", locale)
	return c
}

// Strategy sets the optional parameter "strategy": The analysis
// strategy (desktop or mobile) to use, and desktop is the default
//
// Possible values:
//   "desktop" - Fetch and analyze the URL for desktop browsers
//   "mobile" - Fetch and analyze the URL for mobile devices
func (c *PagespeedapiRunpagespeedCall) Strategy(strategy string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("strategy", strategy)
	return c
}

// UtmCampaign sets the optional parameter "utm_campaign": Campaign name
// for analytics.
func (c *PagespeedapiRunpagespeedCall) UtmCampaign(utmCampaign string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("utm_campaign", utmCampaign)
	return c
}

// UtmSource sets the optional parameter "utm_source": Campaign source
// for analytics.
func (c *PagespeedapiRunpagespeedCall) UtmSource(utmSource string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("utm_source", utmSource)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagespeedapiRunpagespeedCall) Fields(s ...googleapi.Field) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PagespeedapiRunpagespeedCall) IfNoneMatch(entityTag string) *PagespeedapiRunpagespeedCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PagespeedapiRunpagespeedCall) Context(ctx context.Context) *PagespeedapiRunpagespeedCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PagespeedapiRunpagespeedCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PagespeedapiRunpagespeedCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "runPagespeed")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "pagespeedonline.pagespeedapi.runpagespeed" call.
// Exactly one of *PagespeedApiPagespeedResponseV5 or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *PagespeedApiPagespeedResponseV5.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PagespeedapiRunpagespeedCall) Do(opts ...googleapi.CallOption) (*PagespeedApiPagespeedResponseV5, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PagespeedApiPagespeedResponseV5{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs PageSpeed analysis on the page at the specified URL, and returns PageSpeed scores, a list of suggestions to make that page faster, and other information.",
	//   "httpMethod": "GET",
	//   "id": "pagespeedonline.pagespeedapi.runpagespeed",
	//   "parameterOrder": [
	//     "url"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "A Lighthouse category to run; if none are given, only Performance category will be run",
	//       "enum": [
	//         "accessibility",
	//         "best-practices",
	//         "performance",
	//         "pwa",
	//         "seo"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "The locale used to localize formatted results",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+((_|-)[a-zA-Z]+)?",
	//       "type": "string"
	//     },
	//     "strategy": {
	//       "description": "The analysis strategy (desktop or mobile) to use, and desktop is the default",
	//       "enum": [
	//         "desktop",
	//         "mobile"
	//       ],
	//       "enumDescriptions": [
	//         "Fetch and analyze the URL for desktop browsers",
	//         "Fetch and analyze the URL for mobile devices"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The URL to fetch and analyze",
	//       "location": "query",
	//       "pattern": "(?i)http(s)?://.*",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "utm_campaign": {
	//       "description": "Campaign name for analytics.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "utm_source": {
	//       "description": "Campaign source for analytics.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "runPagespeed",
	//   "response": {
	//     "$ref": "PagespeedApiPagespeedResponseV5"
	//   }
	// }

}
