import { FauxButton } from '@/components/atoms/FauxButton/FauxButton';
import { Tab } from '@headlessui/react';

const TabControls = () => {
  return (
    <Tab.List className="flex flex-wrap gap-2 py-2 mx-auto">
      <Tab>
        <FauxButton label="Add Student" />
      </Tab>
      <Tab>
        <FauxButton label="Add Teacher" />
      </Tab>
      <Tab>
        <FauxButton label="Add Subject" />
      </Tab>
      <Tab>
        <FauxButton label="Get Session Logs" />
      </Tab>
    </Tab.List>
  );
};

export default TabControls;
