import Layout from "../src/components/Layout";
import Contact from "../src/pages/Contact";

export default function ContactPage({ ...props }) {
  return (
    <Layout pageTitle="Contact" {...props}>
      <Contact />
    </Layout>
  );
}

// re-export the reusable `getServerSideProps` function
export { getServerSideProps } from "../src/Chakra";
