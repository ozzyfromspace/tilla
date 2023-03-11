import { Meta, StoryObj } from '@storybook/react';
import Header from './Header';

const meta: Meta<typeof Header> = {
  component: Header,
  args: {
    title: 'Khaleesi the Spartan',
    subtitle: "It's all about the dewey decimal system",
  },
};

type Story = StoryObj<typeof Header>;

export const Primary: Story = {
  args: {},
};

export default meta;
