import {
  Flex,
  Image,
  Stack,
  Text,
  Link as ChakraLink,
  Box,
  useColorMode,
} from "@chakra-ui/react";

const Footer = () => {
  // The SVG is permanently black which makes it really hard to see in Dark Mode.
  // I don't want to dynamically change the fill color so I can keep caching the asset.
  // The best option is just to underlay a light gray/white background beneath the text and SVG.
  const { colorMode } = useColorMode();
  const darkMode = colorMode !== "light";

  return (
    <Flex
      justify="center"
      align="center"
      w={"100%"}
      h={"100px"}
      borderTop={"1px solid #eaeaea"}
    >
      <ChakraLink
        href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
        isExternal
      >
        <Stack isInline align={"center"}>
          <Flex justify="center" align="center">
            <Text>Powered by </Text>
            {darkMode && (
              <Image
                src="/vercel-white.svg"
                alt="Vercel Logo"
                ml={"0.5rem"}
                boxSize={20}
              />
            )}
            {!darkMode && (
              <Image
                src="/vercel.svg"
                alt="Vercel Logo"
                ml={"0.5rem"}
                boxSize={20}
              />
            )}
          </Flex>
        </Stack>
      </ChakraLink>
    </Flex>
  );
};

export default Footer;
