import Layout from "../src/components/Layout";
import Home from "../src/pages/Home";

import { initializeApollo, addApolloState } from "../src/graphql";
import { GET_TODOS_QUERY } from "../src/graphql/gql";

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
    query: GET_TODOS_QUERY,
    variables: {
      first: 10,
      after: "",
    },
  });

  console.log(JSON.stringify(res, null, 2));

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
