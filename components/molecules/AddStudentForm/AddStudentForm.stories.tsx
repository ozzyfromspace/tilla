import { Meta, StoryObj } from '@storybook/react';
import AddStudentForm from './AddStudentForm';

const meta: Meta<typeof AddStudentForm> = {
  component: AddStudentForm,
  args: {},
};

type Story = StoryObj<typeof AddStudentForm>;

export const Primary: Story = {
  args: {},
};

export default meta;
