import { useColorMode, Flex, IconButton } from "@chakra-ui/react";
import { SunIcon, MoonIcon } from "@chakra-ui/icons";

export default function ThemeToggler({ ...props }) {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <Flex textAlign="right" {...props}>
      <IconButton
        isRound
        aria-label="Theme toggler"
        size="lg"
        icon={colorMode === "light" ? <MoonIcon /> : <SunIcon />}
        onClick={toggleColorMode}
        variant="ghost"
      />
    </Flex>
  );
}
