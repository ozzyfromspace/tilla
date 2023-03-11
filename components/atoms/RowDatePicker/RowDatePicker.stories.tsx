import { Meta, StoryObj } from '@storybook/react';
import RowDatePicker from './RowDatePicker';

const meta: Meta<typeof RowDatePicker> = {
  component: RowDatePicker,
  args: {
    label: 'Label',
    selectedDate: new Date('01/02/2003'),
    setSelectedDate: undefined,
  },
};

type Story = StoryObj<typeof RowDatePicker>;

export const Primary: Story = {};

export default meta;
