/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Messages
// ====================================================

export interface Messages_messages_channel {
  __typename: "Channel";
  id: string;
  name: string;
}

export interface Messages_messages {
  __typename: "Message";
  id: string;
  content: string;
  channel: Messages_messages_channel;
}

export interface Messages {
  messages: Messages_messages[] | null;
}

export interface MessagesVariables {
  channelName: string;
}
