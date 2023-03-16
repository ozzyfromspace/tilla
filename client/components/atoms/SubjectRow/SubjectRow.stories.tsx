import { Meta, StoryObj } from '@storybook/react';
import SubjectRow from './SubjectRow';

const meta: Meta<typeof SubjectRow> = {
  component: SubjectRow,
  args: {
    rowIndex: 0,
  },
};

type Story = StoryObj<typeof SubjectRow>;

export const Primary: Story = {
  args: {},
};

export default meta;
