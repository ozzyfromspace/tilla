import { Meta, StoryObj } from '@storybook/react';
import { FauxButton } from './FauxButton';

const meta: Meta<typeof FauxButton> = {
  component: FauxButton,
  args: {
    label: 'Fake Button',
    selected: true,
  },
};

type Story = StoryObj<typeof FauxButton>;

export const Selected: Story = {
  args: {},
};

export const DeSelected: Story = {
  args: {
    selected: false,
  },
};

export default meta;
