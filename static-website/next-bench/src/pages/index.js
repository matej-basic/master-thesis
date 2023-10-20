import BenchmarkerNoSSR from "@/components/benchmarkernossr";

export default function Home(props) {
  return (
    <BenchmarkerNoSSR/>
  )
}

export const config = {
  runtime: "nodejs",
  unstable_runtimeJS: false,
};