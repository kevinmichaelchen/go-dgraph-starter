import { Heading, Stack, Box } from "@chakra-ui/react";
import { useIntl } from "react-intl";

export default function Home() {
  const { formatMessage } = useIntl();
  const f = (id) => formatMessage({ id });
  return (
    <Stack
      spacing={"3rem"}
      justify="center"
      align={"center"}
      shouldWrapChildren
      maxW={800}
    >
      <Heading>{f("hello")}</Heading>
      <Box>Your content</Box>
    </Stack>
  );
}
