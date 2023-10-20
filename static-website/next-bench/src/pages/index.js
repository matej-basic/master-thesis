import dynamic from "next/dynamic";

const BenchmarkerNoSSR = dynamic(() => import("../components/benchmarkernossr"), {ssr: false});

export default function Home(props) {
  return (
    <>
    <BenchmarkerNoSSR/>
   </>
  )
}

export const config = {
  runtime: "nodejs",
  unstable_runtimeJS: false,
};