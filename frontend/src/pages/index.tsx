import Coin from '@/components/Coin';
import fetcher from '@/handlers/fetcher';
import {
  Skeleton,
  Stack,
} from '@mantine/core';
import { Reorder } from 'framer-motion';
import { useCallback } from 'react';
import useSWR from 'swr';

interface CoinType {
  name: string;
  symbol: string;
  upvotes: number;
  downvotes: number;
}

const Home = () => {
  const { data, mutate } = useSWR<CoinType[]>(
    new URL('coins', process.env.NEXT_PUBLIC_API_URL).href,
    fetcher,
    {
      fallbackData: [...Array(16)].map((_, i) => ({
        name: '',
        symbol: `${i}`,
        upvotes: 0,
        downvotes: 0,
      })),
    },
  );

  const registerVote = useCallback(
    async (action: 'UP' | 'DOWN', symbol: string) => {
      const optimisticData = data
        .map((coin) => (coin.symbol === symbol ? {
          ...coin,
          upvotes: action === 'UP' ? coin.upvotes + 1 : coin.upvotes,
          downvotes: action === 'DOWN' ? coin.downvotes + 1 : coin.downvotes,
        } : coin))
        .sort((a, b) => (
          (b.upvotes - b.downvotes) - (a.upvotes - a.downvotes)
            || a.symbol.localeCompare(b.symbol)
        ));

      mutate(optimisticData, false);

      await fetch(
        new URL(
          `coins/${symbol}/${action === 'UP' ? 'upvote' : 'downvote'}`,
          process.env.NEXT_PUBLIC_API_URL,
        ),
        {
          method: 'POST',
        },
      );

      mutate(optimisticData);
    },
    [data, mutate],
  );

  return (
    <Stack>
      <Reorder.Group
        values={data}
        onReorder={() => {}}
        style={{
          display: 'contents',
          listStyle: 'none',
        }}
      >
        {data?.map((coin) => (
          <Reorder.Item
            key={coin.symbol}
            value={coin.symbol}
            dragListener={false}
          >
            <Skeleton visible={!coin.name}>
              <Coin
                name={coin.name}
                symbol={coin.symbol}
                upvotes={coin.upvotes}
                downvotes={coin.downvotes}
                registerVote={registerVote}
              />
            </Skeleton>
          </Reorder.Item>
        ))}
      </Reorder.Group>
    </Stack>
  );
};

export default Home;
