import React from 'react';
import ReactDOM from 'react-dom';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from 'react-apollo';
import { ApolloProvider as ApolloProviderHooks } from 'react-apollo-hooks';

import './index.css';
import { App } from './App';
import * as serviceWorker from './serviceWorker';

const client = new ApolloClient({
  // uri: 'https://www.graphqlhub.com/graphql',
  uri: 'http://localhost:8080/graphql',
});

const AppContainer = (): JSX.Element => (
  <ApolloProvider client={client}>
    <ApolloProviderHooks client={client}>
      <App />
    </ApolloProviderHooks>
  </ApolloProvider>
);

ReactDOM.render(<AppContainer />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
