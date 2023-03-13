import type { AppProps } from 'next/app';

import { Box, MantineProvider } from '@mantine/core';
import { JetBrains_Mono as JetBrainsMono } from 'next/font/google';

import Header from '@/components/Header';
import MetaTags from '@/components/MetaTags';
import '@/styles/globals.css';

const jetBrainsMono = JetBrainsMono({
  display: 'block',
  subsets: ['latin'],
});

const App = ({ Component, pageProps }: AppProps) => (
  <>
    <MetaTags />
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
      <Box
        component="main"
        sx={(theme) => ({
          margin: '0 auto',
          maxWidth: theme.breakpoints.sm,
          padding: theme.spacing.md,
          display: 'flex',
          flexDirection: 'column',
          gap: theme.spacing.md,
        })}
      >
        <Header />
        <Component {...pageProps} />
      </Box>
    </MantineProvider>
  </>
);

export default App;
