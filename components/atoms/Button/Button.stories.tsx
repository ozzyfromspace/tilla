import { Meta, StoryObj } from '@storybook/react';
import Button from './Button';

const meta: Meta<typeof Button> = {
  component: Button,
  args: {
    label: 'Button Text',
  },
};

type Story = StoryObj<typeof Button>;

export const Primary: Story = {};

export default meta;