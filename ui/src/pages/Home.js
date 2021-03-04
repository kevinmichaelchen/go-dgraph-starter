import { Heading, Stack, Box } from "@chakra-ui/react";
import { useIntl } from "react-intl";

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
