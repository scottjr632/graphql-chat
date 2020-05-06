/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: CreateMessage
// ====================================================

export interface CreateMessage_createMessage {
  __typename: "Message";
  id: string;
  content: string;
}

export interface CreateMessage {
  createMessage: CreateMessage_createMessage;
}

export interface CreateMessageVariables {
  content: string;
  channelName: string;
}
