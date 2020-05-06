/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL subscription operation: onMessageCreation
// ====================================================

export interface onMessageCreation_messageCreated_channel {
  __typename: "Channel";
  id: string;
  name: string;
}

export interface onMessageCreation_messageCreated {
  __typename: "Message";
  id: string;
  content: string;
  channel: onMessageCreation_messageCreated_channel;
}

export interface onMessageCreation {
  messageCreated: onMessageCreation_messageCreated;
}

export interface onMessageCreationVariables {
  channelName: string;
}
