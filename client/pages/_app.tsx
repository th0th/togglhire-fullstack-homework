import 'focus-visible';
import type { AppProps } from 'next/app';
import Head from 'next/head';
import React from 'react';
import { Provider } from 'urql';
import '../styles/globals.css';
import client from '../utils/client';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <Provider value={client}>
      <Head>
        <title>
          Toggl Hire Homework
        </title>

        <link
          href="https://public-assets.toggl.com/b/assets/@toggl/hire/images/5e8dacfc3d0e4236732aae90_favicon-32x32.9e0c302a.png"
          rel="shortcut icon"
          type="image/x-icon"
        />
      </Head>

      <Component {...pageProps} />
    </Provider>
  );
}
