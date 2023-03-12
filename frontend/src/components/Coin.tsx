import {
  ActionIcon,
  Group, Paper, Stack, Text,
} from '@mantine/core';
import { useTimeout } from '@mantine/hooks';
import { useState } from 'react';
import { FaChevronDown, FaChevronUp } from 'react-icons/fa';

interface CoinProps {
  name: string;
  symbol: string;
  upvotes: number;
  downvotes: number;
  // eslint-disable-next-line no-unused-vars
  registerVote: (action: 'UP' | 'DOWN', symbol: string) => void;
}

const Coin = ({
  name,
  symbol,
  upvotes,
  downvotes,
  registerVote,
}: CoinProps) => {
  // Prevent the user from spamming the vote buttons
  const [ready, setReady] = useState(true);
  const { start } = useTimeout(() => {
    setReady(true);
  }, 250);

  const vote = (action: 'UP' | 'DOWN') => {
    if (!ready) {
      return;
    }

    registerVote(action, symbol);

    setReady(false);
    start();
  };

  return (
    <Paper
      withBorder
      p="sm"
    >
      <Group align="stretch">
        <Stack spacing="xs">
          <ActionIcon
            aria-label="upvote"
            onClick={() => {
              vote('UP');
            }}
            size="sm"
          >
            <Text aria-hidden size="xs">
              <FaChevronUp />
            </Text>
          </ActionIcon>
          <Text size="sm" align="center">
            {upvotes - downvotes}
          </Text>
          <ActionIcon
            aria-label="downvote"
            onClick={() => {
              vote('DOWN');
            }}
            size="sm"
          >
            <Text aria-hidden size="xs">
              <FaChevronDown />
            </Text>
          </ActionIcon>
        </Stack>
        <Group>
          <Text weight="bolder">{`${symbol}`}</Text>
          <Text>{name}</Text>
        </Group>
      </Group>
    </Paper>
  );
};

export default Coin;
