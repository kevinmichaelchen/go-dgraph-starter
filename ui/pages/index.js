import Layout from "../src/components/Layout";
import Home from "../src/pages/Home";

import { initializeApollo, addApolloState } from "../src/graphql";
import { Queries } from "../src/graphql/gql";

export default function HomePage({ ...props }) {
  return (
    <Layout {...props}>
      <Home {...props} />
    </Layout>
  );
}

export async function getServerSideProps({ req }) {
  const apolloClient = initializeApollo();

  const res = await apolloClient.query({
    query: Queries.GET_TODOS_QUERY,
    variables: {
      first: 10,
      after: "",
    },
  });

  return addApolloState(apolloClient, {
    // will be passed to the page component as props
    props: {
      // ChakraUI stores color mode info in cookies.
      // First-time users will not have any cookies,
      // and returning undefined would be invalid.
      cookies: req.headers.cookie ?? "",
      data: res.data,
    },
  });
}
