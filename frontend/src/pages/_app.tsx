import type { AppProps } from 'next/app';

import { MantineProvider } from '@mantine/core';
import { JetBrains_Mono as JetBrainsMono } from 'next/font/google';
import Head from 'next/head';

import '@/styles/globals.css';

const jetBrainsMono = JetBrainsMono({
  display: 'block',
  subsets: ['latin'],
});

const App = ({ Component, pageProps }: AppProps) => (
  <>
    <Head>
      <title>CoinRank</title>
      <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
      <link rel="icon" href="/favicon.ico" />
    </Head>

    <MantineProvider
      withGlobalStyles
      withNormalizeCSS
      theme={{
        colorScheme: 'dark',
        fontFamily: jetBrainsMono.style.fontFamily,
        headings: {
          fontFamily: jetBrainsMono.style.fontFamily,
        },
      }}
    >
      <Component {...pageProps} />
    </MantineProvider>
  </>
);

export default App;
