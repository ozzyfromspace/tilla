'use client';

import Header from '@/components/atoms/Header/Header';
import AddStudentForm from '@/components/molecules/AddStudentForm/AddStudentForm';
import AddSubjectForm from '@/components/molecules/AddSubjectForm/AddSubjectForm';
import AddTeacherForm from '@/components/molecules/AddTeacherForm/AddTeacherForm';
import TabControls from '@/components/molecules/TabControls/TabControls';
import { Tab } from '@headlessui/react';
import AddExcelDownloadForm from '../components/molecules/AddExcelDownloadForm/AddExcelDownloadForm';

const Home = () => {
  return (
    <div className="flex flex-col bg-[hsl(264,72%,30%)] min-h-screen">
      <div className="flex flex-col justify-start items-center bg-violet-900 w-5/6 max-w-5xl mx-auto flex-1">
        <Header title="Eclipse Academy" subtitle="Track Students & Teachers" />
        <Tab.Group>
          <TabControls />
          <Tab.Panels className="flex justify-center items-center pt-6 w-full">
            <Tab.Panel className="w-full">
              <AddStudentForm />
            </Tab.Panel>
            <Tab.Panel className="w-full">
              <AddTeacherForm />
            </Tab.Panel>
            <Tab.Panel className="w-full">
              <AddSubjectForm />
            </Tab.Panel>
            <Tab.Panel>
              <AddExcelDownloadForm />
            </Tab.Panel>
          </Tab.Panels>
        </Tab.Group>
      </div>
      <p className="text-center text-sm font-light text-violet-50 bg-violet-900 w-5/6 max-w-5xl mx-auto pt-16 pb-6">
        Eclipse Academy 2023. All Rights Reserved
      </p>
    </div>
  );
};

export default Home;
