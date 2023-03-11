import axios from 'axios';
import { useCallback, useEffect, useState } from 'react';
import { Student } from './components/molecules/AddStudentForm/AddStudentForm';
import { BASE_URL } from './constants';

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

export const useStudents = () => {
  const [studentIds, setStudentIds] = useState<StudentId[]>(() => []);

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

  return {
    studentDataSlice: studentIds,
    selectedStudent: selectedPersonId,
    setSelectedStudent: setSelectedPersonId,
    fetchStudents,
  };
};
