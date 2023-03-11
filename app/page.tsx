'use client';

import AddStudentForm from '@/components/molecules/AddStudentForm/AddStudentForm';
import AddSubjectForm from '@/components/molecules/AddSubjectForm/AddSubjectForm';
import AddTeacherForm from '@/components/molecules/AddTeacherForm/AddTeacherForm';
import { Tab } from '@headlessui/react';
import AddExcelDownloadForm from '../components/molecules/AddExcelDownloadForm/AddExcelDownloadForm';

const Home = () => {
  return (
    <div>
      <Tab.Group>
        <Tab.List>
          <Tab>Add Student</Tab>
          <Tab>Add Teacher</Tab>
          <Tab>Add Subject</Tab>
          <Tab>Download Excel</Tab>
        </Tab.List>
        <Tab.Panels>
          <Tab.Panel>
            <AddStudentForm />
          </Tab.Panel>
          <Tab.Panel>
            <AddTeacherForm />
          </Tab.Panel>
          <Tab.Panel>
            <AddSubjectForm />
          </Tab.Panel>
          <Tab.Panel>
            <AddExcelDownloadForm />
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
};

export default Home;
