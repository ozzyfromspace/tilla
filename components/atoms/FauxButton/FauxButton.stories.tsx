import { Meta, StoryObj } from '@storybook/react';
import { FauxButton } from './FauxButton';

const meta: Meta<typeof FauxButton> = {
  component: FauxButton,
  args: {
    label: 'Fake Button',
  },
};

type Story = StoryObj<typeof FauxButton>;

export const Primary: Story = {
  args: {},
};

export default meta;
