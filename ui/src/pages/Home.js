import { Heading, Stack, Box } from "@chakra-ui/react";
import { useIntl } from "react-intl";
import { gql, useMutation } from "@apollo/client";
import { Input, Button } from "@chakra-ui/react";
import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  FormHelperText,
} from "@chakra-ui/react";

const CREATE_TODO_MUTATION = gql`
  mutation createTodo($title: String!) {
    createTodo(title: $title) {
      id
      title
      createdAt
      done
    }
  }
`;

export default function Home(props) {
  const { formatMessage } = useIntl();
  const f = (id) => formatMessage({ id });

  const [createTodo, { loading }] = useMutation(CREATE_TODO_MUTATION);

  const edges = props?.data?.todos?.edges ?? [];

  return (
    <Stack
      spacing={"3rem"}
      justify="center"
      align={"center"}
      shouldWrapChildren
      maxW={800}
    >
      <Heading>{f("hello")}</Heading>

      <FormControl id="todo" isRequired onSubmit={() => console.log("submit")}>
        <FormLabel>Todo</FormLabel>
        <Input placeholder="What needs to be done?" />
        <Button mt={4} colorScheme="teal" isLoading={loading} type="submit">
          Submit
        </Button>
      </FormControl>

      <TodoList edges={edges} />
    </Stack>
  );
}

const TodoList = ({ edges }) => {
  return (
    <Box>
      {edges.map((e, i) => (
        <TodoRow key={i} {...e} />
      ))}
    </Box>
  );
};

const TodoRow = ({ cursor, node: { id, createdAt, title, done } }) => {
  return <Box>{title}</Box>;
};
