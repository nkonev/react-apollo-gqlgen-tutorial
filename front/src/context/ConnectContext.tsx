import { createContext, FC, useEffect, useState } from "react";
import { WebSocketLink } from "@apollo/client/link/ws";
import { ApolloClient, ApolloProvider, HttpLink, InMemoryCache, NormalizedCacheObject, split } from "@apollo/client";
import { getMainDefinition } from "@apollo/client/utilities";


interface ContextTypes {
  reconnect: () => void
}

export const ConnectContext = createContext<Partial<ContextTypes>>({})

export const ConnectProvider: FC = ({ children }) => {
  const [state, setState] = useState<ApolloClient<NormalizedCacheObject>>(connect)
  const reconnect = () => setState(connect)
  return(
    <ConnectContext.Provider value={{
      reconnect
    }}>
      <ApolloProvider client={state}>
        { children }
      </ApolloProvider>
    </ConnectContext.Provider>
  )
}

const connect = () => {
  const httpLink = new HttpLink({
    uri: 'http://localhost:2000/graphql',
    credentials: 'include',
  });

  const wsLink = new WebSocketLink({
    uri: 'ws://localhost:2000/graphql',
    options: {
      reconnect: true,
    }
  });

  const splitLink = split(
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === 'OperationDefinition' &&
        definition.operation === 'subscription'
      );
    },
    wsLink,
    httpLink,
  );

  return new ApolloClient({
    link: splitLink,
    cache: new InMemoryCache(),
  });
}