import {
  Stack,
  Box,
  VStack,
  StackDivider,
  IconButton,
  Spinner,
  useToast,
} from "@chakra-ui/react";
import { DeleteIcon } from "@chakra-ui/icons";
import { gql, useMutation } from "@apollo/client";

const TodoList = ({ edges }) => {
  return (
    <VStack spacing={4} divider={<StackDivider borderColor="gray.200" />}>
      {edges.map((e, i) => (
        <TodoRow key={i} {...e} />
      ))}
    </VStack>
  );
};

const DELETE_TODO_MUTATION = gql`
  mutation deleteTodo($id: String!) {
    deleteTodo(id: $id) {
      success
    }
  }
`;

const TodoRow = ({ cursor, node: { id, createdAt, title, done } }) => {
  const [deleteTodo, { loading, error }] = useMutation(DELETE_TODO_MUTATION);
  const toast = useToast();

  if (loading) {
    return <Spinner color="red.500" size="xl" />;
  }

  if (error) {
    toast({
      title: "Oops.",
      position: "top-right",
      description: JSON.stringify(error),
      status: "error",
      duration: 4000,
      isClosable: true,
    });
  }

  // https://stackoverflow.com/questions/63192774/apollo-client-delete-item-from-cache
  const handleDeleteClick = () => {
    deleteTodo({
      variables: { id },
      update(cache) {
        cache.modify({
          fields: {
            allTodos(existingTodos, { readField }) {
              return existingTodos.filter(
                (todo) => id !== readField("id", todo)
              );
            },
          },
        });
      },
    });
  };

  return (
    <Stack isInline justify="space-between" align="center" spacing={5}>
      <Box>{title}</Box>

      <IconButton
        colorScheme="red"
        size="sm"
        aria-label="Delete"
        icon={<DeleteIcon />}
        onClick={handleDeleteClick}
      />
    </Stack>
  );
};

export default TodoList;
