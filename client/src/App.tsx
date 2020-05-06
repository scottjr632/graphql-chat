import React from 'react';
import { InMemoryCache } from 'apollo-boost';
import { ApolloClient } from 'apollo-client';
import { ApolloProvider } from '@apollo/react-hooks';
import { split } from 'apollo-link';
import { HttpLink } from 'apollo-link-http';
import { WebSocketLink } from 'apollo-link-ws';
import { getMainDefinition } from 'apollo-utilities';

import GData from './GData'
import './App.css';

// Create an http link:
const httpLink = new HttpLink({
  uri: 'http://localhost:8080/graphql'
});

// Create a WebSocket link:
const wsLink = new WebSocketLink({
  uri: `ws://localhost:8080/graphql`,
  options: {
    reconnect: true
  }
});


const cache = new InMemoryCache();

// using the ability to split links, you can send data to each link
// depending on what kind of operation is being sent
const link = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,
  httpLink
);

const client = new ApolloClient({ link, cache });

function App() {
  return (
    <ApolloProvider client={client}>
      <GData />
    </ApolloProvider>
  );
}

export default App;
