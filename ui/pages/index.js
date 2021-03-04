import Layout from "../src/components/Layout";
import Home from "../src/pages/Home";

import { initializeApollo, addApolloState } from "../src/graphql";
import { gql } from "@apollo/client";

export default function HomePage({ ...props }) {
  return (
    <Layout {...props}>
      <Home {...props} />
    </Layout>
  );
}

const getTodosQuery = gql`
  query GetTodos {
    todos {
      totalCount
      pageInfo {
        endCursor
        hasNextPage
      }
      edges {
        cursor
        node {
          id
          createdAt
          title
          done
        }
      }
    }
  }
`;

export async function getServerSideProps({ req }) {
  const apolloClient = initializeApollo();

  const res = await apolloClient.query({
    query: getTodosQuery,
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
      revalidate: 1,
    },
  });
}
