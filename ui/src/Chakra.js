import {
  ChakraProvider,
  cookieStorageManager,
  localStorageManager,
  theme,
} from "@chakra-ui/react";

export const headingFont = "Chelsea Market";
export const bodyFont = "Roboto";

const myTheme = {
  ...theme,
  fonts: {
    ...theme.fonts,
    heading: headingFont + `, ` + theme.fonts.heading,
    body: bodyFont + `, ` + theme.fonts.body,
  },
};

export function Chakra({ cookies, children }) {
  // b) Pass `colorModeManager` prop
  const colorModeManager = cookies
    ? cookieStorageManager(cookies)
    : localStorageManager;

  return (
    <ChakraProvider theme={myTheme} colorModeManager={colorModeManager}>
      {children}
    </ChakraProvider>
  );
}

// also export a reusable function getServerSideProps
export function getServerSideProps({ req }) {
  return {
    props: {
      // first time users will not have any cookies and you may not return
      // undefined here, hence ?? is necessary
      cookies: req.headers.cookie ?? "",
    },
  };
}
