import { Meta, StoryObj } from '@storybook/react';
import AddExcelDownloadForm from './AddExcelDownloadForm';

const meta: Meta<typeof AddExcelDownloadForm> = {
  component: AddExcelDownloadForm,
  args: {},
};

type Story = StoryObj<typeof AddExcelDownloadForm>;

export const Primary: Story = {
  args: {},
};

export default meta;
