// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package driveactivity provides access to the Drive Activity API.
//
// See https://developers.google.com/drive/activity/
//
// Usage example:
//
//   import "google.golang.org/api/driveactivity/v2"
//   ...
//   driveactivityService, err := driveactivity.New(oauthHttpClient)
package driveactivity // import "google.golang.org/api/driveactivity/v2"

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

const apiId = "driveactivity:v2"
const apiName = "driveactivity"
const apiVersion = "v2"
const basePath = "https://driveactivity.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and add to the activity record of files in your Google Drive
	DriveActivityScope = "https://www.googleapis.com/auth/drive.activity"

	// View the activity record of files in your Google Drive
	DriveActivityReadonlyScope = "https://www.googleapis.com/auth/drive.activity.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Activity = NewActivityService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Activity *ActivityService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewActivityService(s *Service) *ActivityService {
	rs := &ActivityService{s: s}
	return rs
}

type ActivityService struct {
	s *Service
}

// Action: Information about the action.
type Action struct {
	// Actor: The actor responsible for this action (or empty if all actors
	// are
	// responsible).
	Actor *Actor `json:"actor,omitempty"`

	// Detail: The type and detailed information about the action.
	Detail *ActionDetail `json:"detail,omitempty"`

	// Target: The target this action affects (or empty if affecting all
	// targets). This
	// represents the state of the target immediately after this action
	// occurred.
	Target *Target `json:"target,omitempty"`

	// TimeRange: The action occurred over this time range.
	TimeRange *TimeRange `json:"timeRange,omitempty"`

	// Timestamp: The action occurred at this specific time.
	Timestamp string `json:"timestamp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Actor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Actor") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Action) MarshalJSON() ([]byte, error) {
	type NoMethod Action
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ActionDetail: Data describing the type and additional information of
// an action.
type ActionDetail struct {
	// Comment: A change about comments was made.
	Comment *Comment `json:"comment,omitempty"`

	// Create: An object was created.
	Create *Create `json:"create,omitempty"`

	// Delete: An object was deleted.
	Delete *Delete `json:"delete,omitempty"`

	// DlpChange: A change happened in data leak prevention status.
	DlpChange *DataLeakPreventionChange `json:"dlpChange,omitempty"`

	// Edit: An object was edited.
	Edit *Edit `json:"edit,omitempty"`

	// Move: An object was moved.
	Move *Move `json:"move,omitempty"`

	// PermissionChange: The permission on an object was changed.
	PermissionChange *PermissionChange `json:"permissionChange,omitempty"`

	// Reference: An object was referenced in an application outside of
	// Drive/Docs.
	Reference *ApplicationReference `json:"reference,omitempty"`

	// Rename: An object was renamed.
	Rename *Rename `json:"rename,omitempty"`

	// Restore: A deleted object was restored.
	Restore *Restore `json:"restore,omitempty"`

	// SettingsChange: Settings were changed.
	SettingsChange *SettingsChange `json:"settingsChange,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Comment") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Comment") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ActionDetail) MarshalJSON() ([]byte, error) {
	type NoMethod ActionDetail
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Actor: The actor of a Drive activity.
type Actor struct {
	// Administrator: An administrator.
	Administrator *Administrator `json:"administrator,omitempty"`

	// Anonymous: An anonymous user.
	Anonymous *AnonymousUser `json:"anonymous,omitempty"`

	// Impersonation: An account acting on behalf of another.
	Impersonation *Impersonation `json:"impersonation,omitempty"`

	// System: A non-user actor (i.e. system triggered).
	System *SystemEvent `json:"system,omitempty"`

	// User: An end user.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Administrator") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Administrator") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Actor) MarshalJSON() ([]byte, error) {
	type NoMethod Actor
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Administrator: Empty message representing an administrator.
type Administrator struct {
}

// AnonymousUser: Empty message representing an anonymous user or
// indicating the authenticated
// user should be anonymized.
type AnonymousUser struct {
}

// Anyone: Represents any user (including a logged out user).
type Anyone struct {
}

// ApplicationReference: Activity in applications other than Drive.
type ApplicationReference struct {
	// Type: The reference type corresponding to this event.
	//
	// Possible values:
	//   "UNSPECIFIED_REFERENCE_TYPE" - The type is not available.
	//   "LINK" - The links of one or more Drive items were posted.
	//   "DISCUSS" - Comments were made regarding a Drive item.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApplicationReference) MarshalJSON() ([]byte, error) {
	type NoMethod ApplicationReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Assignment: A comment with an assignment.
type Assignment struct {
	// Subtype: The sub-type of this event.
	//
	// Possible values:
	//   "SUBTYPE_UNSPECIFIED" - Subtype not available.
	//   "ADDED" - An assignment was added.
	//   "DELETED" - An assignment was deleted.
	//   "REPLY_ADDED" - An assignment reply was added.
	//   "REPLY_DELETED" - An assignment reply was deleted.
	//   "RESOLVED" - An assignment was resolved.
	//   "REOPENED" - A resolved assignment was reopened.
	//   "REASSIGNED" - An assignment was reassigned.
	Subtype string `json:"subtype,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Subtype") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Subtype") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Assignment) MarshalJSON() ([]byte, error) {
	type NoMethod Assignment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Comment: A change about comments on an object.
type Comment struct {
	// Assignment: A change on an assignment.
	Assignment *Assignment `json:"assignment,omitempty"`

	// MentionedUsers: Users who are mentioned in this comment.
	MentionedUsers []*User `json:"mentionedUsers,omitempty"`

	// Post: A change on a regular posted comment.
	Post *Post `json:"post,omitempty"`

	// Suggestion: A change on a suggestion.
	Suggestion *Suggestion `json:"suggestion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Assignment") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Assignment") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Comment) MarshalJSON() ([]byte, error) {
	type NoMethod Comment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ConsolidationStrategy: How the individual activities are
// consolidated. A set of activities may be
// consolidated into one combined activity if they are related in some
// way, such
// as one actor performing the same action on multiple targets, or
// multiple
// actors performing the same action on a single target. The strategy
// defines
// the rules for which activities are related.
type ConsolidationStrategy struct {
	// Legacy: The individual activities are consolidated using the legacy
	// strategy.
	Legacy *Legacy `json:"legacy,omitempty"`

	// None: The individual activities are not consolidated.
	None *NoConsolidation `json:"none,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Legacy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Legacy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ConsolidationStrategy) MarshalJSON() ([]byte, error) {
	type NoMethod ConsolidationStrategy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Copy: An object was created by copying an existing object.
type Copy struct {
	// OriginalObject: The the original object.
	OriginalObject *TargetReference `json:"originalObject,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OriginalObject") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OriginalObject") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Copy) MarshalJSON() ([]byte, error) {
	type NoMethod Copy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Create: An object was created.
type Create struct {
	// Copy: If present, indicates the object was created by copying an
	// existing Drive
	// object.
	Copy *Copy `json:"copy,omitempty"`

	// New: If present, indicates the object was newly created (e.g. as a
	// blank
	// document), not derived from a Drive object or external object.
	New *New1 `json:"new,omitempty"`

	// Upload: If present, indicates the object originated externally and
	// was uploaded
	// to Drive.
	Upload *Upload `json:"upload,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Copy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Copy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Create) MarshalJSON() ([]byte, error) {
	type NoMethod Create
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DataLeakPreventionChange: A change in the object's data leak
// prevention status.
type DataLeakPreventionChange struct {
	// Type: The type of Data Leak Prevention (DLP) change.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - An update to the DLP state that is neither
	// FLAGGED or CLEARED.
	//   "FLAGGED" - Document has been flagged as containing sensitive
	// content.
	//   "CLEARED" - Document is no longer flagged as containing sensitive
	// content.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DataLeakPreventionChange) MarshalJSON() ([]byte, error) {
	type NoMethod DataLeakPreventionChange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Delete: An object was deleted.
type Delete struct {
	// Type: The type of delete action taken.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Deletion type is not available.
	//   "TRASH" - An object was put into the trash.
	//   "PERMANENT_DELETE" - An object was deleted permanently.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Delete) MarshalJSON() ([]byte, error) {
	type NoMethod Delete
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeletedUser: A user whose account has since been deleted.
type DeletedUser struct {
}

// Domain: Information about a domain.
type Domain struct {
	// LegacyId: An opaque string used to identify this domain.
	LegacyId string `json:"legacyId,omitempty"`

	// Name: The name of the domain, e.g. "google.com".
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LegacyId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LegacyId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Domain) MarshalJSON() ([]byte, error) {
	type NoMethod Domain
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveActivity: A single Drive activity comprising one or more Actions
// by one or more
// Actors on one or more Targets. Some Action groupings occur
// spontaneously,
// such as moving an item into a shared folder triggering a permission
// change.
// Other groupings of related Actions, such as multiple Actors editing
// one item
// or moving multiple files into a new folder, are controlled by the
// selection
// of a ConsolidationStrategy in the QueryDriveActivityRequest.
type DriveActivity struct {
	// Actions: Details on all actions in this activity.
	Actions []*Action `json:"actions,omitempty"`

	// Actors: All actor(s) responsible for the activity.
	Actors []*Actor `json:"actors,omitempty"`

	// PrimaryActionDetail: Key information about the primary action for
	// this activity. This is either
	// representative, or the most important, of all actions in the
	// activity,
	// according to the ConsolidationStrategy in the request.
	PrimaryActionDetail *ActionDetail `json:"primaryActionDetail,omitempty"`

	// Targets: All Drive objects this activity is about (e.g. file, folder,
	// Team Drive).
	// This represents the state of the target immediately after the
	// actions
	// occurred.
	Targets []*Target `json:"targets,omitempty"`

	// TimeRange: The activity occurred over this time range.
	TimeRange *TimeRange `json:"timeRange,omitempty"`

	// Timestamp: The activity occurred at this specific time.
	Timestamp string `json:"timestamp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Actions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Actions") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveActivity) MarshalJSON() ([]byte, error) {
	type NoMethod DriveActivity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveItem: A Drive item, such as a file or folder.
type DriveItem struct {
	// File: The Drive item is a file.
	File *File `json:"file,omitempty"`

	// Folder: The Drive item is a folder.
	Folder *Folder `json:"folder,omitempty"`

	// MimeType: The MIME type of the Drive item.
	// See
	// https://developers.google.com/drive/v3/web/mime-types.
	MimeType string `json:"mimeType,omitempty"`

	// Name: The target Drive item. The format is "items/ITEM_ID".
	Name string `json:"name,omitempty"`

	// Owner: Information about the owner of this Drive item.
	Owner *Owner `json:"owner,omitempty"`

	// Title: The title of the Drive item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "File") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "File") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveItem) MarshalJSON() ([]byte, error) {
	type NoMethod DriveItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveItemReference: A lightweight reference to a Drive item, such as
// a file or folder.
type DriveItemReference struct {
	// File: The Drive item is a file.
	File *File `json:"file,omitempty"`

	// Folder: The Drive item is a folder.
	Folder *Folder `json:"folder,omitempty"`

	// Name: The target Drive item. The format is "items/ITEM_ID".
	Name string `json:"name,omitempty"`

	// Title: The title of the Drive item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "File") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "File") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveItemReference) MarshalJSON() ([]byte, error) {
	type NoMethod DriveItemReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Edit: An empty message indicating an object was edited.
type Edit struct {
}

// File: A Drive item which is a file.
type File struct {
}

// FileComment: A comment on a file.
type FileComment struct {
	// LegacyCommentId: The comment in the discussion thread. This
	// identifier is an opaque string
	// compatible with the Drive API;
	// see
	// https://developers.google.com/drive/v3/reference/comments/get
	LegacyCommentId string `json:"legacyCommentId,omitempty"`

	// LegacyDiscussionId: The discussion thread to which the comment was
	// added. This identifier is an
	// opaque string compatible with the Drive API and references the
	// first
	// comment in a discussion;
	// see
	// https://developers.google.com/drive/v3/reference/comments/get
	LegacyDiscussionId string `json:"legacyDiscussionId,omitempty"`

	// LinkToDiscussion: The link to the discussion thread containing this
	// comment, for
	// example,
	// "https://docs.google.com/DOCUMENT_ID/edit?disco=THREAD_ID".
	LinkToDiscussion string `json:"linkToDiscussion,omitempty"`

	// Parent: The Drive item containing this comment.
	Parent *DriveItem `json:"parent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LegacyCommentId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LegacyCommentId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FileComment) MarshalJSON() ([]byte, error) {
	type NoMethod FileComment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Folder: A Drive item which is a folder.
type Folder struct {
	// Type: The type of Drive folder.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The folder type is unknown.
	//   "MY_DRIVE_ROOT" - The folder is the root of a user's MyDrive.
	//   "TEAM_DRIVE_ROOT" - The folder is the root of a Team Drive. Note
	// that this folder is
	// a Drive item, and is a distinct entity from the Team Drive itself.
	//   "STANDARD_FOLDER" - The folder is a standard, non-root, folder.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Folder) MarshalJSON() ([]byte, error) {
	type NoMethod Folder
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Group: Information about a group.
type Group struct {
	// Email: The email address of the group.
	Email string `json:"email,omitempty"`

	// Title: The title of the group.
	Title string `json:"title,omitempty"`

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

func (s *Group) MarshalJSON() ([]byte, error) {
	type NoMethod Group
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Impersonation: Information about an impersonation, where an admin
// acts on behalf of an end
// user. Information about the acting admin is not currently available.
type Impersonation struct {
	// ImpersonatedUser: The impersonated user.
	ImpersonatedUser *User `json:"impersonatedUser,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImpersonatedUser") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImpersonatedUser") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Impersonation) MarshalJSON() ([]byte, error) {
	type NoMethod Impersonation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// KnownUser: A known user.
type KnownUser struct {
	// IsCurrentUser: True if this is the user making the request.
	IsCurrentUser bool `json:"isCurrentUser,omitempty"`

	// PersonName: The identifier for this user that can be used with the
	// People API to get
	// more information. The format is "people/ACCOUNT_ID".
	// See
	// https://developers.google.com/people/.
	PersonName string `json:"personName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IsCurrentUser") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IsCurrentUser") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *KnownUser) MarshalJSON() ([]byte, error) {
	type NoMethod KnownUser
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Legacy: A strategy which consolidates activities using the grouping
// rules from the
// legacy V1 Activity API. Similar actions occurring within a window of
// time
// can be grouped across multiple targets (such as moving a set of files
// at
// once) or multiple actors (such as several users editing the same
// item).
// Grouping rules for this strategy are specific to each type of action.
type Legacy struct {
}

// Move: An object was moved.
type Move struct {
	// AddedParents: The added parent object(s).
	AddedParents []*TargetReference `json:"addedParents,omitempty"`

	// RemovedParents: The removed parent object(s).
	RemovedParents []*TargetReference `json:"removedParents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddedParents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddedParents") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Move) MarshalJSON() ([]byte, error) {
	type NoMethod Move
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// New1: An object was created from scratch.
type New1 struct {
}

// NoConsolidation: A strategy which does no consolidation of individual
// activities.
type NoConsolidation struct {
}

// Owner: Information about the owner of a Drive item.
type Owner struct {
	// Domain: The domain of the Drive item owner.
	Domain *Domain `json:"domain,omitempty"`

	// TeamDrive: The Team Drive that owns the Drive item.
	TeamDrive *TeamDriveReference `json:"teamDrive,omitempty"`

	// User: The user that owns the Drive item.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Domain") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Domain") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Owner) MarshalJSON() ([]byte, error) {
	type NoMethod Owner
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Permission: The permission setting of an object.
type Permission struct {
	// AllowDiscovery: If true, the item can be discovered (e.g. in the
	// user's "Shared with me"
	// collection) without needing a link to the item.
	AllowDiscovery bool `json:"allowDiscovery,omitempty"`

	// Anyone: If set, this permission applies to anyone, even logged out
	// users.
	Anyone *Anyone `json:"anyone,omitempty"`

	// Domain: The domain to whom this permission applies.
	Domain *Domain `json:"domain,omitempty"`

	// Group: The group to whom this permission applies.
	Group *Group `json:"group,omitempty"`

	// Role: Indicates the
	// <a href="/drive/web/manage-sharing#roles">Google Drive
	// permissions
	// role</a>. The role determines a user's ability to read, write,
	// and
	// comment on items.
	//
	// Possible values:
	//   "ROLE_UNSPECIFIED" - The role is not available.
	//   "OWNER" - A role granting full access.
	//   "ORGANIZER" - A role granting the ability to manage people and
	// settings.
	//   "FILE_ORGANIZER" - A role granting the ability to contribute and
	// manage content.
	//   "EDITOR" - A role granting the ability to contribute content. This
	// role is sometimes
	// also known as "writer".
	//   "COMMENTER" - A role granting the ability to view and comment on
	// content.
	//   "VIEWER" - A role granting the ability to view content. This role
	// is sometimes also
	// known as "reader".
	//   "PUBLISHED_VIEWER" - A role granting the ability to view content
	// only after it has been
	// published to the web. This role is sometimes also known as
	// "published
	// reader". See https://support.google.com/sites/answer/6372880 for
	// more
	// information.
	Role string `json:"role,omitempty"`

	// User: The user to whom this permission applies.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowDiscovery") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllowDiscovery") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Permission) MarshalJSON() ([]byte, error) {
	type NoMethod Permission
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PermissionChange: A change of the permission setting on an item.
type PermissionChange struct {
	// AddedPermissions: The set of permissions added by this change.
	AddedPermissions []*Permission `json:"addedPermissions,omitempty"`

	// RemovedPermissions: The set of permissions removed by this change.
	RemovedPermissions []*Permission `json:"removedPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddedPermissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddedPermissions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PermissionChange) MarshalJSON() ([]byte, error) {
	type NoMethod PermissionChange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Post: A regular posted comment.
type Post struct {
	// Subtype: The sub-type of this event.
	//
	// Possible values:
	//   "SUBTYPE_UNSPECIFIED" - Subtype not available.
	//   "ADDED" - A post was added.
	//   "DELETED" - A post was deleted.
	//   "REPLY_ADDED" - A reply was added.
	//   "REPLY_DELETED" - A reply was deleted.
	//   "RESOLVED" - A posted comment was resolved.
	//   "REOPENED" - A posted comment was reopened.
	Subtype string `json:"subtype,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Subtype") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Subtype") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Post) MarshalJSON() ([]byte, error) {
	type NoMethod Post
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryDriveActivityRequest: The request message for querying Drive
// activity.
type QueryDriveActivityRequest struct {
	// AncestorName: Return activities for this Drive folder and all
	// children and descendants.
	// The format is "items/ITEM_ID".
	AncestorName string `json:"ancestorName,omitempty"`

	// ConsolidationStrategy: Details on how to consolidate related actions
	// that make up the activity. If
	// not set, then related actions will not be consolidated.
	ConsolidationStrategy *ConsolidationStrategy `json:"consolidationStrategy,omitempty"`

	// Filter: The filtering for items returned from this query request. The
	// format of the
	// filter string is a sequence of expressions, joined by an optional
	// "AND",
	// where each expression is of the form "field operator
	// value".
	//
	// Supported fields:
	//
	//   - <tt>time</tt>: Uses numerical operators on date values either in
	//     terms of milliseconds since Jan 1, 1970 or in RFC 3339 format.
	//     Examples:
	//       - <tt>time > 1452409200000 AND time <= 1492812924310</tt>
	//       - <tt>time >= "2016-01-10T01:02:03-05:00"</tt>
	//
	//   - <tt>detail.action_detail_case</tt>: Uses the "has" operator (:)
	// and
	//     either a singular value or a list of allowed action types
	// enclosed in
	//     parentheses.
	//     Examples:
	//       - <tt>detail.action_detail_case: RENAME</tt>
	//       - <tt>detail.action_detail_case:(CREATE UPLOAD)</tt>
	//       - <tt>-detail.action_detail_case:MOVE</tt>
	Filter string `json:"filter,omitempty"`

	// ItemName: Return activities for this Drive item. The format
	// is
	// "items/ITEM_ID".
	ItemName string `json:"itemName,omitempty"`

	// PageSize: The requested number of activity to return. If not set, a
	// default value
	// will be used.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The next_page_token value returned from a previous
	// QueryDriveActivity
	// request, if any.
	PageToken string `json:"pageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AncestorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AncestorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryDriveActivityRequest) MarshalJSON() ([]byte, error) {
	type NoMethod QueryDriveActivityRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryDriveActivityResponse: Response message for querying Drive
// activity.
type QueryDriveActivityResponse struct {
	// Activities: List of activity requested.
	Activities []*DriveActivity `json:"activities,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or
	// empty if there are no more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Activities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Activities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryDriveActivityResponse) MarshalJSON() ([]byte, error) {
	type NoMethod QueryDriveActivityResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Rename: An object was renamed.
type Rename struct {
	// NewTitle: The new title of the drive object.
	NewTitle string `json:"newTitle,omitempty"`

	// OldTitle: The previous title of the drive object.
	OldTitle string `json:"oldTitle,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NewTitle") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NewTitle") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Rename) MarshalJSON() ([]byte, error) {
	type NoMethod Rename
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Restore: A deleted object was restored.
type Restore struct {
	// Type: The type of restore action taken.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The type is not available.
	//   "UNTRASH" - An object was restored from the trash.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Restore) MarshalJSON() ([]byte, error) {
	type NoMethod Restore
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RestrictionChange: Information about restriction policy changes to a
// feature.
type RestrictionChange struct {
	// Feature: The feature which had a change in restriction policy.
	//
	// Possible values:
	//   "FEATURE_UNSPECIFIED" - The feature which changed restriction
	// settings was not available.
	//   "SHARING_OUTSIDE_DOMAIN" - When restricted, this prevents items
	// from being shared outside the
	// domain.
	//   "DIRECT_SHARING" - When restricted, this prevents direct sharing of
	// individual items.
	//   "ITEM_DUPLICATION" - When restricted, this prevents actions like
	// copy, download, and print
	// that might result in uncontrolled duplicates of items.
	//   "DRIVE_FILE_STREAM" - When restricted, this prevents use of Drive
	// File Stream.
	Feature string `json:"feature,omitempty"`

	// NewRestriction: The restriction in place after the change.
	//
	// Possible values:
	//   "RESTRICTION_UNSPECIFIED" - The type of restriction is not
	// available.
	//   "UNRESTRICTED" - The feature is available without restriction.
	//   "FULLY_RESTRICTED" - The use of this feature is fully restricted.
	NewRestriction string `json:"newRestriction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Feature") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Feature") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RestrictionChange) MarshalJSON() ([]byte, error) {
	type NoMethod RestrictionChange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SettingsChange: Information about settings changes.
type SettingsChange struct {
	// RestrictionChanges: The set of changes made to restrictions.
	RestrictionChanges []*RestrictionChange `json:"restrictionChanges,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RestrictionChanges")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RestrictionChanges") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SettingsChange) MarshalJSON() ([]byte, error) {
	type NoMethod SettingsChange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Suggestion: A suggestion.
type Suggestion struct {
	// Subtype: The sub-type of this event.
	//
	// Possible values:
	//   "SUBTYPE_UNSPECIFIED" - Subtype not available.
	//   "ADDED" - A suggestion was added.
	//   "DELETED" - A suggestion was deleted.
	//   "REPLY_ADDED" - A suggestion reply was added.
	//   "REPLY_DELETED" - A suggestion reply was deleted.
	//   "ACCEPTED" - A suggestion was accepted.
	//   "REJECTED" - A suggestion was rejected.
	//   "ACCEPT_DELETED" - An accepted suggestion was deleted.
	//   "REJECT_DELETED" - A rejected suggestion was deleted.
	Subtype string `json:"subtype,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Subtype") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Subtype") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Suggestion) MarshalJSON() ([]byte, error) {
	type NoMethod Suggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SystemEvent: Event triggered by system operations instead of end
// users.
type SystemEvent struct {
	// Type: The type of the system event that may triggered activity.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The event type is unspecified.
	//   "USER_DELETION" - The event is a consequence of a user account
	// being deleted.
	//   "TRASH_AUTO_PURGE" - The event is due to the system automatically
	// purging trash.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SystemEvent) MarshalJSON() ([]byte, error) {
	type NoMethod SystemEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Target: Information about the target of activity.
type Target struct {
	// DriveItem: The target is a Drive item.
	DriveItem *DriveItem `json:"driveItem,omitempty"`

	// FileComment: The target is a comment on a Drive file.
	FileComment *FileComment `json:"fileComment,omitempty"`

	// TeamDrive: The target is a Team Drive.
	TeamDrive *TeamDrive `json:"teamDrive,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DriveItem") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriveItem") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Target) MarshalJSON() ([]byte, error) {
	type NoMethod Target
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TargetReference: A lightweight reference to the target of activity.
type TargetReference struct {
	// DriveItem: The target is a Drive item.
	DriveItem *DriveItemReference `json:"driveItem,omitempty"`

	// TeamDrive: The target is a Team Drive.
	TeamDrive *TeamDriveReference `json:"teamDrive,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DriveItem") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriveItem") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TargetReference) MarshalJSON() ([]byte, error) {
	type NoMethod TargetReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TeamDrive: Information about a Team Drive.
type TeamDrive struct {
	// Name: The resource name of the Team Drive. The format
	// is
	// "teamDrives/TEAM_DRIVE_ID".
	Name string `json:"name,omitempty"`

	// Root: The root of this Team Drive.
	Root *DriveItem `json:"root,omitempty"`

	// Title: The title of the Team Drive.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TeamDrive) MarshalJSON() ([]byte, error) {
	type NoMethod TeamDrive
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TeamDriveReference: A lightweight reference to a Team Drive.
type TeamDriveReference struct {
	// Name: The resource name of the Team Drive. The format
	// is
	// "teamDrives/TEAM_DRIVE_ID".
	Name string `json:"name,omitempty"`

	// Title: The title of the Team Drive.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TeamDriveReference) MarshalJSON() ([]byte, error) {
	type NoMethod TeamDriveReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeRange: Information about time ranges.
type TimeRange struct {
	// EndTime: The end of the time range.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: The start of the time range.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeRange) MarshalJSON() ([]byte, error) {
	type NoMethod TimeRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UnknownUser: A user about whom nothing is currently known.
type UnknownUser struct {
}

// Upload: An object was uploaded into Drive.
type Upload struct {
}

// User: Information about an end user.
type User struct {
	// DeletedUser: A user whose account has since been deleted.
	DeletedUser *DeletedUser `json:"deletedUser,omitempty"`

	// KnownUser: A known user.
	KnownUser *KnownUser `json:"knownUser,omitempty"`

	// UnknownUser: A user about whom nothing is currently known.
	UnknownUser *UnknownUser `json:"unknownUser,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeletedUser") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeletedUser") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *User) MarshalJSON() ([]byte, error) {
	type NoMethod User
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "driveactivity.activity.query":

type ActivityQueryCall struct {
	s                         *Service
	querydriveactivityrequest *QueryDriveActivityRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// Query: Query past activity in Google Drive.
func (r *ActivityService) Query(querydriveactivityrequest *QueryDriveActivityRequest) *ActivityQueryCall {
	c := &ActivityQueryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.querydriveactivityrequest = querydriveactivityrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ActivityQueryCall) Fields(s ...googleapi.Field) *ActivityQueryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ActivityQueryCall) Context(ctx context.Context) *ActivityQueryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ActivityQueryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ActivityQueryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.querydriveactivityrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/activity:query")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "driveactivity.activity.query" call.
// Exactly one of *QueryDriveActivityResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *QueryDriveActivityResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ActivityQueryCall) Do(opts ...googleapi.CallOption) (*QueryDriveActivityResponse, error) {
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
	ret := &QueryDriveActivityResponse{
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
	//   "description": "Query past activity in Google Drive.",
	//   "flatPath": "v2/activity:query",
	//   "httpMethod": "POST",
	//   "id": "driveactivity.activity.query",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/activity:query",
	//   "request": {
	//     "$ref": "QueryDriveActivityRequest"
	//   },
	//   "response": {
	//     "$ref": "QueryDriveActivityResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/drive.activity",
	//     "https://www.googleapis.com/auth/drive.activity.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ActivityQueryCall) Pages(ctx context.Context, f func(*QueryDriveActivityResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.querydriveactivityrequest.PageToken = pt }(c.querydriveactivityrequest.PageToken) // reset paging to original point
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
		c.querydriveactivityrequest.PageToken = x.NextPageToken
	}
}
