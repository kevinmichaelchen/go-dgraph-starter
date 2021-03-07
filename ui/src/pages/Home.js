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

export default function Home(props) {
  const { formatMessage } = useIntl();
  const f = (id) => formatMessage({ id });

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

      <CreateTodoForm />

      <TodoList edges={edges} />
    </Stack>
  );
}

const CreateTodoForm = () => {
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
        console.log("values", values);

        const title = values[FormFields.title];

        createTodo({
          variables: { title },
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
