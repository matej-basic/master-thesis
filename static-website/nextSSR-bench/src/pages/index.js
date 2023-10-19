
import Benchmarker from "./components/benchmarker";

export default function Home(props) {
  return (
    <>
    <Benchmarker props={props}/>
   </>
  )
}

export const getServerSideProps = () => {
  return { props: { message: new Date().toJSON() } };
};

export const config = {
  runtime: "nodejs",
  unstable_runtimeJS: false,
};