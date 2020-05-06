import { gql } from "apollo-boost";

export default gql`
subscription onMessageCreation($channelName: String!) {
  messageCreated(channelName: $channelName) {
		id
    content
    channel {
      id
      name
    }
  }
}
`