import WordMark from '@/components/WordMark';
import {
  ActionIcon, Group, Space, Text,
} from '@mantine/core';
import { FaGithub } from 'react-icons/fa';

const Header = () => (
  <header>
    <Group>
      <Text
        sx={(theme) => ({
          fontSize: `calc(${theme.fontSizes.xl} * 1.5)`,
        })}
      >
        <WordMark />
      </Text>
      <Space
        sx={{
          flexGrow: 1,
        }}
      />
      <ActionIcon
        component="a"
        aria-label="github"
        href="https://github.com/peterfritz/coinrank"
        target="_blank"
        rel="noopener noreferrer"
        size="lg"
      >
        <Text size="xl">
          <FaGithub />
        </Text>
      </ActionIcon>
    </Group>
  </header>
);

export default Header;
