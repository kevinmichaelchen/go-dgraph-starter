import NextHead from "next/head";
import GoogleFonts from "next-google-fonts";
import { headingFont, bodyFont } from "../Chakra";

export const siteData = {
  siteTitle: "Site Title",
  siteDescription: "Site Description",
  siteImage: "https://live.staticflickr.com/20/70342819_0c9fec17b4_b.jpg",
  twitterHandle: "@yourTwitterHandle",
};

const Head = ({ pageTitle }) => {
  const { siteTitle, siteDescription, siteImage, twitterHandle } = siteData;
  const title = pageTitle ? `${siteTitle} - ${pageTitle}` : siteTitle;
  return (
    <>
      <GoogleFonts
        href={`https://fonts.googleapis.com/css2?family=${headingFont}:wght@400;700&display=swap`}
      />
      <GoogleFonts
        href={`https://fonts.googleapis.com/css2?family=${bodyFont}:wght@400;700&display=swap`}
      />
      <NextHead>
        <title>{title}</title>
        <link rel="icon" href="/favicon.ico" />

        {/* Twitter */}
        <meta name="twitter:card" content="summary" key="twcard" />
        <meta name="twitter:creator" content={twitterHandle} key="twhandle" />

        {/* Open Graph */}
        {/*<meta property="og:url" content={currentURL} key="ogurl" />*/}
        {/*https://search.creativecommons.org/photos/464ae86e-52e2-40c9-abf0-0e29ef985bb1*/}
        <meta property="og:image" content={siteImage} key="ogimage" />
        <meta property="og:site_name" content={siteTitle} key="ogsitename" />
        <meta property="og:title" content={siteTitle} key="ogtitle" />
        <meta
          property="og:description"
          content={siteDescription}
          key="ogdesc"
        />
      </NextHead>
    </>
  );
};

export default Head;
