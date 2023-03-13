import Head from 'next/head';

const MetaTags = () => (
  <Head>
    <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
    <link
      rel="icon"
      href="favicon.ico"
      media="(prefers-color-scheme: light)"
    />
    <link
      rel="icon"
      href="favicon-dark.ico"
      media="(prefers-color-scheme: dark)"
    />
    <link
      rel="icon"
      href="favicon.svg"
      type="image/svg+xml"
    />

    <title>CoinRank</title>
    <meta name="title" content="CoinRank" />
    <meta name="description" content='A powerful tool that allows users to rate and evaluate cryptocurrencies using the familiar "upvote" and "downvote" system.' />

    <meta property="og:type" content="website" />
    <meta property="og:url" content="https://coinrank.ptr.red/" />
    <meta property="og:title" content="CoinRank" />
    <meta property="og:description" content='A powerful tool that allows users to rate and evaluate cryptocurrencies using the familiar "upvote" and "downvote" system.' />
    <meta property="og:image" content="https://coinrank.ptr.red/banner.png" />

    <meta property="twitter:card" content="summary_large_image" />
    <meta property="twitter:url" content="https://coinrank.ptr.red/" />
    <meta property="twitter:title" content="CoinRank" />
    <meta property="twitter:description" content='A powerful tool that allows users to rate and evaluate cryptocurrencies using the familiar "upvote" and "downvote" system.' />
    <meta property="twitter:image" content="https://coinrank.ptr.red/banner.png" />
  </Head>
);

export default MetaTags;
