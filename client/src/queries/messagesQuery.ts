import { gql } from "apollo-boost";

export default gql`
query Messages($channelName: String!) {
  messages(channelName: $channelName) {
		id
    content
    channel {
      id
      name
    }
  }
}
`