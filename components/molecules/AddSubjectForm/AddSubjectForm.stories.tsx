import { Meta, StoryObj } from '@storybook/react';
import AddSubjectForm from './AddSubjectForm';

const meta: Meta<typeof AddSubjectForm> = {
  component: AddSubjectForm,
  args: {},
};

type Story = StoryObj<typeof AddSubjectForm>;

export const Primary: Story = {
  args: {},
};

export default meta;
