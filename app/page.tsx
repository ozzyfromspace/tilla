'use client';

import Button from '@/components/atoms/Button/Button';
import RowDatePicker from '@/components/atoms/RowDatePicker/RowDatePicker';
import { SubjectRowProps } from '@/components/atoms/SubjectRow/SubjectRow';
import { Combobox, Tab } from '@headlessui/react';
import axios from 'axios';
import { Abril_Fatface } from 'next/font/google';
import {
  ChangeEvent,
  Dispatch,
  FormEvent,
  MouseEvent,
  SetStateAction,
  useCallback,
  useEffect,
  useRef,
  useState,
} from 'react';
import { v4 as uuidv4 } from 'uuid';
import RowInput from '../components/atoms/RowInput/RowInput';
import SubjectRowGrid from '../components/molecules/SubjectRowGrid/SubjectRowGrid';

const fatFace = Abril_Fatface({ weight: ['400'], subsets: ['latin'] });

interface Teacher {
  firstName: string;
  lastName: string;
  nickname: string;
}

interface Student extends Teacher {
  calendarId: string;
}

type TeacherField = keyof Teacher;
type StudentField = keyof Student;

type StudentId = Omit<Student, 'calendarId' | 'nickname'> & { id: string };

function toFullName(studentId: StudentId): string {
  return `${studentId.firstName} ${studentId.lastName}`;
}

function toBasicISO(date: Date): string {
  const parts = date.toISOString().split('T');
  const timepart = `${parts[1].split('.')[0]}Z`;
  return `${parts[0]}T${timepart}`;
}

const BASE_URL = 'http://localhost:8080';

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

  const handleStudentChange = (inputField: StudentField) => {
    return (e: ChangeEvent<HTMLInputElement>) => {
      setStudent((s) => ({ ...s, [inputField]: e.target.value }));
    };
  };

  const [teacher, setTeacher] = useState<Teacher>(() => ({
    firstName: '',
    lastName: '',
    nickname: '',
  }));

  const handleTeacherChange = (inputField: TeacherField) => {
    return (e: ChangeEvent<HTMLInputElement>) => {
      setTeacher((t) => ({ ...t, [inputField]: e.target.value }));
    };
  };

  const handleSetSelectedPersonId = (fieldName: string) => {
    for (let i = 0; i < studentIds.length; i++) {
      console.log('in loop', i, studentIds, studentIds[0]);
      const studentName = toFullName(studentIds[i]);

      if (studentName === fieldName) {
        setSelectedPersonId(() => studentIds[i]);
        return;
      }
    }
  };

  const [query, setQuery] = useState('');

  const filteredPeople =
    query === ''
      ? studentIds
      : studentIds.filter((studentId) => {
          return toFullName(studentId)
            .toLowerCase()
            .includes(query.toLowerCase());
        });

  const [startDate, setStartDate] = useState(() => {
    const d = new Date();
    d.setHours(0, 0, 0, 0);
    return d;
  });

  const [endDate, setEndDate] = useState(() => {
    const d = new Date();
    d.setHours(23, 59, 59, 0);
    return d;
  });

  const handleDateSelection = (
    newDateSetter: Dispatch<SetStateAction<Date>>,
    refDate: Date,
    startingDate: boolean
  ) => {
    return (newDate: Date | null) => {
      if (!newDate) return;

      if (startingDate) {
        const refCopy = newDate;
        refCopy.setHours(0, 0, 0, 0);

        if (refDate.getTime() < refCopy.getTime()) return;
        newDateSetter(() => refCopy);
        return;
      }

      const refCopy = newDate;
      refCopy.setHours(23, 59, 59, 0);

      if (refDate.getTime() > refCopy.getTime()) return;

      newDateSetter(() => refCopy);
    };
  };

  const [subjectRows, setSubjectRows] = useState<SubjectRowProps[]>(() => []);

  function handleCreateSubjectRow(
    e: MouseEvent<HTMLButtonElement, globalThis.MouseEvent>
  ) {
    e.preventDefault();
    e.stopPropagation();

    const newId = uuidv4();
    const getNewRowProps = (rowIndex: number): SubjectRowProps => ({
      fieldName: '',
      price: '',
      rowIndex: rowIndex,
      id: newId,
      handleDelete: () => {},
      handleFieldChange: () => () => {},
    });

    setSubjectRows((subs) => {
      return [...subs, getNewRowProps(subs.length + 1)];
    });
  }

  function submitSubjects(
    e: MouseEvent<HTMLFormElement, globalThis.MouseEvent>
  ) {
    interface Subject {
      name: string;
      pricePerHour: number;
    }

    interface Payload {
      studentId: string;
      subjects: Subject[];
    }

    e.preventDefault();
    if (!selectedPersonId.id) {
      console.log('no student id');
      return;
    } else {
      console.log('made it!', selectedPersonId);
    }

    function coerse(si: SubjectRowProps[]): Subject[] {
      const subs: Subject[] = [];

      for (const row of si) {
        if (!row.fieldName || !row.price) continue;

        const sub: Subject = {
          name: row.fieldName,
          pricePerHour: parseFloat(parseFloat(row.price).toFixed(2)),
        };

        subs.push(sub);
      }

      return subs;
    }

    const payload: Payload = {
      studentId: selectedPersonId.id,
      subjects: coerse(subjectRows),
    };

    if (payload.subjects.length === 0) {
      console.log('no valid subject entries');
      return;
    }

    axios.post(`${BASE_URL}/student/subjects`, payload).then((resp) => {
      if (resp.status === 201) {
        setSubjectRows(() => []);
      }
    });
  }

  const handleAddTeacherSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!teacher.firstName) {
      console.log('no teacher firstname was provided');
      return;
    }

    if (!teacher.lastName) {
      console.log('no teacher lastname was provided');
      return;
    }

    if (!teacher.nickname) {
      console.log('no teacher nickname was provided');
      return;
    }

    axios.post(`${BASE_URL}/teacher`, teacher).then((resp) => {
      console.log(resp);

      if (resp.status === 201) {
        setTeacher(() => ({
          firstName: '',
          lastName: '',
          nickname: '',
        }));
      }
    });
  };

  const handleAddStudentSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!student.firstName) {
      console.log('no student firstname was provided');
      return;
    }

    if (!student.lastName) {
      console.log('no student lastname was provided');
      return;
    }

    if (!student.nickname) {
      console.log('no student nickname was provided');
      return;
    }

    if (!student.calendarId) {
      console.log('no student calendar id was provided');
      return;
    }

    axios.post(`${BASE_URL}/student`, student).then((resp) => {
      console.log(resp);

      if (resp.status === 201) {
        fetchStudents();

        setStudent(() => ({
          firstName: '',
          lastName: '',
          nickname: '',
          calendarId: '',
        }));
      }
    });
  };

  const [dynamicLink, setDynamicLink] = useState(() => '');
  const linkRef = useRef<HTMLAnchorElement | null>(null);

  const handleDownloadRequest = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    interface Payload {
      minLocalTime: string;
      maxLocalTime: string;
    }

    interface GoResponse {
      filename: string;
      msg: string;
    }

    const payload: Payload = {
      minLocalTime: toBasicISO(startDate),
      maxLocalTime: toBasicISO(endDate),
    };

    axios.post<GoResponse>(`${BASE_URL}/excel`, payload).then((resp) => {
      if (resp.status === 201) {
        // celebrate!
        console.log(resp.data);
        setDynamicLink(() => `${BASE_URL}/excel/${resp.data.filename}`);
        setTimeout(() => {
          linkRef.current?.click();
        }, 10);
      }
    });
  };

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
            <form onSubmit={handleAddStudentSubmit}>
              <RowInput
                label="First Name"
                inputValue={student.firstName}
                handleChange={handleStudentChange('firstName')}
              />
              <RowInput
                label="Last Name"
                inputValue={student.lastName}
                handleChange={handleStudentChange('lastName')}
              />
              <RowInput
                label="Nickname"
                inputValue={student.nickname}
                handleChange={handleStudentChange('nickname')}
              />
              <RowInput
                label="Calendar Id"
                inputValue={student.calendarId}
                handleChange={handleStudentChange('calendarId')}
              />
              <Button label="Add Student" />
            </form>
          </Tab.Panel>
          <Tab.Panel>
            <form onSubmit={handleAddTeacherSubmit}>
              <RowInput
                label="First Name"
                inputValue={teacher.firstName}
                handleChange={handleTeacherChange('firstName')}
              />
              <RowInput
                label="Last Name"
                inputValue={teacher.lastName}
                handleChange={handleTeacherChange('lastName')}
              />
              <RowInput
                label="Nickname"
                inputValue={teacher.nickname}
                handleChange={handleTeacherChange('nickname')}
              />
              <Button label="Add Teacher" />
            </form>
          </Tab.Panel>
          <Tab.Panel>
            <form onSubmit={submitSubjects}>
              <Combobox
                value={toFullName(selectedPersonId)}
                onChange={handleSetSelectedPersonId}
              >
                <Combobox.Input
                  onChange={(event) => setQuery(event.target.value)}
                />
                <Combobox.Options>
                  {filteredPeople.map((person) => (
                    <Combobox.Option key={person.id} value={toFullName(person)}>
                      {toFullName(person)}
                    </Combobox.Option>
                  ))}
                </Combobox.Options>
              </Combobox>
              <Button
                label="New Subject"
                type="button"
                onClick={handleCreateSubjectRow}
              />
              <SubjectRowGrid
                subjectRows={subjectRows}
                setSubjectRows={setSubjectRows}
              />
              <Button
                label="Save"
                type="submit"
                onClick={() => console.log(subjectRows, selectedPersonId)}
              />
            </form>
          </Tab.Panel>
          <Tab.Panel>
            <form onSubmit={handleDownloadRequest}>
              <RowDatePicker
                label="Start Date"
                selectedDate={startDate}
                setSelectedDate={handleDateSelection(
                  setStartDate,
                  endDate,
                  true
                )}
              />
              <RowDatePicker
                label="End Date"
                selectedDate={endDate}
                setSelectedDate={handleDateSelection(
                  setEndDate,
                  startDate,
                  false
                )}
              />
              <Button label="Download Excel File" />
            </form>
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
      <a href={dynamicLink} ref={linkRef}></a>
    </div>
  );
};

export default Home;
