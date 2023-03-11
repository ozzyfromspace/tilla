'use client';

import { SubjectRowProps } from '@/components/atoms/SubjectRow/SubjectRow';
import AddStudentForm from '@/components/molecules/AddStudentForm/AddStudentForm';
import AddSubjectForm from '@/components/molecules/AddSubjectForm/AddSubjectForm';
import AddTeacherForm from '@/components/molecules/AddTeacherForm/AddTeacherForm';
import { BASE_URL } from '@/constants';
import { Tab } from '@headlessui/react';
import axios from 'axios';
import { useCallback, useEffect, useRef, useState } from 'react';
import AddExcelDownloadForm from '../components/molecules/AddExcelDownloadForm/AddExcelDownloadForm';

export interface Teacher {
  firstName: string;
  lastName: string;
  nickname: string;
}

export interface Student extends Teacher {
  calendarId: string;
}

export type TeacherField = keyof Teacher;
export type StudentField = keyof Student;

export type StudentId = Omit<Student, 'calendarId' | 'nickname'> & {
  id: string;
};

interface RespStudent {
  id: string;
  firstName: string;
  lastName: string;
  calendarId: string;
  nickname: string;
  subjects: Record<string, number>;
}

interface RespStudentArray {
  students: RespStudent[];
}

const Home = () => {
  const [studentIds, setStudentIds] = useState<StudentId[]>(() => []);

  const [student, setStudent] = useState<Student>(() => ({
    firstName: '',
    lastName: '',
    nickname: '',
    calendarId: '',
  }));

  const [selectedPersonId, setSelectedPersonId] = useState<StudentId>(
    studentIds[0] ?? { firstName: '', lastName: '', id: '' }
  );

  const fetchStudents = useCallback(async () => {
    axios.get<RespStudentArray>(`${BASE_URL}/students`).then((resp) => {
      if (resp.status !== 200) {
        console.log('error.', resp.statusText);
        return;
      }

      setStudentIds(() => {
        const newStudentIds: StudentId[] = [];

        for (const studentData of resp.data.students) {
          const idObj: StudentId = {
            firstName: studentData.firstName,
            lastName: studentData.lastName,
            id: studentData.id,
          };

          newStudentIds.push(idObj);
        }

        if (newStudentIds.length) {
          setSelectedPersonId(() => ({ ...newStudentIds[0] }));
        }

        return newStudentIds;
      });
    });
  }, []);

  useEffect(() => {
    fetchStudents();
  }, [fetchStudents]);

  const [teacher, setTeacher] = useState<Teacher>(() => ({
    firstName: '',
    lastName: '',
    nickname: '',
  }));

  const [subjectRows, setSubjectRows] = useState<SubjectRowProps[]>(() => []);

  const [dynamicLink, setDynamicLink] = useState(() => '');
  const linkRef = useRef<HTMLAnchorElement | null>(null);

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
            <AddStudentForm
              student={student}
              fetchStudents={fetchStudents}
              setStudent={setStudent}
            />
          </Tab.Panel>
          <Tab.Panel>
            <AddTeacherForm teacher={teacher} setTeacher={setTeacher} />
          </Tab.Panel>
          <Tab.Panel>
            <AddSubjectForm
              selectedPersonId={selectedPersonId}
              setSelectedPersonId={setSelectedPersonId}
              setSubjectRows={setSubjectRows}
              studentIds={studentIds}
              subjectRows={subjectRows}
            />
          </Tab.Panel>
          <Tab.Panel>
            <AddExcelDownloadForm
              linkRef={linkRef}
              setDynamicLink={setDynamicLink}
            />
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
      <a href={dynamicLink} ref={linkRef}></a>
    </div>
  );
};

export default Home;
