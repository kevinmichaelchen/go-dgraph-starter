import { Heading, Stack, Box } from "@chakra-ui/react";

export default function About() {
  return (
    <Stack spacing={2} maxW={500} lineHeight="taller">
      <Heading>About</Heading>
      <Box>This site is about X.</Box>
    </Stack>
  );
}
