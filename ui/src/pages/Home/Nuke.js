import { Button, Box, useToast } from "@chakra-ui/react";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
} from "@chakra-ui/react";
import { useMutation } from "@apollo/client";
import { Mutations, Queries } from "../../graphql";

export const Nuke = () => {
  const [nukeData, { loading, error }] = useMutation(Mutations.NUKE);
  const toast = useToast();

  const { isOpen, onOpen, onClose } = useDisclosure();

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

  const handleNukeClick = () => {
    // Perform the GraphQL mutation to drop data.
    // The update field should instruct Apollo to modify its cache,
    // overwriting its internal data with an empty object.
    nukeData({
      // An array or function that allows you to specify which queries
      // you want to refetch after a mutation has occurred.
      // Array values can either be queries (with optional variables)
      // or just the string names of queries to be refeteched.
      refetchQueries: [
        {
          query: Queries.GET_TODOS_QUERY,
          variables: {
            first: 10,
            after: "",
          },
        },
      ],
      // A function used to update the cache after a mutation occurs
      update: (cache) => {
        cache.modify({
          fields: {
            allTodos() {
              return {
                edges: [],
                pageInfo: {
                  hasPreviousPage: false,
                  hasNextPage: true,
                  startCursor: "",
                  endCursor: "",
                },
              };
            },
          },
        });
      },
    });

    // Close the modal
    onClose();
  };

  return (
    <Box>
      <Button
        colorScheme="red"
        size="md"
        aria-label="Open modal to delete all Todos"
        onClick={onOpen}
      >
        Delete all Todos
      </Button>

      <Modal onClose={onClose} isOpen={isOpen} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Delete all Todos</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            Caution. This will delete all Todos in the database.
          </ModalBody>
          <ModalFooter>
            <Button onClick={onClose}>Close</Button>
            <Button
              colorScheme="red"
              size="md"
              aria-label="Delete all Todos"
              onClick={handleNukeClick}
            >
              Proceed
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
};
