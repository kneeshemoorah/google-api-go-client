// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package alertcenter provides access to the G Suite Alert Center API.
//
// See https://developers.google.com/admin-sdk/alertcenter/
//
// Usage example:
//
//   import "google.golang.org/api/alertcenter/v1beta1"
//   ...
//   alertcenterService, err := alertcenter.New(oauthHttpClient)
package alertcenter // import "google.golang.org/api/alertcenter/v1beta1"

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

const apiId = "alertcenter:v1beta1"
const apiName = "alertcenter"
const apiVersion = "v1beta1"
const basePath = "https://alertcenter.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Alerts = NewAlertsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Alerts *AlertsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAlertsService(s *Service) *AlertsService {
	rs := &AlertsService{s: s}
	rs.Feedback = NewAlertsFeedbackService(s)
	return rs
}

type AlertsService struct {
	s *Service

	Feedback *AlertsFeedbackService
}

func NewAlertsFeedbackService(s *Service) *AlertsFeedbackService {
	rs := &AlertsFeedbackService{s: s}
	return rs
}

type AlertsFeedbackService struct {
	s *Service
}

// AccountWarning: Alerts for user account warning events.
type AccountWarning struct {
	// Email: Required. The email of the user that this event belongs to.
	Email string `json:"email,omitempty"`

	// LoginDetails: Optional. Details of the login action associated with
	// the warning event.
	// This is only available for:
	//
	// * Suspicious login
	// * Suspicious login (less secure app)
	// * User suspended (suspicious activity)
	LoginDetails *LoginDetails `json:"loginDetails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AccountWarning) MarshalJSON() ([]byte, error) {
	type NoMethod AccountWarning
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Alert: An alert affecting a customer.
// All fields are read-only once created.
type Alert struct {
	// AlertId: Output only. The unique identifier for the alert.
	AlertId string `json:"alertId,omitempty"`

	// CreateTime: Output only. The time this alert was created.
	CreateTime string `json:"createTime,omitempty"`

	// CustomerId: Output only. The unique identifier of the Google account
	// of the customer.
	CustomerId string `json:"customerId,omitempty"`

	// Data: Optional. The data associated with this alert, for
	// example
	// google.apps.alertcenter.type.DeviceCompromised.
	Data googleapi.RawMessage `json:"data,omitempty"`

	// Deleted: Output only. `True` if this alert is marked for deletion.
	Deleted bool `json:"deleted,omitempty"`

	// EndTime: Optional. The time the event that caused this alert ceased
	// being active.
	// If provided, the end time must not be earlier than the start time.
	// If not provided, the end time defaults to the start time.
	EndTime string `json:"endTime,omitempty"`

	// SecurityInvestigationToolLink: Output only. An optional
	// [Security Investigation
	// Tool](https://support.google.com/a/answer/7575955)
	// query for this alert.
	SecurityInvestigationToolLink string `json:"securityInvestigationToolLink,omitempty"`

	// Source: Required. A unique identifier for the system that is reported
	// the alert.
	//
	// Supported sources are any of the following:
	//
	// * Google Operations
	// * Mobile device management
	// * Gmail phishing
	// * Domain wide takeout
	// * Government attack warning
	// * Google identity
	Source string `json:"source,omitempty"`

	// StartTime: Required. The time the event that caused this alert was
	// started or
	// detected.
	StartTime string `json:"startTime,omitempty"`

	// Type: Required. The type of the alert.
	// For a list of available alert types see
	// [G Suite Alert types](/admin-sdk/alertcenter/reference/alert-types).
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AlertId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlertId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Alert) MarshalJSON() ([]byte, error) {
	type NoMethod Alert
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AlertFeedback: A customer feedback about an alert.
type AlertFeedback struct {
	// AlertId: Output only. The alert identifier.
	AlertId string `json:"alertId,omitempty"`

	// CreateTime: Output only. The time this feedback was created.
	CreateTime string `json:"createTime,omitempty"`

	// CustomerId: Output only. The unique identifier of the Google account
	// of the customer.
	CustomerId string `json:"customerId,omitempty"`

	// Email: Output only. The email of the user that provided the feedback.
	Email string `json:"email,omitempty"`

	// FeedbackId: Output only. The unique identifier for the feedback.
	FeedbackId string `json:"feedbackId,omitempty"`

	// Type: Required. The type of the feedback.
	//
	// Possible values:
	//   "ALERT_FEEDBACK_TYPE_UNSPECIFIED" - The feedback type is not
	// specified.
	//   "NOT_USEFUL" - The alert report is not useful.
	//   "SOMEWHAT_USEFUL" - The alert report is somewhat useful.
	//   "VERY_USEFUL" - The alert report is very useful.
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AlertId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlertId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AlertFeedback) MarshalJSON() ([]byte, error) {
	type NoMethod AlertFeedback
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Attachment: Attachment with application-specific information about an
// alert.
type Attachment struct {
	// Csv: A CSV file attachment.
	Csv *Csv `json:"csv,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Csv") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Csv") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Attachment) MarshalJSON() ([]byte, error) {
	type NoMethod Attachment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BadWhitelist: Alert for setting the domain or IP that malicious email
// comes from as
// whitelisted domain or IP in Gmail advanced settings.
type BadWhitelist struct {
	// DomainId: The domain ID.
	DomainId *DomainId `json:"domainId,omitempty"`

	// MaliciousEntity: The entity whose actions triggered a Gmail phishing
	// alert.
	MaliciousEntity *MaliciousEntity `json:"maliciousEntity,omitempty"`

	// Messages: The list of messages contained by this alert.
	Messages []*GmailMessageInfo `json:"messages,omitempty"`

	// SourceIp: The source IP address of the malicious email, for
	// example,
	// `127.0.0.1`.
	SourceIp string `json:"sourceIp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DomainId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DomainId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BadWhitelist) MarshalJSON() ([]byte, error) {
	type NoMethod BadWhitelist
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Csv: A representation of a CSV file attachment, as a list of column
// headers and
// a list of data rows.
type Csv struct {
	// DataRows: The list of data rows in a CSV file, as string arrays
	// rather than as a
	// single comma-separated string.
	DataRows []*CsvRow `json:"dataRows,omitempty"`

	// Headers: The list of headers for data columns in a CSV file.
	Headers []string `json:"headers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataRows") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DataRows") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Csv) MarshalJSON() ([]byte, error) {
	type NoMethod Csv
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CsvRow: A representation of a single data row in a CSV file.
type CsvRow struct {
	// Entries: The data entries in a CSV file row, as a string array rather
	// than a
	// single comma-separated string.
	Entries []string `json:"entries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entries") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CsvRow) MarshalJSON() ([]byte, error) {
	type NoMethod CsvRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceCompromised: A mobile device compromised alert. Derived from
// audit logs.
type DeviceCompromised struct {
	// Email: The email of the user this alert was created for.
	Email string `json:"email,omitempty"`

	// Events: Required. The list of security events.
	Events []*DeviceCompromisedSecurityDetail `json:"events,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeviceCompromised) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceCompromised
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceCompromisedSecurityDetail: Detailed information of a single MDM
// device compromised event.
type DeviceCompromisedSecurityDetail struct {
	// DeviceCompromisedState: The device compromised state. Possible values
	// are "Compromised" or
	// "Not Compromised".
	DeviceCompromisedState string `json:"deviceCompromisedState,omitempty"`

	// DeviceId: Required. The device ID.
	DeviceId string `json:"deviceId,omitempty"`

	// DeviceModel: The model of the device.
	DeviceModel string `json:"deviceModel,omitempty"`

	// DeviceType: The type of the device.
	DeviceType string `json:"deviceType,omitempty"`

	// IosVendorId: Required for iOS, empty for others.
	IosVendorId string `json:"iosVendorId,omitempty"`

	// ResourceId: The device resource ID.
	ResourceId string `json:"resourceId,omitempty"`

	// SerialNumber: The serial number of the device.
	SerialNumber string `json:"serialNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DeviceCompromisedState") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceCompromisedState")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DeviceCompromisedSecurityDetail) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceCompromisedSecurityDetail
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DomainId: Domain ID of Gmail phishing alerts.
type DomainId struct {
	// CustomerPrimaryDomain: The primary domain for the customer.
	CustomerPrimaryDomain string `json:"customerPrimaryDomain,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CustomerPrimaryDomain") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomerPrimaryDomain") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DomainId) MarshalJSON() ([]byte, error) {
	type NoMethod DomainId
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DomainWideTakeoutInitiated: A takeout operation for the entire domain
// was initiated by an admin. Derived
// from audit logs.
type DomainWideTakeoutInitiated struct {
	// Email: The email of the admin who initiated the takeout.
	Email string `json:"email,omitempty"`

	// TakeoutRequestId: The takeout request ID.
	TakeoutRequestId string `json:"takeoutRequestId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DomainWideTakeoutInitiated) MarshalJSON() ([]byte, error) {
	type NoMethod DomainWideTakeoutInitiated
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GmailMessageInfo: Details of a message in phishing spike alert.
type GmailMessageInfo struct {
	// AttachmentsSha256Hash: The `SHA256` hash of email's attachment and
	// all MIME parts.
	AttachmentsSha256Hash []string `json:"attachmentsSha256Hash,omitempty"`

	// Date: The date the malicious email was sent.
	Date string `json:"date,omitempty"`

	// Md5HashMessageBody: The hash of the message body text.
	Md5HashMessageBody string `json:"md5HashMessageBody,omitempty"`

	// Md5HashSubject: The MD5 Hash of email's subject (only available for
	// reported emails).
	Md5HashSubject string `json:"md5HashSubject,omitempty"`

	// MessageBodySnippet: The snippet of the message body text (only
	// available for reported emails).
	MessageBodySnippet string `json:"messageBodySnippet,omitempty"`

	// MessageId: The message ID.
	MessageId string `json:"messageId,omitempty"`

	// Recipient: The recipient of this email.
	Recipient string `json:"recipient,omitempty"`

	// SubjectText: The email subject text (only available for reported
	// emails).
	SubjectText string `json:"subjectText,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AttachmentsSha256Hash") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttachmentsSha256Hash") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GmailMessageInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GmailMessageInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleOperations: An incident reported by Google Operations for a G
// Suite application.
type GoogleOperations struct {
	// AffectedUserEmails: The list of emails which correspond to the users
	// directly affected by the
	// incident.
	AffectedUserEmails []string `json:"affectedUserEmails,omitempty"`

	// AttachmentData: Optional. Application-specific data for an incident,
	// provided when the
	// G Suite application which reported the incident cannot be
	// completely
	// restored to a valid state.
	AttachmentData *Attachment `json:"attachmentData,omitempty"`

	// Description: A detailed, freeform incident description.
	Description string `json:"description,omitempty"`

	// Title: A one-line incident description.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AffectedUserEmails")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AffectedUserEmails") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleOperations) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleOperations
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAlertFeedbackResponse: Response message for an alert feedback
// listing request.
type ListAlertFeedbackResponse struct {
	// Feedback: The list of alert feedback.
	// Feedback entries for each alert are ordered by creation time
	// descending.
	Feedback []*AlertFeedback `json:"feedback,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Feedback") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Feedback") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAlertFeedbackResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAlertFeedbackResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAlertsResponse: Response message for an alert listing request.
type ListAlertsResponse struct {
	// Alerts: The list of alerts.
	Alerts []*Alert `json:"alerts,omitempty"`

	// NextPageToken: The token for the next page. If not empty, indicates
	// that there may be more
	// alerts that match the listing request; this value can be used in
	// a
	// subsequent ListAlertsRequest to get alerts continuing from last
	// result
	// of the current list call.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Alerts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alerts") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAlertsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAlertsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LoginDetails: The details of the login action.
type LoginDetails struct {
	// IpAddress: Optional. The human-readable IP address (for
	// example,
	// `11.22.33.44`) that is associated with the warning event.
	IpAddress string `json:"ipAddress,omitempty"`

	// LoginTime: Optional. The successful login time that is associated
	// with the warning
	// event. This will not be present for blocked login attempts.
	LoginTime string `json:"loginTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IpAddress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IpAddress") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LoginDetails) MarshalJSON() ([]byte, error) {
	type NoMethod LoginDetails
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MailPhishing: Proto for all phishing alerts with common
// payload.
// Supported types are any of the following:
//
// * User reported phishing
// * User reported spam spike
// * Suspicious message reported
// * Phishing reclassification
// * Malware reclassification
type MailPhishing struct {
	// DomainId: The domain ID.
	DomainId *DomainId `json:"domainId,omitempty"`

	// IsInternal: If `true`, the email originated from within the
	// organization.
	IsInternal bool `json:"isInternal,omitempty"`

	// MaliciousEntity: The entity whose actions triggered a Gmail phishing
	// alert.
	MaliciousEntity *MaliciousEntity `json:"maliciousEntity,omitempty"`

	// Messages: The list of messages contained by this alert.
	Messages []*GmailMessageInfo `json:"messages,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DomainId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DomainId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MailPhishing) MarshalJSON() ([]byte, error) {
	type NoMethod MailPhishing
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MaliciousEntity: Entity whose actions triggered a Gmail phishing
// alert.
type MaliciousEntity struct {
	// FromHeader: The sender email address.
	FromHeader string `json:"fromHeader,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FromHeader") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FromHeader") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MaliciousEntity) MarshalJSON() ([]byte, error) {
	type NoMethod MaliciousEntity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PhishingSpike: Alert for a spike in user reported phishing.
// <aside class="warning"><b>Warning</b>: This type has been deprecated.
// Use
// [MailPhishing](/admin-sdk/alertcenter/reference/rest/v1beta1/MailP
// hishing)
// instead.</aside>
type PhishingSpike struct {
	// DomainId: The domain ID.
	DomainId *DomainId `json:"domainId,omitempty"`

	// IsInternal: If `true`, the email originated from within the
	// organization.
	IsInternal bool `json:"isInternal,omitempty"`

	// MaliciousEntity: The entity whose actions triggered a Gmail phishing
	// alert.
	MaliciousEntity *MaliciousEntity `json:"maliciousEntity,omitempty"`

	// Messages: The list of messages contained by this alert.
	Messages []*GmailMessageInfo `json:"messages,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DomainId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DomainId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PhishingSpike) MarshalJSON() ([]byte, error) {
	type NoMethod PhishingSpike
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StateSponsoredAttack: A state-sponsored attack alert. Derived from
// audit logs.
type StateSponsoredAttack struct {
	// Email: The email of the user this incident was created for.
	Email string `json:"email,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StateSponsoredAttack) MarshalJSON() ([]byte, error) {
	type NoMethod StateSponsoredAttack
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuspiciousActivity: A mobile suspicious activity alert. Derived from
// audit logs.
type SuspiciousActivity struct {
	// Email: The email of the user this alert was created for.
	Email string `json:"email,omitempty"`

	// Events: Required. The list of security events.
	Events []*SuspiciousActivitySecurityDetail `json:"events,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuspiciousActivity) MarshalJSON() ([]byte, error) {
	type NoMethod SuspiciousActivity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuspiciousActivitySecurityDetail: Detailed information of a single
// MDM suspicious activity event.
type SuspiciousActivitySecurityDetail struct {
	// DeviceId: Required. The device ID.
	DeviceId string `json:"deviceId,omitempty"`

	// DeviceModel: The model of the device.
	DeviceModel string `json:"deviceModel,omitempty"`

	// DeviceProperty: The device property which was changed.
	DeviceProperty string `json:"deviceProperty,omitempty"`

	// DeviceType: The type of the device.
	DeviceType string `json:"deviceType,omitempty"`

	// IosVendorId: Required for iOS, empty for others.
	IosVendorId string `json:"iosVendorId,omitempty"`

	// NewValue: The new value of the device property after the change.
	NewValue string `json:"newValue,omitempty"`

	// OldValue: The old value of the device property before the change.
	OldValue string `json:"oldValue,omitempty"`

	// ResourceId: The device resource ID.
	ResourceId string `json:"resourceId,omitempty"`

	// SerialNumber: The serial number of the device.
	SerialNumber string `json:"serialNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuspiciousActivitySecurityDetail) MarshalJSON() ([]byte, error) {
	type NoMethod SuspiciousActivitySecurityDetail
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "alertcenter.alerts.delete":

type AlertsDeleteCall struct {
	s          *Service
	alertId    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Marks the specified alert for deletion. An alert that has
// been marked for
// deletion is removed from Alert Center after 30 days.
// Marking an alert for deletion has no effect on an alert which
// has
// already been marked for deletion. Attempting to mark a nonexistent
// alert
// for deletion results in a `NOT_FOUND` error.
func (r *AlertsService) Delete(alertId string) *AlertsDeleteCall {
	c := &AlertsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.alertId = alertId
	return c
}

// CustomerId sets the optional parameter "customerId": The unique
// identifier of the G Suite organization account of the
// customer the alert is associated with.
// Inferred from the caller identity if not provided.
func (c *AlertsDeleteCall) CustomerId(customerId string) *AlertsDeleteCall {
	c.urlParams_.Set("customerId", customerId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlertsDeleteCall) Fields(s ...googleapi.Field) *AlertsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlertsDeleteCall) Context(ctx context.Context) *AlertsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlertsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlertsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/alerts/{alertId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"alertId": c.alertId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "alertcenter.alerts.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AlertsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Marks the specified alert for deletion. An alert that has been marked for\ndeletion is removed from Alert Center after 30 days.\nMarking an alert for deletion has no effect on an alert which has\nalready been marked for deletion. Attempting to mark a nonexistent alert\nfor deletion results in a `NOT_FOUND` error.",
	//   "flatPath": "v1beta1/alerts/{alertId}",
	//   "httpMethod": "DELETE",
	//   "id": "alertcenter.alerts.delete",
	//   "parameterOrder": [
	//     "alertId"
	//   ],
	//   "parameters": {
	//     "alertId": {
	//       "description": "Required. The identifier of the alert to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Optional. The unique identifier of the G Suite organization account of the\ncustomer the alert is associated with.\nInferred from the caller identity if not provided.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/alerts/{alertId}",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "alertcenter.alerts.get":

type AlertsGetCall struct {
	s            *Service
	alertId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the specified alert.
func (r *AlertsService) Get(alertId string) *AlertsGetCall {
	c := &AlertsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.alertId = alertId
	return c
}

// CustomerId sets the optional parameter "customerId": The unique
// identifier of the G Suite organization account of the
// customer the alert is associated with.
// Inferred from the caller identity if not provided.
func (c *AlertsGetCall) CustomerId(customerId string) *AlertsGetCall {
	c.urlParams_.Set("customerId", customerId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlertsGetCall) Fields(s ...googleapi.Field) *AlertsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AlertsGetCall) IfNoneMatch(entityTag string) *AlertsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlertsGetCall) Context(ctx context.Context) *AlertsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlertsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlertsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/alerts/{alertId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"alertId": c.alertId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "alertcenter.alerts.get" call.
// Exactly one of *Alert or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Alert.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AlertsGetCall) Do(opts ...googleapi.CallOption) (*Alert, error) {
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
	ret := &Alert{
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
	//   "description": "Gets the specified alert.",
	//   "flatPath": "v1beta1/alerts/{alertId}",
	//   "httpMethod": "GET",
	//   "id": "alertcenter.alerts.get",
	//   "parameterOrder": [
	//     "alertId"
	//   ],
	//   "parameters": {
	//     "alertId": {
	//       "description": "Required. The identifier of the alert to retrieve.\nReturns a NOT_FOUND error if no such alert.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Optional. The unique identifier of the G Suite organization account of the\ncustomer the alert is associated with.\nInferred from the caller identity if not provided.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/alerts/{alertId}",
	//   "response": {
	//     "$ref": "Alert"
	//   }
	// }

}

// method id "alertcenter.alerts.list":

type AlertsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the alerts.
func (r *AlertsService) List() *AlertsListCall {
	c := &AlertsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// CustomerId sets the optional parameter "customerId": The unique
// identifier of the G Suite organization account of the
// customer the alerts are associated with.
// Inferred from the caller identity if not provided.
func (c *AlertsListCall) CustomerId(customerId string) *AlertsListCall {
	c.urlParams_.Set("customerId", customerId)
	return c
}

// Filter sets the optional parameter "filter": A query string for
// filtering alert results.
// For more details, see
// [Query
// filters](/admin-sdk/alertcenter/guides/query-filters) and
// [Supported
// query filter
// fields](/admin-sdk/alertcenter/reference/filter-fields#alerts.list).
func (c *AlertsListCall) Filter(filter string) *AlertsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": The sort order of the
// list results.
// If not specified results may be returned in arbitrary order.
// You can sort the results in descending order based on the
// creation
// timestamp using `order_by="create_time desc".
// Currently, only sorting by `create_time desc` is supported.
func (c *AlertsListCall) OrderBy(orderBy string) *AlertsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The requested page
// size. Server may return fewer items than
// requested. If unspecified, server picks an appropriate default.
func (c *AlertsListCall) PageSize(pageSize int64) *AlertsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// If empty, a new iteration is started. To continue an iteration, pass
// in
// the value from the previous ListAlertsResponse's
// next_page_token field.
func (c *AlertsListCall) PageToken(pageToken string) *AlertsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlertsListCall) Fields(s ...googleapi.Field) *AlertsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AlertsListCall) IfNoneMatch(entityTag string) *AlertsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlertsListCall) Context(ctx context.Context) *AlertsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlertsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlertsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/alerts")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "alertcenter.alerts.list" call.
// Exactly one of *ListAlertsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListAlertsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlertsListCall) Do(opts ...googleapi.CallOption) (*ListAlertsResponse, error) {
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
	ret := &ListAlertsResponse{
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
	//   "description": "Lists the alerts.",
	//   "flatPath": "v1beta1/alerts",
	//   "httpMethod": "GET",
	//   "id": "alertcenter.alerts.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Optional. The unique identifier of the G Suite organization account of the\ncustomer the alerts are associated with.\nInferred from the caller identity if not provided.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Optional. A query string for filtering alert results.\nFor more details, see [Query\nfilters](/admin-sdk/alertcenter/guides/query-filters) and [Supported\nquery filter fields](/admin-sdk/alertcenter/reference/filter-fields#alerts.list).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Optional. The sort order of the list results.\nIf not specified results may be returned in arbitrary order.\nYou can sort the results in descending order based on the creation\ntimestamp using `order_by=\"create_time desc\"`.\nCurrently, only sorting by `create_time desc` is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The requested page size. Server may return fewer items than\nrequested. If unspecified, server picks an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A token identifying a page of results the server should return.\nIf empty, a new iteration is started. To continue an iteration, pass in\nthe value from the previous ListAlertsResponse's\nnext_page_token field.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/alerts",
	//   "response": {
	//     "$ref": "ListAlertsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AlertsListCall) Pages(ctx context.Context, f func(*ListAlertsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "alertcenter.alerts.feedback.create":

type AlertsFeedbackCreateCall struct {
	s             *Service
	alertId       string
	alertfeedback *AlertFeedback
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Create: Creates new feedback for an alert.
func (r *AlertsFeedbackService) Create(alertId string, alertfeedback *AlertFeedback) *AlertsFeedbackCreateCall {
	c := &AlertsFeedbackCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.alertId = alertId
	c.alertfeedback = alertfeedback
	return c
}

// CustomerId sets the optional parameter "customerId": The unique
// identifier of the G Suite organization account of the
// customer the alert is associated with.
// Inferred from the caller identity if not provided.
func (c *AlertsFeedbackCreateCall) CustomerId(customerId string) *AlertsFeedbackCreateCall {
	c.urlParams_.Set("customerId", customerId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlertsFeedbackCreateCall) Fields(s ...googleapi.Field) *AlertsFeedbackCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlertsFeedbackCreateCall) Context(ctx context.Context) *AlertsFeedbackCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlertsFeedbackCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlertsFeedbackCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.alertfeedback)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/alerts/{alertId}/feedback")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"alertId": c.alertId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "alertcenter.alerts.feedback.create" call.
// Exactly one of *AlertFeedback or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AlertFeedback.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlertsFeedbackCreateCall) Do(opts ...googleapi.CallOption) (*AlertFeedback, error) {
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
	ret := &AlertFeedback{
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
	//   "description": "Creates new feedback for an alert.",
	//   "flatPath": "v1beta1/alerts/{alertId}/feedback",
	//   "httpMethod": "POST",
	//   "id": "alertcenter.alerts.feedback.create",
	//   "parameterOrder": [
	//     "alertId"
	//   ],
	//   "parameters": {
	//     "alertId": {
	//       "description": "Required. The identifier of the alert this feedback belongs to.\nReturns a `NOT_FOUND` error if no such alert.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Optional. The unique identifier of the G Suite organization account of the\ncustomer the alert is associated with.\nInferred from the caller identity if not provided.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/alerts/{alertId}/feedback",
	//   "request": {
	//     "$ref": "AlertFeedback"
	//   },
	//   "response": {
	//     "$ref": "AlertFeedback"
	//   }
	// }

}

// method id "alertcenter.alerts.feedback.list":

type AlertsFeedbackListCall struct {
	s            *Service
	alertId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all the feedback for an alert.
func (r *AlertsFeedbackService) List(alertId string) *AlertsFeedbackListCall {
	c := &AlertsFeedbackListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.alertId = alertId
	return c
}

// CustomerId sets the optional parameter "customerId": The unique
// identifier of the G Suite organization account of the
// customer the alert feedback are associated with.
// Inferred from the caller identity if not provided.
func (c *AlertsFeedbackListCall) CustomerId(customerId string) *AlertsFeedbackListCall {
	c.urlParams_.Set("customerId", customerId)
	return c
}

// Filter sets the optional parameter "filter": A query string for
// filtering alert feedback results.
// For more details, see
// [Query
// filters](/admin-sdk/alertcenter/guides/query-filters) and
// [Supported
// query filter
// fields](/admin-sdk/alertcenter/reference/filter-fields#alerts.feedback
// .list).
func (c *AlertsFeedbackListCall) Filter(filter string) *AlertsFeedbackListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlertsFeedbackListCall) Fields(s ...googleapi.Field) *AlertsFeedbackListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AlertsFeedbackListCall) IfNoneMatch(entityTag string) *AlertsFeedbackListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlertsFeedbackListCall) Context(ctx context.Context) *AlertsFeedbackListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlertsFeedbackListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlertsFeedbackListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/alerts/{alertId}/feedback")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"alertId": c.alertId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "alertcenter.alerts.feedback.list" call.
// Exactly one of *ListAlertFeedbackResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListAlertFeedbackResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlertsFeedbackListCall) Do(opts ...googleapi.CallOption) (*ListAlertFeedbackResponse, error) {
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
	ret := &ListAlertFeedbackResponse{
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
	//   "description": "Lists all the feedback for an alert.",
	//   "flatPath": "v1beta1/alerts/{alertId}/feedback",
	//   "httpMethod": "GET",
	//   "id": "alertcenter.alerts.feedback.list",
	//   "parameterOrder": [
	//     "alertId"
	//   ],
	//   "parameters": {
	//     "alertId": {
	//       "description": "Required. The alert identifier.\nThe \"-\" wildcard could be used to represent all alerts.\nIf alert does not exist returns a `NOT_FOUND` error.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Optional. The unique identifier of the G Suite organization account of the\ncustomer the alert feedback are associated with.\nInferred from the caller identity if not provided.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Optional. A query string for filtering alert feedback results.\nFor more details, see [Query\nfilters](/admin-sdk/alertcenter/guides/query-filters) and [Supported\nquery filter fields](/admin-sdk/alertcenter/reference/filter-fields#alerts.feedback.list).",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/alerts/{alertId}/feedback",
	//   "response": {
	//     "$ref": "ListAlertFeedbackResponse"
	//   }
	// }

}
