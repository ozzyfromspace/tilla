import { FauxButton } from '@/components/atoms/FauxButton/FauxButton';
import { Tab } from '@headlessui/react';

const TabControls = () => {
  return (
    <Tab.List className="flex flex-wrap gap-2 py-2 mx-auto">
      <Tab>
        {({ selected }) => (
          <FauxButton label="Add Student" selected={selected} />
        )}
      </Tab>
      <Tab>
        {({ selected }) => (
          <FauxButton label="Add Teacher" selected={selected} />
        )}
      </Tab>
      <Tab>
        {({ selected }) => (
          <FauxButton label="Add Course" selected={selected} />
        )}
      </Tab>
      <Tab>
        {({ selected }) => (
          <FauxButton label="Get Session Logs" selected={selected} />
        )}
      </Tab>
    </Tab.List>
  );
};

export default TabControls;
