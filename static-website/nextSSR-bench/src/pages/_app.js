import '@/styles/globals.css'

export default function App({ Component, pageProps }) {
  return <Component {...pageProps} />
}

export const config = {
  runtime: "nodejs",
  unstable_runtimeJS: false,
};