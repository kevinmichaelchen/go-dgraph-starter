import Layout from "../src/components/Layout";
import Home from "../src/pages/Home";

export default function HomePage({ ...props }) {
  return (
    <Layout {...props}>
      <Home />
    </Layout>
  );
}

// re-export the reusable `getServerSideProps` function
export { getServerSideProps } from "../src/Chakra";
