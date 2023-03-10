import { Meta, StoryObj } from '@storybook/react';
import RowInput from './RowInput';

const meta: Meta<typeof RowInput> = {
  component: RowInput,
  args: {
    label: 'Label',
    inputValue: 'Some text',
    handleChange: undefined,
  },
};

type Story = StoryObj<typeof RowInput>;

export const Empty: Story = {
  args: {
    inputValue: '',
  },
};

export const Filled: Story = {};

export default meta;
