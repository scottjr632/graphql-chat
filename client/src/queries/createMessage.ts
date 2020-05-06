import { gql } from 'apollo-boost';

export default gql`
mutation CreateMessage($content: String!, $channelName: String!) {
  createMessage(content: $content, channelName: $channelName) {
    id
    content
  }
}
`