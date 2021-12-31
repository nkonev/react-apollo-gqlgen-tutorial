import { createContext, FC, useContext, useEffect } from "react";
import { useQuery, gql, useSubscription } from "@apollo/client";
import { ConnectContext } from "./ConnectContext";


type TAuth = {
  authorized: boolean
  method: string
  reconnect: boolean
}

interface ContextTypes {
  auth: TAuth
  logout: () => void
}

export const AuthorizationContext = createContext<Partial<ContextTypes>>({})

export const AuthorizationProvider: FC = ({ children }) => {
  const { reconnect } = useContext(ConnectContext)

  // GET запрос Auth
  const { loading, error, data } = useQuery(QUERY);
  useEffect(() => {
    if (!loading && !error && data) {
      const { reconnect: rc } = data.auth
      if (rc === true && reconnect) {
        reconnect()
      }
    }
  }, [loading, error, data])

  // Слушаем Auth websocket
  const { loading: wsLoading, error: wsError, data: wsData } = useSubscription(
    SUBSCR
  );
  useEffect(() => {
    console.log("Websocket Auth", wsLoading, wsError, wsData)
  }, [wsLoading, wsError, wsData])

  return (
    <AuthorizationContext.Provider value={{

    }}>
      {children}
    </AuthorizationContext.Provider>
  )
}

const SUBSCR = gql`
    subscription{
        authSubscription{
            authorized,
            method,
            reconnect
        }
    }
`;

const QUERY = gql`
    query {
        auth{
            authorized,
            method,
            reconnect
        }
    }
`;
