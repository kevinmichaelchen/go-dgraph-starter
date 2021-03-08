import { Heading, Stack, Box } from "@chakra-ui/react";
import { useIntl } from "react-intl";
import { gql, useMutation, useQuery, NetworkStatus } from "@apollo/client";
import { GET_TODOS_QUERY } from "../../src/graphql/gql";
import { Input, Button } from "@chakra-ui/react";
import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  FormHelperText,
} from "@chakra-ui/react";
import { Formik, Field, Form } from "formik";

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

const pageSize = 10;

export default function Home(props) {
  const { formatMessage } = useIntl();
  const f = (id) => formatMessage({ id });

  const { loading, error, data, fetchMore, networkStatus } = useQuery(
    GET_TODOS_QUERY,
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

  const loadingMoreTodos = networkStatus === NetworkStatus.fetchMore;

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
      <Heading>{f("hello")}</Heading>

      <CreateTodoForm loadMoreTodos={loadMoreTodos} />

      <TodoList edges={edges} />
    </Stack>
  );
}

const CreateTodoForm = ({ loadMoreTodos }) => {
  const [createTodo, { loading }] = useMutation(CREATE_TODO_MUTATION);
  const FormFields = {
    title: "title",
  };
  function validateName(value) {
    let error;
    if (!value) {
      error = "Value is required";
    }
    return error;
  }
  return (
    <Formik
      initialValues={{
        [FormFields.title]: "",
      }}
      onSubmit={(values, actions) => {
        const { setSubmitting } = actions;
        setSubmitting(false);

        const title = values[FormFields.title];

        // Issue the GraphQL mutation that creates a new Todo
        createTodo({
          variables: { title },
          // TODO this "update cache" function isn't working properly since newly created todo is not being rendered
          update: (cache, { data: { createTodo } }) => {
            cache.modify({
              fields: {
                allTodos(existingTodos = []) {
                  const newTodoRef = cache.writeFragment({
                    data: createTodo,
                    fragment: gql`
                      fragment NewTodo on allTodos {
                        id
                        type
                      }
                    `,
                  });
                  return [newTodoRef, ...existingTodos];
                },
              },
            });
          },
        });

        // Render the Todo we just created
        loadMoreTodos();
      }}
    >
      {(props) => (
        <Form>
          <Field name={FormFields.title} validate={validateName}>
            {({ field, form }) => (
              <FormControl
                isInvalid={
                  form.errors[FormFields.title] &&
                  form.touched[FormFields.title]
                }
              >
                <FormLabel htmlFor={FormFields.title}>Todo</FormLabel>
                <Input
                  {...field}
                  id={FormFields.title}
                  placeholder="What needs to be done?"
                />
                <FormErrorMessage>
                  {form.errors[FormFields.title]}
                </FormErrorMessage>
              </FormControl>
            )}
          </Field>
          <Button mt={4} colorScheme="teal" isLoading={loading} type="submit">
            Submit
          </Button>
        </Form>
      )}
    </Formik>
  );
};

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
