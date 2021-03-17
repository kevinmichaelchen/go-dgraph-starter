import { Button, IconButton, Box } from "@chakra-ui/react";
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
import { Mutations } from "../../graphql";

export const Nuke = () => {
  const [nukeData, { loading, error }] = useMutation(Mutations.NUKE);

  const { isOpen, onOpen, onClose } = useDisclosure();

  const handleNukeClick = () => {
    nukeData();
    onClose();

    // TODO Todo List is not currently re-rendering. You'd expect it to re-render with zero items.
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
