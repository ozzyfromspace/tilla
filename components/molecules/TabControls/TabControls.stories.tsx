import { Tab } from '@headlessui/react';
import { Meta, StoryObj } from '@storybook/react';
import TabControls from './TabControls';

const MockTabGroup = () => {
  return (
    <Tab.Group>
      <TabControls />
      <Tab.Panels className="flex justify-center items-center pt-6">
        <Tab.Panel>1</Tab.Panel>
        <Tab.Panel>2</Tab.Panel>
        <Tab.Panel>3</Tab.Panel>
        <Tab.Panel>4</Tab.Panel>
      </Tab.Panels>
    </Tab.Group>
  );
};

const meta: Meta<typeof MockTabGroup> = {
  component: MockTabGroup,
  args: {},
};

type Story = StoryObj<typeof MockTabGroup>;

export const Primary: Story = {
  args: {},
};

export default meta;
