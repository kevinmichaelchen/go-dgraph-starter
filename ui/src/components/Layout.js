import Navbar from "./Navbar";
import { Box, Flex } from "@chakra-ui/react";
import Head from "./Head";
import Footer from "./Footer";
import { Chakra } from "../Chakra";

const Layout = ({ children, pageTitle, cookies }) => {
  return (
    <Chakra cookies={cookies}>
      <Box>
        <Head pageTitle={pageTitle} />
        <Navbar />
        <Flex
          minH="100vh"
          p="0 0.5rem"
          direction="column"
          justifyContent="center"
          alignItems="center"
        >
          <Flex
            direction="column"
            justifyContent="center"
            alignItems="center"
            p="5rem 0"
            maxW={800}
            flex={1}
          >
            {children}
          </Flex>
          <Footer />
        </Flex>
      </Box>
    </Chakra>
  );
};

export default Layout;
