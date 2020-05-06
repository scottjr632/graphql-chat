import { gql } from 'apollo-boost';
export default gql`
query Channels {
	channels {
    name
    id
  }
}
`