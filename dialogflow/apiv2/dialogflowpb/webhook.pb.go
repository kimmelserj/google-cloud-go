// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/dialogflow/v2/webhook.proto

package dialogflowpb

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for a webhook call.
type WebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`, or
	// `projects/<Project ID>/agent/environments/<Environment ID>/users/<User
	// ID>/sessions/<Session ID>`.
	Session string `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	// The unique identifier of the response. Contains the same value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	// The result of the conversational query or event processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *QueryResult `protobuf:"bytes,2,opt,name=query_result,json=queryResult,proto3" json:"query_result,omitempty"`
	// Optional. The contents of the original request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *OriginalDetectIntentRequest `protobuf:"bytes,3,opt,name=original_detect_intent_request,json=originalDetectIntentRequest,proto3" json:"original_detect_intent_request,omitempty"`
}

func (x *WebhookRequest) Reset() {
	*x = WebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookRequest) ProtoMessage() {}

func (x *WebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookRequest.ProtoReflect.Descriptor instead.
func (*WebhookRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *WebhookRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *WebhookRequest) GetResponseId() string {
	if x != nil {
		return x.ResponseId
	}
	return ""
}

func (x *WebhookRequest) GetQueryResult() *QueryResult {
	if x != nil {
		return x.QueryResult
	}
	return nil
}

func (x *WebhookRequest) GetOriginalDetectIntentRequest() *OriginalDetectIntentRequest {
	if x != nil {
		return x.OriginalDetectIntentRequest
	}
	return nil
}

// The response message for a webhook call.
//
// This response is validated by the Dialogflow server. If validation fails,
// an error will be returned in the [QueryResult.diagnostic_info][google.cloud.dialogflow.v2.QueryResult.diagnostic_info] field.
// Setting JSON fields to an empty value with the wrong type is a common error.
// To avoid this error:
//
// - Use `""` for empty strings
// - Use `{}` or `null` for empty objects
// - Use `[]` or `null` for empty arrays
//
// For more information, see the
// [Protocol Buffers Language
// Guide](https://developers.google.com/protocol-buffers/docs/proto3#json).
type WebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The text response message intended for the end-user.
	// It is recommended to use `fulfillment_messages.text.text[0]` instead.
	// When provided, Dialogflow uses this field to populate
	// [QueryResult.fulfillment_text][google.cloud.dialogflow.v2.QueryResult.fulfillment_text] sent to the integration or API caller.
	FulfillmentText string `protobuf:"bytes,1,opt,name=fulfillment_text,json=fulfillmentText,proto3" json:"fulfillment_text,omitempty"`
	// Optional. The rich response messages intended for the end-user.
	// When provided, Dialogflow uses this field to populate
	// [QueryResult.fulfillment_messages][google.cloud.dialogflow.v2.QueryResult.fulfillment_messages] sent to the integration or API caller.
	FulfillmentMessages []*Intent_Message `protobuf:"bytes,2,rep,name=fulfillment_messages,json=fulfillmentMessages,proto3" json:"fulfillment_messages,omitempty"`
	// Optional. A custom field used to identify the webhook source.
	// Arbitrary strings are supported.
	// When provided, Dialogflow uses this field to populate
	// [QueryResult.webhook_source][google.cloud.dialogflow.v2.QueryResult.webhook_source] sent to the integration or API caller.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. This field can be used to pass custom data from your webhook to the
	// integration or API caller. Arbitrary JSON objects are supported.
	// When provided, Dialogflow uses this field to populate
	// [QueryResult.webhook_payload][google.cloud.dialogflow.v2.QueryResult.webhook_payload] sent to the integration or API caller.
	// This field is also used by the
	// [Google Assistant
	// integration](https://cloud.google.com/dialogflow/docs/integrations/aog)
	// for rich response messages.
	// See the format definition at [Google Assistant Dialogflow webhook
	// format](https://developers.google.com/assistant/actions/build/json/dialogflow-webhook-json)
	Payload *_struct.Struct `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional. The collection of output contexts that will overwrite currently
	// active contexts for the session and reset their lifespans.
	// When provided, Dialogflow uses this field to populate
	// [QueryResult.output_contexts][google.cloud.dialogflow.v2.QueryResult.output_contexts] sent to the integration or API caller.
	OutputContexts []*Context `protobuf:"bytes,5,rep,name=output_contexts,json=outputContexts,proto3" json:"output_contexts,omitempty"`
	// Optional. Invokes the supplied events.
	// When this field is set, Dialogflow ignores the `fulfillment_text`,
	// `fulfillment_messages`, and `payload` fields.
	FollowupEventInput *EventInput `protobuf:"bytes,6,opt,name=followup_event_input,json=followupEventInput,proto3" json:"followup_event_input,omitempty"`
	// Optional. Additional session entity types to replace or extend developer
	// entity types with. The entity synonyms apply to all languages and persist
	// for the session. Setting this data from a webhook overwrites
	// the session entity types that have been set using `detectIntent`,
	// `streamingDetectIntent` or [SessionEntityType][google.cloud.dialogflow.v2.SessionEntityType] management methods.
	SessionEntityTypes []*SessionEntityType `protobuf:"bytes,10,rep,name=session_entity_types,json=sessionEntityTypes,proto3" json:"session_entity_types,omitempty"`
}

func (x *WebhookResponse) Reset() {
	*x = WebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookResponse) ProtoMessage() {}

func (x *WebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookResponse.ProtoReflect.Descriptor instead.
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_webhook_proto_rawDescGZIP(), []int{1}
}

func (x *WebhookResponse) GetFulfillmentText() string {
	if x != nil {
		return x.FulfillmentText
	}
	return ""
}

func (x *WebhookResponse) GetFulfillmentMessages() []*Intent_Message {
	if x != nil {
		return x.FulfillmentMessages
	}
	return nil
}

func (x *WebhookResponse) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *WebhookResponse) GetPayload() *_struct.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *WebhookResponse) GetOutputContexts() []*Context {
	if x != nil {
		return x.OutputContexts
	}
	return nil
}

func (x *WebhookResponse) GetFollowupEventInput() *EventInput {
	if x != nil {
		return x.FollowupEventInput
	}
	return nil
}

func (x *WebhookResponse) GetSessionEntityTypes() []*SessionEntityType {
	if x != nil {
		return x.SessionEntityTypes
	}
	return nil
}

// Represents the contents of the original request that was passed to
// the `[Streaming]DetectIntent` call.
type OriginalDetectIntentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source of this request, e.g., `google`, `facebook`, `slack`. It is set
	// by Dialogflow-owned servers.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. The version of the protocol used for this request.
	// This field is AoG-specific.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Optional. This field is set to the value of the `QueryParameters.payload`
	// field passed in the request. Some integrations that query a Dialogflow
	// agent may provide additional information in the payload.
	//
	// In particular, for the Dialogflow Phone Gateway integration, this field has
	// the form:
	//
	//	<pre>{
	//	 "telephony": {
	//	   "caller_id": "+18558363987"
	//	 }
	//	}</pre>
	//
	// Note: The caller ID field (`caller_id`) will be redacted for Trial
	// Edition agents and populated with the caller ID in [E.164
	// format](https://en.wikipedia.org/wiki/E.164) for Essentials Edition agents.
	Payload *_struct.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *OriginalDetectIntentRequest) Reset() {
	*x = OriginalDetectIntentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalDetectIntentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalDetectIntentRequest) ProtoMessage() {}

func (x *OriginalDetectIntentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalDetectIntentRequest.ProtoReflect.Descriptor instead.
func (*OriginalDetectIntentRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *OriginalDetectIntentRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *OriginalDetectIntentRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *OriginalDetectIntentRequest) GetPayload() *_struct.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_google_cloud_dialogflow_v2_webhook_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_v2_webhook_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x2f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x7c, 0x0a, 0x1e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x1b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xef, 0x03, 0x0a, 0x0f, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x5d,
	0x0a, 0x14, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x13, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4c, 0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x14, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x75, 0x70, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x12, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x75, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x5f, 0x0a, 0x14, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x82, 0x01, 0x0a, 0x1b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x9b, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0xf8,
	0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_v2_webhook_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_v2_webhook_proto_rawDescData = file_google_cloud_dialogflow_v2_webhook_proto_rawDesc
)

func file_google_cloud_dialogflow_v2_webhook_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_v2_webhook_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_v2_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_v2_webhook_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_v2_webhook_proto_rawDescData
}

var file_google_cloud_dialogflow_v2_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_dialogflow_v2_webhook_proto_goTypes = []interface{}{
	(*WebhookRequest)(nil),              // 0: google.cloud.dialogflow.v2.WebhookRequest
	(*WebhookResponse)(nil),             // 1: google.cloud.dialogflow.v2.WebhookResponse
	(*OriginalDetectIntentRequest)(nil), // 2: google.cloud.dialogflow.v2.OriginalDetectIntentRequest
	(*QueryResult)(nil),                 // 3: google.cloud.dialogflow.v2.QueryResult
	(*Intent_Message)(nil),              // 4: google.cloud.dialogflow.v2.Intent.Message
	(*_struct.Struct)(nil),              // 5: google.protobuf.Struct
	(*Context)(nil),                     // 6: google.cloud.dialogflow.v2.Context
	(*EventInput)(nil),                  // 7: google.cloud.dialogflow.v2.EventInput
	(*SessionEntityType)(nil),           // 8: google.cloud.dialogflow.v2.SessionEntityType
}
var file_google_cloud_dialogflow_v2_webhook_proto_depIdxs = []int32{
	3, // 0: google.cloud.dialogflow.v2.WebhookRequest.query_result:type_name -> google.cloud.dialogflow.v2.QueryResult
	2, // 1: google.cloud.dialogflow.v2.WebhookRequest.original_detect_intent_request:type_name -> google.cloud.dialogflow.v2.OriginalDetectIntentRequest
	4, // 2: google.cloud.dialogflow.v2.WebhookResponse.fulfillment_messages:type_name -> google.cloud.dialogflow.v2.Intent.Message
	5, // 3: google.cloud.dialogflow.v2.WebhookResponse.payload:type_name -> google.protobuf.Struct
	6, // 4: google.cloud.dialogflow.v2.WebhookResponse.output_contexts:type_name -> google.cloud.dialogflow.v2.Context
	7, // 5: google.cloud.dialogflow.v2.WebhookResponse.followup_event_input:type_name -> google.cloud.dialogflow.v2.EventInput
	8, // 6: google.cloud.dialogflow.v2.WebhookResponse.session_entity_types:type_name -> google.cloud.dialogflow.v2.SessionEntityType
	5, // 7: google.cloud.dialogflow.v2.OriginalDetectIntentRequest.payload:type_name -> google.protobuf.Struct
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_v2_webhook_proto_init() }
func file_google_cloud_dialogflow_v2_webhook_proto_init() {
	if File_google_cloud_dialogflow_v2_webhook_proto != nil {
		return
	}
	file_google_cloud_dialogflow_v2_context_proto_init()
	file_google_cloud_dialogflow_v2_intent_proto_init()
	file_google_cloud_dialogflow_v2_session_proto_init()
	file_google_cloud_dialogflow_v2_session_entity_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_dialogflow_v2_webhook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalDetectIntentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dialogflow_v2_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_v2_webhook_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_v2_webhook_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_v2_webhook_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_v2_webhook_proto = out.File
	file_google_cloud_dialogflow_v2_webhook_proto_rawDesc = nil
	file_google_cloud_dialogflow_v2_webhook_proto_goTypes = nil
	file_google_cloud_dialogflow_v2_webhook_proto_depIdxs = nil
}
