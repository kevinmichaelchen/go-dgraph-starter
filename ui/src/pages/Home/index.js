import { Heading, Stack, Box } from "@chakra-ui/react";
import { useIntl } from "react-intl";
import { useQuery, NetworkStatus } from "@apollo/client";
import { Queries } from "../../graphql";
import CreateTodoForm from "./CreateTodoForm";
import TodoList from "./TodoList";
import { Nuke } from "./Nuke";

const pageSize = 10;

export default function Home(props) {
  // Hooks for i18n
  const { formatMessage } = useIntl();
  const f = (id) => formatMessage({ id });

  // Hooks to query a page of Todos
  const { loading, error, data, fetchMore, networkStatus } = useQuery(
    Queries.GET_TODOS_QUERY,
    {
      variables: {
        first: pageSize,
        after: "",
      },
      // Setting this value to true will make the component rerender when
      // the "networkStatus" changes, so we are able to know if it is fetching
      // more data
      notifyOnNetworkStatusChange: true,
    }
  );

  // Whether we're "fetching more" Todos (beyond the initial page)
  const loadingMoreTodos = networkStatus === NetworkStatus.fetchMore;

  // The end cursor, so we can continue to paginate when we successfully create a Todo
  const endCursor = data?.todos?.pageInfo?.endCursor ?? "";

  const loadMoreTodosFactory = (endCursor) => () => {
    console.log(`Fetching more Todos after cursor ${endCursor}`);
    fetchMore({
      variables: {
        first: pageSize,
        after: endCursor,
      },
    });
  };
  const loadMoreTodos = loadMoreTodosFactory(endCursor);

  if (error) {
    return <Box>Failed to load posts: {JSON.stringify(error)}</Box>;
  }

  if (loading && !loadingMoreTodos) {
    return <Box>Loading...</Box>;
  }

  const dataEdges = data?.todos?.edges ?? [];
  const propEdges = props?.data?.todos?.edges ?? [];
  console.log("dataEdges", dataEdges);
  console.log("propEdges", propEdges);

  const edges = dataEdges || propEdges;

  return (
    <Stack
      spacing={"3rem"}
      justify="center"
      align={"center"}
      shouldWrapChildren
      maxW={800}
    >
      <Nuke />
      <Heading>{f("hello")}</Heading>

      <CreateTodoForm loadMoreTodos={loadMoreTodos} />

      <TodoList edges={edges} />
    </Stack>
  );
}
