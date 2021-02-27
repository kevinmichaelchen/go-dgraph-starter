import Layout from "../src/components/Layout";
import About from "../src/pages/About";

export default function AboutPage({ ...props }) {
  return (
    <Layout pageTitle="About" {...props}>
      <About />
    </Layout>
  );
}

// re-export the reusable `getServerSideProps` function
export { getServerSideProps } from "../src/Chakra";
