import { gql } from "@apollo/client";

export const GET_TODOS_QUERY = gql`
  query GetTodos($first: Int!, $after: String!) {
    todos(first: $first, after: $after) {
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
