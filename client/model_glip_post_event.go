/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ringcentral

import (
	"time"
)

type GlipPostEvent struct {
	// Internal identifier of a post
	Id string `json:"id,omitempty"`
	// Type of a post event
	EventType string `json:"eventType,omitempty"`
	// Internal identifier of a group a post belongs to
	GroupId string `json:"groupId,omitempty"`
	// Type of a post. 'TextMessage' - an incoming text message; 'PersonJoined' - a message notifying that person has joined a conversation; 'PersonsAdded' - a message notifying that a person(s) were added to a conversation
	Type string `json:"type,omitempty"`
	// For 'TextMessage' post type only. Message text
	Text string `json:"text,omitempty"`
	// Internal identifier of a user - author of a post
	CreatorId string `json:"creatorId,omitempty"`
	// For PersonsAdded post type only. Identifiers of persons added to a group
	AddedPersonIds []string `json:"addedPersonIds,omitempty"`
	// For PersonsRemoved post type only. Identifiers of persons removed from a group
	RemovedPersonIds []string `json:"removedPersonIds,omitempty"`
	// List of at mentions in post text with names.
	Mentions []GlipMentionsInfo `json:"mentions,omitempty"`
	// Post creation datetime in ISO 8601 format
	CreationTime time.Time `json:"creationTime,omitempty"`
	// Post last change datetime in ISO 8601 format
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
}
