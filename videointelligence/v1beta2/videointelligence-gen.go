// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package videointelligence provides access to the Cloud Video Intelligence API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/videointelligence/apiv1 instead.
//
// See https://cloud.google.com/video-intelligence/docs/
//
// Usage example:
//
//   import "google.golang.org/api/videointelligence/v1beta2"
//   ...
//   videointelligenceService, err := videointelligence.New(oauthHttpClient)
package videointelligence // import "google.golang.org/api/videointelligence/v1beta2"

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

const apiId = "videointelligence:v1beta2"
const apiName = "videointelligence"
const apiVersion = "v1beta2"
const basePath = "https://videointelligence.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Videos = NewVideosService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Videos *VideosService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewVideosService(s *Service) *VideosService {
	rs := &VideosService{s: s}
	return rs
}

type VideosService struct {
	s *Service
}

// GoogleCloudVideointelligenceV1AnnotateVideoProgress: Video annotation
// progress. Included in the `metadata`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1AnnotateVideoProgress struct {
	// AnnotationProgress: Progress metadata for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationProgress []*GoogleCloudVideointelligenceV1VideoAnnotationProgress `json:"annotationProgress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationProgress")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationProgress") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1AnnotateVideoProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1AnnotateVideoProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1AnnotateVideoResponse: Video annotation
// response. Included in the `response`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1AnnotateVideoResponse struct {
	// AnnotationResults: Annotation results for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationResults []*GoogleCloudVideointelligenceV1VideoAnnotationResults `json:"annotationResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationResults")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1AnnotateVideoResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1AnnotateVideoResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1Entity: Detected entity from video
// analysis.
type GoogleCloudVideointelligenceV1Entity struct {
	// Description: Textual description, e.g. `Fixed-gear bicycle`.
	Description string `json:"description,omitempty"`

	// EntityId: Opaque entity ID. Some IDs may be available in
	// [Google Knowledge Graph
	// Search
	// API](https://developers.google.com/knowledge-graph/).
	EntityId string `json:"entityId,omitempty"`

	// LanguageCode: Language code for `description` in BCP-47 format.
	LanguageCode string `json:"languageCode,omitempty"`

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

func (s *GoogleCloudVideointelligenceV1Entity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1Entity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1ExplicitContentAnnotation: Explicit
// content annotation (based on per-frame visual signals only).
// If no explicit content has been detected in a frame, no annotations
// are
// present for that frame.
type GoogleCloudVideointelligenceV1ExplicitContentAnnotation struct {
	// Frames: All video frames where explicit content was detected.
	Frames []*GoogleCloudVideointelligenceV1ExplicitContentFrame `json:"frames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Frames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Frames") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1ExplicitContentAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1ExplicitContentAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1ExplicitContentFrame: Video frame level
// annotation results for explicit content.
type GoogleCloudVideointelligenceV1ExplicitContentFrame struct {
	// PornographyLikelihood: Likelihood of the pornography content..
	//
	// Possible values:
	//   "LIKELIHOOD_UNSPECIFIED" - Unspecified likelihood.
	//   "VERY_UNLIKELY" - Very unlikely.
	//   "UNLIKELY" - Unlikely.
	//   "POSSIBLE" - Possible.
	//   "LIKELY" - Likely.
	//   "VERY_LIKELY" - Very likely.
	PornographyLikelihood string `json:"pornographyLikelihood,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "PornographyLikelihood") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PornographyLikelihood") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1ExplicitContentFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1ExplicitContentFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1LabelAnnotation: Label annotation.
type GoogleCloudVideointelligenceV1LabelAnnotation struct {
	// CategoryEntities: Common categories for the detected entity.
	// E.g. when the label is `Terrier` the category is likely `dog`. And in
	// some
	// cases there might be more than one categories e.g. `Terrier` could
	// also be
	// a `pet`.
	CategoryEntities []*GoogleCloudVideointelligenceV1Entity `json:"categoryEntities,omitempty"`

	// Entity: Detected entity.
	Entity *GoogleCloudVideointelligenceV1Entity `json:"entity,omitempty"`

	// Frames: All video frames where a label was detected.
	Frames []*GoogleCloudVideointelligenceV1LabelFrame `json:"frames,omitempty"`

	// Segments: All video segments where a label was detected.
	Segments []*GoogleCloudVideointelligenceV1LabelSegment `json:"segments,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CategoryEntities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CategoryEntities") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1LabelAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1LabelAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1LabelFrame: Video frame level
// annotation results for label detection.
type GoogleCloudVideointelligenceV1LabelFrame struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1LabelFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1LabelFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1LabelFrame) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1LabelFrame
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1LabelSegment: Video segment level
// annotation results for label detection.
type GoogleCloudVideointelligenceV1LabelSegment struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// Segment: Video segment where a label was detected.
	Segment *GoogleCloudVideointelligenceV1VideoSegment `json:"segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1LabelSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1LabelSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1LabelSegment) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1LabelSegment
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1SpeechRecognitionAlternative:
// Alternative hypotheses (a.k.a. n-best list).
type GoogleCloudVideointelligenceV1SpeechRecognitionAlternative struct {
	// Confidence: The confidence estimate between 0.0 and 1.0. A higher
	// number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is typically provided only for the top
	// hypothesis, and
	// only for `is_final=true` results. Clients should not rely on
	// the
	// `confidence` field as it is not guaranteed to be accurate or
	// consistent.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// Transcript: Transcript text representing the words that the user
	// spoke.
	Transcript string `json:"transcript,omitempty"`

	// Words: A list of word-specific information for each recognized word.
	Words []*GoogleCloudVideointelligenceV1WordInfo `json:"words,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1SpeechRecognitionAlternative) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1SpeechRecognitionAlternative
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1SpeechRecognitionAlternative) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1SpeechRecognitionAlternative
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1SpeechTranscription: A speech
// recognition result corresponding to a portion of the audio.
type GoogleCloudVideointelligenceV1SpeechTranscription struct {
	// Alternatives: May contain one or more recognition hypotheses (up to
	// the maximum specified
	// in `max_alternatives`).  These alternatives are ordered in terms
	// of
	// accuracy, with the top (first) alternative being the most probable,
	// as
	// ranked by the recognizer.
	Alternatives []*GoogleCloudVideointelligenceV1SpeechRecognitionAlternative `json:"alternatives,omitempty"`

	// LanguageCode: Output only.
	// The
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag
	// of the
	// language in this result. This language code was detected to have the
	// most
	// likelihood of being spoken in the audio.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alternatives") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alternatives") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1SpeechTranscription) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1SpeechTranscription
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1VideoAnnotationProgress: Annotation
// progress for a single video.
type GoogleCloudVideointelligenceV1VideoAnnotationProgress struct {
	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// ProgressPercent: Approximate percentage processed thus far.
	// Guaranteed to be
	// 100 when fully processed.
	ProgressPercent int64 `json:"progressPercent,omitempty"`

	// StartTime: Time when the request was received.
	StartTime string `json:"startTime,omitempty"`

	// UpdateTime: Time of the most recent update.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InputUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1VideoAnnotationProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1VideoAnnotationProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1VideoAnnotationResults: Annotation
// results for a single video.
type GoogleCloudVideointelligenceV1VideoAnnotationResults struct {
	// Error: If set, indicates an error. Note that for a single
	// `AnnotateVideoRequest`
	// some videos may succeed and some may fail.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// ExplicitAnnotation: Explicit content annotation.
	ExplicitAnnotation *GoogleCloudVideointelligenceV1ExplicitContentAnnotation `json:"explicitAnnotation,omitempty"`

	// FrameLabelAnnotations: Label annotations on frame level.
	// There is exactly one element for each unique label.
	FrameLabelAnnotations []*GoogleCloudVideointelligenceV1LabelAnnotation `json:"frameLabelAnnotations,omitempty"`

	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// SegmentLabelAnnotations: Label annotations on video level or user
	// specified segment level.
	// There is exactly one element for each unique label.
	SegmentLabelAnnotations []*GoogleCloudVideointelligenceV1LabelAnnotation `json:"segmentLabelAnnotations,omitempty"`

	// ShotAnnotations: Shot annotations. Each shot is represented as a
	// video segment.
	ShotAnnotations []*GoogleCloudVideointelligenceV1VideoSegment `json:"shotAnnotations,omitempty"`

	// ShotLabelAnnotations: Label annotations on shot level.
	// There is exactly one element for each unique label.
	ShotLabelAnnotations []*GoogleCloudVideointelligenceV1LabelAnnotation `json:"shotLabelAnnotations,omitempty"`

	// SpeechTranscriptions: Speech transcription.
	SpeechTranscriptions []*GoogleCloudVideointelligenceV1SpeechTranscription `json:"speechTranscriptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1VideoAnnotationResults) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1VideoAnnotationResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1VideoSegment: Video segment.
type GoogleCloudVideointelligenceV1VideoSegment struct {
	// EndTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the end of the segment (inclusive).
	EndTimeOffset string `json:"endTimeOffset,omitempty"`

	// StartTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the start of the segment (inclusive).
	StartTimeOffset string `json:"startTimeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTimeOffset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTimeOffset") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1VideoSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1VideoSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1WordInfo: Word-specific information for
// recognized words. Word information is only
// included in the response when certain request parameters are set,
// such
// as `enable_word_time_offsets`.
type GoogleCloudVideointelligenceV1WordInfo struct {
	// Confidence: Output only. The confidence estimate between 0.0 and 1.0.
	// A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is set only for the top alternative.
	// This field is not guaranteed to be accurate and users should not rely
	// on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// EndTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the end of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	EndTime string `json:"endTime,omitempty"`

	// SpeakerTag: Output only. A distinct integer value is assigned for
	// every speaker within
	// the audio. This field specifies which one of those speakers was
	// detected to
	// have spoken this word. Value ranges from 1 up to
	// diarization_speaker_count,
	// and is only set if speaker diarization is enabled.
	SpeakerTag int64 `json:"speakerTag,omitempty"`

	// StartTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the start of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	StartTime string `json:"startTime,omitempty"`

	// Word: The word corresponding to this set of information.
	Word string `json:"word,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1WordInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1WordInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1WordInfo) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1WordInfo
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1beta2AnnotateVideoProgress: Video
// annotation progress. Included in the `metadata`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1beta2AnnotateVideoProgress struct {
	// AnnotationProgress: Progress metadata for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationProgress []*GoogleCloudVideointelligenceV1beta2VideoAnnotationProgress `json:"annotationProgress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationProgress")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationProgress") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2AnnotateVideoProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2AnnotateVideoProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2AnnotateVideoRequest: Video
// annotation request.
type GoogleCloudVideointelligenceV1beta2AnnotateVideoRequest struct {
	// Features: Requested video annotation features.
	//
	// Possible values:
	//   "FEATURE_UNSPECIFIED" - Unspecified.
	//   "LABEL_DETECTION" - Label detection. Detect objects, such as dog or
	// flower.
	//   "SHOT_CHANGE_DETECTION" - Shot change detection.
	//   "EXPLICIT_CONTENT_DETECTION" - Explicit content detection.
	//   "SPEECH_TRANSCRIPTION" - Speech transcription.
	Features []string `json:"features,omitempty"`

	// InputContent: The video data bytes.
	// If unset, the input video(s) should be specified via `input_uri`.
	// If set, `input_uri` should be unset.
	InputContent string `json:"inputContent,omitempty"`

	// InputUri: Input video location. Currently, only
	// [Google Cloud Storage](https://cloud.google.com/storage/) URIs
	// are
	// supported, which must be specified in the following
	// format:
	// `gs://bucket-id/object-id` (other URI formats
	// return
	// google.rpc.Code.INVALID_ARGUMENT). For more information, see
	// [Request URIs](/storage/docs/reference-uris).
	// A video URI may include wildcards in `object-id`, and thus
	// identify
	// multiple videos. Supported wildcards: '*' to match 0 or more
	// characters;
	// '?' to match 1 character. If unset, the input video should be
	// embedded
	// in the request as `input_content`. If set, `input_content` should be
	// unset.
	InputUri string `json:"inputUri,omitempty"`

	// LocationId: Optional cloud region where annotation should take place.
	// Supported cloud
	// regions: `us-east1`, `us-west1`, `europe-west1`, `asia-east1`. If no
	// region
	// is specified, a region will be determined based on video file
	// location.
	LocationId string `json:"locationId,omitempty"`

	// OutputUri: Optional location where the output (in JSON format) should
	// be stored.
	// Currently, only [Google Cloud
	// Storage](https://cloud.google.com/storage/)
	// URIs are supported, which must be specified in the following
	// format:
	// `gs://bucket-id/object-id` (other URI formats
	// return
	// google.rpc.Code.INVALID_ARGUMENT). For more information, see
	// [Request URIs](/storage/docs/reference-uris).
	OutputUri string `json:"outputUri,omitempty"`

	// VideoContext: Additional video context and/or feature-specific
	// parameters.
	VideoContext *GoogleCloudVideointelligenceV1beta2VideoContext `json:"videoContext,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Features") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Features") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2AnnotateVideoRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2AnnotateVideoRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2AnnotateVideoResponse: Video
// annotation response. Included in the `response`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1beta2AnnotateVideoResponse struct {
	// AnnotationResults: Annotation results for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationResults []*GoogleCloudVideointelligenceV1beta2VideoAnnotationResults `json:"annotationResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationResults")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2AnnotateVideoResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2AnnotateVideoResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2Entity: Detected entity from video
// analysis.
type GoogleCloudVideointelligenceV1beta2Entity struct {
	// Description: Textual description, e.g. `Fixed-gear bicycle`.
	Description string `json:"description,omitempty"`

	// EntityId: Opaque entity ID. Some IDs may be available in
	// [Google Knowledge Graph
	// Search
	// API](https://developers.google.com/knowledge-graph/).
	EntityId string `json:"entityId,omitempty"`

	// LanguageCode: Language code for `description` in BCP-47 format.
	LanguageCode string `json:"languageCode,omitempty"`

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

func (s *GoogleCloudVideointelligenceV1beta2Entity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2Entity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2ExplicitContentAnnotation:
// Explicit content annotation (based on per-frame visual signals
// only).
// If no explicit content has been detected in a frame, no annotations
// are
// present for that frame.
type GoogleCloudVideointelligenceV1beta2ExplicitContentAnnotation struct {
	// Frames: All video frames where explicit content was detected.
	Frames []*GoogleCloudVideointelligenceV1beta2ExplicitContentFrame `json:"frames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Frames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Frames") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2ExplicitContentAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2ExplicitContentAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2ExplicitContentDetectionConfig:
// Config for EXPLICIT_CONTENT_DETECTION.
type GoogleCloudVideointelligenceV1beta2ExplicitContentDetectionConfig struct {
	// Model: Model to use for explicit content detection.
	// Supported values: "builtin/stable" (the default if unset)
	// and
	// "builtin/latest".
	Model string `json:"model,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Model") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Model") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2ExplicitContentDetectionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2ExplicitContentDetectionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2ExplicitContentFrame: Video frame
// level annotation results for explicit content.
type GoogleCloudVideointelligenceV1beta2ExplicitContentFrame struct {
	// PornographyLikelihood: Likelihood of the pornography content..
	//
	// Possible values:
	//   "LIKELIHOOD_UNSPECIFIED" - Unspecified likelihood.
	//   "VERY_UNLIKELY" - Very unlikely.
	//   "UNLIKELY" - Unlikely.
	//   "POSSIBLE" - Possible.
	//   "LIKELY" - Likely.
	//   "VERY_LIKELY" - Very likely.
	PornographyLikelihood string `json:"pornographyLikelihood,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "PornographyLikelihood") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PornographyLikelihood") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2ExplicitContentFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2ExplicitContentFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2LabelAnnotation: Label annotation.
type GoogleCloudVideointelligenceV1beta2LabelAnnotation struct {
	// CategoryEntities: Common categories for the detected entity.
	// E.g. when the label is `Terrier` the category is likely `dog`. And in
	// some
	// cases there might be more than one categories e.g. `Terrier` could
	// also be
	// a `pet`.
	CategoryEntities []*GoogleCloudVideointelligenceV1beta2Entity `json:"categoryEntities,omitempty"`

	// Entity: Detected entity.
	Entity *GoogleCloudVideointelligenceV1beta2Entity `json:"entity,omitempty"`

	// Frames: All video frames where a label was detected.
	Frames []*GoogleCloudVideointelligenceV1beta2LabelFrame `json:"frames,omitempty"`

	// Segments: All video segments where a label was detected.
	Segments []*GoogleCloudVideointelligenceV1beta2LabelSegment `json:"segments,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CategoryEntities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CategoryEntities") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2LabelAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2LabelAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2LabelDetectionConfig: Config for
// LABEL_DETECTION.
type GoogleCloudVideointelligenceV1beta2LabelDetectionConfig struct {
	// LabelDetectionMode: What labels should be detected with
	// LABEL_DETECTION, in addition to
	// video-level labels or segment-level labels.
	// If unspecified, defaults to `SHOT_MODE`.
	//
	// Possible values:
	//   "LABEL_DETECTION_MODE_UNSPECIFIED" - Unspecified.
	//   "SHOT_MODE" - Detect shot-level labels.
	//   "FRAME_MODE" - Detect frame-level labels.
	//   "SHOT_AND_FRAME_MODE" - Detect both shot-level and frame-level
	// labels.
	LabelDetectionMode string `json:"labelDetectionMode,omitempty"`

	// Model: Model to use for label detection.
	// Supported values: "builtin/stable" (the default if unset)
	// and
	// "builtin/latest".
	Model string `json:"model,omitempty"`

	// StationaryCamera: Whether the video has been shot from a stationary
	// (i.e. non-moving) camera.
	// When set to true, might improve detection accuracy for moving
	// objects.
	// Should be used with `SHOT_AND_FRAME_MODE` enabled.
	StationaryCamera bool `json:"stationaryCamera,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LabelDetectionMode")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LabelDetectionMode") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2LabelDetectionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2LabelDetectionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2LabelFrame: Video frame level
// annotation results for label detection.
type GoogleCloudVideointelligenceV1beta2LabelFrame struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2LabelFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2LabelFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1beta2LabelFrame) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1beta2LabelFrame
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1beta2LabelSegment: Video segment level
// annotation results for label detection.
type GoogleCloudVideointelligenceV1beta2LabelSegment struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// Segment: Video segment where a label was detected.
	Segment *GoogleCloudVideointelligenceV1beta2VideoSegment `json:"segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2LabelSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2LabelSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1beta2LabelSegment) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1beta2LabelSegment
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1beta2ShotChangeDetectionConfig: Config
// for SHOT_CHANGE_DETECTION.
type GoogleCloudVideointelligenceV1beta2ShotChangeDetectionConfig struct {
	// Model: Model to use for shot change detection.
	// Supported values: "builtin/stable" (the default if unset)
	// and
	// "builtin/latest".
	Model string `json:"model,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Model") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Model") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2ShotChangeDetectionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2ShotChangeDetectionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2SpeechContext: Provides "hints" to
// the speech recognizer to favor specific words and phrases
// in the results.
type GoogleCloudVideointelligenceV1beta2SpeechContext struct {
	// Phrases: *Optional* A list of strings containing words and phrases
	// "hints" so that
	// the speech recognition is more likely to recognize them. This can be
	// used
	// to improve the accuracy for specific words and phrases, for example,
	// if
	// specific commands are typically spoken by the user. This can also be
	// used
	// to add additional words to the vocabulary of the recognizer.
	// See
	// [usage limits](https://cloud.google.com/speech/limits#content).
	Phrases []string `json:"phrases,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Phrases") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Phrases") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2SpeechContext) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2SpeechContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative:
// Alternative hypotheses (a.k.a. n-best list).
type GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative struct {
	// Confidence: The confidence estimate between 0.0 and 1.0. A higher
	// number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is typically provided only for the top
	// hypothesis, and
	// only for `is_final=true` results. Clients should not rely on
	// the
	// `confidence` field as it is not guaranteed to be accurate or
	// consistent.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// Transcript: Transcript text representing the words that the user
	// spoke.
	Transcript string `json:"transcript,omitempty"`

	// Words: A list of word-specific information for each recognized word.
	Words []*GoogleCloudVideointelligenceV1beta2WordInfo `json:"words,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1beta2SpeechTranscription: A speech
// recognition result corresponding to a portion of the audio.
type GoogleCloudVideointelligenceV1beta2SpeechTranscription struct {
	// Alternatives: May contain one or more recognition hypotheses (up to
	// the maximum specified
	// in `max_alternatives`).  These alternatives are ordered in terms
	// of
	// accuracy, with the top (first) alternative being the most probable,
	// as
	// ranked by the recognizer.
	Alternatives []*GoogleCloudVideointelligenceV1beta2SpeechRecognitionAlternative `json:"alternatives,omitempty"`

	// LanguageCode: Output only.
	// The
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag
	// of the
	// language in this result. This language code was detected to have the
	// most
	// likelihood of being spoken in the audio.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alternatives") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alternatives") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2SpeechTranscription) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2SpeechTranscription
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2SpeechTranscriptionConfig: Config
// for SPEECH_TRANSCRIPTION.
type GoogleCloudVideointelligenceV1beta2SpeechTranscriptionConfig struct {
	// AudioTracks: *Optional* For file formats, such as MXF or MKV,
	// supporting multiple audio
	// tracks, specify up to two tracks. Default: track 0.
	AudioTracks []int64 `json:"audioTracks,omitempty"`

	// DiarizationSpeakerCount: *Optional*
	// If set, specifies the estimated number of speakers in the
	// conversation.
	// If not set, defaults to '2'.
	// Ignored unless enable_speaker_diarization is set to true.
	DiarizationSpeakerCount int64 `json:"diarizationSpeakerCount,omitempty"`

	// EnableAutomaticPunctuation: *Optional* If 'true', adds punctuation to
	// recognition result hypotheses.
	// This feature is only available in select languages. Setting this
	// for
	// requests in other languages has no effect at all. The default 'false'
	// value
	// does not add punctuation to result hypotheses. NOTE: "This is
	// currently
	// offered as an experimental service, complimentary to all users. In
	// the
	// future this may be exclusively available as a premium feature."
	EnableAutomaticPunctuation bool `json:"enableAutomaticPunctuation,omitempty"`

	// EnableSpeakerDiarization: *Optional* If 'true', enables speaker
	// detection for each recognized word in
	// the top alternative of the recognition result using a speaker_tag
	// provided
	// in the WordInfo.
	// Note: When this is true, we send all the words from the beginning of
	// the
	// audio for the top alternative in every consecutive responses.
	// This is done in order to improve our speaker tags as our models learn
	// to
	// identify the speakers in the conversation over time.
	EnableSpeakerDiarization bool `json:"enableSpeakerDiarization,omitempty"`

	// EnableWordConfidence: *Optional* If `true`, the top result includes a
	// list of words and the
	// confidence for those words. If `false`, no word-level
	// confidence
	// information is returned. The default is `false`.
	EnableWordConfidence bool `json:"enableWordConfidence,omitempty"`

	// FilterProfanity: *Optional* If set to `true`, the server will attempt
	// to filter out
	// profanities, replacing all but the initial character in each filtered
	// word
	// with asterisks, e.g. "f***". If set to `false` or omitted,
	// profanities
	// won't be filtered out.
	FilterProfanity bool `json:"filterProfanity,omitempty"`

	// LanguageCode: *Required* The language of the supplied audio as
	// a
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language
	// tag.
	// Example: "en-US".
	// See [Language
	// Support](https://cloud.google.com/speech/docs/languages)
	// for a list of the currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// MaxAlternatives: *Optional* Maximum number of recognition hypotheses
	// to be returned.
	// Specifically, the maximum number of `SpeechRecognitionAlternative`
	// messages
	// within each `SpeechTranscription`. The server may return fewer
	// than
	// `max_alternatives`. Valid values are `0`-`30`. A value of `0` or `1`
	// will
	// return a maximum of one. If omitted, will return a maximum of one.
	MaxAlternatives int64 `json:"maxAlternatives,omitempty"`

	// SpeechContexts: *Optional* A means to provide context to assist the
	// speech recognition.
	SpeechContexts []*GoogleCloudVideointelligenceV1beta2SpeechContext `json:"speechContexts,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioTracks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioTracks") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2SpeechTranscriptionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2SpeechTranscriptionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2VideoAnnotationProgress:
// Annotation progress for a single video.
type GoogleCloudVideointelligenceV1beta2VideoAnnotationProgress struct {
	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// ProgressPercent: Approximate percentage processed thus far.
	// Guaranteed to be
	// 100 when fully processed.
	ProgressPercent int64 `json:"progressPercent,omitempty"`

	// StartTime: Time when the request was received.
	StartTime string `json:"startTime,omitempty"`

	// UpdateTime: Time of the most recent update.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InputUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2VideoAnnotationProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2VideoAnnotationProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2VideoAnnotationResults: Annotation
// results for a single video.
type GoogleCloudVideointelligenceV1beta2VideoAnnotationResults struct {
	// Error: If set, indicates an error. Note that for a single
	// `AnnotateVideoRequest`
	// some videos may succeed and some may fail.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// ExplicitAnnotation: Explicit content annotation.
	ExplicitAnnotation *GoogleCloudVideointelligenceV1beta2ExplicitContentAnnotation `json:"explicitAnnotation,omitempty"`

	// FrameLabelAnnotations: Label annotations on frame level.
	// There is exactly one element for each unique label.
	FrameLabelAnnotations []*GoogleCloudVideointelligenceV1beta2LabelAnnotation `json:"frameLabelAnnotations,omitempty"`

	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// SegmentLabelAnnotations: Label annotations on video level or user
	// specified segment level.
	// There is exactly one element for each unique label.
	SegmentLabelAnnotations []*GoogleCloudVideointelligenceV1beta2LabelAnnotation `json:"segmentLabelAnnotations,omitempty"`

	// ShotAnnotations: Shot annotations. Each shot is represented as a
	// video segment.
	ShotAnnotations []*GoogleCloudVideointelligenceV1beta2VideoSegment `json:"shotAnnotations,omitempty"`

	// ShotLabelAnnotations: Label annotations on shot level.
	// There is exactly one element for each unique label.
	ShotLabelAnnotations []*GoogleCloudVideointelligenceV1beta2LabelAnnotation `json:"shotLabelAnnotations,omitempty"`

	// SpeechTranscriptions: Speech transcription.
	SpeechTranscriptions []*GoogleCloudVideointelligenceV1beta2SpeechTranscription `json:"speechTranscriptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2VideoAnnotationResults) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2VideoAnnotationResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2VideoContext: Video context and/or
// feature-specific parameters.
type GoogleCloudVideointelligenceV1beta2VideoContext struct {
	// ExplicitContentDetectionConfig: Config for
	// EXPLICIT_CONTENT_DETECTION.
	ExplicitContentDetectionConfig *GoogleCloudVideointelligenceV1beta2ExplicitContentDetectionConfig `json:"explicitContentDetectionConfig,omitempty"`

	// LabelDetectionConfig: Config for LABEL_DETECTION.
	LabelDetectionConfig *GoogleCloudVideointelligenceV1beta2LabelDetectionConfig `json:"labelDetectionConfig,omitempty"`

	// Segments: Video segments to annotate. The segments may overlap and
	// are not required
	// to be contiguous or span the whole video. If unspecified, each video
	// is
	// treated as a single segment.
	Segments []*GoogleCloudVideointelligenceV1beta2VideoSegment `json:"segments,omitempty"`

	// ShotChangeDetectionConfig: Config for SHOT_CHANGE_DETECTION.
	ShotChangeDetectionConfig *GoogleCloudVideointelligenceV1beta2ShotChangeDetectionConfig `json:"shotChangeDetectionConfig,omitempty"`

	// SpeechTranscriptionConfig: Config for SPEECH_TRANSCRIPTION.
	SpeechTranscriptionConfig *GoogleCloudVideointelligenceV1beta2SpeechTranscriptionConfig `json:"speechTranscriptionConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ExplicitContentDetectionConfig") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "ExplicitContentDetectionConfig") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2VideoContext) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2VideoContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2VideoSegment: Video segment.
type GoogleCloudVideointelligenceV1beta2VideoSegment struct {
	// EndTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the end of the segment (inclusive).
	EndTimeOffset string `json:"endTimeOffset,omitempty"`

	// StartTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the start of the segment (inclusive).
	StartTimeOffset string `json:"startTimeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTimeOffset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTimeOffset") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2VideoSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2VideoSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1beta2WordInfo: Word-specific
// information for recognized words. Word information is only
// included in the response when certain request parameters are set,
// such
// as `enable_word_time_offsets`.
type GoogleCloudVideointelligenceV1beta2WordInfo struct {
	// Confidence: Output only. The confidence estimate between 0.0 and 1.0.
	// A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is set only for the top alternative.
	// This field is not guaranteed to be accurate and users should not rely
	// on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// EndTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the end of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	EndTime string `json:"endTime,omitempty"`

	// SpeakerTag: Output only. A distinct integer value is assigned for
	// every speaker within
	// the audio. This field specifies which one of those speakers was
	// detected to
	// have spoken this word. Value ranges from 1 up to
	// diarization_speaker_count,
	// and is only set if speaker diarization is enabled.
	SpeakerTag int64 `json:"speakerTag,omitempty"`

	// StartTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the start of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	StartTime string `json:"startTime,omitempty"`

	// Word: The word corresponding to this set of information.
	Word string `json:"word,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1beta2WordInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1beta2WordInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1beta2WordInfo) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1beta2WordInfo
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p1beta1AnnotateVideoProgress: Video
// annotation progress. Included in the `metadata`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1p1beta1AnnotateVideoProgress struct {
	// AnnotationProgress: Progress metadata for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationProgress []*GoogleCloudVideointelligenceV1p1beta1VideoAnnotationProgress `json:"annotationProgress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationProgress")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationProgress") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1AnnotateVideoProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1AnnotateVideoProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1AnnotateVideoResponse: Video
// annotation response. Included in the `response`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1p1beta1AnnotateVideoResponse struct {
	// AnnotationResults: Annotation results for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationResults []*GoogleCloudVideointelligenceV1p1beta1VideoAnnotationResults `json:"annotationResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationResults")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1AnnotateVideoResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1AnnotateVideoResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1Entity: Detected entity from
// video analysis.
type GoogleCloudVideointelligenceV1p1beta1Entity struct {
	// Description: Textual description, e.g. `Fixed-gear bicycle`.
	Description string `json:"description,omitempty"`

	// EntityId: Opaque entity ID. Some IDs may be available in
	// [Google Knowledge Graph
	// Search
	// API](https://developers.google.com/knowledge-graph/).
	EntityId string `json:"entityId,omitempty"`

	// LanguageCode: Language code for `description` in BCP-47 format.
	LanguageCode string `json:"languageCode,omitempty"`

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

func (s *GoogleCloudVideointelligenceV1p1beta1Entity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1Entity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1ExplicitContentAnnotation:
// Explicit content annotation (based on per-frame visual signals
// only).
// If no explicit content has been detected in a frame, no annotations
// are
// present for that frame.
type GoogleCloudVideointelligenceV1p1beta1ExplicitContentAnnotation struct {
	// Frames: All video frames where explicit content was detected.
	Frames []*GoogleCloudVideointelligenceV1p1beta1ExplicitContentFrame `json:"frames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Frames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Frames") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1ExplicitContentAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1ExplicitContentAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1ExplicitContentFrame: Video
// frame level annotation results for explicit content.
type GoogleCloudVideointelligenceV1p1beta1ExplicitContentFrame struct {
	// PornographyLikelihood: Likelihood of the pornography content..
	//
	// Possible values:
	//   "LIKELIHOOD_UNSPECIFIED" - Unspecified likelihood.
	//   "VERY_UNLIKELY" - Very unlikely.
	//   "UNLIKELY" - Unlikely.
	//   "POSSIBLE" - Possible.
	//   "LIKELY" - Likely.
	//   "VERY_LIKELY" - Very likely.
	PornographyLikelihood string `json:"pornographyLikelihood,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "PornographyLikelihood") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PornographyLikelihood") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1ExplicitContentFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1ExplicitContentFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1LabelAnnotation: Label
// annotation.
type GoogleCloudVideointelligenceV1p1beta1LabelAnnotation struct {
	// CategoryEntities: Common categories for the detected entity.
	// E.g. when the label is `Terrier` the category is likely `dog`. And in
	// some
	// cases there might be more than one categories e.g. `Terrier` could
	// also be
	// a `pet`.
	CategoryEntities []*GoogleCloudVideointelligenceV1p1beta1Entity `json:"categoryEntities,omitempty"`

	// Entity: Detected entity.
	Entity *GoogleCloudVideointelligenceV1p1beta1Entity `json:"entity,omitempty"`

	// Frames: All video frames where a label was detected.
	Frames []*GoogleCloudVideointelligenceV1p1beta1LabelFrame `json:"frames,omitempty"`

	// Segments: All video segments where a label was detected.
	Segments []*GoogleCloudVideointelligenceV1p1beta1LabelSegment `json:"segments,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CategoryEntities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CategoryEntities") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1LabelAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1LabelAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1LabelFrame: Video frame level
// annotation results for label detection.
type GoogleCloudVideointelligenceV1p1beta1LabelFrame struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1LabelFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1LabelFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p1beta1LabelFrame) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1LabelFrame
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p1beta1LabelSegment: Video segment
// level annotation results for label detection.
type GoogleCloudVideointelligenceV1p1beta1LabelSegment struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// Segment: Video segment where a label was detected.
	Segment *GoogleCloudVideointelligenceV1p1beta1VideoSegment `json:"segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1LabelSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1LabelSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p1beta1LabelSegment) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1LabelSegment
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative:
// Alternative hypotheses (a.k.a. n-best list).
type GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative struct {
	// Confidence: The confidence estimate between 0.0 and 1.0. A higher
	// number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is typically provided only for the top
	// hypothesis, and
	// only for `is_final=true` results. Clients should not rely on
	// the
	// `confidence` field as it is not guaranteed to be accurate or
	// consistent.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// Transcript: Transcript text representing the words that the user
	// spoke.
	Transcript string `json:"transcript,omitempty"`

	// Words: A list of word-specific information for each recognized word.
	Words []*GoogleCloudVideointelligenceV1p1beta1WordInfo `json:"words,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p1beta1SpeechTranscription: A speech
// recognition result corresponding to a portion of the audio.
type GoogleCloudVideointelligenceV1p1beta1SpeechTranscription struct {
	// Alternatives: May contain one or more recognition hypotheses (up to
	// the maximum specified
	// in `max_alternatives`).  These alternatives are ordered in terms
	// of
	// accuracy, with the top (first) alternative being the most probable,
	// as
	// ranked by the recognizer.
	Alternatives []*GoogleCloudVideointelligenceV1p1beta1SpeechRecognitionAlternative `json:"alternatives,omitempty"`

	// LanguageCode: Output only.
	// The
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag
	// of the
	// language in this result. This language code was detected to have the
	// most
	// likelihood of being spoken in the audio.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alternatives") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alternatives") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1SpeechTranscription) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1SpeechTranscription
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1VideoAnnotationProgress:
// Annotation progress for a single video.
type GoogleCloudVideointelligenceV1p1beta1VideoAnnotationProgress struct {
	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// ProgressPercent: Approximate percentage processed thus far.
	// Guaranteed to be
	// 100 when fully processed.
	ProgressPercent int64 `json:"progressPercent,omitempty"`

	// StartTime: Time when the request was received.
	StartTime string `json:"startTime,omitempty"`

	// UpdateTime: Time of the most recent update.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InputUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1VideoAnnotationProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1VideoAnnotationProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1VideoAnnotationResults:
// Annotation results for a single video.
type GoogleCloudVideointelligenceV1p1beta1VideoAnnotationResults struct {
	// Error: If set, indicates an error. Note that for a single
	// `AnnotateVideoRequest`
	// some videos may succeed and some may fail.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// ExplicitAnnotation: Explicit content annotation.
	ExplicitAnnotation *GoogleCloudVideointelligenceV1p1beta1ExplicitContentAnnotation `json:"explicitAnnotation,omitempty"`

	// FrameLabelAnnotations: Label annotations on frame level.
	// There is exactly one element for each unique label.
	FrameLabelAnnotations []*GoogleCloudVideointelligenceV1p1beta1LabelAnnotation `json:"frameLabelAnnotations,omitempty"`

	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// SegmentLabelAnnotations: Label annotations on video level or user
	// specified segment level.
	// There is exactly one element for each unique label.
	SegmentLabelAnnotations []*GoogleCloudVideointelligenceV1p1beta1LabelAnnotation `json:"segmentLabelAnnotations,omitempty"`

	// ShotAnnotations: Shot annotations. Each shot is represented as a
	// video segment.
	ShotAnnotations []*GoogleCloudVideointelligenceV1p1beta1VideoSegment `json:"shotAnnotations,omitempty"`

	// ShotLabelAnnotations: Label annotations on shot level.
	// There is exactly one element for each unique label.
	ShotLabelAnnotations []*GoogleCloudVideointelligenceV1p1beta1LabelAnnotation `json:"shotLabelAnnotations,omitempty"`

	// SpeechTranscriptions: Speech transcription.
	SpeechTranscriptions []*GoogleCloudVideointelligenceV1p1beta1SpeechTranscription `json:"speechTranscriptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1VideoAnnotationResults) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1VideoAnnotationResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1VideoSegment: Video segment.
type GoogleCloudVideointelligenceV1p1beta1VideoSegment struct {
	// EndTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the end of the segment (inclusive).
	EndTimeOffset string `json:"endTimeOffset,omitempty"`

	// StartTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the start of the segment (inclusive).
	StartTimeOffset string `json:"startTimeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTimeOffset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTimeOffset") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1VideoSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1VideoSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p1beta1WordInfo: Word-specific
// information for recognized words. Word information is only
// included in the response when certain request parameters are set,
// such
// as `enable_word_time_offsets`.
type GoogleCloudVideointelligenceV1p1beta1WordInfo struct {
	// Confidence: Output only. The confidence estimate between 0.0 and 1.0.
	// A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is set only for the top alternative.
	// This field is not guaranteed to be accurate and users should not rely
	// on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// EndTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the end of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	EndTime string `json:"endTime,omitempty"`

	// SpeakerTag: Output only. A distinct integer value is assigned for
	// every speaker within
	// the audio. This field specifies which one of those speakers was
	// detected to
	// have spoken this word. Value ranges from 1 up to
	// diarization_speaker_count,
	// and is only set if speaker diarization is enabled.
	SpeakerTag int64 `json:"speakerTag,omitempty"`

	// StartTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the start of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	StartTime string `json:"startTime,omitempty"`

	// Word: The word corresponding to this set of information.
	Word string `json:"word,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p1beta1WordInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1WordInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p1beta1WordInfo) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p1beta1WordInfo
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1AnnotateVideoProgress: Video
// annotation progress. Included in the `metadata`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1p2beta1AnnotateVideoProgress struct {
	// AnnotationProgress: Progress metadata for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationProgress []*GoogleCloudVideointelligenceV1p2beta1VideoAnnotationProgress `json:"annotationProgress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationProgress")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationProgress") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1AnnotateVideoProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1AnnotateVideoProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1AnnotateVideoResponse: Video
// annotation response. Included in the `response`
// field of the `Operation` returned by the `GetOperation`
// call of the `google::longrunning::Operations` service.
type GoogleCloudVideointelligenceV1p2beta1AnnotateVideoResponse struct {
	// AnnotationResults: Annotation results for all videos specified in
	// `AnnotateVideoRequest`.
	AnnotationResults []*GoogleCloudVideointelligenceV1p2beta1VideoAnnotationResults `json:"annotationResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnnotationResults")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnnotationResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1AnnotateVideoResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1AnnotateVideoResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1Entity: Detected entity from
// video analysis.
type GoogleCloudVideointelligenceV1p2beta1Entity struct {
	// Description: Textual description, e.g. `Fixed-gear bicycle`.
	Description string `json:"description,omitempty"`

	// EntityId: Opaque entity ID. Some IDs may be available in
	// [Google Knowledge Graph
	// Search
	// API](https://developers.google.com/knowledge-graph/).
	EntityId string `json:"entityId,omitempty"`

	// LanguageCode: Language code for `description` in BCP-47 format.
	LanguageCode string `json:"languageCode,omitempty"`

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

func (s *GoogleCloudVideointelligenceV1p2beta1Entity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1Entity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1ExplicitContentAnnotation:
// Explicit content annotation (based on per-frame visual signals
// only).
// If no explicit content has been detected in a frame, no annotations
// are
// present for that frame.
type GoogleCloudVideointelligenceV1p2beta1ExplicitContentAnnotation struct {
	// Frames: All video frames where explicit content was detected.
	Frames []*GoogleCloudVideointelligenceV1p2beta1ExplicitContentFrame `json:"frames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Frames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Frames") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1ExplicitContentAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1ExplicitContentAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1ExplicitContentFrame: Video
// frame level annotation results for explicit content.
type GoogleCloudVideointelligenceV1p2beta1ExplicitContentFrame struct {
	// PornographyLikelihood: Likelihood of the pornography content..
	//
	// Possible values:
	//   "LIKELIHOOD_UNSPECIFIED" - Unspecified likelihood.
	//   "VERY_UNLIKELY" - Very unlikely.
	//   "UNLIKELY" - Unlikely.
	//   "POSSIBLE" - Possible.
	//   "LIKELY" - Likely.
	//   "VERY_LIKELY" - Very likely.
	PornographyLikelihood string `json:"pornographyLikelihood,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "PornographyLikelihood") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PornographyLikelihood") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1ExplicitContentFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1ExplicitContentFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1LabelAnnotation: Label
// annotation.
type GoogleCloudVideointelligenceV1p2beta1LabelAnnotation struct {
	// CategoryEntities: Common categories for the detected entity.
	// E.g. when the label is `Terrier` the category is likely `dog`. And in
	// some
	// cases there might be more than one categories e.g. `Terrier` could
	// also be
	// a `pet`.
	CategoryEntities []*GoogleCloudVideointelligenceV1p2beta1Entity `json:"categoryEntities,omitempty"`

	// Entity: Detected entity.
	Entity *GoogleCloudVideointelligenceV1p2beta1Entity `json:"entity,omitempty"`

	// Frames: All video frames where a label was detected.
	Frames []*GoogleCloudVideointelligenceV1p2beta1LabelFrame `json:"frames,omitempty"`

	// Segments: All video segments where a label was detected.
	Segments []*GoogleCloudVideointelligenceV1p2beta1LabelSegment `json:"segments,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CategoryEntities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CategoryEntities") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1LabelAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1LabelAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1LabelFrame: Video frame level
// annotation results for label detection.
type GoogleCloudVideointelligenceV1p2beta1LabelFrame struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// TimeOffset: Time-offset, relative to the beginning of the video,
	// corresponding to the
	// video frame for this location.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1LabelFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1LabelFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1LabelFrame) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1LabelFrame
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1LabelSegment: Video segment
// level annotation results for label detection.
type GoogleCloudVideointelligenceV1p2beta1LabelSegment struct {
	// Confidence: Confidence that the label is accurate. Range: [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// Segment: Video segment where a label was detected.
	Segment *GoogleCloudVideointelligenceV1p2beta1VideoSegment `json:"segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1LabelSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1LabelSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1LabelSegment) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1LabelSegment
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox:
// Normalized bounding box.
// The normalized vertex coordinates are relative to the original
// image.
// Range: [0, 1].
type GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox struct {
	// Bottom: Bottom Y coordinate.
	Bottom float64 `json:"bottom,omitempty"`

	// Left: Left X coordinate.
	Left float64 `json:"left,omitempty"`

	// Right: Right X coordinate.
	Right float64 `json:"right,omitempty"`

	// Top: Top Y coordinate.
	Top float64 `json:"top,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bottom") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bottom") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox
	var s1 struct {
		Bottom gensupport.JSONFloat64 `json:"bottom"`
		Left   gensupport.JSONFloat64 `json:"left"`
		Right  gensupport.JSONFloat64 `json:"right"`
		Top    gensupport.JSONFloat64 `json:"top"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Bottom = float64(s1.Bottom)
	s.Left = float64(s1.Left)
	s.Right = float64(s1.Right)
	s.Top = float64(s1.Top)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingPoly:
// Normalized bounding polygon for text (that might not be aligned with
// axis).
// Contains list of the corner points in clockwise order starting
// from
// top-left corner. For example, for a rectangular bounding box:
// When the text is horizontal it might look like:
//         0----1
//         |    |
//         3----2
//
// When it's clockwise rotated 180 degrees around the top-left corner
// it
// becomes:
//         2----3
//         |    |
//         1----0
//
// and the vertex order will still be (0, 1, 2, 3). Note that values can
// be less
// than 0, or greater than 1 due to trignometric calculations for
// location of
// the box.
type GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingPoly struct {
	// Vertices: Normalized vertices of the bounding polygon.
	Vertices []*GoogleCloudVideointelligenceV1p2beta1NormalizedVertex `json:"vertices,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Vertices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Vertices") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingPoly) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingPoly
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1NormalizedVertex: A vertex
// represents a 2D point in the image.
// NOTE: the normalized vertex coordinates are relative to the original
// image
// and range from 0 to 1.
type GoogleCloudVideointelligenceV1p2beta1NormalizedVertex struct {
	// X: X coordinate.
	X float64 `json:"x,omitempty"`

	// Y: Y coordinate.
	Y float64 `json:"y,omitempty"`

	// ForceSendFields is a list of field names (e.g. "X") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "X") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1NormalizedVertex) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1NormalizedVertex
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1NormalizedVertex) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1NormalizedVertex
	var s1 struct {
		X gensupport.JSONFloat64 `json:"x"`
		Y gensupport.JSONFloat64 `json:"y"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.X = float64(s1.X)
	s.Y = float64(s1.Y)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation:
// Annotations corresponding to one tracked object.
type GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation struct {
	// Confidence: Object category's labeling confidence of this track.
	Confidence float64 `json:"confidence,omitempty"`

	// Entity: Entity to specify the object category that this track is
	// labeled as.
	Entity *GoogleCloudVideointelligenceV1p2beta1Entity `json:"entity,omitempty"`

	// Frames: Information corresponding to all frames where this object
	// track appears.
	Frames []*GoogleCloudVideointelligenceV1p2beta1ObjectTrackingFrame `json:"frames,omitempty"`

	// Segment: Each object track corresponds to one video segment where it
	// appears.
	Segment *GoogleCloudVideointelligenceV1p2beta1VideoSegment `json:"segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1ObjectTrackingFrame: Video frame
// level annotations for object detection and tracking. This
// field
// stores per frame location, time offset, and confidence.
type GoogleCloudVideointelligenceV1p2beta1ObjectTrackingFrame struct {
	// NormalizedBoundingBox: The normalized bounding box location of this
	// object track for the frame.
	NormalizedBoundingBox *GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingBox `json:"normalizedBoundingBox,omitempty"`

	// TimeOffset: The timestamp of the frame in microseconds.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "NormalizedBoundingBox") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NormalizedBoundingBox") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1ObjectTrackingFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1ObjectTrackingFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative:
// Alternative hypotheses (a.k.a. n-best list).
type GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative struct {
	// Confidence: The confidence estimate between 0.0 and 1.0. A higher
	// number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is typically provided only for the top
	// hypothesis, and
	// only for `is_final=true` results. Clients should not rely on
	// the
	// `confidence` field as it is not guaranteed to be accurate or
	// consistent.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// Transcript: Transcript text representing the words that the user
	// spoke.
	Transcript string `json:"transcript,omitempty"`

	// Words: A list of word-specific information for each recognized word.
	Words []*GoogleCloudVideointelligenceV1p2beta1WordInfo `json:"words,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1SpeechTranscription: A speech
// recognition result corresponding to a portion of the audio.
type GoogleCloudVideointelligenceV1p2beta1SpeechTranscription struct {
	// Alternatives: May contain one or more recognition hypotheses (up to
	// the maximum specified
	// in `max_alternatives`).  These alternatives are ordered in terms
	// of
	// accuracy, with the top (first) alternative being the most probable,
	// as
	// ranked by the recognizer.
	Alternatives []*GoogleCloudVideointelligenceV1p2beta1SpeechRecognitionAlternative `json:"alternatives,omitempty"`

	// LanguageCode: Output only.
	// The
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag
	// of the
	// language in this result. This language code was detected to have the
	// most
	// likelihood of being spoken in the audio.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alternatives") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alternatives") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1SpeechTranscription) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1SpeechTranscription
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1TextAnnotation: Annotations
// related to one detected OCR text snippet. This will contain
// the
// corresponding text, confidence value, and frame level information for
// each
// detection.
type GoogleCloudVideointelligenceV1p2beta1TextAnnotation struct {
	// Segments: All video segments where OCR detected text appears.
	Segments []*GoogleCloudVideointelligenceV1p2beta1TextSegment `json:"segments,omitempty"`

	// Text: The detected text.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Segments") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Segments") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1TextAnnotation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1TextAnnotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1TextFrame: Video frame level
// annotation results for text annotation (OCR).
// Contains information regarding timestamp and bounding box locations
// for the
// frames containing detected OCR text snippets.
type GoogleCloudVideointelligenceV1p2beta1TextFrame struct {
	// RotatedBoundingBox: Bounding polygon of the detected text for this
	// frame.
	RotatedBoundingBox *GoogleCloudVideointelligenceV1p2beta1NormalizedBoundingPoly `json:"rotatedBoundingBox,omitempty"`

	// TimeOffset: Timestamp of this frame.
	TimeOffset string `json:"timeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RotatedBoundingBox")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RotatedBoundingBox") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1TextFrame) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1TextFrame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1TextSegment: Video segment level
// annotation results for text detection.
type GoogleCloudVideointelligenceV1p2beta1TextSegment struct {
	// Confidence: Confidence for the track of detected text. It is
	// calculated as the highest
	// over all frames where OCR detected text appears.
	Confidence float64 `json:"confidence,omitempty"`

	// Frames: Information related to the frames where OCR detected text
	// appears.
	Frames []*GoogleCloudVideointelligenceV1p2beta1TextFrame `json:"frames,omitempty"`

	// Segment: Video segment where a text snippet was detected.
	Segment *GoogleCloudVideointelligenceV1p2beta1VideoSegment `json:"segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1TextSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1TextSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1TextSegment) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1TextSegment
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudVideointelligenceV1p2beta1VideoAnnotationProgress:
// Annotation progress for a single video.
type GoogleCloudVideointelligenceV1p2beta1VideoAnnotationProgress struct {
	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// ProgressPercent: Approximate percentage processed thus far.
	// Guaranteed to be
	// 100 when fully processed.
	ProgressPercent int64 `json:"progressPercent,omitempty"`

	// StartTime: Time when the request was received.
	StartTime string `json:"startTime,omitempty"`

	// UpdateTime: Time of the most recent update.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InputUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1VideoAnnotationProgress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1VideoAnnotationProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1VideoAnnotationResults:
// Annotation results for a single video.
type GoogleCloudVideointelligenceV1p2beta1VideoAnnotationResults struct {
	// Error: If set, indicates an error. Note that for a single
	// `AnnotateVideoRequest`
	// some videos may succeed and some may fail.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// ExplicitAnnotation: Explicit content annotation.
	ExplicitAnnotation *GoogleCloudVideointelligenceV1p2beta1ExplicitContentAnnotation `json:"explicitAnnotation,omitempty"`

	// FrameLabelAnnotations: Label annotations on frame level.
	// There is exactly one element for each unique label.
	FrameLabelAnnotations []*GoogleCloudVideointelligenceV1p2beta1LabelAnnotation `json:"frameLabelAnnotations,omitempty"`

	// InputUri: Video file location in
	// [Google Cloud Storage](https://cloud.google.com/storage/).
	InputUri string `json:"inputUri,omitempty"`

	// ObjectAnnotations: Annotations for list of objects detected and
	// tracked in video.
	ObjectAnnotations []*GoogleCloudVideointelligenceV1p2beta1ObjectTrackingAnnotation `json:"objectAnnotations,omitempty"`

	// SegmentLabelAnnotations: Label annotations on video level or user
	// specified segment level.
	// There is exactly one element for each unique label.
	SegmentLabelAnnotations []*GoogleCloudVideointelligenceV1p2beta1LabelAnnotation `json:"segmentLabelAnnotations,omitempty"`

	// ShotAnnotations: Shot annotations. Each shot is represented as a
	// video segment.
	ShotAnnotations []*GoogleCloudVideointelligenceV1p2beta1VideoSegment `json:"shotAnnotations,omitempty"`

	// ShotLabelAnnotations: Label annotations on shot level.
	// There is exactly one element for each unique label.
	ShotLabelAnnotations []*GoogleCloudVideointelligenceV1p2beta1LabelAnnotation `json:"shotLabelAnnotations,omitempty"`

	// SpeechTranscriptions: Speech transcription.
	SpeechTranscriptions []*GoogleCloudVideointelligenceV1p2beta1SpeechTranscription `json:"speechTranscriptions,omitempty"`

	// TextAnnotations: OCR text detection and tracking.
	// Annotations for list of detected text snippets. Each will have list
	// of
	// frame information associated with it.
	TextAnnotations []*GoogleCloudVideointelligenceV1p2beta1TextAnnotation `json:"textAnnotations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1VideoAnnotationResults) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1VideoAnnotationResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1VideoSegment: Video segment.
type GoogleCloudVideointelligenceV1p2beta1VideoSegment struct {
	// EndTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the end of the segment (inclusive).
	EndTimeOffset string `json:"endTimeOffset,omitempty"`

	// StartTimeOffset: Time-offset, relative to the beginning of the
	// video,
	// corresponding to the start of the segment (inclusive).
	StartTimeOffset string `json:"startTimeOffset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTimeOffset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTimeOffset") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1VideoSegment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1VideoSegment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudVideointelligenceV1p2beta1WordInfo: Word-specific
// information for recognized words. Word information is only
// included in the response when certain request parameters are set,
// such
// as `enable_word_time_offsets`.
type GoogleCloudVideointelligenceV1p2beta1WordInfo struct {
	// Confidence: Output only. The confidence estimate between 0.0 and 1.0.
	// A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is set only for the top alternative.
	// This field is not guaranteed to be accurate and users should not rely
	// on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// EndTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the end of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	EndTime string `json:"endTime,omitempty"`

	// SpeakerTag: Output only. A distinct integer value is assigned for
	// every speaker within
	// the audio. This field specifies which one of those speakers was
	// detected to
	// have spoken this word. Value ranges from 1 up to
	// diarization_speaker_count,
	// and is only set if speaker diarization is enabled.
	SpeakerTag int64 `json:"speakerTag,omitempty"`

	// StartTime: Time offset relative to the beginning of the audio,
	// and
	// corresponding to the start of the spoken word. This field is only set
	// if
	// `enable_word_time_offsets=true` and only in the top hypothesis. This
	// is an
	// experimental feature and the accuracy of the time offset can vary.
	StartTime string `json:"startTime,omitempty"`

	// Word: The word corresponding to this set of information.
	Word string `json:"word,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudVideointelligenceV1p2beta1WordInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1WordInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudVideointelligenceV1p2beta1WordInfo) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudVideointelligenceV1p2beta1WordInfo
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleLongrunningOperation: This resource represents a long-running
// operation that is the result of a
// network API call.
type GoogleLongrunningOperation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleLongrunningOperation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleLongrunningOperation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleRpcStatus: The `Status` type defines a logical error model that
// is suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type GoogleRpcStatus struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
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

func (s *GoogleRpcStatus) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleRpcStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "videointelligence.videos.annotate":

type VideosAnnotateCall struct {
	s                                                       *Service
	googlecloudvideointelligencev1beta2Annotatevideorequest *GoogleCloudVideointelligenceV1beta2AnnotateVideoRequest
	urlParams_                                              gensupport.URLParams
	ctx_                                                    context.Context
	header_                                                 http.Header
}

// Annotate: Performs asynchronous video annotation. Progress and
// results can be
// retrieved through the `google.longrunning.Operations`
// interface.
// `Operation.metadata` contains `AnnotateVideoProgress`
// (progress).
// `Operation.response` contains `AnnotateVideoResponse` (results).
func (r *VideosService) Annotate(googlecloudvideointelligencev1beta2Annotatevideorequest *GoogleCloudVideointelligenceV1beta2AnnotateVideoRequest) *VideosAnnotateCall {
	c := &VideosAnnotateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.googlecloudvideointelligencev1beta2Annotatevideorequest = googlecloudvideointelligencev1beta2Annotatevideorequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VideosAnnotateCall) Fields(s ...googleapi.Field) *VideosAnnotateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *VideosAnnotateCall) Context(ctx context.Context) *VideosAnnotateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *VideosAnnotateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *VideosAnnotateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudvideointelligencev1beta2Annotatevideorequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/videos:annotate")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "videointelligence.videos.annotate" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *VideosAnnotateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
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
	//   "description": "Performs asynchronous video annotation. Progress and results can be\nretrieved through the `google.longrunning.Operations` interface.\n`Operation.metadata` contains `AnnotateVideoProgress` (progress).\n`Operation.response` contains `AnnotateVideoResponse` (results).",
	//   "flatPath": "v1beta2/videos:annotate",
	//   "httpMethod": "POST",
	//   "id": "videointelligence.videos.annotate",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta2/videos:annotate",
	//   "request": {
	//     "$ref": "GoogleCloudVideointelligenceV1beta2_AnnotateVideoRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning_Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
