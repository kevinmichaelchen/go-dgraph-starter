import { Box } from "@chakra-ui/react";

const hoverStyles = {
  borderColor: "#b7e4cf",
  // color: "#355e4b",
  color: "#36473f",
  borderWidth: "2px",
  fontWeight: "bolder",
};

const Card = ({ children, ...restProps }) => {
  return (
    <Box
      m={"1rem"}
      p={"1.5rem"}
      textAlign={"left"}
      color={"inherit"}
      textDecoration={"none"}
      border={"1px solid #eaeaea"}
      borderRadius={"10px"}
      transition={"color 0.15s ease, border-color 0.15s ease"}
      _hover={hoverStyles}
      _active={hoverStyles}
      _focus={hoverStyles}
      {...restProps}
    >
      {children}
    </Box>
  );
};

export default Card;
