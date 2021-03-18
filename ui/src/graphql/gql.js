import { gql } from "@apollo/client";

const TodoFragment = gql`
  fragment TodoFields on Todo {
    id
    createdAt
    title
    done
  }
`;

const GET_TODOS_QUERY = gql`
  ${TodoFragment}
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
          ...TodoFields
        }
      }
    }
  }
`;

const NUKE_DATA_MUTATION = gql`
  mutation {
    nuke
  }
`;

export const Queries = {
  GET_TODOS_QUERY,
};

export const Fragments = {
  TodoFragment,
};

export const Mutations = {
  NUKE: NUKE_DATA_MUTATION,
};
