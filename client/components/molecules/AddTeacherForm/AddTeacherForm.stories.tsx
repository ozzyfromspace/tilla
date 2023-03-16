import { Meta, StoryObj } from '@storybook/react';
import AddTeacherForm from './AddTeacherForm';

const meta: Meta<typeof AddTeacherForm> = {
  component: AddTeacherForm,
  args: {},
};

type Story = StoryObj<typeof AddTeacherForm>;

export const Primary: Story = {
  args: {},
};

export default meta;
