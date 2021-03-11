import { gql, useMutation } from "@apollo/client";
import { Input, Button, useToast, Spinner } from "@chakra-ui/react";
import { FormControl, FormLabel, FormErrorMessage } from "@chakra-ui/react";
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

const CreateTodoForm = ({ loadMoreTodos }) => {
  const [createTodo, { loading, error }] = useMutation(CREATE_TODO_MUTATION);
  const toast = useToast();

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
                        title
                        createdAt
                        done
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

export default CreateTodoForm;
