import { IntlProvider } from "react-intl";
import { useRouter } from "next/router";
import * as locales from "../content/locale";
import "../styles/globals.css";

import { ApolloProvider } from "@apollo/client";
import { useApollo } from "../src/graphql";

function MyApp({ Component, pageProps }) {
  // Apollo Client for GraphQL
  const apolloClient = useApollo(pageProps);

  // i18n
  const router = useRouter();
  const { locale, defaultLocale, pathname } = router;
  const localeCopy = locales[locale];
  const messages = { ...localeCopy.global, ...localeCopy[pathname] };
  return (
    <ApolloProvider client={apolloClient}>
      <IntlProvider
        locale={locale}
        defaultLocale={defaultLocale}
        messages={messages}
      >
        <Component {...pageProps} />
      </IntlProvider>
    </ApolloProvider>
  );
}

export default MyApp;
